package events

import (
	"encoding/json"
	"errors"
	"net/http"
)

type DimensionInput struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

type EventInput struct {
	EventType  string         `json:"eventType"`
	WebsiteUrl string         `json:"websiteUrl"`
	SessionId  string         `json:"sessionId"`
	ResizeFrom DimensionInput `json:"resizeFrom"`
	ResizeTo   DimensionInput `json:"resizeTo"`
	Pasted     bool           `json:"pasted"`
	FormId     string         `json:"formId"`
	TimeTaken  int            `json:"timeTaken"`
}

func (i EventInput) validate() error {
	err := i.validateEventType()
	if err != nil {
		return err
	}

	if len(i.WebsiteUrl) == 0 {
		return errors.New("invalid value specified for websiteUrl")
	}
	if len(i.SessionId) == 0 {
		return errors.New("invalid value specified for sessionId")
	}

	return nil
}

func parseEventInput(request *http.Request) (EventInput, error) {
	var input EventInput
	err := json.NewDecoder(request.Body).Decode(&input)
	return input, err
}

func (i EventInput) validateEventType() error {
	switch i.EventType {
	case copyPasteEventType:
		if len(i.FormId) == 0 {
			return errors.New("invalid value specified for formId")
		}
		if i.FormId != formIdInputEmail && i.FormId != formIdInputCardNumber && i.FormId != formIdInputCvv {
			return errors.New("invalid value specified for formId")
		}
	case screenResizeEventType:
		if len(i.ResizeFrom.Width) == 0 || len(i.ResizeFrom.Height) == 0 {
			return errors.New("invalid value specified for resizeFrom")
		}
		if len(i.ResizeTo.Width) == 0 || len(i.ResizeTo.Height) == 0 {
			return errors.New("invalid value specified for resizeTo")
		}
	case timeTakenEventType:
		if i.TimeTaken <= 0 {
			return errors.New("invalid value specified for timeTaken")
		}
	default:
		return errors.New("invalid value specified for eventType")
	}
	return nil
}
