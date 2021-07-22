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
	// given a struct pointer this function populates its elements with enumerated flags
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
		// fieldName := structTypeObj.Field(i).Name
		// FieldValue := structValueObj.FieldByName(fieldName)
		fieldValue := structValueObj.Field(i)
		if !fieldValue.CanSet() {
			return fmt.Errorf("Invalid Struct Object")
		}
		if fieldValue.Type() != ptype {
			return fmt.Errorf("Inconsistent types of str")
		}
		fieldValue.Set(reflect.ValueOf(constArray[i]).Convert(ptype))
	}
	return err
}

func GetFlagComponents(testSum interface{}) (I []interface{}) {
	// Given a sum value of any integer types this function returns an array
	// of its components
	origType := reflect.TypeOf(testSum)
	// kind := origType.Kind()
	int_type := reflect.TypeOf(constArray[0])
	fmt.Println()
	if !reflect.ValueOf(testSum).Type().ConvertibleTo(int_type) {
		return I
	}
	testSumRef := reflect.ValueOf(testSum).Convert(int_type)
	cumulative := 0
	for i := 0; i < 64; i++ {
		testSumVal := testSumRef.Interface().(int)
		current := constArray[i]
		hasFlag := testSumVal&current == current
		if hasFlag == true {
			I = append(I, reflect.ValueOf(current).Convert(origType).Interface())
			cumulative |= current
			if cumulative >= testSumVal {
				break
			}
		}
	}
	return I
}

func FlagInSum(flag interface{}, flagSum interface{}) (bool, interface{}) {
	// this function checks wheter a flag is present in a sum of flags
	// return  value contains boolean result for presence of flag in flagSum and
	// Leftover of FlagSum if True else nil
	int_type := reflect.TypeOf(0)
	fv := reflect.ValueOf(flag)
	fsv := reflect.ValueOf(flagSum)
	if !fv.Type().ConvertibleTo(int_type) || !fsv.Type().ConvertibleTo(int_type) {
		panic("the type of inputs must be any integer types")
	}
	fvI := fv.Convert(int_type).Interface().(int)
	fsvI := fsv.Convert(int_type).Interface().(int)
	truth := (fvI & fsvI) == fvI
	if truth == true {
		leftover := reflect.ValueOf(fvI ^ fsvI)
		return true, leftover.Convert(fsv.Type()).Interface()
	}
	return false, nil
}
