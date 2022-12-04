package events

import (
	"sync"
)

var (
	copyPasteEventType    = "copyAndPaste"
	screenResizeEventType = "screenResize"
	timeTakenEventType    = "timeTaken"

	formIdInputEmail      = "inputEmail"
	formIdInputCardNumber = "inputCardNumber"
	formIdInputCvv        = "inputCVV"
)

type Event struct {
	sync.Mutex
	WebsiteUrl         string          `json:"websiteUrl,omitempty"`
	SessionId          string          `json:"sessionId,omitempty"`
	ResizeFrom         Dimension       `json:"resizeFrom,omitempty"`
	ResizeTo           Dimension       `json:"resizeTo,omitempty"`
	CopyAndPaste       map[string]bool `json:"copyAndPaste,omitempty"`       // map[fieldId]true
	FormCompletionTime int             `json:"formCompletionTime,omitempty"` // Seconds
}

type Dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

func NewEvent(input EventInput) *Event {
	return &Event{SessionId: input.SessionId, WebsiteUrl: input.WebsiteUrl, CopyAndPaste: make(map[string]bool)}
}
