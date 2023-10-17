package model

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const MysqlConfig = "gacha:ackm4WD6HDDPRx8A@(13.229.215.81:3306)/gacha?charset=utf8mb4&parseTime=True&loc=Local"
const Instance = "gacha"

func TestCreateModelFile(t *testing.T) {

	//第一个参数为表后缀
	//第二个参数为表名称；单表直接表名;多表中间使用,分隔;所有表填空
	CreateModelFile("", "", "prize")
}

type ModelInfo struct {
	BDName          string
	TablePrefixName string
	TableName       string
	PackageName     string
	ModelName       string
	TableSchema     *[]TABLE_SCHEMA
	SpecialStr      string
	IsHaveId        bool
}

type TABLE_SCHEMA struct {
	Field string `db:"Field" json:"Field"`
	Type  string `db:"Type" json:"Type"`
	Key   string `db:"Key" json:"Key"`
	Extra string `db:"Extra" json:"Extra"`
}

func (m *ModelInfo) ColumnNames() []string {
	result := make([]string, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		result = append(result, t.Field)
	}
	return result
}

func (m *ModelInfo) ColumnCount() int {
	return len(*m.TableSchema)
}

func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}

func FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func Tags(ts TABLE_SCHEMA) template.HTML {

	//判断是否为主键
	keyStr := ""
	typeStr := ts.Type
	if ts.Key == "PRI" {
		keyStr = "primaryKey;"
		typeStr = ""
	}
	return template.HTML("`json:" + `"` + ts.Field + `"` +
		" gorm:" + `"` + keyStr + `column:` + ts.Field + ";" + typeStr + "\"`")
}

func ExportColumn(columnName string) string {
	return MarshalFirstLetterToUpper(columnName)
}

func inArray(slice []string, str string, ret string) (string, error) {
	for _, v := range slice {
		if strings.Contains(str, v) {
			return ret, nil
		}
	}
	return "", errors.New("Not found value")
}

func TypeConvert(ts TABLE_SCHEMA) string {

	str := ts.Type
	if ts.Field == "deleted_at" {
		return "gorm.DeletedAt"
	}
	sliceBool := []string{"tinyint(1)"}
	if value, ok := inArray(sliceBool, str, "bool"); ok == nil {
		return value
	}
	sliceInt8 := []string{"smallint", "tinyint"}
	if value, ok := inArray(sliceInt8, str, "int8"); ok == nil {
		return value
	}

	sliceDate := []string{"timestamp", "datetime", "datetime(3)"}
	if value, ok := inArray(sliceDate, str, "time.Time"); ok == nil {
		return value
	}

	sliceStr := []string{"varchar", "text", "longtext", "char", "date", "enum"}
	if value, ok := inArray(sliceStr, str, "string"); ok == nil {
		return value
	}

	sliceBig := []string{"bigint"}
	if value, ok := inArray(sliceBig, str, "int64"); ok == nil {
		return value
	}

	sliceFlo := []string{"float", "double", "decimal"}
	if value, ok := inArray(sliceFlo, str, "float64"); ok == nil {
		return value
	}

	sliceInt := []string{"int"}
	if value, ok := inArray(sliceInt, str, "int"); ok == nil {
		return value
	}

	return str
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func ColumnAndType(table_schema []TABLE_SCHEMA) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.Field+" "+TypeConvert(t))
	}
	return strings.Join(result, ",")
}

func ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+Postfix)
	}
	return strings.Join(result, sep)
}

func (m *ModelInfo) CheckFirstTable() string {

	getTablesNameSql := "show tables from " + Instance
	tablaNames, _ := NewEngineInstance().QueryString(getTablesNameSql)
	tableFirst := tablaNames[0]["Tables_in_"+Instance]
	return tableFirst

}

func NewEngineInstance() *xorm.Engine {

	NewEngine, _ := xorm.NewEngine("mysql", MysqlConfig)
	NewEngine.DB().SetConnMaxLifetime(time.Duration(86400) * time.Second)
	return NewEngine
}

func createModelFile(render *template.Template, dbName, tableName, tablePreFix, tableLastFix string) {
	tableSchema := []TABLE_SCHEMA{}
	err := NewEngineInstance().SQL(
		"show columns from " + tableName + " from " + dbName).Find(&tableSchema)

	if err != nil {
		log.Fatal("sql error :", err.Error())
		return
	}
	//配置项：表后缀
	if tableLastFix != "" {
		count := len(tableName) - len(tableLastFix)
		tableName = tableName[:count]
	}
	if tablePreFix != "" {
		count := len(tablePreFix)
		tableName = tableName[count:]
	}
	AppPath, _ := os.Getwd()
	//配置项：生成model文件的文件夹
	modelFolder := AppPath + "/"
	fileName := modelFolder + strings.ToLower(tableName) + ".go"
	_ = os.Remove(fileName)
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal("create file error :", err.Error())
	}
	defer f.Close()
	//判断表是否存在id
	isHaveId := false
	for _, v := range tableSchema {
		if v.Field == "id" {
			isHaveId = true
		}
	}

	newTableName := MarshalFirstLetterToUpper(tableName)
	model := &ModelInfo{
		PackageName:     "model",
		BDName:          dbName,
		TablePrefixName: tablePreFix + tableName + tableLastFix,
		TableName:       tableName,
		ModelName:       newTableName,
		TableSchema:     &tableSchema,
		SpecialStr:      "<",
		IsHaveId:        isHaveId,
	}

	if err := render.Execute(f, model); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileName)
	cmd := exec.Command("goimports", "-w", fileName)
	cmd.Run()
}

/*
转换为大驼峰命名法则
首字母大写，“_” 忽略后大写
*/
func MarshalFirstLetterToUpper(name string) string {
	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	return s
}

// 创建model文件
func CreateModelFile(tablePreFix, tableLastFix, createTableName string) {

	//模板数据加载
	data := make([]byte, 100000)
	count := 0
	AppPath, _ := os.Getwd()
	configPath := AppPath + "/model.tpl"
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Open error :", err.Error())
		return
	}
	count, err = file.Read(data)
	if err != nil {
		log.Fatal("read tplFile error :", err.Error())
		return
	}

	//渲染
	render := template.Must(template.New("model").
		Funcs(template.FuncMap{
			"unescaped":            unescaped,
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"Tags":                 Tags,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
		}).Parse(string(data[:count])))
	if createTableName == "" {
		//创建该实例的所有表文件
		getTablesNameSql := "show tables from " + Instance
		tablaNames, err := NewEngineInstance().QueryString(getTablesNameSql)
		if err != nil {
			log.Fatal("sql error :", err.Error())
		}
		for _, table := range tablaNames {
			tableCol := "Tables_in_" + Instance
			tablePrefixName := table[tableCol]
			createModelFile(render, Instance, tablePrefixName, tablePreFix, tableLastFix)
		}
		return
	}
	tableNameSlice := strings.Split(createTableName, ",")
	for _, v := range tableNameSlice {
		if tableLastFix != "" {
			v = v + tableLastFix
		}
		if tablePreFix != "" {
			v = tablePreFix + v
		}
		createModelFile(render, Instance, v, tablePreFix, tableLastFix)
	}
}
func unescaped(x string) interface{} { return template.HTML(x) }
