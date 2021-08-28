package mysql

import (
	"api-gmr/store/repository"
	"fmt"
	"reflect"
)

type billingFilter struct {
	Year   string `field:"year"`
	Month  string `field:"month"`
	UserID string `field:"user_id"`
	Status string `field:"status"`
}

func buildBillingFilter(rfilter repository.BillingFilter) ([]string, []interface{}) {
	filter := billingFilter{
		Year:   zeroToemptyStr(rfilter.GetYear()),
		Month:  zeroToemptyStr(rfilter.GetMonth()),
		UserID: zeroToemptyStr(rfilter.GetUserID()),
		Status: fmt.Sprint(rfilter.GetStatus()),
	}

	s := reflect.ValueOf(&filter).Elem()
	typeOfs := s.Type()

	var keys []string
	var args []interface{}

	for i := 0; i < s.NumField(); i++ {
		field := typeOfs.Field(i).Tag.Get("field")
		f := s.Field(i)
		value := fmt.Sprint(f.Interface())
		if len(value) > 0 {
			keys = append(keys, fmt.Sprintf("%s = ?", field))
			args = append(args, value)
		}
	}
	return keys, args
}

func zeroToemptyStr(v int) string {
	if v > 0 {
		return fmt.Sprint(v)
	}
	return ""
}
