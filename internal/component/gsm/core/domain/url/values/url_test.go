package values

import (
	"reflect"
	"testing"
)

func TestNewUrl(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *Url
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				value: "http://google.com",
			},
			want:    &Url{Value: "http://google.com"},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				value: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUrl(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
