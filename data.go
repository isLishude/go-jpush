package jpush

// DefaultOption is default push option
var DefaultOption = &Option{
	TTL:       86400 * 10,
	IsIOSProd: false,
	DeferTime: 0,
}

// Data is push message includes meta info
// `OS` field as platform. supports `android` `ios` `winphone`
// use []string type to select OS channel. use string "all" to push all os channel
// `Aud` field as `audience`.
// use string `all` to push all user within app. use `Audience` type to define by self
//
type Data struct {
	OS    interface{}   `json:"platform"`
	Aud   interface{}   `json:"audience"` // Post to all user then be “all”
	Msg   string        `json:"message,omitempty"`
	Opt   *Option       `json:"options,omitempty"`
	Notif *Notification `json:"notification"`
	SMS   SMS           `json:"sms_message,omitempty"`
	CID   string        `json:"cid,omitempty"`
}

// Audience is object for `audience` field within `Data` above
type Audience struct {
	Alias  []string `json:"alias"`
	Tag    []string `json:"tag,omitempty"`
	TagAnd []string `json:"tag_and,omitempty"`
	TagOr  []string `json:"tag_or,omitempty"`
	TagNot []string `json:"tag_not,omitempty"`
}

// Option is object for `option` field within `Data` above
type Option struct {
	SendNo    int  `json:"sendno,omitempty"`
	TTL       int  `json:"time_to_live,omitempty"`      // msg pushed life perrid
	IsIOSProd bool `json:"apns_production,omitempty"`   // production or not for IOS
	DeferTime int  `json:"big_push_duration,omitempty"` // defer time of pushing to user endpoint
}

// Notification is `notif` field for `jpush.Data`
type Notification struct {
	Alert   string   `json:"alert,omitempty"`
	Android *Android `json:"android"`
	IOS     *IOS     `json:"ios"`
}

// Android is object for `Android` field within `Notification` above
type Android struct {
	Platform
	BuildID    int         `json:"builder_id,omitempty"`
	Priority   int         `json:"priority,omitempty"`
	Style      int         `json:"style,omitempty"`
	AlertType  int         `json:"alert_type,omitempty"`
	BigText    string      `json:"big_text,omitempty"`
	Inbox      interface{} `json:"inbox,omitempty"`
	BigPicPath string      `json:"big_pic_path,omitempty"`
	LargeIcon  string      `json:"large_icon,omitempty"`
	Intent     interface{} `json:"intent,omitempty"`
}

// IOS is object for `IOS` field within `Notification` above
type IOS struct {
	Platform
	Sound            interface{} `json:"sound,omitempty"`
	Badge            int         `json:"badge,omitempty"`
	ContentAvailable bool        `json:"content-available,omitempty"`
	MutableContent   bool        `json:"mutable-content,omitempty"`
	ThreadID         string      `json:"thread-id,omitempty"`
}

// Platform is object for `Android` or `IOS` field within `Notification` above
// english doc: https://docs.jiguang.cn/en/jpush/server/push/rest_api_v3_push/
// chinese doc: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/
type Platform struct {
	Alert    string      `json:"alert,omitempty"`  // Short message under `title`
	Title    string      `json:"title,omitempty"`  // Message title
	Extra    interface{} `json:"extras,omitempty"` // Extra data
	Category string      `json:"category,omitempty"`
}

// Message is message within app
type Message struct {
	Content     string      `json:"msg_content"`
	Title       string      `json:"title,omitempty"`
	ContentType string      `json:"content_type,omitempty"`
	Extra       interface{} `json:"extras,omitempty"`
}

// SMS is send sms
type SMS struct {
	Delay    int         `json:"delay_time"`
	SignID   int         `json:"signid,omitempty"`
	TempID   int         `json:"temp_id,omitempty"`
	TempPara interface{} `json:"temp_data,omitempty"`
}
