package test

import (
	 "video/golang_ffmpeg"
	"testing"
)

func TestGetInfo(t *testing.T) {
	CommandPath := "/usr/bin/ffprobe"
	Config := golang_ffmpeg.Config{TimeOut:10}
	new(golang_ffmpeg.Video).Create(Config).Info(CommandPath,"/mnt/d/video/oceans.mp4")
}