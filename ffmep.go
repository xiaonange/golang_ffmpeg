package golang_ffmpeg
const (
 TYPE_STREAMS = "streams"
 TYPE_FORMAT = "format"

)

type Config struct {
	TimeOut int    `json:"time_out"`
}

//获取文件信息参数
func GetInfoParams(videoPath string)[]string  {
	return []string{"-v","quiet","-print_format","json","-show_format",videoPath}
}

//获取合并文件参数
func GetConcatParams(videos[]string,output string)[]string  {
	params :=[]string{}
	params = append(params,)
	for _,v:=range videos{
		params = append(params,"-i")
		params = append(params,v)
	}
	params = append(params,"-c:s")
	params = append(params,"copy")
	params = append(params,output)
	return params
}

//给视频增加文字
func GetVideoAddTextParams(input,output,text string)[]string {
	return []string{"-i", input, "-vf", "drawtext=fontcolor=white:fontsize=40:x=50:y=50:text="+text, output}
}
//给视频增加图片
func GetVideoAddImgParams(input,output,img string)[]string {
	return []string{"-i", input, "-vf", `"movie=`+ img +`[wm];[in][wm]overlay=10:10[out]"`, output}
}
//给视频增加图片
func GetVideocoverParams(input,output ,startTime string)[]string {
	return []string{"-i", input, "-ss", startTime, "-vframes", "1", output}
}