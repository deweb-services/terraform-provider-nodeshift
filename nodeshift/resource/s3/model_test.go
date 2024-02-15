package s3

import (
	"reflect"
	"testing"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBucketResourceModel_FromClientResponse(t *testing.T) {
	type fields struct {
		Key types.String
	}
	type args struct {
		c *client.S3BucketConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "bucket resource model from client response",
			fields: fields{
				Key: types.StringValue(""),
			},
			args: args{
				c: &client.S3BucketConfig{
					Key: "key",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BucketResourceModel{
				Key: tt.fields.Key,
			}
			if err := m.FromClientResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketResourceModel_ToClientRequest(t *testing.T) {
	type fields struct {
		Key types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.S3BucketConfig
		wantErr bool
	}{
		{
			name: "bucket resource model to client request",
			fields: fields{
				Key: types.StringValue(""),
			},
			want:    &client.S3BucketConfig{},
			wantErr: false,
		},
		{
			name:    "bucket resource model to client request error",
			fields:  fields{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BucketResourceModel{
				Key: tt.fields.Key,
			}
			got, err := m.ToClientRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClientRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToClientRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
