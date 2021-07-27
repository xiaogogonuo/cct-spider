package insertdb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetQuotesAndValues(v interface{}) (insertValues []interface{}, strLen int) {
	insertValues = make([]interface{}, 0)

	elem := reflect.ValueOf(v)

	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	for i := 0; i < elem.NumField(); i++ {
		curField := elem.Field(i)
		val, valLen := _parseField(curField)
		strLen += valLen
		insertValues = append(insertValues, val)
	}
	return
}

func _parseField(v reflect.Value) (fieldValue interface{}, valLen int) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldValue = v.Int()
		valLen = len(strconv.FormatInt(v.Int(), 10))

	case reflect.String:
		fieldValue = v.String()
		valLen = v.Len()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldValue = v.Uint()
		valLen = len(strconv.FormatUint(v.Uint(), 10))

	case reflect.Float32:
		fieldValue = v.Float()
		valLen = len(strconv.FormatFloat(v.Float(), 'f', -1, 32))

	case reflect.Float64:
		fieldValue = v.Float()
		valLen = len(strconv.FormatFloat(v.Float(), 'f', -1, 64))

	}
	return

}

func GetInsertBaseSQLCode(v interface{}, dbName string) (preambleSql string, epilogueSql string, oneQuoteSql string) {

	elem := reflect.ValueOf(v)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	elemType := elem.Type()
	numFields := elem.NumField()

	quotes := strings.Repeat("?,", numFields)
	quotes = quotes[0 : len(quotes)-1]

	insertFields := make([]string, 0, numFields)
	epilogues := make([]string, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := elemType.Field(i).Name
		insertFields = append(insertFields, field)
		epilogues = append(epilogues, fmt.Sprintf("%s = VALUES(%s)", field, field))

	}
	oneQuoteSql = fmt.Sprintf("(%s)", quotes)
	preambleSql = fmt.Sprintf("INSERT INTO %s (%s) VALUES ", dbName, strings.Join(insertFields, ", "))
	epilogueSql = fmt.Sprintf("ON DUPLICATE KEY UPDATE %s", strings.Join(epilogues, ", "))
	return
}

func GetUpdateBaseSQLCode(v interface{}) (updateFields [][]string, epilogue string, fieldLen int) {

	elem := reflect.ValueOf(v)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	elemType := elem.Type()
	fieldLen = elem.NumField()
	updateFields = make([][]string, 0, fieldLen)

	idField := elemType.Field(0).Name
	for i := 1; i < fieldLen; i++ {
		field := elemType.Field(i).Name
		updateFields = append(updateFields, []string{fmt.Sprintf(`%s = CASE %s`, field, idField)})
	}
	epilogue = fmt.Sprintf(`WHERE %s IN `, idField)
	fieldLen--
	return

}

func GetWhenAndThen(v interface{}) (updateValues []string) {
	updateValues = make([]string, 0)

	elem := reflect.ValueOf(v)

	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	id := elem.Field(0)
	for i := 1; i < elem.NumField(); i++ {
		curField := elem.Field(i)
		val, _ := _parseField(curField)
		updateValues = append(updateValues, fmt.Sprintf(`WHEN '%s' THEN '%s'`, id, val))
	}

	return
}
