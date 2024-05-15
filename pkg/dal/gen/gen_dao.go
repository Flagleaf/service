//go:generate go run gen_dao.go -type=example --package=cmd --output=example_generated.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Column struct {
	Type  string
	Name  string
	Field string
	Param string
}
type Schema struct {
	TableName  string
	ColumnList []Column
}

type Option struct {
	Schema Schema
}

func main() {
	generateDao()
}

// 定义生成 example API 的代码
func generateDao() {
	// 读取模板文件的内容
	templateContent, err := ioutil.ReadFile("param.tmpl")
	if err != nil {
		log.Fatalf("failed to read template file: %v", err)
	}
	tmpl, err := template.New("example").Parse(string(templateContent))
	fmt.Println(tmpl)
	if err != nil {
		fmt.Println(err)
	}
	// 定义代码参数
	data := Option{Schema: Schema{TableName: "User", ColumnList: []Column{
		Column{Field: "user_id", Param: "userId", Name: "UserId", Type: "int"},
		Column{Field: "name", Param: "name", Name: "Name", Type: "string"},
		Column{Field: "email", Param: "email", Name: "Email", Type: "string"}}}}

	// 生成代码
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
