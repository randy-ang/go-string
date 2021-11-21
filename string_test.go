package gostring_test

import (
	"testing"

	gostring "github.com/randy-ang/go-string/cmd/go-string"
)

func TestFormat(t *testing.T) {
	type args struct {
		format  string
		mapArgs map[string]interface{}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test map equivalent map argument",
			args: args{
				mapArgs: map[string]interface{}{
					"testInt":   123,
					"testFloat": 1.2,
				},
				format: "some-format{testNoArg},{testFloat}",
			},
			want: "some-format{testNoArg},1.2",
		},
		{
			name: "replace value using map value",
			args: args{
				mapArgs: map[string]interface{}{
					"testInt":   123,
					"testFloat": 1.2,
				},
				format: "some-format{testInt},{testFloat}",
			},
			want: "some-format123,1.2",
		},
		{
			name: "map arguments is nil",
			args: args{
				mapArgs: nil,
				format:  "some-format{test}",
			},
			want: "some-format{test}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gostring.Format(tt.args.format, tt.args.mapArgs); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
