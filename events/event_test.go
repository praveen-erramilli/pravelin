package events

import (
	"reflect"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		input EventInput
	}

	tests := []struct {
		name string
		args args
		want *Event
	}{
		{
			name: "shouldCreateNewEventWithCopyAndPasteInput",
			args: args{input: getCopyPasteInput()},
			want: &Event{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", CopyAndPaste: make(map[string]bool)},
		},
		{
			name: "shouldCreateNewEventWithScreenResizeInput",
			args: args{input: getScreenResizeInput()},
			want: &Event{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", CopyAndPaste: make(map[string]bool)},
		},
		{
			name: "shouldCreateNewEventWithSubmitInput",
			args: args{input: getSubmitInput()},
			want: &Event{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", CopyAndPaste: make(map[string]bool)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvent(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getCopyPasteInput() EventInput {
	return EventInput{
		EventType:  "copyAndPaste",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "123123-123123-123123123",
		Pasted:     true,
		FormId:     "inputCardNumber"}
}

func getScreenResizeInput() EventInput {
	return EventInput{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "123123-123123-123123123",
		ResizeFrom: DimensionInput{Width: "1920", Height: "1080"},
		ResizeTo:   DimensionInput{Width: "1280", Height: "720"},
	}
}

func getSubmitInput() EventInput {
	return EventInput{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "123123-123123-123123123",
		TimeTaken:  72,
	}
}
