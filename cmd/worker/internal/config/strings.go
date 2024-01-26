package config

import (
	"fmt"
	"reflect"

	"github.com/alexflint/go-arg"
)

func validateStrings(cfg config, p *arg.Parser) {
	valuesSlice := reflect.ValueOf(cfg)
	for i := 0; i < valuesSlice.NumField(); i++ {
		if valuesSlice.Field(i).Type().Kind() == reflect.String && valuesSlice.Field(i).Interface() == "" {
			p.Fail(fmt.Sprintf("error processing environment variable %q: string: missing value", valuesSlice.Type().Field(i).Name))
		}
	}
}
