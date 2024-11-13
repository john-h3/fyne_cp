package main

import (
	cp "golang.design/x/clipboard"
	"log"
	"time"
)

func InitClipboard() {
	err := cp.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadText() []byte {
	return cp.Read(cp.FmtText)
}

func WriteText(content []byte) {
	cp.Write(cp.FmtText, content)
}

func WatchClipboard(fn func([]byte)) {
	tick := time.NewTicker(1000 * time.Millisecond)
	defer tick.Stop()
	for range tick.C {
		content := ReadText()
		if len(content) > 0 {
			fn(content)
		}
	}
}
