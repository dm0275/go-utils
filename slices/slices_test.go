package slices

import "testing"

func TestStringInSlice(t *testing.T) {
	type args struct {
		wantedString    string
		Stringlist []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Positive_UseCase", args: args{wantedString: "one", Stringlist: []string{"one", "two"},}, want: true},
		{name: "Negative_UseCase", args: args{wantedString: "one", Stringlist: []string{"two", "three"},}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.wantedString, tt.args.Stringlist); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}