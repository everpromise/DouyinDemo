package utils

import (
	"bufio"
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// 将字符串写入文件
func WriteResult(data string, outfile string) error {

	file, err := os.OpenFile(outfile, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("writer", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(data)
	writer.Flush()
	return err
}

// 读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func ReadFollowers(userFile string) string {
	f, err := os.OpenFile(userFile, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}
