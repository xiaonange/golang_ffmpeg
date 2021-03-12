package test

import (
	 "github.com/xiaonange/golang_ffmpeg"
	"testing"
)

func TestGetInfo(t *testing.T) {
	Config := golang_ffmpeg.Config{}
	new(golang_ffmpeg.Video).Create(Config).Info("/mnt/d/video/oceans.mp4")
}