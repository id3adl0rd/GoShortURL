package validator

import "testing"

func TestUrlParse(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				value: "http://google.com",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UrlParse(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UrlParse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
