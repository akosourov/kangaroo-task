package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x1, v1, x2, v2, err := getArgs(os.Args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if getResult(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func getArgs(args ...string) (x1, v1, x2, v2 int, err error) {
	newErr := func() error { return errors.New("Bad input args") }

	if len(args) < 5 {
		err = newErr()
		return
	}
	x1, err = strconv.Atoi(args[1])
	if err != nil {
		err = newErr()
		return
	}
	v1, err = strconv.Atoi(args[2])
	if err != nil {
		err = newErr()
		return
	}
	x2, err = strconv.Atoi(args[3])
	if err != nil {
		err = newErr()
		return
	}
	v2, err = strconv.Atoi(args[4])
	if err != nil {
		err = newErr()
		return
	}
	return
}

func getResult(x1, v1, x2, v2 int) bool {
	if x1 == x2 {
		return true
	}
	if x1 < x2 && v1 <= v2 {
		return false
	}
	if x2 < x1 && v2 <= v1 {
		return false
	}

	i, j := x1, x2
	for i < j {
		i += v1
		j += v2
	}
	if i == j {
		return true
	}
	return false
}
