Use Case Diagram
================

A use case is a situation where your system is used to fulfill one or more of
your customer's requirements; a use case captures a piece of functionality that 
the system provides.

Use cases are the starting point of your model since they affect and guide all
the other elements within your system's design. They describe a system's
requirements strictly from the outside looking in; They specify the value that
the system delivers to customers.

# Components of a Use Case Diagram

## Actors - Outside Your System
An actor interacts with the system and is not part of the system.

Actors don't have to be actual people. While an actor might be a person, it
could also be a third party's system, such as in a business-to-business
application. Think of an actor as a black box: you cannot change an actor and
you are not interested in how it works, but it must interact with your system.

Sometimes identifying an actor can be tricky, here are a couple questions that
might help identify if you have an actor or not:
  1. Is the "thing" an actual person interacting with the system?
  2. Is the "thing" something that I can change within the system's design?

If you answered "Yes" to the first, you probably have an actor. If you answered
"No" to the second question you probably have an Actor.

Don't for get to include auditors, installers, maintainers, and so on. If you
focus on only the obvious users of your system, then you might forget about some
of these other stakeholders.

Some actors are just a special kind of another actor. For example, FooActor can
do everything that BarActor can do but with a couple additional interactions
with the system. Use a generalization arrow to show this.

#### How to create an Actor
```Go
// creating two actors
user := usecase.NewActor("User")
admin := usecase.NewActor("Administrator")

// adding the actors to your diagram
// creating a diagram
fooSystem := usecase.NewDiagram("Foo System")

// add a basic user
fooSystem.AddActor(user)

// add a user that can do everything that "user" can do plus more
fooSystem.AddActor(admin, WithGeneralizaion(user))
// OR you can add a generalization after adding both actors
fooSystem.AddActor(admin)
fooSystem.AddGeneralization(admin, user)
```

## Use Cases - System Requirements
Use cases must have clear pass/fail criteria. A use case, or job, might be as
simple as allowing the user the log in or as complex as executing a distributed
transaction across multiple global databases.

The important thing to remember is that a use case, from the user's perspective,
is a complete use of the system; there is some interaction with the system, as
well as some output from that interaction. An example use case could be "Create
a new customer account".

A rule of thumb to help determine if you have a use case is:
  - A use case is something that provide some measurable result to the user or
    an external system

#### How to create a UseCase
```Go
// create a use case
createUserAccount := usecase.NewUseCase("Create a user account")

// add the use case to the diagram
fooSystem.AddUseCase(createUserAccount)
```

### Use Case Descriptions
Every use case should be accompanied by a description to help system designers
understand how the system's concerns will be met, how to know the most important
actor is, and the steps that are involved.

There are no rules for what exactly goes into a use case description but some
examples are:
  - **Use case description detail**:
      What it means and why it's useful.
  - **Related requirements**:
      Some indication as to which requirements this use case partially or
      completely fufills.
  - **Goal in context**:
      The use case's place within the system and why this use case is important.
  - **Preconditions**:
      What needs to happen before the use case can be executed.
  - **Successful end condition**:
      What the system's condition should be if the sue case executes
      successfully.
  - **Failed end condition**:
      What the system's condition should be if the use case fails to execute
      successfully.
  - **Primary actors**:
      The main actors that participate in the use case. Often includes the
      actors that trigger or directly receive information from the use case's
      execution.
  - **Secondary actors**:
      Actors that participate but are not the main players in a use case's
      execution.
  - **Included use cases**:
      Use cases that it depends on
  - **Base use cases**:
      Use cases that are inherited
  - **Trigger**:
      The event triggered by an actor that causes the use case to execute.
  - **Main flow**:
      The place to describe each of the important steps in a use case's normal
      execution.
  - **Extensions**:
      A description of any alternative steps from the ones described in the main
      flow.

#### How to add UseCaseDescription
TODO
```Go
```

## Communication Lines
A communication line connects and actor and a use case to show the actor
participating in the use case. There is a potential to have any number of actors
involved in a use case, there is no real limit.

The purpose of a communication line is to show that an actor is simply involved
in a use case, not to imply an information exchange in any particular direction
or that the actor starts the use case. That type of information can be added to
the use case's description.

#### How to create a Communication
```Go
// create a communication line between admin and createUseAccount
fooSystem.AddCommunication(admin, createUserAccount)

// add another communication line to createUseAccount
fooSystem.AddCommunication(user, createUserAccount)
```

## Use Case Communications
### The Include Relationship
When there is repetitive behavior shared between two use cases then that
behavior is best separated and captured within a new use case. This new use case
can then be reused by other use cases using and include relationship. For
example if two use cases need to check a credential database (that's an actor)
then a new use case of "Check Identity" should be created and then any use case
that needs to check identity then includes that use case.

The include relationship declares that the use case including the other
completely reuses all of the steps from the use case being included.

#### How to add an Include relationship
```Go
fooSystem.AddInclude(
  fooSystem.UseCase("Create a user account"),
  fooSystem.UseCase("Another UseCase Child"),
)
```

### The Generalization Relationship (inheritance)
Sometimes there are use cases whose behavior can be applied to several different
cases, but with small changes. Unlike the include relationship that uses
everything as is without changes, this relationship allows you to reuse a subset
of behavior with small changes for a collection of specific situations. In
object-oriented terms, you have a number of specialized cases of a generalized
use case.

For example, currently you have a "Create Account" use case but what if you want
to have different types of accounts that differ slightly from each other? You
would want to describe the general behavior use case of "Create Account" and
then define specialized use cases in which the account being created is a
specific type, such as a regular account or a management account.

By using inheritance, you are saying that every step in the general use case
must occur in the specialized use case.

#### How to add an Inheritance
```Go
fooSystem.AddInheritance(
  fooSystem.UseCase("Create a user account"),
  fooSystem.UseCase("Another UseCase Child"),
)
```
## Complete Example
```Go
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
```

Run the Go code to generate DOT code:

```Bash
go run example/useCase/readmeExample.go | dot -Tsvg > out.html
```

And that will produce an svg image of the following:

![output graph](/examples/useCase/example.svg)
