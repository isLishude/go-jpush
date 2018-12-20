package jpush

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/json-iterator/go"
)

// Endpoint
var (
	APIEndpoint        = "https://api.jpush.cn/v3/push" // default api endpoint
	BeijingAPIEndpoint = "https://bjapi.push.jiguang.cn/v3/push"
)

// Endpoint
var (
	AppKey        string // username for basic auth
	MasterSecret  string // password for basic auth
	DefaultClient *JPush // client which created by ENV variables
)

// JPush is jpush service struct
type JPush struct {
	APIEndpoint string
	Username    string // AppKey
	Password    string // MasterSecret
}

// NewJPushClient create JPush instance
func NewJPushClient(api, appKey, masterSecret string) *JPush {
	return &JPush{APIEndpoint: api, Username: appKey, Password: masterSecret}
}

// Send is push message to user related
func (j *JPush) Send(data []byte) error {
	// create a new request
	req, err := http.NewRequest(http.MethodPost, j.APIEndpoint, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("[JPush] Create a resquest Error:%v", err)
	}

	// Set header and basic auth
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(j.Username, j.Password)

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("[JPush] Resquest Error:%v", err)
	}
	defer resp.Body.Close()

	// Check error within response data
	// if http status code is not 200 then returns error
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 300) {
		var res Result
		if err := jsoniter.NewDecoder(resp.Body).Decode(&res); err != nil {
			return fmt.Errorf("[JPush] Decode response data error: %v", err)
		}
		if err := res.Error; err != nil {
			return fmt.Errorf("[JPush] Response Error: Code %v Msg %v", err["code"], err["message"])
		}
	}
	return nil
}
