package repository

import (
	"reflect"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r repository) queryWhere(where string, args ...interface{}) *gorm.DB {
	return r.db.Where(where, args...)
}

func (r repository) queryFirst(dest interface{}, where string, args ...interface{}) *gorm.DB {
	return r.queryWhere(where, args...).First(dest)
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
		campo := typ.Field(i).Name
		campoValor := val.Field(i)
		if campoValor.Interface() != reflect.Zero(campoValor.Type()).Interface() {
			resultado[campo] = campoValor.Interface()
		}
	}

	return resultado
}
