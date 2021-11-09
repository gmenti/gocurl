package gocurl

import (
	"testing"

	"gotest.tools/assert"
)

func TestExec(t *testing.T) {
	type args struct {
		filepath string
		v        *map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		assert  func(args)
	}{
		{
			name: "should execute a curl",
			args: args{
				filepath: "./test_curl.sh",
				v:        &map[string]interface{}{},
			},
			wantErr: false,
			assert: func(a args) {
				data := *a.v
				assert.Equal(t, float64(200), data["code"])
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(tt.args.filepath, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				tt.assert(tt.args)
			}
		})
	}
}
