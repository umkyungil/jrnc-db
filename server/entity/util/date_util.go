package util

import (
	"regexp"
	"time"
)

// 日付整合性チェック（未来日含む）
func CheckDate(dateStr string) error {
	// パラメータが空白の場合はスキップ
	if dateStr == "" {
		return nil
	}

	// 削除する文字列を定義
	reg := regexp.MustCompile(`[-|/|:| |　]`)

	// 指定文字を削除
	str := reg.ReplaceAllString(dateStr, "")

	// 数値の値に対してフォーマットを定義
	format := string([]rune(FORMAT_DATE)[:len(str)])

	// パース処理 → 日付ではない場合はエラー
	_, error := time.Parse(format, str)
	return error
}
