package codex

import (
	"path/filepath"
	"strings"

	"github.com/quanxiaoxuan/go-utils/codex/template"
	"github.com/quanxiaoxuan/go-utils/filex"
	"github.com/quanxiaoxuan/go-utils/stringx"
)

const (
	DEMO       = "demo"
	TXT        = ".txt"
	Java       = ".java"
	Sql        = ".sql"
	GO         = ".go"
	Code       = "code"
	Controller = "controller"
	Service    = "service"
	Dao        = "dao"
	Params     = "params"
	LineSep    = "\n"
	Tab        = "\t"
)

// 字段配置
type FieldList []*Field
type Field struct {
	Name    string `json:"name"`    // 字段名
	Origin  string `json:"origin"`  // 原始字段名
	Type    string `json:"type"`    // 字段类型
	Comment string `json:"comment"` // 备注
	Default string `json:"default"` // 默认值
}

func checkDirAndName(saveDir, name string) (string, string) {
	saveDir = filepath.Join(saveDir, Code)
	if !filex.Exists(saveDir) {
		filex.Create(saveDir)
	}
	if name == "" {
		name = DEMO
	}
	name = strings.ToLower(name)
	return saveDir, name
}

// 生成通用代码
func GenCodeByFieldList(saveDir, name string, fieldList FieldList) (outPath string, err error) {
	saveDir, name = checkDirAndName(saveDir, name)
	// 合并所有
	all := strings.Builder{}
	all.WriteString(LineSep)
	all.WriteString(BuildJavaClass(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildGoStruct(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildGormStruct(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildSelectSql(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildSelectSqlAlias(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildInsertSql(name, fieldList))
	all.WriteString(LineSep)
	all.WriteString(BuildUpdateSql(name, fieldList))
	// 写入文件
	outPath = filepath.Join(saveDir, name+TXT)
	err = filex.WriteFile(outPath, all.String())
	if err != nil {
		return
	}
	return
}

// 生成Java-Class
func GenJavaClassByFieldList(saveDir, name string, fieldList FieldList) (outPath string, err error) {
	saveDir, name = checkDirAndName(saveDir, name)
	// 合并所有
	content := strings.Builder{}
	content.WriteString(LineSep)
	content.WriteString(BuildJavaClass(name, fieldList))
	content.WriteString(LineSep)
	// 写入文件
	outPath = filepath.Join(saveDir, stringx.UpperCamelCase(name)+Java)
	err = filex.WriteFile(outPath, content.String())
	if err != nil {
		return
	}
	return
}

// 生成SQL样例
func GenSqlByFieldList(saveDir, name string, fieldList FieldList) (outPath string, err error) {
	saveDir, name = checkDirAndName(saveDir, name)
	// 合并所有
	content := strings.Builder{}
	content.WriteString(LineSep)
	content.WriteString(BuildSelectSql(name, fieldList))
	content.WriteString(LineSep)
	content.WriteString(BuildSelectSqlAlias(name, fieldList))
	content.WriteString(LineSep)
	content.WriteString(BuildInsertSql(name, fieldList))
	content.WriteString(LineSep)
	content.WriteString(BuildUpdateSql(name, fieldList))
	// 写入文件
	outPath = filepath.Join(saveDir, name+Sql)
	err = filex.WriteFile(outPath, content.String())
	if err != nil {
		return
	}
	return
}

// 生成go结构体
func GenGoStructByFieldList(saveDir, name string, fieldList FieldList) (outPath string, err error) {
	saveDir, name = checkDirAndName(saveDir, name)
	// 合并所有
	content := strings.Builder{}
	content.WriteString(LineSep)
	content.WriteString(BuildGoStruct(name, fieldList))
	content.WriteString(LineSep)
	content.WriteString(LineSep)
	content.WriteString(BuildGormStruct(name, fieldList))
	content.WriteString(LineSep)
	// 写入文件
	outPath = filepath.Join(saveDir, name+GO)
	err = filex.WriteFile(outPath, content.String())
	if err != nil {
		return
	}
	return
}

// 生成go模板
func GenGoTemplateByName(saveDir, name string) (err error) {
	saveDir, name = checkDirAndName(saveDir, name)
	modelName := stringx.UpperCamelCase(name)
	paramsFile := filepath.Join(saveDir, Params, name+"_"+Params+GO)
	controllerFile := filepath.Join(saveDir, Controller, name+"_"+Controller+GO)
	serviceFile := filepath.Join(saveDir, Service, name+"_"+Service+GO)
	daoFile := filepath.Join(saveDir, Dao, name+"_"+Dao+GO)
	if err = filex.WriteFile(paramsFile, strings.ReplaceAll(template.GoParam, `{modelName}`, modelName)); err != nil {
		return err
	}
	if err = filex.WriteFile(controllerFile, strings.ReplaceAll(template.GoController, `{modelName}`, modelName)); err != nil {
		return err
	}
	if err = filex.WriteFile(serviceFile, strings.ReplaceAll(template.GoService, `{modelName}`, modelName)); err != nil {
		return err
	}
	if err = filex.WriteFile(daoFile, strings.ReplaceAll(template.GoDao, `{modelName}`, modelName)); err != nil {
		return err
	}
	return
}
