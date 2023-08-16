package account

import (
	"testing"
)

func TestGetAssetStatusReply_SymbolReChargeFrozen(t *testing.T) {
	type fields struct {
		Items []*AssetStatusVoItem
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "all frozen",
			fields: fields{
				Items: []*AssetStatusVoItem{
					{
						Symbol:   "",
						Function: FrozenFunction_ALL,
					},
				},
			},
			args: args{
				symbol: "BTC",
			},
			want: true,
		},
		{
			name: "recharge frozen",
			fields: fields{
				Items: []*AssetStatusVoItem{
					{
						Symbol:   "BTC",
						Function: FrozenFunction_RECHARGE,
					},
				},
			},
			args: args{
				symbol: "BTC",
			},
			want: true,
		},
		{
			name: "not frozen",
			fields: fields{
				Items: []*AssetStatusVoItem{},
			},
			args: args{
				symbol: "BTC",
			},
			want: false,
		},
		{
			name: "not frozen recharge",
			fields: fields{
				Items: []*AssetStatusVoItem{
					{
						Symbol:   "BTC",
						Function: FrozenFunction_WITHDRAW,
					},
				},
			},
			args: args{
				symbol: "BTC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := &GetAssetStatusReply{
				Items: tt.fields.Items,
			}
			if got := reply.SymbolReChargeFrozen(tt.args.symbol); got != tt.want {
				t.Errorf("SymbolReChargeFrozen() = %v, want %v", got, tt.want)
			}
		})
	}
}
