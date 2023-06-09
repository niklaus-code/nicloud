package utils

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

var paths = make(map[int]string)

func Bind(e *gin.Engine, controller interface{}) {
	Register(controller)
	pathInit()
	for _, path := range paths {
		e.GET(path, match(path))
	}
}

func pathInit() {
	i := 0
	for class, value := range Routes {
		for method, _ := range value {
			path := "/" + class + "/" + method
			paths[i] = path
			i += 1
		}
	}
}

type Route struct {
	Method      reflect.Value
	Args        []reflect.Type
	Method_type string
}

var aa = make(map[string]Route)

var Routes = make(map[string]map[string]Route)

func Register(controller interface{}) bool {
	v := reflect.ValueOf(controller)
	if v.NumMethod() == 0 {
		return false
	}

	tmp := reflect.TypeOf(controller).String()
	module := tmp
	if strings.Contains(tmp, ".") {
		module = tmp[strings.Index(tmp, ".")+1:]
	}

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		aa[action] = Route{method, params, "post"}

		Routes[module] = aa
	}
	return true
}

func match(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		if len(fields) < 3 {
			return
		}
		v, ok := Routes[fields[1]][fields[2]]

		if ok {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			v.Method.Call(arguments)
		}
	}
}
