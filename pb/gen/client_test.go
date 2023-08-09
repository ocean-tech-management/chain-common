package gen

import (
	context "context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	cli := NewClient("16.162.237.77:8010", "testapi", []byte(TestPri), "Test")
	resp, err := cli.SendToken(context.Background(), &SendTokenRequest{
		ApiKey:    "Test",
		AccountId: "01bccb93-4e5f-4b5f-8c60-123dd62aa43f",
		ToAddress: "0x0992ee761a4f7b8FC5a5ACAD011a2b6FdC9e8D72",
		Amount:    "100",
		Chain:     ChainType_BSC,
	})
	fmt.Println(resp, err)
	// fmt.Println(resp, err)
}
