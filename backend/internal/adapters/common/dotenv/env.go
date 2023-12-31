package dotenv

import (
	"fmt"
	"os"
	"reflect"
)

func Unmarshal[T interface{}](c *T) error {
	t := reflect.TypeOf(*c)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}

		st := reflect.ValueOf(c).Elem()
		ft := st.FieldByName(t.Field(i).Name)
		vl := os.Getenv(tag)
		if vl == "" {
			continue
		}

		switch tp := ft.Type().String(); tp {
		case "string":
			ft.SetString(vl)
		default:
			return fmt.Errorf("env parser does not support type %s", tp)
		}
	}
	return nil
}
