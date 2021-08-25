package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
		inputValue   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default value",
			args: args{"test", "value", ""},
			want: "value",
		}, {
			name: "os env value",
			args: args{"test", "value", "valueos"},
			want: "valueos",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.inputValue != "" {
				os.Setenv(tt.args.key, tt.args.inputValue)
			}
			got := GetValue(tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
