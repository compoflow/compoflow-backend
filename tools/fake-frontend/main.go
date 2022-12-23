package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
	agent  = gorequest.New()
)

type FakeData struct {
	Xml string `json:"xml"`
}

func main() {
	data, err := ioutil.ReadFile("../../resource/xmlStr.xml")
	if err != nil {
		logger.Errorln("Failed to read xml:", err)
		return
	}

	fakeData := FakeData{
		Xml: string(data),
	}

	msg, err := json.Marshal(fakeData)
	if err != nil {
		logger.Errorln("Failed to parse xml to json:", err)
		return
	}

	_, body, errs := agent.Post("http://localhost:8080/deploy").Send(string(msg)).End()
	if len(errs) > 0 {
		logger.Errorln("http errors:", errs)
		return
	}
	logger.Infoln(body)
}
