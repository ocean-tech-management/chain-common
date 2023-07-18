package gen

import (
	context "context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	cli := NewClient("3.39.236.151:8010", "testapi", []byte(TestPri), "Test")
	resp, err := cli.CreateChainAccount(context.Background(), &CreateAccountRequest{ApiKey: "testkey", Uid: "testuid"})
	fmt.Println(resp, err)
}
