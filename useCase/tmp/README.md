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
```

#### How to add UseCaseDescription
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
```
## Complete Example
```Go
```

Run the Go code to generate DOT code:

```Bash
```

And that will produce an svg image of the following:

![output graph](/examples/useCase/example.svg)
