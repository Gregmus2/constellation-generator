package main

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/scenes"
	"github.com/go-gl/glfw/v3.3/glfw"
	"math"
	"math/rand"
	"neurons-graph/objects"
	"sort"
	"time"
)

const size = 1

type Scene struct {
	scenes.Base

	factory *objects.ObjectFactory
	con     *graphics.PercentToPosConverter

	neurons []*Neuron
}

type Neuron struct {
	circle      *objects.Circle
	connections []connection
}

type connection struct {
	neuron *Neuron
	line   *objects.Line
}

func NewScene(base scenes.Base, factory *objects.ObjectFactory, con *graphics.PercentToPosConverter) *Scene {
	return &Scene{
		Base:    base,
		factory: factory,
		con:     con,
		neurons: make([]*Neuron, 0),
	}
}

type Sibling struct {
	neuron   *Neuron
	distance float64
}

func (s *Scene) Init() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < size*50; i++ {
		neuron := s.factory.NewCircle(objects.CircleModel{
			Radius: 2.0,
			Color:  graphics.Yellow(),
		})
		s.DrawObjects.Put(neuron)
		s.neurons = append(s.neurons, &Neuron{
			circle: neuron,
		})
	}

	s.placeNeurons()
}

func (s *Scene) placeNeurons() {
	for _, neuron := range s.neurons {
		neuron.circle.X = float32(rand.Intn(s.Cfg.Window.W*size)) - float32(s.Cfg.Window.W)*0.5*float32(size-1)
		neuron.circle.Y = float32(rand.Intn(s.Cfg.Window.H*size)) - float32(s.Cfg.Window.H)*0.5*float32(size-1)

		for _, conn := range neuron.connections {
			s.DrawObjects.Remove(conn.line)
		}
		neuron.connections = make([]connection, 0)
	}

	for i, neuron := range s.neurons {
		x, y := neuron.circle.X, neuron.circle.Y
		siblings := make([]Sibling, 0, len(s.neurons)-1)
		for j := 0; j < len(s.neurons); j++ {
			if i == j {
				continue
			}

			someX, someY := s.neurons[j].circle.X, s.neurons[j].circle.Y
			distance := math.Sqrt(math.Pow(float64(x-someX), 2) + math.Pow(float64(y-someY), 2))
			siblings = append(siblings, Sibling{
				neuron:   s.neurons[j],
				distance: distance,
			})
		}

		sort.Slice(siblings, func(i, j int) bool {
			return siblings[i].distance < siblings[j].distance
		})

		connections := rand.Intn(3)
		for j := 0; j < connections; j++ {
			line := s.factory.NewLine(objects.LineModel{
				neuron.circle.X, neuron.circle.Y,
				siblings[j].neuron.circle.X, siblings[j].neuron.circle.Y,
				graphics.White(),
			})
			neuron.connections = append(neuron.connections, connection{
				neuron: siblings[j].neuron,
				line:   line,
			})
			s.DrawObjects.Put(line)
		}
	}
}

func (s *Scene) Callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	s.Base.Callback(w, key, scancode, action, mods)

	if action == glfw.Press {
		switch key {
		case glfw.KeySpace:
			s.placeNeurons()
		}
	}
}

/*
todo
	create neurons (box)
	create connections (line)
	create graph
	random generation
	logic
*/
