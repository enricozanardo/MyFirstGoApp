package model

import "errors"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type Person struct {
	Name string
	Surname string
	Email string
}

func (p *Person) FullName(name string, surname string) (string, error){

	if (name == "" || surname == "") {
		return "", errors.New("Name or Surname not provided")
	}

	fullname := name + " " + surname

	return fullname, nil
}


