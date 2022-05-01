package main

import (
	"errors"
	"fmt"

	"example.com/errorhandling/stdpw"
)

func main() {
	err := stdpw.PasswordValidation("#")
	if err == nil {
		fmt.Println("Done")
		return
	}

	fmt.Printf("%v\n", err)
	// fmt.Printf("%T\n", err)
	fmt.Println("------------")
	if errors.Is(err, stdpw.ErrInvalidLength) {
		fmt.Println("custom : Invalid Length")
		errInvalidLength := &stdpw.ErrInvalidLengthType{}
		if errors.As(err, errInvalidLength) {
			fmt.Println("\t--->", errInvalidLength.ActualLength)
		}

	}
	if errors.Is(err, stdpw.ErrMissingSmallLetter) {
		fmt.Println("custom : MissingSmallLetter")
		errMissingSmallLetter := &stdpw.ErrMissingSmallLetterType{}
		if errors.As(err, errMissingSmallLetter) {
			fmt.Println("\t--->", errMissingSmallLetter.Desc)
		}
	}
	if errors.Is(err, stdpw.ErrMissingCapitalLetter) {
		fmt.Println("custom : MissingCapitalLetter")
		errMissingCapitalLetter := &stdpw.ErrMissingCapitalLetterType{}
		if errors.As(err, errMissingCapitalLetter) {
			fmt.Println("\t--->", errMissingCapitalLetter.Desc)
		}
	}
	if errors.Is(err, stdpw.ErrMissingDigit) {
		fmt.Println("custom : MissingDigit")
		errMissingDigit := &stdpw.ErrMissingDigitType{}
		if errors.As(err, errMissingDigit) {
			fmt.Println("\t--->", errMissingDigit.Desc)
		}
	}
	fmt.Println("Bad Password")
}
