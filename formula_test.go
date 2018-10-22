package formula_test

import (
	"testing"

	"github.com/BuiltCloud/formula"
)

func TestStack(t *testing.T) {

	type User struct {
		Id        int    `json:"id"`
		Username    float64    `json:"username"`
		Password    float64    `json:"password"`
	}
	user := User{5, 6.1, 3.2}
	if	r,error := formula.Formula(&user,"Username + 1"); error == nil{
		 f,_ :=r.Float64()
		assertTrue(t, f == 7.1, "2.计算错误")
	}else {
		t.Errorf("1.计算错误")
	}

	formula.Equation(&user,"Username=Password + 1")
	assertTrue(t, user.Username == 4.2, "3.计算错误")
}

func assertTrue(t *testing.T, condition bool, message string) {
	if !condition {
		t.Errorf("%s",  message)
	}
}

