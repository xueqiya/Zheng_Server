package model

import (
	"net/http"
	"strings"
)

func SendSms(content string) (bool, error) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=5708dcd5a145f3c13853799d76fd8b508d000b7c8c6a9c64f83a19f402a36761`
	json := `{"msgtype": "text",
		"text": {"content": "` + content + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(json))
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return false, err
	}
	return true, err
}
