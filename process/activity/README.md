Activity Diagram
================

Use cases show what your system should do. Activity diagrams allow you to
specify how your system will accomplish its goals. Activity diagrams show
high-level actions chained together to represent a process occurring in your
systems. For example, you can use an activity diagram to model the steps
involved with creating a blog account.

# Components of an Activity Diagram

## Actions
An Action is a step in the overall activity. For example, for the activity of
"Was Car" there could be the Actions of Lather, Rinse, and Dry.

#### How to create an action
```Go
```

### Linking to other activity diagrams
Sometimes and activity gets too complex to be useful, in that case they can be
broken down into other activity diagrams and linked to each other.

## Decisions and Merges

### Decisions
Decisions are used when you want to execute a different sequence of actions
depending on a condition. Each branch from a decision must have a Guard
Condition that determine which edge is taken after a decision node. Guard
Conditions must be mutually exclusive.

#### How to add a decision
```Go
```

### Merges
A Merge node marks the end of the conditional behavior started at the decision
node.

## Forks and Joins

### Forks
Forks represent parallel actions

### Joins
Joins mean that all incoming actions (started by a fork) must finish before the
flow can proceed past the join.

## Time events

### Timeout
A timeout represents a waiting period during and activity. It has an arrow in
and out.

### Cron
A cron represents a timing event that initiates the activity in regular
intervals. This is represented with just one arrow leaving the time symbol.

## Objects
Objects are data objects that are an important aspect of the process you're
modeling. They are represented as rectangular nodes in the flow of the actions.

### Input Object Pin
Instead of having a dedicated node for an object you can also just represent the
input to an action.

### Output Object Pin
Same as input but for the output of an action. The pins can show transformations
of the object. For example, and Ouput Pin of Order might have a transformation
of Order.Cost as the Input Pin of another activity.

### Send and Receiving Signals
Activities may involve interactions with external people, systems, or processes.
For example, when authorizing a credit card payment, you need to verify the card
by interacting with an approval service provided by the credit card company.

#### Send
Send signals are signals sent to any external participant. Send signals are
non-blocking, they instantly go to the next step after sending the signal.

#### Receive
A receive signal has the effect of waking up an action in your activity diagram.
Receive signals are blocking, they must get a response before moving to the next
step.

## Starting an Activity
There are a few different ways an activity can start:
  - basic initial node
  - receiving input data object
  - time events
  - receive signal

## Interrupts
Actions can be placed into an interrupt group to signify that they will respond
to a receive signal and continue down it's activity path.

## Ending activity paths
Sometimes and activity has many paths that end at different times and that can
be represented individually.

## Partitions
Actions can be grouped into partitions to represent different involvement of
participants. For example, an activity for support might have some actions taken
by 1st support, some by 2nd support, and some by engineering.

## Expansion Regions
Expansion regions show that actions in a region are performed for each item in
an input collection. For example, an expansion region could be used to model a
software function that takes a list of files as input and seraches each file for
a search term.
