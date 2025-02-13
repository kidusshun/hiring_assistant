package jobposting

import "strconv"

func ConvertStrToInt(str string) (int, error) {
	intVal, err := strconv.Atoi(str)
	return intVal, err
}