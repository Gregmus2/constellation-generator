package objects

import (
	"github.com/Gregmus2/simple-engine/graphics"
)

type ObjectFactory struct {
	shape   *graphics.ShapeHelper
	program *graphics.Program
}

func NewObjectFactory(s *graphics.ShapeHelper, p *graphics.Program) *ObjectFactory {
	return &ObjectFactory{shape: s, program: p}
}
