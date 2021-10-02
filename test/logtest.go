package main

import (
	"go-2021video/ploglib"
	"time"
)

func main() {
	logger := ploglib.NewLog("info")
	for {
		id := 507
		name := "hushengs"
		logger.Debug("这是一条Debug日志", id, name)
		logger.Trace("这是一条Trace日志", 2)
		logger.Info("这是一条Info日志", 3)
		logger.Warning("这是一条Warning日志", 4)
		logger.Error("这是一条Error日志", 5)
		logger.Fatal("这是一条Fatal日志", 6)
		time.Sleep(2 * time.Second)
	}
}
