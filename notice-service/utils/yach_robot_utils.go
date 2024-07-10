package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RobotSendTextMessage(body string) (err error) {
	//数据参数
	data := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": body,
		},
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reqUrl := "https://yach-oapi.zhiyinlou.com/robot/send?access_token=cWc1T2k5cDJJdDVlMUlNYi9lVnpqZ2pQN0JYb0ZDMk1QMFZJcEc4Y2pvcTNRRWNqanhreUpJY05zQkNDSlhlbg&timestamp=" + Int64ToString(SystemCurrentMills())
	post, err := http.Post(reqUrl, "application/json;charset=utf-8", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	all, err := io.ReadAll(post.Body)
	if err != nil {
		return err
	}
	log.Println("RobotSendTextMessage res:", string(all))
	return nil
}
