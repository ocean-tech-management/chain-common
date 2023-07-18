package gen

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/suiguo/hwlib/logger"
	"github.com/suiguo/hwlib/rsa"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const TestPri = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDhUAYhmbx6X80S
LBMwTiD3xPTzuIfUjvq8itWzIXqpHJrc26wjV9Dwqt3XXeO44eeUA6gcbIzvoHU/
+K5yaIZb+6X22tu9eEoLwwgH5yd+zKZHgizUDRlR7/L25mRm3OiSywJ28W+B1PVI
2i/qV49uvhwyspRkQ4ZOR5An76EiGMa6GZmYAvX4J2AAnKjDnITNVZbKcY184ZJB
wLFHX9PyGVYvMpPMFTkFuSw2Dah90xKXuIN17Yp+jBXiKllLdyuH20GEpnhrvmQX
K+V9yZQF/SZdbKwxPIaFQ+IlahHoRWIsBYLHUMvCii7/nfBMbJc8I5KlNzWWEb98
lwd+5v1VAgMBAAECggEAX2D/53NsHSW26w1rZhR80kY3J4EjRvDr5aqkAjuW8EXi
rMH1YdTH50l9tbDSOK6w2LCflsDA/KOhXt5IAriKwB+MRy6ovNFSDx6VpLrOtlet
wDG0BmQxJsV7xdcMA5tafOGhfnaALKbY2uk6RWqhllC2ISQFu1f2X+bdeHpxbLvi
7T5RUvTDzomb+FP1+zxNJkeTAAP/ecbvysCXRWBLz4ujkX6Znk5DJbjYkbVaPXjF
CMEb2XJmf2LYa7LxgprZdfwn3FmM7HyiWdjlCEeF7ztFbPYu7xreJM2CI5ln+VBB
ZsC31/sLClgocgvSw4F8P+x4QR9AxuBMZrIerU2U4QKBgQDrxWxRqvnNmbgooIaW
GXUMxPHs+WPaQ/G7afkvdvQ9JfqL5uBiZPLzVzN1mdRwG0NBIbOSvW3cyCqAzJmi
MNX/aSw/LjhE9lv3YEz2+el4DaG44YSLyWqrwvhg+8bgkL0b6R2tt4kv8ZFunABF
GdUjh4yqFO6wLQzRil0xlhhyOQKBgQD0pOJMGCCr83Sg0bSp/Ek4Px7XeBc36E/b
G+AIoQW3E6hHl2irlfSg58w8oEjL/ZWnkoM90i09eG72zugUbSquREx11ZBb3/d/
TaypPgbmYPfPZ5deNYNRCBOwB0TfZQVY61zZwIQePoKYOpst1pmO4HW5eqrsuIE0
Hy2BDuLz/QKBgQDI47nMyFL69Wyt6UFj6aMLU1ATq9eB5Xy2RLCW2dN3usGFrR4p
mHxODVICdSVGtGQUvgOFF7ThdwiIIE0TnmroqpOR7e7yC7wGxt4tXnmo2mye3EEU
3nTujz0VXdJyC8GmY3XvS8AvwrQ5O+Ea/8zU7i1TymwqXuhaLnwDwTFpOQKBgEMM
9eD5M+ss3Kg/EY8NUFwUILXZejOeflSFPU/gIhrdTl/gxZVRkiyPm1B6dblDFUNK
SqIk+rVATtliOAVxLiN5IKOjFt+3cLP5a/suvuFhbknwEHKHpCgPWKYEOAIqQ97t
ExOzSgKoC08BsQpNKOUZ6+ocDsC2iOLqMKSQme9ZAoGATU8m0clvHOQiV1QL1h5T
b9FCMQApxP3O5Z423JBE0H+OL+WL4/riUb/RD59QOD88LhiihqY+6wudMeG+W8PW
/WY5zeBX0bmfrmKXK5e50ljrtOg+TppQ2895oqYL+0w6i3KP5qjwzc8KFtScCB0v
XQuvzNTylHVFrTBPyE9aRZ0=
-----END PRIVATE KEY-----
`

type ChainClient struct {
	ChainServiceClient
	pri    []byte
	apikey string
	name   string
}

func newSignData(req interface{}, apikey string, nonce string) ([]byte, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	extra := "api_key=" + apikey + "&nonce=" + nonce
	data = append(data, []byte(extra)...)
	return data, nil
}

// 签名等
func (c *ChainClient) interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	requestid := uuid.NewString()
	signdata, err := newSignData(req, c.apikey, requestid)
	sign, err := rsa.SignRSA(signdata, c.pri)
	if err != nil {
		return err
	}
	begin := time.Now()
	defer func() {
		logger.NewStdLogger(1).Info("Grpc", "nonce", requestid, "method", method, "Error", err, "cost ", fmt.Sprintf("%15s", time.Since(begin)))
	}()
	header := metadata.New(map[string]string{
		"nonce":   requestid,
		"api_key": c.apikey,
		"sign":    sign,
		"name":    c.name,
	})
	ctx = metadata.NewOutgoingContext(ctx, header)
	return invoker(ctx, method, req, reply, cc, opts...)
}

// 数据签名 url grpc 连接  apiky  通信的rsa私钥 name区分使用哪个公钥验签
func NewClient(URL string, apikey string, pri []byte, name string) *ChainClient {
	chaincli := &ChainClient{pri: pri, apikey: apikey, name: name}
	conn, err := grpc.Dial(URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(chaincli.interceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := NewChainServiceClient(conn)
	chaincli.ChainServiceClient = client
	return chaincli
}
