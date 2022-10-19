package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "regexp"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	readFileDeep("./百度脑图")
}

// 读取整个目录，进行替换
func readFileDeep(folderName string) {
	files, _ := ioutil.ReadDir(folderName)
	for _, f := range files {
		fileOrFolderName := fmt.Sprintf("%s/%s", folderName, f.Name())
		if f.IsDir() {
			readFileDeep(fileOrFolderName)
		} else {
			processButtomLink(fileOrFolderName)
		}
	}
}

func processButtomLink(fileName string) {
	// 底部的链接字符串
	var buttomLink string
	// 读取整个文件内容到 content 中
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("读取文件错误")
	}
	// 新的 content，进行替换
	var newContent string = string(content)

	// 查找底部链接
	re := regexp2.MustCompile(`(\[\d*\]:)(?<=\[\d*\]:).*`, 0)
	m, _ := re.FindStringMatch(string(content))
	for m != nil {
		// 逐行替换
		newContent = strings.Replace(newContent, m.String(), "", 1)
		buttomLink += m.String() + "\n"
		m, _ = re.FindNextMatch(m)
	}
	// 查找文章所有的 [num] 标记
	re2 := regexp2.MustCompile(`\[\d*\]`, 0)
	m1, _ := re2.FindStringMatch(newContent)
	for m1 != nil {
		// 根据 [1] 找到它对应的链接，将 [1] 替换为 (具体的url)
		regx := fmt.Sprintf("(?<=%s:).*", m1.String())
		regx = strings.Replace(regx, "[", "\\[", -1)
		regx = strings.Replace(regx, "]", "\\]", -1)
		re3 := regexp2.MustCompile(regx, 0)
		if m3, _ := re3.FindStringMatch(buttomLink); m3 != nil {
			// 进行替换
			replaceStr := fmt.Sprintf("(%s)", strings.Trim(m3.String(), " "))
			newContent = strings.Replace(newContent, m1.String(), replaceStr, 1)
		}
		m1, _ = re2.FindNextMatch(m1)
	}
	err = writeToFile(fileName, []byte(newContent))
	if err != nil {
		log.Fatal("文件写入失败")
	}
}

// 写入内容到文件中
func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
