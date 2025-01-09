package model

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func modelTestFail(msg string, name string) {
	// logger 还未初始化，只能直接先打印
	fmt.Fprintf(os.Stderr, "modelTestFail [%s]: %s\n", name, msg)
	panic("bad database model")
}

func modelTest[S any, D any]() bool {
	var s S
	var d D

	src := reflect.TypeOf(s)
	dest := reflect.TypeOf(d)
	name := src.Name()

	if fmt.Sprintf("%sM", src.Name()) != dest.Name() {
		modelTestFail(name, fmt.Sprintf("name error src: %s dest: %s", src.Name(), dest.Name()))
		return false
	}

	if src.Kind() != reflect.Struct || dest.Kind() != reflect.Struct {
		modelTestFail(name, "not struct")
		return false
	}

	for i := 0; i < src.NumField(); i++ {
		fieldSrc := src.Field(i)
		if fieldSrc.Type.Kind() == reflect.Ptr && fieldSrc.Type.Elem().Kind() == reflect.Struct {
			continue
		}

		fieldDest, ok := dest.FieldByName(fieldSrc.Name)
		if !ok {
			modelTestFail(name, fmt.Sprintf("dest filed[%s] not found", fieldSrc.Name))
			return false
		}

		if fieldDest.Type.Kind() != fieldSrc.Type.Kind() {
			modelTestFail(name, fmt.Sprintf("dest filed[%s] kind error", fieldDest.Name))
			return false
		} else if fieldDest.Tag.Get("gorm") != fieldSrc.Tag.Get("gorm") {
			modelTestFail(name, fmt.Sprintf("dest filed[%s] gorm error", fieldDest.Name))
			return false
		}
	}

	for i := 0; i < dest.NumField(); i++ {
		fieldDest := dest.Field(i)
		if fieldDest.Type.Kind() == reflect.Ptr && fieldDest.Type.Elem().Kind() == reflect.Struct {
			modelTestFail(name, fmt.Sprintf("dest filed[%s] kind error", fieldDest.Name))
			return false
		}

		if strings.Contains(fieldDest.Tag.Get("gorm"), "foreignKey:") {
			modelTestFail(name, fmt.Sprintf("dest filed[%s] gorm has foreignKey", fieldDest.Name))
			return false
		}
	}

	tablename := reflect.ValueOf(&d).MethodByName("TableName")
	if tablename.Kind() != reflect.Func {
		modelTestFail(name, fmt.Sprintf("table name error"))
		return false
	}

	fun, ok := tablename.Interface().(func() string)
	if !ok {
		modelTestFail(name, fmt.Sprintf("table name error"))
		return false
	}

	if fun() != strings.ToLower(src.Name()) {
		modelTestFail(name, fmt.Sprintf("table name error"))
		return false
	}

	return true
}

var AutoCreateModelList = []interface{}{
	&BagM{},
	&BuyRecordM{},
	&ClassM{},
	&ConfigM{},
	&ImageM{},
	&MsgM{},
	&UserM{},
	&WupinM{},
	XieyiM{},
}
