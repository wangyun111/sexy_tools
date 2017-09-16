package noopsyche

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
)

var engine *xorm.Engine
var dataBaseUrl = "root:123456@tcp(127.0.0.1:3306)/wangyun?charset=utf8mb4&loc=Asia%2FShanghai"

//@Summary 通过数据库获取全部表或指定表的结构体
//@Param   tableName 需要获取的表名
func InitializeDateBaseStruct(tableName ...string) (initializeStruct, errmsg string) {
	engine, err := xorm.NewEngine("mysql", dataBaseUrl)
	dataBase := dataBaseUrl[strings.Index(dataBaseUrl, "/")+1 : strings.Index(dataBaseUrl, "?")]
	if err != nil {
		errmsg += "初始化数据库失败,errmsg:=" + err.Error()
		return
	} else {
		tablesMap, err := engine.QueryString("show tables from " + dataBase)
		if err != nil {
			errmsg += "获取数据库表列表失败,errmsg:=" + err.Error()
			return
		}
		if len(tablesMap) == 0 {
			errmsg += "数据库表为空,请创建后重试。"
			return
		}
		tNumber := len(tableName)
		if tNumber > 0 {
			for _, val := range tableName {
				tSmap, err := engine.QueryString("SHOW CREATE TABLE " + val)
				if err != nil {
					errmsg += "查询表创建sql失败:errmsg:=" + err.Error()
					continue
				}
				createSql := tSmap[0]["Create Table"]
				createSql = strings.Replace(createSql, "`", "", -1)
				maxColumnSql := GetMaxColumns(dataBase, val)
				var maxColumn string
				_, err = engine.Sql(maxColumnSql).Get(&maxColumn)
				if err != nil {
					errmsg += "获取" + val + "表最长列失败,errmsg:=" + err.Error()
				}
				_, _, _, _, structStr := NewNoopsycheCreateStruct(createSql, maxColumn, false, false)
				initializeStruct += structStr + lineFeed
			}
		} else {
			for _, t := range tablesMap {
				for _, val := range t {
					tSmap, err := engine.QueryString("SHOW CREATE TABLE " + val)
					if err != nil {
						errmsg += "查询表创建sql失败:errmsg:=" + err.Error()
						continue
					}
					createSql := tSmap[0]["Create Table"]
					createSql = strings.Replace(createSql, "`", "", -1)
					maxColumnSql := GetMaxColumns(dataBase, val)
					var maxColumn string
					_, err = engine.Sql(maxColumnSql).Get(&maxColumn)
					if err != nil {
						errmsg += "获取" + val + "表最长列失败,errmsg:=" + err.Error()
					}
					_, _, _, _, structStr := NoopsycheCreateStruct(createSql, maxColumn, false, false)
					initializeStruct += structStr + lineFeed
				}
			}
		}
	}
	return
}

func NoopsycheCreateStruct(sqlStr, maxColunm string, isBson, isJson bool) (tebaleRemark, structName, tableName, allColunm, resultStruct string) {
	defer func() {
		if allColunm != "" {
			allColunm = allColunm[:len(allColunm)-1]
		}
	}()
	spaceStr := " "
	spaceNumber := len(maxColunm)
	if spaceNumber == 0 {
		spaceNumber = 15
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
	if lastCommentIndex != -1 {
		tebaleRemark = lastStr[lastCommentIndex+9 : strings.LastIndex(lastStr, "'")]
		resultStruct += "//" + tebaleRemark + "\n"
	}
	resultStruct += "type " + structName + " struct { \n"
	endIndex := strings.Index(sqlStr, "PRIMARY KEY")
	if endIndex == -1 {
		endIndex = strings.LastIndex(sqlStr, ")")
	}
	sqlStr = sqlStr[beginIndex+1 : endIndex]
	sqlStr = strings.TrimSpace(sqlStr)
	sqlStr = sqlStr[:len(sqlStr)-1]
	arr := strings.Split(sqlStr, ",\n")
	var rowCount int
	for i := 0; i < len(arr); i++ {
		columns := arr[i]
		spacingStr := strings.TrimSpace(columns)
		spaceArr := strings.Fields(spacingStr) //以空格分割
		var sname, stype, sremark string
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
			cIndex := strings.Index(columns, "COMMENT")
			if cIndex != -1 {
				sremark = strings.Replace(strings.TrimSpace(columns[cIndex+7:]), "'", "", -1)
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
