package api

import (
	"errors"
	"testing"
)

func Test_errorStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "HTTP error",
			args: args{
				err: errors.New("404 NotFound"),
			},
			want: 404,
		},
		{
			name: "Not HTTP error",
			args: args{
				err: errors.New("foo"),
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorStatusCode(tt.args.err); got != tt.want {
				t.Errorf("errorStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
