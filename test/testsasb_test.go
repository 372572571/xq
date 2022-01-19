package test

import (
	"testing"
)

func Test_sbab(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			est_sbab(tt.args.t)
		})
	}
}
