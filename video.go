package ffmpeg_go

import (
	"errors"
)

type Video struct {
	Config  Config
	Params []string
}

func  (v *Video)Create(Config  Config  ) *Video {
	v.Config = Config
	return v
}

//设置参数
func (v *Video) SetParams( Params []string) *Video {
    v.Params = Params
	return v
}


//打开多个
func (v *Video) OpenMulti() *Video {

	return v
}

//打开单个
func (v *Video) Open() *Video {

	return v
}

//获取video信息
func (v *Video)Info(VideoPath string) {
	if VideoPath == ""{
		panic("VideoPath can't null")
	}
	if len(v.Params) < 1 {
		v.Params = GetInfoParams()
	}
	v.Params = append(v.Params,VideoPath)
	_,_,_ = Run("/usr/bin/ffmpeg",v.Params)
}

//合并video
func (v *Video)Concat(videos...string) error {
	if len(videos) < 1 {
		return errors.New("Error: video can't be null")
	}
    return nil
}