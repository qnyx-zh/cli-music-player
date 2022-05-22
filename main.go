package main

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	audioFile, err := os.Open("/Users/apricity/Documents/codespace/cli-music-player/static/陈奕迅-孤勇者.mp3")
	ErrProcess(err)
	defer audioFile.Close()
	auditStream, format, err := mp3.Decode(audioFile)
	ErrProcess(err)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	ErrProcess(err)
	done := make(chan bool)
	speaker.Play(beep.Seq(auditStream, beep.Callback(func() {
		done <- true
	})))
	<-done
}
func ErrProcess(err error) {
	if err != nil {
		panic("出现异常")
	}
}
