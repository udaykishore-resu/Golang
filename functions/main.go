package main

import (
	functiontypes "functions/function_types"
)

func main() {
	functiontypes.Func_With_NoArgs()
	functiontypes.Func_With_Args(20, 40)
	a := 30
	b := 40
	functiontypes.Func_With_RefArgs(&a, &b)
	//functiontypes.Anonymous_fun()
	//functiontypes.Closure_fun()
	//functiontypes.Variadic_fun()
	//functiontypes.Callback_fun()
}
