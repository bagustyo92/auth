package service

import "testing"

func Test_isPhoneNumber(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"#Case1",
			args{"test"},
			false,
		},
		{
			"#Case2",
			args{"081282222234"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPhoneNumber(tt.args.phone); got != tt.want {
				t.Errorf("isPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
