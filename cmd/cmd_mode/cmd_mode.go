package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	addHour()
}

func addHour() {
	//go PrintRunTime()

	for {
		fmt.Println("请输入时间（格式：2006-01-02 15:04:05），输入 'exit' 退出：")

		// 从标准输入读取用户输入
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// 去掉输入中的换行符
		//input = input[:len(input)-1]

		if input == "exit" {
			break // 用户输入 'exit' 时退出循环
		}
		fmt.Println(input)
		// 解析输入时间字符串
		//inputTime, err := time.Parse("2006-01-02 15:04:05", inputTimeStr)
		//if err != nil {
		//	log.Println("无效的时间格式，请重新输入")
		//	continue // 继续等待用户输入
		//}

		// 添加一个小时
		//outputTime := inputTime.Add(time.Hour)

		// 输出结果
		//fmt.Println("输入时间:", inputTime.Format("2006-01-02 15:04:05"))
		//fmt.Println("一个小时后的时间:", outputTime.Format("2006-01-02 15:04:05"))
	}
}
