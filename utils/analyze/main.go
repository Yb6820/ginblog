package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	//获取当前项目的工作路径
	dirpath, err := os.Getwd()
	if err != nil {
		log.Println("Get file Pwd error!", err)
	}
	fmt.Println(dirpath)
	AnalyzeDBLog(dirpath + "\\dblog\\log20230612.log")
}

// 分析数据库日志
// 输入参数：日志文件名
// 输出参数：输出为excel或则发送统计信息到邮箱
func AnalyzeDBLog(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Println("AnalyzeDBlog Open File error!\t", err)
		return
	}
	InfoCount := 0
	buf := bufio.NewScanner(file)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		re, err := regexp.Compile(`level=info`)
		if err != nil {
			log.Println("InfoReg Compile error!", err)
		}
		//如果时info类的日志只进行简单的技术操作
		if re.MatchString(line) {
			InfoCount++
			//
			re, err = regexp.Compile(`time="(.*?)"`)
			if err != nil {
				log.Println("WarnReg Compile error!", err)
			}
			str := re.FindStringIndex(line)
			fmt.Println(str)
		} else {
			re, err = regexp.Compile(`level=warn`)
			if err != nil {
				log.Println("WarnReg Compile error!", err)
			}
			//把warn类和error类的信息分别进行解析存放
			if re.MatchString(line) {
				//获取出时间信息
				re, err = regexp.Compile(`time="(.*?)"`)
				if err != nil {
					log.Println("Warn Time Reg Compile error!", err)
				}
				timeInfo := re.FindString(line)
				//获取出执行的sql msg
				re, err = regexp.Compile(`msg="(.*?)"`)
				if err != nil {
					log.Println("Warn Msg Reg Compile error!", err)
				}
				msgInfo := re.MatchString(line)
				//记录warn日志
				fmt.Println(timeInfo, msgInfo)
			} else {
				fmt.Println("error")
			}
		}
	}
}

// 分析http日志
// 输入参数：日志文件名
// 输出参数：输出为excel或则发送统计信息到邮箱
func AnalyzeHttpLog(filename string) {

}
