package utils

import "testing"

func TestGetAwsAcctMatches(t *testing.T) {
	type args struct {
		awsCreds []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAwsAcctMatches(tt.args.awsCreds)
		})
	}
}
