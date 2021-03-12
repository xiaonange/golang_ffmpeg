package golang_ffmpeg
const (
 TYPE_STREAMS = "streams"
 TYPE_FORMAT = "format"

)

type Config struct {
	TimeOut int    `json:"time_out"`
}

//获取文件信息参数
func GetInfoParams()[]string  {
	return []string{"-v","quiet","-print_format","json","-show_format"}
}

//获取合并文件信息参数
func GetConcatParams()[]string  {
	return []string{"-v","quiet","-print_format","json","-show_format"}
}