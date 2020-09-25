package scripts

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unsafe"
)

func TestResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// From https://medium.com/@ferencfbin/golang-own-structscan-method-for-sql-rows-978c5c80f9b5
// Can't get to work properly.
func StructScan(rows *sql.Rows, model interface{}) error {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination") // @todo add new error message
	}

	v = reflect.Indirect(v)
	t := v.Type()

	cols, _ := rows.Columns()

	var m map[string]interface{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		m = make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

	}

	for i := 0; i < v.NumField(); i++ {
		field := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]

		if item, ok := m[field]; ok {
			if v.Field(i).CanSet() {
				if item != nil {
					switch v.Field(i).Kind() {
					case reflect.String:
						v.Field(i).SetString(b2s(item.([]uint8)))
					case reflect.Float32, reflect.Float64:
						v.Field(i).SetFloat(item.(float64))
					case reflect.Ptr:
						if reflect.ValueOf(item).Kind() == reflect.Bool {
							itemBool := item.(bool)
							v.Field(i).Set(reflect.ValueOf(&itemBool))
						}
					case reflect.Struct:
						v.Field(i).Set(reflect.ValueOf(item))
					default:
						fmt.Println(t.Field(i).Name, ": ", v.Field(i).Kind(), " - > - ", reflect.ValueOf(item).Kind()) // @todo remove after test out the Get methods
					}
				}
			}
		}
	}

	return nil
}

func b2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}
