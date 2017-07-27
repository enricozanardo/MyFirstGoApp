package test

import (
	"testing"
	"github.com/onezerobinary/MyGoApp/model"
	"fmt"
)


func TestFullName(t *testing.T){
	p := model.Person {
		Name: "Enrico",
		Surname: "Zanardo",
		Email: "ezanardo@onezerobinary.com",
	}
	fmt.Println("name: " + p.Name + " Surname: " + p.Surname)

	expectedResult := p.Name + " " + p.Surname

	result, error := p.FullName("Enrico", "Zanardo")

	fmt.Println("result: " + result )

	if error != nil {
		t.Error("Expeceted Enrico Zanardo, got", result)
	}

	if expectedResult != result {
		t.Error("Expeceted Enrico Zanardo, got", result)
	}
}
