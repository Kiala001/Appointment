package events

import "github.com/stretchr/testify/mock"

type EventBus struct {
	mock.Mock
}

func (e EventBus) Publish(){
	
}