//tencentai.go
package tencentai

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

type AiClient struct {
	Client *http.Client
	Appid  string
	Appkey string
	Mutext sync.Mutex
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	databytes := []byte(str)
	byteslen := len(databytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, databytes[r.Intn(byteslen)])
	}
	return string(result)
}

func NewAiClient(appid, appkey string) *AiClient {
	client := &AiClient{Appid: appid, Appkey: appkey}
	client.Client = &http.Client{}

	return client
}

func (this *AiClient) getMd5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return strings.ToUpper(string(md5str))
}

func (this *AiClient) GetBody(options map[string]string) io.Reader {
	urlstr := string("")

	keys := make([]string, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if key == "sign" {
			continue
		}

		urlstr += key
		urlstr += "="
		urlstr += url.QueryEscape(options[key])
		urlstr += "&"

	}

	urlstr += "app_key="
	urlstr += url.QueryEscape(this.Appkey)
	sign := this.getMd5(urlstr)
	urlstr += "&sign="
	urlstr += url.QueryEscape(sign)
	return strings.NewReader(urlstr)

}

func (this *AiClient) Post(url string, body io.Reader) (*http.Response, error) {
	reqest, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return this.Client.Do(reqest)

}
