Class Diagrams
==============

Classes are usually represented with a box with three sections, the first
section is the class name, the middle section are the attributes, and the last
section are the operations.

## Visibility
How does a class selectively reveal its operations and data to other classes? By
using visibility.

### Public
The most accessible of visibility. This is represented with a "+" symbol. The
public interface is the part of your class that other classes will depend on the
most. It is important that your public interfaces change as little as possible.

### Protected
Protected are elements that aren't accessible unless inherited from the parent.
Inheritance is a bad design pattern in my opinion and composition is a better
way to go. Protected attributes/operations are represented with a "#" symbol.

### Package
Package visibility is for attributes that are accessible to all members of a
package but not accessible outside of the package. They are represented by the
"~" symbol.

### Private
Private visibility is the most tightly constrained only the class where it's
defined can access it, not even other classes in the same package have access.
This level of visibility is represented with the "-" symbol.

## Attributes
Attributes can be represented as just a name or better yet with the pattern
<visibility><name>: <type>[multiplicity] but they can also be represented by an attribute
association.

### Multiplicity
You can represent that an attribute is actually a collection by using
multiplicity. For example, if a class has a package attribute "orders" that is a
list of Order objects then that is represented like "~orders: Orders[\*]{unique}"
But if you wanted to represent a list that is limited then instead of a start
then it would be a number for a fixed length or a range by using "1..5" for
example for a list that is 1-5 in length.

The additional multiplicity of "{unique}" means that there are duplicates in the
list and this is the default if it's not specified. Other options include:
  - ordered
  - readOnly
  - union
  - subset

## Operations
A class's operations describe what a class can do but not necessarily how it is
going to do it. An operation is more like a minimal contract. A class's
operations should totally encompass all of the behavior that the class contains,
including all the work that maintains the class's attributes.

An operations is represented as <visibility><name>([parameters]): <return type>

## Static
Attributes and Operations can be set as static and are represented an underline.

## Class Relationships
The types of supported relationships are, from weakest to strongest
relationship:
  - Dependency
  - Association
  - Aggregation
  - Composition
  - Inheritance

### Dependency
### Association
### Aggregation
### Composition
### Inheritance
