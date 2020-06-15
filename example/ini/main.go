package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini 配置文件解析器

// MYSQL 配置
type MysqlConfig struct {
	Address string `ini:"address"`
	Port int `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// Redis 配置
type RedisConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database"`
	Test bool `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数校验
	// 传进来的data必须是指针类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr{
		err = errors.New("data param should a pointer")
		return
	}
	// 传入的data参数必须为结构体指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should a struct pointer")
		return
	}
	// 读取文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line,"#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			for i:= 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		}else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line,"=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			for i:=0; i< sValue.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == key{
					fieldName = field.Name
					fieldType = field
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int64,reflect.Int8,reflect.Int16,reflect.Int32:
				var intval int64
				intval, err = strconv.ParseInt(value, 10,64)
				if err != nil {
					return
				}
				fileObj.SetInt(intval)
			case reflect.Bool:
				var parseBool bool
				parseBool, err = strconv.ParseBool(value)
				if err != nil {
					return
				}
				fileObj.SetBool(parseBool)
			case reflect.Float32,reflect.Float64:
				var f float64
				f, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return
				}
				fileObj.SetFloat(f)

			}
		}
	}
	return
}


func main() {
	var cfg Config
	err := loadIni("./my.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
