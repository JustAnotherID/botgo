package main

import (
	"fmt"

	"github.com/JustAnotherID/botgo/dto"
	"github.com/JustAnotherID/botgo/event"
)

// ThreadEventHandler 论坛主贴事件
func ThreadEventHandler() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		fmt.Println(event, data)
		return nil
	}
}
