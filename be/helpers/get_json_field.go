package helpers

import (
	"reflect"

	"ewallet-server-v1/dtos"
)

func jsonFieldName(fieldName string) string {
	t := reflect.TypeOf(dtos.RequestValidationMiddleware{})
	field, found := t.FieldByName(fieldName)
	if !found {
		return ""
	}
	jsonTag := field.Tag.Get("json")
	return jsonTag
}

func msgForTag(tag string, nameSpace string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid format email"
	case "min":
		return "The minimum password character is 5"
	case "max":
		return "The maximum description character is 35"
	case "gte":
		if nameSpace == "RequestTransferFund.Amount" {
			return "The minimum transfer is Rp.1000"
		}
		return "The minimum top up is is Rp.50000"
	case "lte":
		return "The maximum top up is is Rp.10000000"
	}
	return ""
}
