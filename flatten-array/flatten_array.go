package flatten

import "reflect"

func Flatten(input interface{}) []interface{} {
	var res = []interface{}{}
	RecFlatten(input, &res)
	return res
}

func RecFlatten(input interface{}, output *[]interface{}) interface{} {
	var parse = input.([]interface{})
	for _, v := range parse {
		if v != nil {
			if reflect.TypeOf(v).String() == "int" {
				*output = append(*output, v)
			} else {
				RecFlatten(v, output)
			}
		}
	}
	return output
}