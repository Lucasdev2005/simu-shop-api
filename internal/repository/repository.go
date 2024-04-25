package repository

import (
	"fmt"
	"log"
	"reflect"
	"simushop/internal/core"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r repository) paginate(paginate core.Paginate) *gorm.DB {
	return r.db.Offset(paginate.GetOffset()).Limit(paginate.GetLimit())
}

func (r repository) queryWhere(where string, args ...interface{}) *gorm.DB {
	return r.db.Where(where, args...)
}

func (r repository) queryFirst(dest interface{}, where string, args ...interface{}) *gorm.DB {
	return r.queryWhere(where, args...).First(dest)
}

func (r repository) update(pointer any, where string, args ...interface{}) *gorm.DB {
	log.Println("[update] mapFields: ", r.mapFields(pointer))
	return r.db.Model(pointer).Select("*").Where(where, args...).Updates(r.mapFields(pointer))
}

func (r repository) mapFields(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("Esperava-se uma struct")
	}

	resultado := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		campo := typ.Field(i)
		gormTag := campo.Tag.Get("gorm")
		columnName := extractColumnName(gormTag)
		if columnName != "" {
			campoValor := val.Field(i)
			if campoValor.Interface() != reflect.Zero(campoValor.Type()).Interface() {
				resultado[columnName] = campoValor.Interface()
			}
		}
	}

	return resultado
}

func extractColumnName(tag string) string {
	parts := splitTag(tag, ";")
	for _, part := range parts {
		if len(part) > 7 && part[0:7] == "column:" {
			return part[7:]
		}
	}
	return ""
}

func splitTag(tag, delimiter string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(tag); i++ {
		if tag[i] == delimiter[0] {
			parts = append(parts, tag[start:i])
			start = i + len(delimiter)
		}
	}
	parts = append(parts, tag[start:])
	return parts
}

func newError(text string) error {
	return fmt.Errorf(text)
}
