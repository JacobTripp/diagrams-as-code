Diagrams-as-Code
================
This project was created because I wanted to learn Graphviz/DOT, gain
deeper understanding of UML notation, as well as to improve communication about
software design, and to more easily model systems so I can learn more quickly.

This is a work in progress but suggestions and contributions are welcome. I
originally didn't intend to open source this project but I thought "Why not?" so
you are warned that there is some sloppy Saturday morning careless code. But I
do see a great future for this code.

## Description
When modeling a system there are many levels of detail that can be represented
with diagrams. There are 5 groups of diagrams, Use Case, Logical, Process,
development, and physical. A group can have many different kinds of diagrams.

For an example on planning new software see the "SPOMit" project in the examples
directory.

## TODO
- [ ] Documentation
- [ ] CLI program to guide the design process.
- [ ] Make into a server (automate with docker)
- [ ] Add a .Write(io.Writer) method to replace the use of .String()
- [ ] Get away from DOT and inject a rendering dependency, like to produce
  straight SVGs or even use DOT still.
- [ ] Diagram generating code, a better way to define new diagrams
- [ ] Add code generation to make the boiler plate code based on the definitions
  of diagrams.
- More diagrams!
  - [ ] Class diagram
  - [ ] Object diagram
  - [ ] Sequence diagram
  - [ ] Communication Diagram
  - [ ] Timing Diagram
  - [ ] Interaction Diagram
  - [ ] Composite Structure Diagram
  - [ ] Component Diagram
  - [ ] Package Diagram
  - [ ] State Machine Diagram
  - [ ] Deployment Diagram
