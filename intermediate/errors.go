package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: square root of negative number.")
	}
	return 0, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("processError: %w", &myError{"error while processing"})
	}
	return nil
}

// Create Error
type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

func errorEx() {
	fmt.Println(sqrt(16))
	result, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error:", err)
	}

	if err := eProcess(); err != nil {
		fmt.Println("Error: %w", err)
	}
}
