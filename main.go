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
	// 今日の曜日の取得
	weekday := []string{"日", "月", "火", "水", "木", "金", "土"}
	day := time.Now().Weekday()
	return weekday[day]
}

func get_message() string {
	today := get_weekday()
	var msg string

	switch today {
	case "日":
		msg = "なし"
	case "月":
		msg = "資源プラスチック・ダンボール・剪定枝・落ち葉・下草 or 資源プラスチック"
	case "火":
		msg = "もやせるごみ or もやせるごみ・ビン・カン"
	case "水":
		msg = "ペットボトル"
	case "木":
		msg = "衣類・布類 or もやせないごみ・新聞紙・牛乳パック類"
	case "金":
		msg = "もやせるごみ"
	case "土":
		msg = "なし"
	}

	msg = today + "曜日は、" + msg + "の回収日です。"
	return msg
}

func main() {
	loadenv()

	// Webhook URL
	url := os.Getenv("WEBHOOK_URL")

	// Slackメッセージ
	message := map[string]string{
		"text": get_message(),
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
