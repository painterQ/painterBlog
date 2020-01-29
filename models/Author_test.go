package models

import "testing"

func Test_pwd2key(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1",args{pwd:"123456"},"123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pwd2key(tt.args.pwd); got != tt.want {
				t.Errorf("pwd2key() = %v, want %v", got, tt.want)
			}
		})
	}
}