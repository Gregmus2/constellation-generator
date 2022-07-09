package objects

import "github.com/Gregmus2/simple-engine/graphics"
import "github.com/go-gl/gl/v4.6-core/gl"

type LineModel struct {
	X1, Y1, X2, Y2 float32
	Color          graphics.Color
}

type Line struct {
	LineModel
	prog  *graphics.Program
	Shape *graphics.ShapeHelper
}

func (m *ObjectFactory) NewLine(model LineModel) *Line {
	return &Line{
		LineModel: model,
		prog:      m.Prog,
		Shape:     m.Shape,
	}
}

func (o *Line) Draw(scale, offsetX, offsetY float32) error {
	o.prog.ApplyProgram(o.Color)
	o.Shape.Line((o.X1+offsetX)*scale, (o.Y1+offsetY)*scale, (o.X2+offsetX)*scale, (o.Y2+offsetY)*scale)
	gl.UseProgram(0)

	return nil
}
