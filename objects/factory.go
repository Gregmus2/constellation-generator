package objects

import (
	"github.com/Gregmus2/simple-engine/objects"
)

type ObjectFactory struct {
	*objects.ObjectFactory
}

func NewObjectFactory(f *objects.ObjectFactory) *ObjectFactory {
	return &ObjectFactory{ObjectFactory: f}
}
