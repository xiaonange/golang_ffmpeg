package golang_ffmpeg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func  Run(commandName string, params []string) (bool, error, *os.Process) {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false, err, cmd.Process
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if (strings.Contains(line, "100%: Done")){
			return true, nil, cmd.Process
		}
		if err2 != nil || io.EOF == err2{
			cmd.Wait()
			return false, err2, cmd.Process
		}
		fmt.Println(line)
	}
	//cmd.Wait()
	return true, nil, cmd.Process
}
