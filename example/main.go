package main

import (
	"fmt"

	"github.com/playscale/errors-with-stack"
)

func exampleErrorsWithStack() {
	errorWithStack := errors.New("This error should contain a stack Trace")
	fmt.Println(errorWithStack.Error())

	errors.DisableStackTraces()

	plainOldGoError := errors.New("This is just a plain old golang error")
	fmt.Println(plainOldGoError.Error())
}

func main() {
	exampleErrorsWithStack()
}
