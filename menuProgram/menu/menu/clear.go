package menu

import (
	"os"
	"os/exec" //외부 프로그램 실행을 위한 exec 모듈
	"runtime"
)

func ClearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout //cmd.Stdout 속성에 os.Stdout을 할당하여 명령어 실행 결과를 현재 터미널 창에 출력
	cmd.Run() 				//cmd.Run()을 호출하여 설정한 명령어를 실행
							//cmd.Output() 메서드를 호출하여 실행 결과 확인 가능
}

