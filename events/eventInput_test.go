package events

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_parseEventInput(t *testing.T) {
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    EventInput
		wantErr bool
	}{
		{
			name:    "shouldParseCopyPasteEvent",
			args:    args{request: getCopyPasteRequest()},
			want:    EventInput{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", EventType: copyPasteEventType, Pasted: true, FormId: "inputCardNumber"},
			wantErr: false,
		},
		{
			name:    "shouldParseResizeEvent",
			args:    args{request: getResizeRequest()},
			want:    EventInput{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", EventType: screenResizeEventType, ResizeFrom: DimensionInput{"1920", "1080"}, ResizeTo: DimensionInput{"1280", "720"}},
			wantErr: false,
		},
		{
			name:    "shouldParseTimeTakenEvent",
			args:    args{request: getTimeTakenRequest()},
			want:    EventInput{SessionId: "123123-123123-123123123", WebsiteUrl: "https://ravelin.com", EventType: timeTakenEventType, TimeTaken: 72},
			wantErr: false,
		},
		{
			name:    "shouldFailForInvalidInput",
			args:    args{request: getInvalidRequest()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEventInput(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEventInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEventInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventInput_validate(t *testing.T) {
	type fields struct {
		EventType  string
		WebsiteUrl string
		SessionId  string
		ResizeFrom DimensionInput
		ResizeTo   DimensionInput
		Pasted     bool
		FormId     string
		TimeTaken  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "shouldThrowErrorForInvalidCopyPasteEvent",
			fields: fields{
				EventType: copyPasteEventType,
			},
			wantErr: true,
		},
		{
			name: "shouldValidateForCorrectCopyPasteEvent",
			fields: fields{
				EventType:  copyPasteEventType,
				FormId:     formIdInputEmail,
				WebsiteUrl: "www.example.com",
				SessionId:  "13234",
			},
			wantErr: false,
		},
		{
			name: "shouldThrowErrorForMissingWebsiteURL",
			fields: fields{
				EventType: copyPasteEventType,
				FormId:    formIdInputEmail,
				SessionId: "13234",
			},
			wantErr: true,
		},
		{
			name: "shouldThrowErrorForInvalidFormId",
			fields: fields{
				EventType:  copyPasteEventType,
				FormId:     "button",
				WebsiteUrl: "www.example.com",
				SessionId:  "13234",
			},
			wantErr: true,
		},
		{
			name: "shouldThrowErrorForMissingSessionID",
			fields: fields{
				EventType:  copyPasteEventType,
				FormId:     "submitButton",
				WebsiteUrl: "www.example.com",
			},
			wantErr: true,
		},
		{
			name: "shouldThrowErrorForInvalidResizeEvent",
			fields: fields{
				EventType: screenResizeEventType,
				ResizeFrom: DimensionInput{
					"123", "123",
				},
			},
			wantErr: true,
		},
		{
			name: "shouldThrowErrorForInvalidResizeEvent",
			fields: fields{
				EventType: screenResizeEventType,
				ResizeTo: DimensionInput{
					"123", "123",
				},
			},
			wantErr: true,
		},
		{
			name: "shouldValidateForValidResizeEvent",
			fields: fields{
				EventType: screenResizeEventType,
				ResizeFrom: DimensionInput{
					"123", "123",
				},
				ResizeTo: DimensionInput{
					"123", "123",
				},
				WebsiteUrl: "www.example.com",
				SessionId:  "13234",
			},
			wantErr: false,
		},

		{
			name: "shouldThrowErrorForInvalidTimeTakenEvent",
			fields: fields{
				EventType:  timeTakenEventType,
				TimeTaken:  -10,
				WebsiteUrl: "www.example.com",
				SessionId:  "13234",
			},
			wantErr: true,
		},
		{
			name: "shouldThrowErrorForInvalidTimeTakenEvent",
			fields: fields{
				EventType:  timeTakenEventType,
				TimeTaken:  10,
				WebsiteUrl: "www.example.com",
				SessionId:  "13234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := EventInput{
				EventType:  tt.fields.EventType,
				WebsiteUrl: tt.fields.WebsiteUrl,
				SessionId:  tt.fields.SessionId,
				ResizeFrom: tt.fields.ResizeFrom,
				ResizeTo:   tt.fields.ResizeTo,
				Pasted:     tt.fields.Pasted,
				FormId:     tt.fields.FormId,
				TimeTaken:  tt.fields.TimeTaken,
			}
			if err := i.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getInvalidRequest() *http.Request {
	req, _ := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("}}}"))
	return req
}

func getTimeTakenRequest() *http.Request {
	req, _ := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("{\n  \"eventType\": \"timeTaken\",\n  \"websiteUrl\": \"https://ravelin.com\",\n  \"sessionId\": \"123123-123123-123123123\", \"timeTaken\": 72}"))
	return req
}

func getCopyPasteRequest() *http.Request {
	req, _ := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("{\n  \"eventType\": \"copyAndPaste\",\n  \"websiteUrl\": \"https://ravelin.com\",\n  \"sessionId\": \"123123-123123-123123123\",\n  \"pasted\": true,\n  \"formId\": \"inputCardNumber\"\n}"))
	return req
}

func getResizeRequest() *http.Request {
	req, _ := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("{\n  \"eventType\": \"screenResize\",\n  \"websiteUrl\": \"https://ravelin.com\",\n  "+
			"\"sessionId\": \"123123-123123-123123123\",\n  \"resizeFrom\": {\n    \"width\": \"1920\",\n    "+
			"\"height\": \"1080\"\n  },\n  \"resizeTo\": {\n    \"width\": \"1280\",\n    \"height\": \"720\"\n  }\n}"))
	return req
}
