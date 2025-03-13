package main

import (
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"reflect"
	"testing"
)

func Test_createMessagesFromTemplate(t *testing.T) {
	tests := []struct {
		name string
		want []*schema.Message
	}{
		// TODO: Add test cases.
		{
			name: "Test_createMessagesFromTemplate",
			want: []*schema.Message{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createMessagesFromTemplate()
			t.Log(got)
		})
	}
}

func Test_createTemplate(t *testing.T) {
	tests := []struct {
		name string
		want prompt.ChatTemplate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTemplate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
