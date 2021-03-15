#视频支持
* 获取视频信息
* 视频增加图片水印
* 视频增加文字水印
* 获取帧图片
* 视频合并
* 视频格式转换
* 视频转化为gif图片
```
	ffprobe := "/usr/bin/ffprobe"
	ffmpeg := "/usr/bin/ffmpeg"
	Config := golang_ffmpeg.Config{TimeOut:10}
	//获取音视频信息
	new(golang_ffmpeg.Video).Create(Config).Info(ffprobe,"/mnt/d/video/oceans.mp4")
	//给视频添加图片水印
	err,result :=new(golang_ffmpeg.Video).Create(Config).AddImg(ffmpeg,"/mnt/d/video/oceans.mp4","/mnt/d/video/test.mp4","/mnt/d/video/babymarkt-logo.png")
	//给视频添加文本水印
	err,result :=new(golang_ffmpeg.Video).Create(Config).AddText(ffmpeg,"/mnt/d/video/oceans.mp4","/mnt/d/video/test1.mp4","jdjjdjj")
	//获取视频帧图片
	err,result :=new(golang_ffmpeg.Video).Create(Config).GetVideoCover(ffmpeg,"/mnt/d/video/oceans.mp4","/mnt/d/video/test1.jpg","00:00:20")
	//视频合并
	err,result :=new(golang_ffmpeg.Video).Create(Config).Concat(ffmpeg,[]string{"/mnt/d/video/oceans.mp4","/mnt/d/video/big_buck_bunny.mp4"},"/mnt/d/video/con.mp4")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
```
#音频支持
* 获取音频信息
* 音频合并
* 音频格式转换


#图像支持

