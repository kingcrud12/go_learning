package main

import (
	"go_learning/internal/concurrency"
	"go_learning/internal/constantes"
	"go_learning/internal/data_structures"
	"go_learning/internal/error_handling"
	"go_learning/internal/flow"
	"go_learning/internal/generics"
	"go_learning/internal/interfaces"
	"go_learning/internal/learn_functions"
	"go_learning/internal/shadowing"
	"go_learning/internal/structs_and_pointers"
	"go_learning/internal/tools"
	"go_learning/internal/types"
	"go_learning/internal/variables"
)

func main() {
	variables.Run()
	constantes.Run()
	shadowing.Run()
	flow.Run()
	learn_functions.Run()
	types.Run()
	structs_and_pointers.Run()
	error_handling.Run()
	interfaces.Run()
	generics.Run()
	data_structures.Run()
	tools.Run()
	concurrency.Run()
}
