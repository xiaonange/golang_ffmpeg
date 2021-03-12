package test

import (
	ffmpeg_go "github.com/xiaonange/go-ffmpeg"
	"testing"
)

func TestGetInfo(t *testing.T) {
	Config := ffmpeg_go.Config{}
	new(ffmpeg_go.Video).Create(Config).Info("/mnt/d/video/oceans.mp4")
}