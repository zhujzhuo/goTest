package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// SourceTableNameExp := "test_cold"
	// multiLineStr := fmt.Sprintf("这是第一行内容\n这是第二行内容，包含一些其他信息。\n最后一行内容。")
	// fmt.Println(multiLineStr)
	// SourceTableNameExpeee := fmt.Sprintf("%s_%d", SourceTableNameExp, 0)
	// fmt.Println(SourceTableNameExpeee)

	fmt.Println(generateString())

}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const scheduleidLength = 32

func generateString() (randomStr string) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, scheduleidLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	randomStr = string(b)
	return randomStr

}
