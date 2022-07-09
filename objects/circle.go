package objects

import "github.com/Gregmus2/simple-engine/graphics"
import "github.com/go-gl/gl/v4.6-core/gl"

type CircleModel struct {
	X, Y   float32
	Radius float32
	Color  graphics.Color
}

type Circle struct {
	CircleModel
	prog  *graphics.Program
	Shape *graphics.ShapeHelper
}

func (m *ObjectFactory) NewCircle(model CircleModel) *Circle {
	return &Circle{
		CircleModel: model,
		prog:        m.Prog,
		Shape:       m.Shape,
	}
}

func (o *Circle) Draw(scale, offsetX, offsetY float32) error {
	o.prog.ApplyProgram(o.Color)
	o.Shape.Circle((o.X+offsetX)*scale, (o.Y+offsetY)*scale, o.Radius*scale)
	gl.UseProgram(0)

	return nil
}
