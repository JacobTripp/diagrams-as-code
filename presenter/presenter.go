package presenter

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("vim-go")
}

/*
Things to present:
	diagram
	node
	edge
*/
type writeFn func(io.Writer, Diagram) error
type Presenter struct {
	w writeFn
}

type Diagram interface {
}

type Interface interface {
	Write(io.Writer, Diagram) error
}

func New(w writeFn) Interface {
	return Presenter{w: w}
}

func (p Presenter) Write(out io.Writer, d Diagram) error {
	fn := p.w
	_ = fn(out, d)
	return nil
}
