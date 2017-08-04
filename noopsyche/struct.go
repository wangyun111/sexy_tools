package noopsyche

import (
	"fmt"
	"strings"
)

//根据 sql 创建结构体
func NoopsycheCreateStruct(str, maxColunm string, isBson, isJson bool) (tebaleRemark, tableName, allColunm, resultStruct string) {
	defer func() {
		if allColunm != "" {
			allColunm = allColunm[:len(allColunm)-1]
		}
	}()
	spaceStr := " "
	spaceNumber := len(maxColunm)
	var lastCommentIndex int
	beginIndex := strings.IndexAny(str, "(")
	titleStr := str[:beginIndex]
	tableArr := strings.Fields(titleStr)
	tableName = tableArr[len(tableArr)-1]
	tableArr = strings.Split(tableName, "_")
	var finalTableName string
	for i := 0; i < len(tableArr); i++ {
		finalTableName += strings.ToUpper(tableArr[i][:1]) + tableArr[i][1:]
	}
	tableName = finalTableName
	lastBracketIndex := strings.LastIndexAny(str, ")")
	lastStr := str[lastBracketIndex+1:]
	lastCommentIndex = strings.Index(lastStr, "COMMENT='")
	if lastCommentIndex == -1 {
		lastCommentIndex = strings.Index(lastStr, "comment='")
	}
	if lastCommentIndex != -1 {
		tebaleRemark = lastStr[lastCommentIndex+9 : strings.LastIndex(lastStr, "'")]
	}
	resultStruct += "//" + tebaleRemark + "\n"
	resultStruct += "type " + tableName + " struct { \n"
	endIndex := strings.Index(str, "PRIMARY KEY")
	str = str[beginIndex+1 : endIndex]
	str = strings.TrimSpace(str)
	str = str[:len(str)-1]
	arr := strings.Split(str, ",\n")
	var rowCount int
	for i := 0; i < len(arr); i++ {
		columns := arr[i]
		spacingStr := strings.TrimSpace(columns)
		spaceArr := strings.Fields(spacingStr) //以空格分割
		var sname, stype, scomment, sremark string
		spaceArrLen := len(spaceArr)
		if spaceArrLen < 4 {
			fmt.Println("不是字段,跳出循环")
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
