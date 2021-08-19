package testing

import "testing"

func TestIsValidEmailAddress(t *testing.T) {
	type args struct {
		emailAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{" correct", args{"lakshmi.mothadi@gmail.com"}, true},
		{"not correct ", args{"lakshmi xebia@gmail.com"}, false},
		{"this is correct", args{"lakshmi.mothadi@xebia.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmailAddress(tt.args.emailAddress); got != tt.want {
				t.Errorf("IsValidEmailAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
