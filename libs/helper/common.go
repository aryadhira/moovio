package helper

import (
	"moovio/libs/constant"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func MongodbURIGenerator() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if username != "" && password != "" {
		return "mongodb://" + username + ":" + password + "@" + host + ":" + port
	} else {
		return "mongodb://" + host + ":" + port
	}
}

func ConvertDateTimetoDateInt(dt time.Time) int {
	dateintstr := dt.Format(constant.DateIntFormat)

	dateint, err := strconv.Atoi(dateintstr)
	if err != nil {
		return 0
	}

	return dateint
}

func InterfaceToString(input interface{}) string {
	out := ""

	if input == nil {
		return out
	}

	out = input.(string)

	return out
}

func InterfaceToInt(input interface{}) int {
	out := 0

	if input == nil {
		return out
	}

	switch input.(type){
	case string:
		data := input.(string)
		out, _ = strconv.Atoi(data)
	}

	return out
}

func InterfaceToFloat64(input interface{}) float64 {
	out := 0.0

	if input == nil {
		return out
	}

	out = input.(float64)

	return out
}

func ArrayinterfaceToArrayString(input []interface{}) []string {
	out := []string{}

	if input == nil {
		return out
	}

	for _, each := range input {
		val := InterfaceToString(each)
		out = append(out, val)
	}

	return out
}

func StructToMap(obj interface{}) map[string]interface{}{
	result := make(map[string]interface{})
    value := reflect.ValueOf(obj)
    t := value.Type()

    if t.Kind() != reflect.Struct {
        return nil
    }

    for i := 0; i < value.NumField(); i++ {
        field := t.Field(i)
        fieldValue := value.Field(i).Interface()
		key := strings.ToLower(field.Name)
        result[key] = fieldValue
    }

    return result
}