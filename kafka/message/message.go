package message

import(
	"encoding/json"
	"github.com/sukstar76/model"
)

var Message []byte

type LogMessage struct {
	User model.User `json:"user" form:"user" query:"user"`
	Msg string  `json:"msg" form: "msg" query: "msg"`
}

func SetMessage(lm LogMessage) {
	msg, _ := json.Marshal(lm)
	Message = msg
}


