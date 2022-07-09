package main

import "github.com/Gregmus2/simple-engine/common"

type Init struct{}

func NewInit() common.Init {
	return &Init{}
}

func (i *Init) OpenGL() error {
	return nil
}
