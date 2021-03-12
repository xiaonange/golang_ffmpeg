package test

import (
	"Go-FFMpeg"
	"Go-FFMpeg/Media"
	"testing"
)

func TestGetInfo(t *testing.T) {
	Config := Go_FFMpeg.Config{}
	new(media.Video).Create(Config).Info("/mnt/d/video/oceans.mp4")
}