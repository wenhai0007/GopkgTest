package pkg

import "fmt"

func Prt(str string) {
	fmt.Print(str)
}

type PhoneInterface interface {
	Mprint()
	Mycall()
}
