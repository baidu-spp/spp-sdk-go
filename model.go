package push

type PushResponse struct {
	RequestID uint32      `json:"request_id,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Result    interface{} `json:"result,omitempty"`
}

type Message struct {
	MessageType    int           `json:"message_type"`
	Notification   *Notification `json:"notification,omitempty"`
	Transmission   *Transmission `json:"transmission,omitempty"`
	Aps            *Aps          `json:"aps,omitempty"`
	Condition      []*Condition  `json:"condition,omitempty"`
	Option         *Option       `json:"option,omitempty"`
	MsgId          string        `json:"msg_id,omitempty"`
	PushTime       int64         `json:"push_time,omitempty"`
	CallBackParams string        `json:"cbparams,omitempty"`
}

type UnicastMessage struct {
	*Message
	PushID string `json:"push_id"`
}

type MuticastMessage struct {
	*Message
	PushIDs []string ` json:"push_ids"`
}

type CuidsMessage struct {
	*Message
	Cuids []string `json:"cuids"`
}

type Notification struct {
	ID      uint32  `json:"id,omitempty"`
	Notify  *Notify `json:"notify,omitempty"`
	Title   string  `json:"title,omitempty"`
	Content string  `json:"content,omitempty"`
	Style   int     `json:"style,omitempty"`
	Image   string  `json:"image,omitempty"`
	Icon    string  `json:"icon,omitempty"`
	Action  *Action `json:"action,omitempty"`
}

type Notify struct {
	Vibrate int `json:"vibrate,omitempty"`
	Sound   int `json:"sound,omitempty"`
	Lights  int `json:"lights,omitempty"`
}

type Action struct {
	ActionType int               `json:"action_type,omitempty"`
	URL        string            `json:"url,omitempty"`
	ClassName  string            `json:"class_name,omitempty"`
	Param      map[string]string `json:"params,omitempty"`
}

type Transmission struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Condition struct {
	Key     string   `json:"key,omitempty"`
	Values  []string `json:"values,omitempty"`
	Operate string   `json:"operate,omitempty"`
}

type Option struct {
	Expire uint32 `json:"expire,omitempty"`
	Speed  uint32 `json:"speed,omitempty"`
	Limit  uint32 `json:"limit,omitempty"`
}

type Aps struct {
	Alert            Alert                  `json:"alert,omitempty"`
	Badge            int32                  `json:"badge,omitempty"`
	Sound            string                 `json:"sound,omitempty"`
	ContentAvailable int32                  `json:"content-available,omitempty"`
	MutableContent   int32                  `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	ApnsCollapseId   string                 `json:"apns-collapse-id,omitempty"`
	Env              int32                  `json:"_env"`
	BadgeConf        string                 `json:"_badge_conf"`
	NoTitle          string                 `json:"_no_title,omitempty"`
	TraceId          string                 `json:"id,omitempty"`
	Extras           map[string]interface{} `json:"extras"`
}

type Alert struct {
	Title           string   `json:"title,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	Body            string   `json:"body,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	ActionLocKey    string   `json:"action-loc-key,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	SubtitleLocKey  string   `json:"subtitle-loc-key,omitempty"`
	SubtitleLocArgs []string `json:"subtitle-loc-args,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
}
