package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func loadenv() {
	// .envファイルの読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic("could not read .env file")
	}
}

func get_weekday() string {
	weekday := []string{"日", "月", "火", "水", "木", "金", "土"}
	day := time.Now().Weekday()
	return weekday[day]
}

func main() {
	loadenv()

	// Webhook URL
	url := os.Getenv("WEBHOOK_URL")

	// Slackメッセージ
	message := map[string]string{
		"text": get_weekday(),
	}

	// jsonエンコードする
	jsonValue, _ := json.Marshal(message)

	// HTTPリクエストを作成する
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")

	// リクエストを送る
	client := &http.Client{}
	res, _ := client.Do(req)

	// レスポンスを出力する
	fmt.Println(res.Status)
}
