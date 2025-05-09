package main

import (
	"flag"
	"fmt"
	"log"

	groundtruth "github.com/ForrestSu/train_lstm_tesseract/ground_truth"
)

func main() {
	var modeUsage = "[mode] 模式选择:\n" +
		"- gen: 生成训练数据【模棱两可字母】\n" +
		"- gen_random: 生成训练数据\n" +
		"- eval: 识别率评估\n" +
		"- single: 单个测试用例\n"
	var mode string
	flag.StringVar(&mode, "mode", "real", modeUsage)
	// 生成数据
	var font string
	var n int
	var force bool
	flag.StringVar(&font, "font", "Arial Regular", "使用字体 [gen]")
	flag.IntVar(&n, "n", 1, "随机用例个数 [gen]")
	flag.BoolVar(&force, "force", false, "根据索引生成图片 [gen]")
	// 模型识别
	var lang, psm, testDir string
	flag.StringVar(&lang, "lang", "eng", "模型选择：eng")
	flag.StringVar(&psm, "psm", "13", "PageMode: 7 (单行文本) 13 (多行文本)")
	flag.StringVar(&testDir, "test_dir", "img/", "用于测试的图片目录")
	flag.Parse()

	switch mode {
	case "gen": // 生成训练数据（模棱两可的字母）
		tpl := groundtruth.NewAmbiguous("ground_truth.txt", font,
			"tesstrain/data/arial-ground-truth/", force)
		tpl.Gen(n)
	case "gen_random": // 生成随机用例
		tpl := groundtruth.NewRandom("random_case.txt", font, "img/", force)
		tpl.Gen(n)
	case "eval": // 识别率评估
		fmt.Printf(">>> Lang:%s PSM=%s TestDir=[%s]...\n", lang, psm, testDir)
		if err := groundtruth.PassRate(tpl.OutDir, lines, lang, psm); err != nil {
			log.Fatal(err)
		}
	case "real":
		tpl := groundtruth.NewH160("real.txt", "tesstrain/data/arial-ground-truth/")
		tpl.Gen(n)
	case "single": // 单个测试用例
		text, err := groundtruth.OcrText("train/0.tif", lang, psm)
		log.Printf("err:%v text=%s\n", err, text)
	default:
		log.Println("unknown mode:", mode)
	}
}
