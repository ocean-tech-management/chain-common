package gen

import (
	context "context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	cli := NewClient("16.162.237.77:8010", "testapi", []byte(TestPri), "Test")
	resp, err := cli.SendToken(context.Background(), &SendTokenRequest{
		ApiKey:    "7955ab59-7c02-4191-9fc3-3f89c007b74d",
		AccountId: "e781d8e7-6bbd-4748-a294-3c1c71ade320",
		ToAddress: "0x0992ee761a4f7b8FC5a5ACAD011a2b6FdC9e8D72",
		Amount:    "100",
		Chain:     ChainType_BSC,
	})
	t.Log(resp)
	fmt.Println(resp, err)
	// fmt.Println(resp, err)
}
