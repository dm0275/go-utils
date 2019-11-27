package errors

import (
	"testing"
)

var err error

func Test_checkError(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Nil_Error", args: args{e: nil}},
		{name: "Validate_Error", args: args{e: err}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}