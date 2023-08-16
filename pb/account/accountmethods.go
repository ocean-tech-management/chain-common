package account

import "strings"

func (x *GetAssetStatusReply) SymbolFunctionFrozen(symbol string, funct FrozenFunction) bool {
	for _, v := range x.Items {
		if v.Symbol == "" || strings.EqualFold(v.Symbol, symbol) {
			if v.Function == FrozenFunction_ALL || v.Function == funct {
				return true
			}
		}
	}
	return false
}

// 是否禁用充值
func (x *GetAssetStatusReply) SymbolReChargeFrozen(symbol string) bool {
	return x.SymbolFunctionFrozen(symbol, FrozenFunction_RECHARGE)
}
