package events

import (
	"reflect"
	"testing"
)

func TestEvent_ProcessEvent(t *testing.T) {

	type args struct {
		input EventInput
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		postFunction func(*Event)
	}{
		{
			name:    "shouldProcessSuccessfullyForCopyPasteInput",
			args:    args{input: getCopyPasteInput()},
			wantErr: false,
			postFunction: func(event *Event) {
				if !event.CopyAndPaste["inputCardNumber"] {
					t.Errorf("Failed to set copyAndPaste")
				}
			},
		},
		{
			name:    "shouldProcessSuccessfullyForResizeInput",
			args:    args{input: getScreenResizeInput()},
			wantErr: false,
			postFunction: func(event *Event) {
				if want := getScreenResizeInput().ResizeFrom; !(reflect.DeepEqual(event.ResizeFrom.Width, want.Width) && reflect.DeepEqual(event.ResizeFrom.Height, want.Height)) {
					t.Errorf("Failed to set resizeFrom. Got = %v, want = %v", event.ResizeFrom, want)
				}
				if want := getScreenResizeInput().ResizeTo; !(reflect.DeepEqual(event.ResizeTo.Width, want.Width) && reflect.DeepEqual(event.ResizeTo.Height, want.Height)) {
					t.Errorf("Failed to set resizeTo. Got = %v, want = %v", event.ResizeTo, want)
				}
			},
		},
		{
			name:    "shouldProcessSuccessfullyForTimeTakenInput",
			args:    args{input: getSubmitInput()},
			wantErr: false,
			postFunction: func(event *Event) {
				if want := getSubmitInput().TimeTaken; !reflect.DeepEqual(event.FormCompletionTime, want) {
					t.Errorf("Failed to set FormCompletionTime. Got = %v, want = %v", event.FormCompletionTime, want)
				}
			},
		},
		{
			name:    "shouldFailForAlreadySubmittedSession",
			args:    args{input: getSubmitInput()},
			wantErr: false,
			postFunction: func(event *Event) {
				if err := event.ProcessEvent(getSubmitInput()); err == nil {
					t.Errorf("ProcessEvent() error = %v, wantErr %v", err, true)
				}
			},
		},
		{
			name:    "shouldFailForInvalidInput",
			args:    args{input: getInvalidInput()},
			wantErr: true,
			postFunction: func(event *Event) {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := NewEvent(tt.args.input)
			if err := event.ProcessEvent(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ProcessEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.postFunction(event)
		})
	}
}

func getInvalidInput() EventInput {
	return EventInput{
		EventType:  "xyz",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "123123-123123-123123123",
		TimeTaken:  72,
	}
}
