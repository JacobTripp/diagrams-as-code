//
// Auto-generated code, DO NOT EDIT
//
package activity

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/diagram"
)

var Action = diagram.NodeType{
	Name: "Action",
	Description: "An Action is a step in the overall activity. For example, for the activity of 'Was Car' there could be the Actions of Lather, Rinse, and Dry.",
	Attributes: map[string]string{
		"shape": "rect",
		"style": "rounded",
	},
}

func (d Activity) AddAction(name string) error {
	if err := d.d.AddNode(Action, name); err != nil {
		return err
	}
	return nil
}

var Object = diagram.NodeType{
	Name: "Object",
	Description: "Objects are data objects that are an important aspect of the process you're modeling. They are represented as rectangular nodes in the flow of the actions.",
	Attributes: map[string]string{
		"shape": "rect",
	},
}

func (d Activity) AddObject(name string) error {
	if err := d.d.AddNode(Object, name); err != nil {
		return err
	}
	return nil
}

var Decision = diagram.NodeType{
	Name: "Decision",
	Description: "Decisions are used when you want to execute a different sequence of actions depending on a condition. Each branch from a decision must have a Guard Condition that determine which edge is taken after a decision node. Guard Conditions must be mutually exclusive.",
	Attributes: map[string]string{
		"label": "",
		"orientation": "90",
		"shape": "utr",
	},
}

func (d Activity) AddDecision(name string) error {
	if err := d.d.AddNode(Decision, name); err != nil {
		return err
	}
	return nil
}

var Expansion = diagram.NodeType{
	Name: "Expansion",
	Description: "Expansion regions show that actions in a region are performed for each item in an input collection. For example, an expansion region could be used to model a software function that takes a list of files as input and seraches each file for a search term.",
	Attributes: map[string]string{
		"label": "",
		"shape": "noverhang",
	},
}

func (d Activity) AddExpansion(name string) error {
	if err := d.d.AddNode(Expansion, name); err != nil {
		return err
	}
	return nil
}

var ActivityEnd = diagram.NodeType{
	Name: "Activity End",
	Description: "This represents and end to the entire activity.",
	Attributes: map[string]string{
		"fillcolor": "black",
		"label": "",
		"shape": "doublecircle",
	},
}

func (d Activity) AddActivityEnd(name string) error {
	if err := d.d.AddNode(ActivityEnd, name); err != nil {
		return err
	}
	return nil
}

var PathEnd = diagram.NodeType{
	Name: "Path End",
	Description: "Sometimes and activity has many paths that end at different times and that can be represented individually.",
	Attributes: map[string]string{
		"label": "X",
		"shape": "doublecircle",
	},
}

func (d Activity) AddPathEnd(name string) error {
	if err := d.d.AddNode(PathEnd, name); err != nil {
		return err
	}
	return nil
}

var Fork = diagram.NodeType{
	Name: "fork",
	Description: "Forks represent parallel actions",
	Attributes: map[string]string{
		"label": "",
		"orientation": "90",
		"shape": "restrictionsite",
	},
}

func (d Activity) AddFork(name string) error {
	if err := d.d.AddNode(Fork, name); err != nil {
		return err
	}
	return nil
}

var Init = diagram.NodeType{
	Name: "Init",
	Description: "This represents a basic start to an activity flow.",
	Attributes: map[string]string{
		"fillcolor": "black",
		"label": "",
		"shape": "circle",
	},
}

func (d Activity) AddInit(name string) error {
	if err := d.d.AddNode(Init, name); err != nil {
		return err
	}
	return nil
}

var Join = diagram.NodeType{
	Name: "Join",
	Description: "Joins mean that all incoming actions (started by a fork) must finish before the flow can proceed past the join.",
	Attributes: map[string]string{
		"label": "",
		"orientation": "270",
		"shape": "restrictionsite",
	},
}

func (d Activity) AddJoin(name string) error {
	if err := d.d.AddNode(Join, name); err != nil {
		return err
	}
	return nil
}

var Merge = diagram.NodeType{
	Name: "Merge",
	Description: "A Merge node marks the end of the conditional behavior started at the decision node.",
	Attributes: map[string]string{
		"label": "",
		"shape": "diamond",
	},
}

func (d Activity) AddMerge(name string) error {
	if err := d.d.AddNode(Merge, name); err != nil {
		return err
	}
	return nil
}

var Partition = diagram.NodeType{
	Name: "Partition",
	Description: "Actions can be grouped into partitions to represent different involvement of participants. For example, an activity for support might have some actions taken by 1st support, some by 2nd support, and some by engineering.",
	Attributes: map[string]string{
		"cluster": "true",
		"label": "",
	},
}

func (d Activity) AddPartition(name string) error {
	if err := d.d.AddNode(Partition, name); err != nil {
		return err
	}
	return nil
}

var SendSignal = diagram.NodeType{
	Name: "Send Signal",
	Description: "Send signals are signals sent to any external participant. Send signals are non-blocking, they instantly go to the next step after sending the signal.",
	Attributes: map[string]string{
		"orientation": "90",
		"shape": "house",
	},
}

func (d Activity) AddSendSignal(name string) error {
	if err := d.d.AddNode(SendSignal, name); err != nil {
		return err
	}
	return nil
}

var RecieveSignal = diagram.NodeType{
	Name: "Recieve Signal",
	Description: "A receive signal has the effect of waking up an action in your activity diagram. Receive signals are blocking, they must get a response before moving to the next step.",
	Attributes: map[string]string{
		"orientation": "270",
		"shape": "house",
	},
}

func (d Activity) AddRecieveSignal(name string) error {
	if err := d.d.AddNode(RecieveSignal, name); err != nil {
		return err
	}
	return nil
}

var TimeEvent = diagram.NodeType{
	Name: "Time Event",
	Description: "A timeout represents a waiting period during and activity.",
	Attributes: map[string]string{
		"shape": "invtriangle",
	},
}

func (d Activity) AddTimeEvent(name string) error {
	if err := d.d.AddNode(TimeEvent, name); err != nil {
		return err
	}
	return nil
}

var Transformation = diagram.NodeType{
	Name: "Transformation",
	Description: "Transformations represent how objects change during the transmissions. For example, one action might output and object but only the transformation of that object as Object.Attribute.",
	Attributes: map[string]string{
		"shape": "Msquare",
	},
}

func (d Activity) AddTransformation(name string) error {
	if err := d.d.AddNode(Transformation, name); err != nil {
		return err
	}
	return nil
}
var adders = map[string]func(Activity, string) error {
"Action": Activity.AddAction,
"Object": Activity.AddObject,
"Decision": Activity.AddDecision,
"Expansion": Activity.AddExpansion,
"ActivityEnd": Activity.AddActivityEnd,
"PathEnd": Activity.AddPathEnd,
"Fork": Activity.AddFork,
"Init": Activity.AddInit,
"Join": Activity.AddJoin,
"Merge": Activity.AddMerge,
"Partition": Activity.AddPartition,
"SendSignal": Activity.AddSendSignal,
"RecieveSignal": Activity.AddRecieveSignal,
"TimeEvent": Activity.AddTimeEvent,
"Transformation": Activity.AddTransformation,
}
var Default = diagram.EdgeType{
	Name: "Default",
	Description: "The main wait activities are connected.",
	Attributes: map[string]string{
		"arrowhead": "vee",
	},
}

var Interrupt = diagram.EdgeType{
	Name: "Interrupt",
	Description: "The action to take for an inturrupt.",
	Attributes: map[string]string{
		"arrowhead": "curve",
	},
}

type Activity struct {
	d *diagram.Diagram
	name string
}

func New(name string) Activity {
	diagram := diagram.New(
		diagram.DiagramType{
			Name: "Activity",
			Description: "Use cases show what your system should do. Activity diagrams allow you to specify how your system will accomplish its goals. Activity diagrams show high-level actions chained together to represent a process occurring in your systems. For example, you can use an activity diagram to model the steps involved with creating a blog account.",
			Attributes: map[string]string{
				"splines": "ortho",
			},
		},
		diagram.WithNodeTypes(Action,Object,Decision,Expansion,ActivityEnd,PathEnd,Fork,Init,Join,Merge,Partition,SendSignal,RecieveSignal,TimeEvent,Transformation),
		diagram.WithEdgeTypes(Default,Interrupt),
	)
	return Activity{name: name, d: &diagram}
}

func (d Activity) Connect(t diagram.EdgeType, from, to string) error {
	if err := d.d.AddEdge(t, from, to); err != nil {
		return err
	}
	return nil
}

func (d Activity) Add(nodeType, name string) error {
	fn, ok := adders[nodeType]
	if !ok {
		return fmt.Errorf("'%s' is not a valid node type", nodeType)
	}
	if err := fn(d, name); err != nil {
		return err
	}
	return nil
}