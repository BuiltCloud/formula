package formula

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/BuiltCloud/formula/evaler"
)

func struct2Map(obj interface{}) map[string]string {
	var data = make(map[string]string)
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		data[typeOfT.Field(i).Name] = fmt.Sprint(f.Interface())
	}
	return data
}
// Formula 计算公式 "x+1"
func Formula(obj interface{},formula string) (result *big.Rat, err error){
	smap := struct2Map(obj)
	return evaler.EvalWithVariables(formula,  smap)
}

// Equation 计算等式 "y=x+1"
func Equation(obj interface{}, formula string) {
	fs := strings.Split(formula, "=")
	if len(fs)==2 {
		if r,error := Formula(obj,fs[1]); error==nil{
			s := reflect.ValueOf(obj).Elem()  // 反射获取测试对象对应的struct枚举类型
			//strings.Replace(str, " ", "", -1)
			field := s.FieldByName(fs[0])
			switch t := field.Interface().(type) {
			case int:
				field.SetInt(r.Num().Int64())
			case float64:
				r,_:=r.Float64()
				field.SetFloat(r)
			default:
				_ = t
				panic("非数字")
			}
		}else {
			panic("计算公式出错")
		}
	}else {
		panic("非等式")
	}

}
