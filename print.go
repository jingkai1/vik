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
		default:
			fmt.Printf("%v\t", in[i])
		}
	}
	fmt.Printf("\n")
}

func Logln(in ...interface{}) {
	for i := 0; i < len(in); i++ {
		switch reflect.TypeOf(in[i]).Kind() { //指针
		case reflect.Array:
			log.Printf("%s", in[i])
		case reflect.Chan:
			log.Printf("%v", in[i])
		case reflect.Map:
			log.Printf("%c", in[i])
		case reflect.Ptr: // 结构指针
			log.Printf("%+v", in[i])
		case reflect.Slice:
			log.Printf("%s", in[i])
		default:
			log.Printf("%v", in[i])
		}
	}
}
