package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func fileExist(fileName string) bool {
	s, err := os.Stat(fileName)
	if err != nil {
		return false
	} else if s.IsDir() {
		return false
	} else {
		return true
	}
}

func DirExist(fileName string) bool {
	s, err := os.Stat(fileName)
	if err != nil {
		return false
	} else if s.IsDir() {
		return true
	} else {
		return false
	}
}

func main() {
	var (
		cmdInput []string
		source   []string
		target   string
	)

	// 1. 输入过滤与预处理
	cmdInput = os.Args
	if len(cmdInput) < 2 {
		fmt.Printf("bjmCp: missing file operand\n")
		fmt.Println("Try 'bjmCp --help' for more information.")
		return
	} else if cmdInput[1] == "--help" {
		fmt.Println("Usage: bjmCp File... DIRECTORY")
		return
	} else if len(cmdInput) < 3 {
		fmt.Printf("bjmCp: missing destination file operand after '%v'\n", cmdInput[1])
		fmt.Println("Try 'bjmCp --help' for more information.")
		return
	}
	source = cmdInput[1 : len(cmdInput)-1]
	target = cmdInput[len(cmdInput)-1]

	// 2. 文件存在判断
	// 2.1 所有 source 需要是文件且存在
	for _, file := range source {
		if !fileExist(file) {
			fmt.Printf("bjmCp: cannot stat '%v': No such file.\n", file)
			return
		}
	}
	// 2.2 若 source 长度只有1，则 target 可以是存在的目录，或者是文件，否则必须是存在目录
	if len(source) == 1 {
		if os.IsPathSeparator(target[len(target)-1]) && !DirExist(target) {
			fmt.Printf("bjmCp: target '%v' is not a directory\n", target)
			return
		}
	} else {
		if !DirExist(target) {
			fmt.Printf("bjmCp: target '%v' is not a directory\n", target)
			return
		}
	}

	// 3. 复制操作
	if len(source) == 1 {
		if DirExist(target) {
			buff, err := ioutil.ReadFile(source[0])
			if err != nil {
				fmt.Printf("bjmCp: read file '%v' error\n", source[0])
				return
			}

			// write
			err = ioutil.WriteFile(path.Join(target, path.Base(source[0])), buff, 0666)
			if err != nil {
				fmt.Printf("bjmCp: write file '%v' error\n", source[0])
				return
			}
		} else {
			buff, err := ioutil.ReadFile(source[0])
			if err != nil {
				fmt.Printf("bjmCp: read file '%v' error\n", source[0])
				return
			}

			// write
			err = ioutil.WriteFile(target, buff, 0666)
			if err != nil {
				fmt.Printf("bjmCp: write file '%v' error\n", source[0])
				return
			}
		}
	} else {
		for _, file := range source {
			// read
			buff, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("bjmCp: read file '%v' error\n", file)
				break
			}

			// write
			err = ioutil.WriteFile(path.Join(target, path.Base(file)), buff, 0666)
			if err != nil {
				fmt.Printf("bjmCp: write file '%v' error\n", file)
				break
			}
		}
	}

	return
}
