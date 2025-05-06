package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// 加载已有的测试用例
func loadExistedCase(fileName string) []string {
	// 文件存在，读取文件内容
	_, err := os.Stat(fileName)
	if err != nil {
		return nil
	}
	// 读取文件内容
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("读取文件失败:%v", err)
		return nil
	}
	list := strings.Split(string(content), "\n")
	items := skipEmpty(list)
	log.Printf("加载已有的测试用例: %d 条\n", len(items))
	return items
}

// 可选：过滤空行
func skipEmpty(list []string) []string {
	items := make([]string, 0, len(list))
	for _, item := range list {
		trim := strings.TrimSpace(item)
		if trim != "" {
			items = append(items, trim)
		}
	}
	return items
}

// SeqGenerator 序列生成器
type SeqGenerator interface {
	Gen(n int) []string
}

// 生成训练数据(4个字符组合) 243个用例
type groundTruthSeq struct{}

// Gen implements SeqGenerator
func (groundTruthSeq) Gen(randCnt int) []string {
	// 易混淆的字符: 0O、5S (2^4=16）
	var tmp = make([]string, 0, 100)
	dfs("", 4, "0O", &tmp)
	dfs("", 4, "5S", &tmp)
	if len(tmp) != 32 {
		panic("len(tmp) should be 16")
	}
	myFont := myFontCharset()
	if len(myFont) != 9 {
		panic("len(MyFont) should be 9")
	}
	rnd := randomSeq{}.Gen(randCnt)
	if len(rnd) != randCnt {
		panic(fmt.Errorf("len(rnd) should be %d", randCnt))
	}
	all := []string{"FPH0", "O0JC"}
	all = append(all, tmp...)
	all = append(all, myFont...)
	all = append(all, rnd...)
	log.Printf("生成新训练数据: %d 条\n", len(all))
	return all
}

// 模式1：生成长度为 N 的字符串
func dfs(prefix string, n int, dict string, results *[]string) {
	if n <= 0 {
		*results = append(*results, prefix)
		return
	}
	for _, c := range dict {
		dfs(prefix+string(c), n-1, dict, results)
	}
}

// 36个字符
const allowChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 模式2：全部雅黑的字符
func myFontCharset() []string {
	var results []string
	end := len(allowChars) - 4
	for i := 0; i <= end; i += 4 {
		results = append(results, allowChars[i:i+4])
	}
	return results
}

// 模式3：随机N个用例
type randomSeq struct{}

func (randomSeq) Gen(n int) []string {
	var results []string
	const total = len(allowChars)
	for i := 0; i < n; i++ {
		var one string
		for k := 0; k < 4; k++ {
			c := allowChars[rand.Intn(total)]
			one += string(c)
		}
		results = append(results, one)
	}
	return results
}
