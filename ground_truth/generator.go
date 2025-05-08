package groundtruth

import (
	"fmt"
	"log"
	"math/rand"
)

// IGenerator 序列生成器
type IGenerator interface {
	Gen(n int) []string
}

var (
	GenAmbiguous IGenerator = ambiguousGen{}
	GenRandom    IGenerator = randomGen{}
)

// 生成训练数据(4个字符组合) 259个用例
type ambiguousGen struct{}

// Gen implements SeqGenerator
func (ambiguousGen) Gen(randCnt int) []string {
	// 全量字符集
	myFont := myFontCharset()
	if len(myFont) != 9 {
		panic("len(MyFont) should be 9")
	}
	// 易混淆的字符: 0O、5S (2^4=16）
	var tmp = make([]string, 0, 100)
	dfs("", 4, "0O", &tmp)
	dfs("", 4, "5S", &tmp)
	dfs("", 4, "1I", &tmp)
	if len(tmp) != 48 {
		panic("len(tmp) should be 48")
	}
	rnd := randomGen{}.Gen(randCnt)
	if len(rnd) != randCnt {
		panic(fmt.Errorf("len(rnd) should be %d", randCnt))
	}
	all := append(myFont, tmp...)
	all = append(all, rnd...)
	all = append(all, "FPH0", "O0JC")
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
type randomGen struct{}

func (randomGen) Gen(n int) []string {
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
