package jpush

import (
	"os"
	"strconv"
)

func init() {
	AppKey = os.Getenv("JPUSH_USERNAME")
	MasterSecret = os.Getenv("JPUSH_PASSWORD")
	APIEndpoint = os.Getenv("JPUSH_URL")
	if isDev, err := strconv.ParseBool(os.Getenv("JPUSH_MODE")); err == nil {
		DefaultOption.IsIOSProd = isDev
	}
	DefaultClient = NewJPushClient(APIEndpoint, AppKey, MasterSecret)
}
