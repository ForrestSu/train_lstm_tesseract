package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

const fontYahei = "Microsoft YaHei"

func main() {
	var modeUsage = "[mode] 模式选择:\n" +
		"- gen: 生成训练数据\n" +
		"- pass: 识别率测试\n" +
		"- single: 单个测试用例\n"
	var mode string
	flag.StringVar(&mode, "mode", "gen_test", modeUsage)
	var font, lang, psm string
	flag.StringVar(&font, "font", fontYahei, "字体 SimSun")
	flag.StringVar(&lang, "lang", "eng", "模型选择：lang,yahei")
	flag.StringVar(&psm, "psm", "13", "PageMode")
	var limit int
	flag.IntVar(&limit, "limit", 200, "检查数据集")
	flag.Parse()

	switch mode {
	case "prepare_pass":
		var ans = genRandomNCase(200)
		fmt.Println("len(ans) = ", len(ans))
		if err := writeNCase(ans[0:limit]); err != nil {
			log.Fatal(err)
		}
	case "pass":
		if err := passRate(lang, psm); err != nil {
			log.Fatal(err)
		}
	case "single": // 单个测试用例
		text, err := ocrText("train/0.tif", lang, psm)
		log.Printf("err:%v text=%s\n", err, text)
	case "gen":
		lines := genGroundTruth()
		fmt.Println("len(lines) = ", len(lines))
		if err := genImgByFont(font, lines); err != nil {
			log.Fatal(err)
		}
	case "load_store":
		loadStore()
	default:
		log.Println("unknown mode:", mode)
	}
}

// 生成训练数据(4个字符组合) 243个用例
func genGroundTruth() []string {
	// 易混淆的字符: 0O、5S (2^4=16）
	var tmp = make([]string, 0, 100)
	dfs("", 4, "0O", &tmp)
	dfs("", 4, "5S", &tmp)
	if len(tmp) != 32 {
		panic("len(tmp) should be 16")
	}
	yahei := yaheiCharset()
	if len(yahei) != 9 {
		panic("len(yahei) should be 9")
	}
	rnd := genRandomNCase(200)
	if len(rnd) != 200 {
		panic("len(rnd) should be 200")
	}
	all := []string{"FPH0", "O0JC"}
	all = append(all, tmp...)
	all = append(all, yahei...)
	all = append(all, rnd...)
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
func yaheiCharset() []string {
	var results []string
	end := len(allowChars) - 4
	for i := 0; i <= end; i += 4 {
		// fmt.Println(allowChars[i : i+4])
		results = append(results, allowChars[i:i+4])
	}
	return results
}

// 模式3：随机N个用例
func genRandomNCase(n int) []string {
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
