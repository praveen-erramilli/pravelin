package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"pravelin/logging"
)

func (event *Event) Print() {
	jsonData, err := json.Marshal(event)
	if err != nil {
		logging.ErrorLog.Println("unable to marshal event")
		logging.ErrorLog.Println(err.Error())
	} else {
		fmt.Printf("%s \n", jsonData)
	}
}

func (event *Event) handleCopyPaste(input EventInput) {
	event.CopyAndPaste[input.FormId] = input.Pasted
}

func (event *Event) handleScreenResize(input EventInput) {
	event.ResizeFrom = Dimension{Width: input.ResizeFrom.Width, Height: input.ResizeFrom.Height}
	event.ResizeTo = Dimension{Width: input.ResizeTo.Width, Height: input.ResizeTo.Height}
}

func (event *Event) handleTimeTaken(input EventInput) {
	event.FormCompletionTime = input.TimeTaken
}

func (event *Event) validate() error {
	if event.FormCompletionTime != 0 {
		return errors.New("unable to process event for already submitted form")
	}
	return nil
}

func (event *Event) ProcessEvent(input EventInput) error {
	event.Lock()
	defer event.Unlock()

	if err := event.validate(); err != nil {
		return err
	}

	logging.InfoLog.Printf("Processing event %s \n", input.EventType)

	switch input.EventType {
	case copyPasteEventType:
		event.handleCopyPaste(input)
	case screenResizeEventType:
		event.handleScreenResize(input)
	case timeTakenEventType:
		event.handleTimeTaken(input)
	default:
		return errors.New("invalid value specified for EventType")
	}
	event.Print()
	return nil
}
