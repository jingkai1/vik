package vik

import (
	"fmt"
	"log"
	"reflect"
)

func Println(in ...interface{}) {
	for i := 0; i < len(in); i++ {
		switch reflect.TypeOf(in[i]).Kind() { //指针
		case reflect.Array:
			fmt.Printf("%s\t", in[i])
		case reflect.Chan:
			fmt.Printf("%v\t", in[i])
		case reflect.Map:
			fmt.Printf("%c\t", in[i])
		case reflect.Ptr: // 结构指针
			fmt.Printf("%+v\t", in[i])
		case reflect.Slice:
			fmt.Printf("%s\t", in[i])
		case reflect.Struct: //结构
			fmt.Printf("%+v\t", in[i])
		default:
			fmt.Printf("%v\t", in[i])
		}
	}
	fmt.Printf("\n")
}

func Logln(in ...interface{}) {
	var typ = make(map[reflect.Kind]string)
	typ[reflect.Array] = "%s"
	typ[reflect.Chan] = "%v"
	typ[reflect.Map] = "%c"
	typ[reflect.Ptr] = "%+v"
	typ[reflect.Slice] = "%s"
	typ[reflect.Struct] = "%+v"
	var stringSlice = make([]string, 0)
	for i := 0; i < len(in); i++ {
		if v, ok := typ[reflect.TypeOf(in[i]).Kind()]; ok {
			stringSlice = append(stringSlice, v)
			stringSlice = append(stringSlice, "\t")
		} else {
			stringSlice = append(stringSlice, "%v")
			stringSlice = append(stringSlice, "\t")
		}
	}
	stringPormat := ""
	for i := 0; i < len(stringSlice); i++ {
		stringPormat = stringPormat + stringSlice[i]
	}
	log.Printf(stringPormat, in...)
}
