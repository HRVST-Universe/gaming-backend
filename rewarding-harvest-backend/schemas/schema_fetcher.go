package schemas

import (
  "log"

  "gorm.io/gorm"
)

type TableSchema struct {
  TableName  string `json:"tableName"`
  ColumnName string `json:"columnName"`
  DataType   string `json:"dataType"`
  IsNullable string `json:"isNullable"`
}

func FetchSchema(DB *gorm.DB) []TableSchema {
  var schemas []TableSchema

  query := `
  SELECT 
    table_name, 
    column_name, 
    data_type, 
    is_nullable
  FROM information_schema.columns
  WHERE table_schema = 'public'
  ORDER BY table_name, ordinal_position;
  `

  if err := DB.Raw(query).Scan(&schemas).Error; err != nil {
    log.Fatalf("Failed to fetch schema: %v", err)
  }

  return schemas
}
