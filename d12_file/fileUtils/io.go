package fileUtils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// ReadAllByRead 读方式1
func ReadAllByRead(fileName string) (content string, err error) {
	var (
		buff = make([]byte, 128)
	)

	// 1. open file && defer close
	pFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer func() { _ = pFile.Close() }()

	// 2. for { read bytes buff }
	for {
		_, err = pFile.Read(buff)

		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			content = ""
			break
		} else {
			content += string(buff[:127])
		}
	}

	return
}

// ReadAllByBufIO 读方式2
func ReadAllByBufIO(fileName string) (content string, err error) {
	var (
		line string
	)

	// 1. open file && defer close
	pFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer func() { _ = pFile.Close() }()

	// 2. init Reader
	pReader := bufio.NewReader(pFile)

	// 3. for { read line }
	for {
		line, err = pReader.ReadString('\n')
		if err == io.EOF {
			err = nil
			content += line
			break
		} else if err != nil {
			content = ""
			break
		} else {
			content += line
		}
	}

	return
}

// ReadAllByIOUtil 读方式3
func ReadAllByIOUtil(fileName string) (content string, err error) {
	// 1. read all file
	buff, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	content = string(buff)

	return
}

// WriteByWrite 写方式1 - 不追加1
func WriteByWrite(fileName string, content string) (err error) {
	// 1. open file && defer close
	pFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer func() { _ = pFile.Close() }()

	// write string
	_, err = pFile.WriteString(content) // 字节流则用：pFile.Write([]byte)

	return
}

// WriteByWrite 写方式1 - 追加1
func WriteAppendByWrite(fileName string, content string) (err error) {
	// 1. open file && defer close
	pFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer func() { _ = pFile.Close() }()

	// write string
	_, err = pFile.WriteString(content) // 字节流则用：pFile.Write([]byte)

	return
}

// WriteByBufIO 写方式2 - 追加
func WriteByBufIO(fileName string, content string) (err error) {
	// 1. open file && defer close
	pFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer func() { _ = pFile.Close() }()

	// 2. init Writer
	pWriter := bufio.NewWriter(pFile)

	// 3. write to cache
	_, err = pWriter.WriteString(content)
	if err != nil {
		return
	}

	// 4. write to file
	err = pWriter.Flush()

	// 3. write to cache
	_, err = pWriter.WriteString(content)
	if err != nil {
		return
	}

	// 4. write to file
	err = pWriter.Flush()

	return
}

// WriteByIOUtil 写方式3 - 不追加
func WriteByIOUtil(fileName string, content string) (err error) {
	err = ioutil.WriteFile(fileName, []byte(content), 0666)
	return
}

func test() {
	os.Stdin
}
