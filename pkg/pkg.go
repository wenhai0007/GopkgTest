package pkg

import "fmt"

func Prt(str string) {
	fmt.Print(str)
}

type PhoneInterface interface {
	Mprint()
	Mycall()
}

type Mystruct struct {
}

func (ms Mystruct) Mprint() {

}

func (ms Mystruct) Mycall() {

}

func InitInterface() Mystruct {
	aStruct := Mystruct{}
	return aStruct
}
