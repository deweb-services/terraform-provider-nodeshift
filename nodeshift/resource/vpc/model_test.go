package vpc

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestVPCResourceModel_ToClientRequest(t *testing.T) {
	t.Parallel()

	type fields struct {
		ID          types.String
		IPRange     types.String
		Name        types.String
		Description types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.CreateVPCRequest
		wantErr bool
	}{
		{
			name: "empty range; vpc resource model to client request; expect ok",
			fields: fields{
				ID:          types.String{},
				IPRange:     types.String{},
				Name:        types.String{},
				Description: types.String{},
			},
			want: &client.CreateVPCRequest{
				Name:        "",
				Description: "",
				IPRange:     "",
			},
		},
		{
			name: "filled range; vpc resource model to client request; expect ok",
			fields: fields{
				ID:          types.String{},
				IPRange:     basetypes.NewStringValue("127.0"),
				Name:        types.String{},
				Description: types.String{},
			},
			want: &client.CreateVPCRequest{
				Name:        "",
				Description: "",
				IPRange:     "127.0",
			},
		},
		{
			name: "vpc resource model to client request; expect error",
			fields: fields{
				ID:          types.String{},
				IPRange:     basetypes.NewStringValue("127.0.0.1/24"),
				Name:        types.String{},
				Description: types.String{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := &ResourceModel{
				UUID:        tt.fields.ID,
				IPRange:     tt.fields.IPRange,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
			}
			got, err := m.ToClientRequest()
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
