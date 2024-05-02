package utils

import "strconv"

func ParseStrToUint(str string) (*uint, error) {
	res, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil, err
	}

	uintRes := uint(res)

	return &uintRes, nil
}

func ParseStrToInt(str string) (*int, error) {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, err
	}

	uintRes := int(res)

	return &uintRes, nil
}
