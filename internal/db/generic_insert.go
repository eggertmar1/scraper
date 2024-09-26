package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func Insert(db *sql.DB, tableName string, model interface{}) error {
	// Reflect on the struct to get field names and values
	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %T", model)
	}

	// Prepare the query parts
	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		column := field.Tag.Get("db") // Get column name from struct tag
		if column == "" {
			continue // Skip fields without the db tag
		}

		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)+1))
		values = append(values, v.Field(i).Interface())
	}

	// Build and execute the query
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ","), strings.Join(placeholders, ","))
	_, err := db.Exec(query, values...)
	return err
}
