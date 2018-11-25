package main

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/piex/gowechat"
)

func main() {

	tuling := Tuling{
		URL: "http://openapi.tuling123.com/openapi/api/v2",
		Keys: map[string]Rebot{
			"Mervyn": Rebot{
				Name: "asmr",
				Key:  "dc54d5efed42466f9f0f57def6e7dead",
			},
		},
	}

	wx := gowechat.Start()

	for msg := range wx.MsgChan {
		for _, m := range msg.AddMsgList {

			m.Content = strings.Replace(m.Content, "&lt;", "<", -1)
			m.Content = strings.Replace(m.Content, "&gt;", ">", -1)

			switch m.MsgType {
			case 1:
				if m.FromUserName[:2] == "@@" { // 群消息

					content := strings.Split(m.Content, ":<br/>")[1]
					fmt.Println(content)
					if (wx.User.NickName != "" && strings.Contains(content, "@"+wx.User.NickName)) || (wx.User.RemarkName != "" && strings.Contains(content, "@"+wx.User.RemarkName)) {
						content = strings.Replace(content, "@"+wx.User.NickName, "", -1)
						content = strings.Replace(content, "@"+wx.User.RemarkName, "", -1)
						fmt.Printf("[群消息] %s", content)
						reply, _ := tuling.getReply(content, m.FromUserName)
						wx.SendMessage(reply, m.FromUserName)
					}
				} else {
					// 普通消息
					fmt.Printf("[好友消息] %s", m.Content)
					reply, _ := tuling.getReply(m.Content, m.FromUserName)
					wx.SendMessage(reply, m.FromUserName)
				}
			}
		}

		glog.Flush()
	}
}
