package golang_ffmpeg

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
func (v *Video)Info(CommandName,VideoPath string)(error,string) {
	if VideoPath == ""{
		return errors.New("输入视频不能为空"),""
	}
	if len(v.Params) < 1 {
		v.Params = GetInfoParams(VideoPath)
	}
	return Run(CommandName,v.Params)
}

//合并video
func (v *Video)Concat(CommandName string,videos[]string, SavePath string) (error,string)  {
	if len(videos) < 2 {
		return errors.New("合并视频数量必须大于2"),""
	}
	if len(v.Params) < 1 {
		v.Params = GetConcatParams(videos,SavePath)
	}
	return Run(CommandName,v.Params)
}

//视频图片水印
func (v *Video)AddImg(CommandName,InVideoPath,OutVideoPath,Img string)(error,string) {
	if InVideoPath == "" || OutVideoPath == "" || Img == "" {
		return errors.New("请校正参数"),""
	}
	if len(v.Params) < 1 {
		v.Params = GetVideoAddImgParams(InVideoPath, OutVideoPath, Img)
	}
	return Run(CommandName, v.Params)
}

//视频文本水印
func (v *Video)AddText(CommandName,InVideoPath,OutVideoPath,Text string)(error,string) {
	if InVideoPath == "" || OutVideoPath == "" || Text == "" {
		return errors.New("请校正参数"),""
	}
	if len(v.Params) < 1 {
		v.Params = GetVideoAddTextParams(InVideoPath, OutVideoPath, Text)
	}
	return Run(CommandName, v.Params)
}

//视频
func (v *Video)GetVideoCover(CommandName,InVideoPath,OutVideoPath ,StartTime string)(error,string) {
	if InVideoPath == "" || OutVideoPath == "" || StartTime == "" {
		return errors.New("请校正参数"),""
	}
	if len(v.Params) < 1 {
		v.Params = GetVideocoverParams(InVideoPath, OutVideoPath, StartTime)
	}
	return Run(CommandName, v.Params)
}