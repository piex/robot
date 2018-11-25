package main

import (
	"testing"
)

func TestGetReply(t *testing.T) {

	tuling := Tuling{
		URL: "http://openapi.tuling123.com/openapi/api/v2",
		Keys: map[string]Rebot{
			"Mervyn": Rebot{
				Name: "asmr",
				Key:  "dc54d5efed42466f9f0f57def6e7dead",
			},
		},
	}

	reply, err := tuling.getReply("几点了", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}
