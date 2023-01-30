package presenter

import (
	"bytes"
	"io"
	"testing"

	"github.com/JacobTripp/diagrams-as-code/process/activity"
	"github.com/stretchr/testify/assert"
)

func TestDiagramName(t *testing.T) {
	d := activity.New("Foo Activity")
	p := New(func(out io.Writer, d Diagram) error {
		out.Write([]byte("hello"))
		return nil
	})
	var buf bytes.Buffer
	p.Write(&buf, d)
	assert.NotEmpty(t, buf.String())
}
