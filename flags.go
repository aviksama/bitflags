package bitflags

import (
	"fmt"
	"reflect"
)

const (
	a int = 1 << iota
	b int = 1 << iota
	c int = 1 << iota
	d int = 1 << iota
	e int = 1 << iota
	f int = 1 << iota
	g int = 1 << iota
	h int = 1 << iota
	i int = 1 << iota
	j int = 1 << iota
	k int = 1 << iota
	l int = 1 << iota
	m int = 1 << iota
	n int = 1 << iota
	o int = 1 << iota
	p int = 1 << iota
)

var constArray = [16]int{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func kindCheck(typeOf reflect.Type, kind reflect.Kind) (err error) {
	if typeOf.Kind() != kind {
		err = fmt.Errorf("error: Type %s is not a %s", typeOf.String(), kind.String())
	}
	return err
}

func validPointerToStruct(flags interface{}) (r reflect.Type, err error) {
	topType := reflect.TypeOf(flags)
	err = kindCheck(topType, reflect.Ptr)
	if err != nil {
		return r, err
	}
	structTypeObj := topType.Elem() // use Elem to get the underlying type of the pointer
	// bootomType := structTypeObj.Kind()
	err = kindCheck(structTypeObj, reflect.Struct)
	if err != nil {
		return r, err
	}
	return structTypeObj, err
}

func BuildFlagsStruct(flags interface{}) (err error) {
	structTypeObj, err := validPointerToStruct(flags)
	if err != nil {
		return err
	}
	length := structTypeObj.NumField()
	minimum := getMin(len(constArray), length)
	ptype := structTypeObj.Field(0).Type
	switch ptype.Kind() {
	case reflect.Int8:
		if minimum > 7 {
			return fmt.Errorf("the datatype %s doesn't support more than 7 elements", ptype)
		}
	case reflect.Uint8:
		if minimum > 8 {
			return fmt.Errorf("the datatype %s doesn't support more than 8 elements", ptype)
		}
	case reflect.Int16:
		if minimum > 15 {
			return fmt.Errorf("the datatype %s doesn't support more than 15 elements", ptype)
		}
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		return fmt.Errorf("Invalid type for flags")
	}
	structValueObj := reflect.ValueOf(flags).Elem()
	for i := 0; i < minimum; i++ {
		fieldName := structTypeObj.Field(i).Name
		FieldValue := structValueObj.FieldByName(fieldName)
		if !FieldValue.CanSet() {
			return fmt.Errorf("Invalid Struct Object")
		}
		if FieldValue.Type() != ptype {
			return fmt.Errorf("Inconsistent types of str")
		}
		FieldValue.Set(reflect.ValueOf(constArray[i]).Convert(ptype))
	}
	return err
}

func GetFlagComponents(flagStruct interface{}, flagSum interface{}) []interface{} {
	// tbd
	return []interface{}{}
}

func ValidateComponentsInSum(flagStruct interface{}, flagSum interface{}) (err error) {
	//tbd
	return err
}
