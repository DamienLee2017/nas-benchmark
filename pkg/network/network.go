package network

import (
	"fmt"
	"net/http"
	"time"
)

func Test() (int, error) {
	fmt.Println("Network Test")
	start := time.Now()
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	fmt.Printf("Network Test took %v\n", duration)

	// 评分逻辑需要根据实际情况设计
	score := int(duration.Milliseconds())
	score = 1000 - score
	return score, nil
}
