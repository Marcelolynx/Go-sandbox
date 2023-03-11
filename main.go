package main

import (
	"errors"
	"fmt"
)

//exercicio ninja nível 1 - curso aprenda GO Youtube

var a int = 2023
var b string = "James Bond"
var c bool = true

func main() {
	fmt.Println(Soma(9, 10))
}

func Soma(x int, y int) (int, error) {
	IsAdult := x + y
	if IsAdult < 18 {
		return IsAdult, errors.New("não pode ser menor que 18")
	}
	return x + y, nil
}
