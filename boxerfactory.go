package main

import "errors"

func NewBoxer(weightClass, name string) (IBoxer, error) {
	switch weightClass {
	case "Welterweight":
		return NewWelterweight(name), nil
	case "Junior Middleweight":
		return NewJuniorMiddleweight(name), nil
	}
	return nil, errors.New("boxing weight class not found")
}
