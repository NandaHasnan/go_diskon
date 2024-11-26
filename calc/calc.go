package calc

import "errors"

func Sum(num1 int, num2 int) (int, error) {
	if num2 > 15 {
		return 0, errors.New("angka harus kurang dari 15")
	}
	return num1 * num2, nil
}
