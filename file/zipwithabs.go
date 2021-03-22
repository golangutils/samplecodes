package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//ZipFilescompressesoneormanyfilesintoasingleziparchivefile.
//压缩多个文件到一个文件里面
//Param1:输出的zip文件的名字
//Param2:需要添加到zip文件里面的文件
//Param3:由于file是绝对路径，打包后可能不是想要的目录，oldform就是filename中需要被替换的掉的路径
//Param4:要替换成的路径
func ZipFiles(filename string, files []string, oldform, newform string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	//把files添加到zip中
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		//获取file的基础信息
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//使用上面的FileInforHeader()就可以把文件保存的路径替换成我们自己想要的了，如下面
		fmt.Println(file, oldform, newform)
		header.Name = strings.Replace(file, oldform, newform, -1)

		//优化压缩
		//更多参考seehttp://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fileList := []string{
		"aop/List.txt",
		"aop/02code/json.txt",
	}
	//比如线上的程序的目录为game,"/Users/Jack/Downloads/QQreceiver/72bian_log/"下面的文件为我们线game下的程序文件内容
	cur, _ := os.Getwd()
	fmt.Println(cur)

	for i, f := range fileList {
		fileList[i] = filepath.Join(cur, f)
	}

	err := ZipFiles("test.zip", fileList, filepath.Join(cur, "aop")+string(filepath.Separator), "")
	if err != nil {
		fmt.Println(err)
	}
}
