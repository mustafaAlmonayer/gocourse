package intermediate

import "fmt"

type customError struct {
	code    int
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func doSomething() error {
	return &customError{
		code:    500,
		message: "Not Found",
	}
}

func customErrors() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
	}
}
