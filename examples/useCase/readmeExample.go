package main

import (
	"fmt"

	usecase "github.com/JacobTripp/diagrams-as-code/diagrams/useCase"
)

func main() {
	fooSystem := usecase.NewDiagram("Foo System")

	fooSystem.AddGeneralization(
		fooSystem.Actor("User"),
		fooSystem.Actor("Administrator"),
	)

	fooSystem.AddCommunication(
		fooSystem.Actor("Administrator"),
		fooSystem.UseCase("Create Account"),
	)

	fooSystem.AddCommunication(
		fooSystem.Actor("User"),
		fooSystem.UseCase("Create Account"),
	)
	fooSystem.AddInheritance(
		fooSystem.UseCase("Create Account"),
		fooSystem.UseCase("Create User Account"),
	)
	fooSystem.AddInheritance(
		fooSystem.UseCase("Create Account"),
		fooSystem.UseCase("Create Admin Account"),
	)

	fmt.Println(fooSystem.String())
}
