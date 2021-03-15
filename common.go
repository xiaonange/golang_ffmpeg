package golang_ffmpeg

import "os"

//判断文件是否存在
func FileIsExist(path string)(bool,error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return false, nil
	}
	return false, err
}

