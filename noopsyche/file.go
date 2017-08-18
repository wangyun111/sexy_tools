package noopsyche

import (
	"os"
	"strings"
)

const (
	lineFeed = "\n"
	// show  create table `users`
)

//@Sumary    获取表最长字段
//@Param     dataBase   数据库名称
//@Param     tableName  表名称
//@Success   查询 sql
func GetMaxColumns(dataBase, tableName string) (maxColumnsSql string) {
	maxColumnsSql = `SELECT COLUMN_NAME  FROM information_schema.COLUMNS  ` + lineFeed +
		`    WHERE table_schema='` + dataBase + `' AND TABLE_NAME = '` + tableName + `' ` + lineFeed +
		`    ORDER BY length(COLUMN_NAME) DESC LIMIT 1`
	return maxColumnsSql
}

//@Sumary         自动构建表的 model 层
//@Description    最长字段定义 json 与 bson 或与 description 的空格，默认30
//@Param          filePath   需要创建的路径(注意带上/)
//@Param          sqlStr     创建表的 sql
//@Param          maxColunm  最长字段
//@Param          isBson     是否显示 bson 列
//@Param          isJson     是否显示 json 列
//@Success        查询 sql
func NoppsycheModelFile(filePath, sqlStr, maxColunm string, isBson, isJson bool) (msg string) {
	_, structName, tableName, _, resultStruct := NewNoopsycheCreateStruct(sqlStr, maxColunm, isBson, isJson)
	fileName := filePath + tableName + ".go"
	_, err := os.Stat(fileName)
	if err == nil {
		msg = "file already exists"
		return
	}

	fileContent := `
	    package models
		import (
			"github.com/astaxie/beego/orm"
			"sexy_tools/time"
		)`
	fileContent += lineFeed + resultStruct
	structParams := "_" + structName
	addMethod := `//添加
		func Add` + structName + `(` + structParams + " " + structName + `) (id int64, err error) {
			o := orm.NewOrm()
			` + structParams + `.Id = 0
			` + structParams + `.CreateTime = time.GetNowTime()
			id, err = o.Insert(&` + structParams + `)
			return
		}`
	fileContent += lineFeed + addMethod
	deleteByIdMethod := `//根据id删除
		func DeleteById` + structName + `(id int64) (number int64, err error) {
			o := orm.NewOrm()
			sql := "delete from ` + tableName + ` where id=? "
			res, err := o.Raw(sql, id).Exec()
			if err == nil {
				number, _ = res.RowsAffected()
			}
			return
		}`
	fileContent += lineFeed + deleteByIdMethod
	deleteMethod := `//条件删除
		func Delete` + structName + `(where string, args ...interface{}) (number int64, err error) {
			o := orm.NewOrm()
			sql := "delete from ` + tableName + ` where 1=1 "
			sql += where
			res, err := o.Raw(sql, args).Exec()
			if err == nil {
				number, _ = res.RowsAffected()
			}
			return
		}`
	fileContent += lineFeed + deleteMethod
	updateByIdMethod := `//根据 id 进行修改
		func UpdateById` + structName + `(` + structParams + ` ` + structName + `) (number int64, err error) {
			o := orm.NewOrm()
			number, err = o.Update(&` + structParams + `)
			return
		}`
	fileContent += lineFeed + updateByIdMethod
	updateMethod := `//条件修改
		func Update` + structName + `(set, where string, args ...interface{}) (number int64, err error) {
			o := orm.NewOrm()
			sql := "update ` + tableName + ` set 1=1 "
			sql += set
			sql += where
			res, err := o.Raw(sql, args).Exec()
			if err == nil {
				number, _ = res.RowsAffected()
			}
			return
		}`
	fileContent += lineFeed + updateMethod
	queryMethod := `//条件查询列表
		func Query` + structName + `(pageNo, pageSize int64, sort, where string, args ...interface{}) (list []` + structName + `, err error) {
			o := orm.NewOrm()
			sql := "select *  from ` + tableName + ` where 1=1 "
			sql += where
			if sort != "" {
				sql += " order by ? desc"
			}
			sql += " limit ?,? "
			_, err = o.Raw(sql, args, (pageNo-1)*pageSize, pageSize).QueryRows(&list)
			return
		}`
	fileContent += lineFeed + queryMethod
	QueryCountMethod := `//条件查询总条数
		func Query` + structName + `Count(where string, args ...interface{}) (count int64, err error) {
			o := orm.NewOrm()
			sql := "select count(id) from ` + tableName + ` where 1=1 "
			sql += where
			err = o.Raw(sql, args).QueryRow(&count)
			return
		}`
	fileContent += lineFeed + QueryCountMethod
	FindConditionMethod := `//条件查询单个
		func FindByCondition` + structName + `(where string, args ...interface{}) (` + structParams + " " + structName + `, err error) {
			o := orm.NewOrm()
			sql := "select * from ` + tableName + ` where 1=1 "
			sql += where
			err = o.Raw(sql, args).QueryRow(&` + structParams + `)
			return
		}`
	fileContent += lineFeed + FindConditionMethod
	FindByIdMethod := `//根据 id查询单个
		func FindById` + structName + `(id int64) (` + structParams + " " + structName + `, err error) {
			o := orm.NewOrm()
			qs := o.QueryTable("` + tableName + `")
			err = qs.Filter("id", id).One(&` + structParams + `)
			return
		}`
	fileContent += lineFeed + FindByIdMethod
	f, _ := os.Create(fileName)
	f.WriteString(fileContent)
	f.Sync()
	f.Close()
	msg = "create file success"
	return
}

//@Sumary         根据 sql 创建结构体
//@Description    最长字段定义 json 与 bson 或与 description 的空格，默认30
//@Param          filePath   需要创建的路径(注意带上/)
//@Param          sqlStr     创建表的 sql
//@Param          maxColunm  最长字段
//@Param          isBson     是否显示 bson 列
//@Param          isJson     是否显示 json 列
//@Success        查询 sql
func NewNoopsycheCreateStruct(sqlStr, maxColunm string, isBson, isJson bool) (tebaleRemark, structName, tableName, allColunm, resultStruct string) {
	defer func() {
		if allColunm != "" {
			allColunm = allColunm[:len(allColunm)-1]
		}
	}()
	spaceStr := " "
	spaceNumber := len(maxColunm)
	if spaceNumber == 0 {
		spaceNumber = 30
	}
	var lastCommentIndex int
	beginIndex := strings.IndexAny(sqlStr, "(")
	titleStr := sqlStr[:beginIndex] //表名
	tableArr := strings.Fields(titleStr)
	tableName = tableArr[len(tableArr)-1]
	tableArr = strings.Split(tableName, "_")
	for i := 0; i < len(tableArr); i++ {
		structName += strings.ToUpper(tableArr[i][:1]) + tableArr[i][1:]
	}
	lastBracketIndex := strings.LastIndexAny(sqlStr, ")")
	lastStr := sqlStr[lastBracketIndex+1:]
	lastCommentIndex = strings.Index(lastStr, "COMMENT='")
	if lastCommentIndex == -1 {
		lastCommentIndex = strings.Index(lastStr, "comment='")
	}
	if lastCommentIndex != -1 {
		tebaleRemark = lastStr[lastCommentIndex+9 : strings.LastIndex(lastStr, "'")]
	}
	resultStruct += "//" + tebaleRemark + "\n"
	resultStruct += "type " + structName + " struct { \n"
	endIndex := strings.Index(sqlStr, "PRIMARY KEY")
	sqlStr = sqlStr[beginIndex+1 : endIndex]
	sqlStr = strings.TrimSpace(sqlStr)
	sqlStr = sqlStr[:len(sqlStr)-1]
	arr := strings.Split(sqlStr, ",\n")
	var rowCount int
	for i := 0; i < len(arr); i++ {
		columns := arr[i]
		spacingStr := strings.TrimSpace(columns)
		spaceArr := strings.Fields(spacingStr) //以空格分割
		var sname, stype, scomment, sremark string
		spaceArrLen := len(spaceArr)
		if spaceArrLen < 4 {
			continue
		} else {
			sname = spaceArr[0]
			allColunm += sname + ","
			allColunmLen := len(allColunm)
			if allColunmLen-rowCount > 70 {
				allColunm += "\n"
				rowCount = allColunmLen
			}
			snameArr := strings.Split(sname, "_")
			snameArrLen := len(snameArr)
			if snameArrLen > 0 {
				snameStr := ""
				for i := 0; i < snameArrLen; i++ {
					obj := snameArr[i]
					snameStr += strings.ToUpper(obj[:1]) + obj[1:]
				}
				sname = snameStr
			}
			stype = spaceArr[1]
			sIndex := strings.Index(stype, "(")
			if sIndex != -1 {
				stype = stype[:sIndex]
			}
			scomment = spaceArr[spaceArrLen-2]
			if strings.EqualFold("COMMENT", scomment) {
				sremark = strings.Replace(spaceArr[spaceArrLen-1], "'", "", -1)
			}
		}
		if strings.EqualFold("TINYINT", stype) || strings.EqualFold("SMALLINT", stype) || strings.EqualFold("MEDIUMINT", stype) ||
			strings.EqualFold("INT", stype) || strings.EqualFold("INTEGER", stype) || strings.EqualFold("BIGINT", stype) {
			stype = "int"
		} else if strings.EqualFold("FLOAT", stype) || strings.EqualFold("DECIMAL", stype) || strings.EqualFold("DOUBLE", stype) {
			stype = "float64"
		} else {
			stype = "string"
		}
		snameLen := len(sname)
		diffNum := spaceNumber - snameLen
		var addSpace string = spaceStr
		if diffNum > 0 {
			for i := 0; i < diffNum; i++ {
				addSpace += spaceStr
			}
		}
		jsonSname := strings.ToLower(sname[:1]) + sname[1:]
		rows := sname + " " + stype + "`"
		if isBson {
			rows += "bson:\"" + jsonSname + "\"" + addSpace
		}
		if isJson {
			rows += "json:\"" + jsonSname + "\"" + addSpace
		}
		rows += "description:\"" + sremark + "\"`"
		resultStruct += rows + "\n"
	}
	resultStruct += "}"
	return
}
