package helper

import (
	"errors"
	"regexp"
)

func CleanPhone(phone string) (string, error) {
	p := []rune(phone)
	p1 := regexp.MustCompile(`^(\+62)8[1-9][0-9]{6,9}$`)
	p2 := regexp.MustCompile(`^(62)8[1-9][0-9]{6,9}$`)
	p3 := regexp.MustCompile(`^(0)8[1-9][0-9]{6,9}$`)

	if p1.MatchString(phone) {
		phone = "0" + string(p[3:len(p)])
		return phone, nil
	}
	if p2.MatchString(phone) {
		phone = "0" + string(p[2:len(p)])
		return phone, nil
	}
	if p3.MatchString(phone) {
		return phone, nil
	}
	return "", errors.New("Wrong format number")
}
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
