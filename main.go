package main

import (
	"go/token"
	"go/types"
	"log"
	"strconv"
)

func main() {
}

func calc(formula string) int {
	answer, err := eval(formula)
	if err != nil {
		log.Fatal(err)
	}
	intAnswer, err := strconv.Atoi(answer.Value.ExactString())
	if err != nil {
		log.Fatal(err)
	}
	return intAnswer
}

func eval(expr string) (types.TypeAndValue, error) {
	return types.Eval(token.NewFileSet(), types.NewPackage("main", "calc"), token.NoPos, expr)
}
