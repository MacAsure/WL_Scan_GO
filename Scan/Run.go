package Scan

import (
	"bufio"
	"fmt"
	"neiwang/Common"
	"os"
	"strings"
	"sync"
	"time"
)

func CheckURL(cmd Common.Cmd) {
	if cmd.Target != "" {
		var scan Scan
		fmt.Printf("开始扫描...\n")
		startTime := time.Now()
		if cmd.Dns {
			attack_list(Common.Initurl(cmd.Target), &scan)
		} else {
			attack_list(Common.Init(cmd.Target), &scan)
		}
		endTime := time.Since(startTime)
		fmt.Printf("扫描结束,耗时: %v\n", endTime)
		fmt.Printf("\n存在漏洞的连接:\n")
		fmt.Println("---------------------------------------------")
		for _, i := range scan.result {
			fmt.Println(i)
		}
	}
	if cmd.Targets != "" {
		var scan Scan
		fmt.Printf("开始批量扫描...\n")
		startTime := time.Now()
		targets := ReadFiles(cmd.Targets)
		var wg sync.WaitGroup
		var taskChan = make(chan string, len(targets))
		for _, target := range targets {
			taskChan <- target
		}
		close(taskChan)

		for i := 0; i <= cmd.Thread; i++ {
			wg.Add(1)
			go func() {
				for {
					target, ok := <-taskChan
					if !ok {
						break
					}
					attack_list(Common.Initurl(target), &scan)
				}
				defer wg.Done()
			}()
		}
		wg.Wait()
		WriteFiles(cmd.Output, scan.result)
		endTime := time.Since(startTime)
		fmt.Printf("\n[*]扫描结束，共耗时: %v\n", endTime)
		fmt.Printf("\n存在漏洞的连接:\n")
		fmt.Println("---------------------------------------------")
		for _, i := range scan.result {
			fmt.Println(i)
		}
		fmt.Printf("\n输出结果保存至:%v", cmd.Output)
	}
}

// 读取文件
func ReadFiles(filename string) []string {
	var readcontent []string
	r, _ := os.Open(filename)
	defer r.Close()
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		readcontent = append(readcontent, line)
	}
	return readcontent
}

// 写入文件
func WriteFiles(filename string, data []string) {
	if !strings.HasSuffix(filename, ".txt") {
		filename = filename + ".txt"
	}
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	write := bufio.NewWriter(f)
	for _, r := range data {
		write.WriteString(r)
	}
	write.Flush()
}
