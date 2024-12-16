package utils

import (
  "fmt"
  "log"
  "os"
  "strings"
  "text/template"

  "github.com/rewarding-harvest-backend/schemas"
)

const modelTemplate = `
package models

import (
  "time"
  "gorm.io/gorm"
)

type {{.TableName}} struct {
  {{range .Columns}}
  {{.GoFieldName}} {{.GoType}} ` + "`gorm:\"column:{{.DBFieldName}}\" json:\"{{.JsonFieldName}}\"`" + `
  {{end}}
}
`

type Column struct {
  DBFieldName   string
  GoFieldName   string
  JsonFieldName string
  GoType        string
}

type ModelData struct {
  TableName string
  Columns   []Column
}

func GenerateModels(schemas []schemas.TableSchema) {
  modelsDir := "./models"
  if err := os.MkdirAll(modelsDir, os.ModePerm); err != nil {
    log.Fatalf("Failed to create models directory: %v", err)
  }

  for tableName, columns := range GroupSchemaByTable(schemas) {
    data := ModelData{
      TableName: toGoStructName(tableName),
      Columns:   MapColumns(columns),
    }

    file, err := os.Create(fmt.Sprintf("%s/%s.go", modelsDir, tableName))
    if err != nil {
      log.Fatalf("Failed to create model file for %s: %v", tableName, err)
    }
    defer file.Close()

    t := template.Must(template.New("model").Parse(modelTemplate))
    if err := t.Execute(file, data); err != nil {
      log.Fatalf("Failed to execute template: %v", err)
    }
    fmt.Printf("✅ Model created for table: %s\n", tableName)
  }
}

func GroupSchemaByTable(schemas []schemas.TableSchema) map[string][]schemas.TableSchema {
  grouped := make(map[string][]schemas.TableSchema)
  for _, schema := range schemas {
    grouped[schema.TableName] = append(grouped[schema.TableName], schema)
  }
  return grouped
}

func MapColumns(columns []schemas.TableSchema) []Column {
  var mapped []Column
  for _, col := range columns {
    mapped = append(mapped, Column{
      DBFieldName:   col.ColumnName,
      GoFieldName:   toGoStructName(col.ColumnName),
      JsonFieldName: strings.ToLower(col.ColumnName),
      GoType:        toGoType(col.DataType, col.IsNullable),
    })
  }
  return mapped
}
