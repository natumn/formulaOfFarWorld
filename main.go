package main

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
)

func main() {
}

func calc(formula string) interface{} {
	answer, err := eval(formula)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer.Type)
	fmt.Println(answer.Value)
	return answer.Value
}

func eval(expr string) (types.TypeAndValue, error) {
	return types.Eval(token.NewFileSet(), types.NewPackage("main", "calc"), token.NoPos, expr)
}
