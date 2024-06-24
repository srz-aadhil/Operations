package division

import "errors"

func Divide(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Division by zero is impossible")
	}
	{
		return float64(num1 / num2), nil
	}
}
