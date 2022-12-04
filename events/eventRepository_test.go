package events

import (
	"reflect"
	"sync"
	"testing"
)

func TestEventRepository_GetOrCreateEvent(t *testing.T) {
	type fields struct {
		Mutex   sync.Mutex
		storage map[string]*Event
	}
	type args struct {
		input EventInput
	}

	copyPasteEvent := &Event{
		SessionId:    "123123-123123-123123123",
		WebsiteUrl:   "https://ravelin.com",
		CopyAndPaste: make(map[string]bool),
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Event
	}{
		{
			name: "shouldGetAlreadyAvailableEvent",
			fields: fields{
				storage: map[string]*Event{
					"123123-123123-123123123": copyPasteEvent,
				},
			},
			args: args{
				input: getCopyPasteInput(),
			},
			want: copyPasteEvent,
		},
		{
			name: "shouldCreateNewEvent",
			fields: fields{
				storage: map[string]*Event{
					"123123": copyPasteEvent,
				},
			},
			args: args{
				input: getCopyPasteInput(),
			},
			want: copyPasteEvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventRepository := &EventRepository{
				Mutex:   tt.fields.Mutex,
				storage: tt.fields.storage,
			}
			if got := eventRepository.GetOrCreateEvent(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrCreateEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
