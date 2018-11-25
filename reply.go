package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

var client = &http.Client{}

func (t *Tuling) getReply(msg string, uid string) (string, error) {
	params := make(map[string]interface{})
	params["reqType"] = 0
	params["userInfo"] = UserInfo{
		APIKey: "xxx",
		UserID: uid[2:32],
	}
	params["perception"] = Perception{InputText{Text: msg}}

	data, err := json.Marshal(params)
	if err != nil {
		glog.Error(err)
	}

	glog.Info("[$] tuling post ...", string(data))

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", t.URL, body)

	resp, _ := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)

	var reply Reply

	if err := json.Unmarshal(r, &reply); err != nil {
		return "", err
	}

	glog.Info("[$] Tuling reply ...", reply)

	if reply.Intent.Code > 10000 {
		return reply.Results[0].Values.Text, nil
	}

	// switch reply.Intent.Code {
	// case 10008:
	// case 10004:
	// 	return reply.Results[0].Values.Text, nil
	// default:
	// 	return "不知道你在说啥～", nil
	// }

	return "不知道你在说啥～", nil

}
