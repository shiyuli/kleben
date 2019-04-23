package core

import (
	. "../templates"
	"fmt"
	"html/template"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
)

/*
Get the name of a function.
https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
 */
func GetFunctionName(foo interface{}) string {
	rawName := runtime.FuncForPC(reflect.ValueOf(foo).Pointer()).Name()
	temp := strings.Split(rawName, ".")
	originName := temp[len(temp) - 1]

	return strings.ToLower(originName)
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func print(message string) {
	fmt.Println(message)
}

func throw(message string) {
	panic(message)
}
