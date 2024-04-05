package restyHttp

import (
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	once   sync.Once
	client *resty.Client
)

// GetHttpClient
// @Description: 实例化
// @param proxys
// @return *resty.Client
func GetHttpClient(proxys ...string) *resty.Client {
	once.Do(func() {
		client = resty.New()
		if len(proxys) > 0 {
			proxy := proxys[0]
			client.SetProxy(proxy)
		}
		client.SetTimeout(time.Second * 5)
	})
	return client
}

// GetMobileHttpRequest
// @Description: 设置用户代理
// @return *resty.Request
func GetMobileHttpRequest() *resty.Request {
	return GetHttpClient().SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1").R()
}
