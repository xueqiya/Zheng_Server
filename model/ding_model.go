package model

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func SendSms(content string) (string, error) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=5708dcd5a145f3c13853799d76fd8b508d000b7c8c6a9c64f83a19f402a36761`
	json := `{"msgtype": "text","text": {"content": "【正事多通知】\n` + content + `"}}`

	log.Println("请求钉钉json：" + json)
	req, err := http.NewRequest(http.MethodPost, webHook, strings.NewReader(json))
	if err != nil {
		return "构建钉钉请求体失败", err
	}

	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "请求钉钉失败", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	body := string(respBody[:])
	log.Println("钉钉Body：" + body)
	return body, err
}
