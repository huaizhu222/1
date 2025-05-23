package response

import (
	"encoding/json"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j).Format("2005-05-01"))
	// var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2005-05-01"))
	// return []byte(stmp), nil
}

type UerResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
	Gender   string   `json:"gender"`
	Birthday JsonTime `json:"birthday"`
	Mobile   string   `json:"mobile"`
}
