package vpc

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestVPCResourceModel_FromClientResponse(t *testing.T) {
	type fields struct {
		ID          types.String
		IPRange     types.String
		Name        types.String
		Description types.String
	}
	type args struct {
		c *client.VPCConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "vpc resource model from client response",
			fields: fields{
				ID:          types.String{},
				IPRange:     types.String{},
				Name:        types.String{},
				Description: types.String{},
			},
			args: args{
				c: &client.VPCConfig{
					ID:          "",
					Name:        "",
					Description: "",
					IPRange:     "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VPCResourceModel{
				ID:          tt.fields.ID,
				IPRange:     tt.fields.IPRange,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
			}
			if err := m.FromClientResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVPCResourceModel_ToClientRequest(t *testing.T) {
	type fields struct {
		ID          types.String
		IPRange     types.String
		Name        types.String
		Description types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.VPCConfig
		wantErr bool
	}{
		{
			name: "vpc resource model to client request ",
			fields: fields{
				ID:          types.String{},
				IPRange:     basetypes.NewStringValue("127.0.0.1/24"),
				Name:        types.String{},
				Description: types.String{},
			},
			want: &client.VPCConfig{
				ID:          "",
				Name:        "",
				Description: "",
				IPRange:     "127.0.0.1/24",
			},
			wantErr: false,
		},
		{
			name: "vpc resource model to client request error",
			fields: fields{
				ID:          types.String{},
				IPRange:     types.String{},
				Name:        types.String{},
				Description: types.String{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VPCResourceModel{
				ID:          tt.fields.ID,
				IPRange:     tt.fields.IPRange,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
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
