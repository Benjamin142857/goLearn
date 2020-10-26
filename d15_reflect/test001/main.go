package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

type mapList []map[string]int

func test1() {
	a := 142857
	b := "Benjamin"
	c := mapList{
		{"c1": 111},
		{"c2": 222},
	}
	d := Test{"Benjamin", 22}

	va := reflect.TypeOf(a)
	vb := reflect.TypeOf(b)
	vc := reflect.TypeOf(c)
	vd := reflect.TypeOf(d)

	fmt.Printf("%T, %v\n", va, va)
	fmt.Printf("%T, %v\n", vb, vb)
	fmt.Printf("%T, %v\n", vc, vc)
	fmt.Printf("%T, %v\n", vd, vd)
}

func test2() {
	a := 142857
	b := "Benjamin"
	c := mapList{
		{"c1": 111},
		{"c2": 222},
	}

	va := reflect.TypeOf(a)
	vb := reflect.TypeOf(b)
	vc := reflect.TypeOf(c)

	fmt.Println(va.Size(), va.String())
	fmt.Println(vb.Size(), vb.String())
	fmt.Println(vc.Size(), vc.String())
}

func test3() {
	var (
		myInt    int
		myString string
		myMap    mapList
		myStruct Test
	)

	tf_myInt := reflect.TypeOf(myInt)
	tf_myString := reflect.TypeOf(myString)
	tf_myMap := reflect.TypeOf(myMap)
	tf_myStruct := reflect.TypeOf(myStruct)

	fmt.Printf("%T, %v\n", tf_myInt, tf_myInt)
	fmt.Printf("%T, %v\n", tf_myString, tf_myString)
	fmt.Printf("%T, %v\n", tf_myMap, tf_myMap)
	fmt.Printf("%T, %v\n", tf_myStruct, tf_myStruct)

	fmt.Println(tf_myInt.Name(), tf_myInt.Kind())
	fmt.Println(tf_myString.Name(), tf_myString.Kind())
	fmt.Println(tf_myMap.Name(), tf_myMap.Kind())
	fmt.Println(tf_myStruct.Name(), tf_myStruct.Kind())
}

func test4() {
	a := 142857
	b := mapList{
		{"c1": 111},
		{"c2": 222},
	}
	c := Test{Name: "a", Age: 20}

	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	cValue := reflect.ValueOf(c)

	fmt.Printf("%T, %v\n", aValue, aValue)
	fmt.Printf("%T, %v\n", bValue, bValue)
	fmt.Printf("%T, %v\n", cValue, cValue)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("%T, %v\n", aValue.Interface(), aValue.Interface())
	//cc := cValue.Interface()
	//fmt.Printf("%p, %p\n", &c, &cc)
	//fmt.Println(c.Age)
	//fmt.Println(c == cc)
}

func getValueKindAndValue(i interface{}) (t reflect.Kind, v interface{}) {
	valueObj := reflect.ValueOf(i)
	t = valueObj.Kind()
	switch t {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = valueObj.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = valueObj.Uint()
	case reflect.Float32, reflect.Float64:
		v = valueObj.Float()
	case reflect.Bool:
		v = valueObj.Bool()
	case reflect.String:
		v = valueObj.String()
	default:
		v = valueObj.Interface()
	}
	return
}

type myField struct {
	name    string
	theType string
	tag     string
}

func getStructFieldInfo(i interface{}) (fieldSlice []myField) {
	typeObj := reflect.TypeOf(i)
	if typeObj.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < typeObj.NumField(); i++ {
		fieldObj := typeObj.Field(i)
		fieldSlice = append(fieldSlice, myField{
			name:    fieldObj.Name,
			theType: fieldObj.Type.Name(),
			tag:     string(fieldObj.Tag),
		})
	}
	return
}

func main() {
	//t1 := Test{Name: "Benjamin", Age: 22}
	for _, field := range getStructFieldInfo(20) {
		fmt.Println(field)
	}
}
