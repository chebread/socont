// socont는 터미널에서 음향조절하는 프로그램임.

// TODO: https://github.com/itchyny/volume-go로 cmd 구현 공부해보자.

// 10단계로 나눠서 10 포인트씩 증감하기. -> 현재 사용자가 39다. 그러면 -9 해서 30으로 하자.
// socont: 1단계 증가 및 현재의 사운드 값 반환(단계적으로)
// 10단계 나누서 증가 socont -i
// '' 감소 socont -d

package main

import (
	"fmt"
	"os"
	"strconv"

	volume "github.com/itchyny/volume-go"
)

type Env struct {
	Name    string
	Version string
}

var ProgramEnv Env

func main() {
	ProgramEnv.Name = os.Args[0]
	ProgramEnv.Version = "development"

	curVol, err := volume.GetVolume() // Vol is Volumn
	if err != nil {
		fmt.Printf("get volume failed: %+v", err)
		return
	}
	curVolStr := strconv.Itoa(curVol)

	// 보호 구문
	if len(os.Args) < 2 {
		help()
		return
	}

	flagArg := os.Args[1]

	// TODO: kong 으로 -i foo 하면 foo 단계 더 높게 설정하게 (즉, 지금 5단계면 -i 2 하면 7단계로, -i 1 하면 -i랑 똑같이) 하기.

	var newVol int
	switch flagArg {
	// increase
	case "--increase", "-i":
		if curVol < 10 {
			newVol = 10
		} else {
			firstDigit, err := strconv.Atoi(curVolStr[0:1])
			if err != nil {
				fmt.Printf("Error: Atoi\n")
				return
			}

			newVol = firstDigit*10 + 10

		}

		volume.SetVolume(newVol)
		fmt.Printf("Changed volume value: %d (%d level)\n", newVol, newVol/10)
	case "--decrease", "-d":
		if curVol < 10 {
			newVol = 0
		} else {
			firstDigit, err := strconv.Atoi(curVolStr[0:1])
			if err != nil {
				fmt.Printf("Error: Atoi\n")
				return
			}

			newVol = firstDigit*10 - 10
		}

		volume.SetVolume(newVol)
		fmt.Printf("Changed volume value: %d (%d level)\n", newVol, newVol/10)
	case "--info", "-I":
		var level int
		if curVol < 10 {
			level = 0
		} else {
			firstDigit, err := strconv.Atoi(curVolStr[0:1])
			if err != nil {
				fmt.Printf("Error: Atoi\n")
				return
			}
			level = firstDigit
		}
		fmt.Printf("Current volume value: %d (%d level)\n", curVol, level)
	case "--help", "-h":
		help()
	case "--version", "-V":
		showVersion()
	default:
		fmt.Printf("Error: '%s' is an invalid command\n", flagArg)
	}
}

func help() {
	fmt.Println("help")
}

func showVersion() {
	fmt.Printf("%s %s\n", ProgramEnv.Name, ProgramEnv.Version)
}
