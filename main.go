package main

import (
	"flag"
	"fmt"
	"log"

	groundtruth "github.com/ForrestSu/train_lstm_tesseract/ground_truth"
)

func main() {
	var modeUsage = "[mode] 模式选择:\n" +
		"- gen: 生成训练数据\n" +
		"- eval: 识别率评估\n" +
		"- single: 单个测试用例\n"
	var mode string
	flag.StringVar(&mode, "mode", "gen", modeUsage)
	var font string
	flag.StringVar(&font, "font", "Arial Regular", "使用字体 [gen]")
	var n int
	flag.IntVar(&n, "n", 200, "随机用例个数 [gen]")
	var force bool
	flag.BoolVar(&force, "force", false, "根据索引生成图片 [gen]")
	// 模型识别
	var lang, psm string
	flag.StringVar(&lang, "lang", "eng", "模型选择：eng")
	flag.StringVar(&psm, "psm", "13", "PageMode: 7 (单行文本) 13 (多行文本)")
	flag.Parse()

	switch mode {
	case "gen": // 生成训练数据
		tpl := groundtruth.NewTemplate("ground_truth.txt", font,
			groundtruth.GenAmbiguous, "tesstrain/data/arial-ground-truth/", force)
		tpl.Gen(n)
	case "eval": // 识别率评估
		fmt.Printf(">>> Lang:%s PSM=%s...\n", lang, psm)
		tpl := groundtruth.NewTemplate("random_case.txt", font,
			groundtruth.GenRandom, "img/", force)
		lines := tpl.Gen(n)
		if err := groundtruth.PassRate(tpl.OutDir, lines, lang, psm); err != nil {
			log.Fatal(err)
		}
	case "single": // 单个测试用例
		text, err := groundtruth.OcrText("train/0.tif", lang, psm)
		log.Printf("err:%v text=%s\n", err, text)
	default:
		log.Println("unknown mode:", mode)
	}
}
