package parser

import (
	"testing"
)

type (
	mockClient struct{}
)

func (c mockClient) fetch(url string) (*Response, error) {
	r := &Response{
		LongURL: "https://github.com/smith-30/go-goo.gl-parser",
	}

	return r, nil
}

func TestParser_DecodeURL(t *testing.T) {
	type fields struct {
		APIKey  string
		Fetcher Fetcher
	}
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				APIKey:  "a",
				Fetcher: mockClient{},
			},
			args: args{
				target: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "2",
			fields: fields{
				APIKey:  "a",
				Fetcher: mockClient{},
			},
			args: args{
				target: "https://github.com/smith-30/go-goo.gl-parser",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "3",
			fields: fields{
				APIKey:  "a",
				Fetcher: mockClient{},
			},
			args: args{
				"https://goo.gl/kMxDaw",
			},
			want:    "https://github.com/smith-30/go-goo.gl-parser",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				APIKey:  tt.fields.APIKey,
				Fetcher: tt.fields.Fetcher,
			}
			got, err := p.DecodeURL(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.DecodeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parser.DecodeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
