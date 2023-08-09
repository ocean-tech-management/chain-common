package kafka

const NotifyDepositTopic = "NotifyDeposit"
const ConsumerOrderGroup = "ConsumerOrderGroup"

type NotifyDeposit struct {
	Chain     string `json:"chain"`
	AccountId string `json:"account_id"`
	ApiKey    string `json:"api_key"`
	Addr      string `json:"addr"`
	Symbol    string `json:"symbol"`
	Contract  string `json:"contract"`
	Hash      string `json:"hash"`
	Amount    string `json:"amount"` //精度处理后还是带精度的？
}
