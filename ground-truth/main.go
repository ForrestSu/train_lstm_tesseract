package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
		tpl := template{
			idxFile: "ground_truth.txt",
			font:    font,
			Seq:     groundTruthSeq{},
			OutDir:  "tesstrain/data/arial-ground-truth/",
			Force:   force,
		}
		tpl.Gen(n)
	case "eval": // 识别率评估
		fmt.Printf(">>> Lang:%s PSM=%s...\n", lang, psm)
		tpl := template{
			idxFile: "random_case.txt",
			font:    font,
			Seq:     randomSeq{},
			OutDir:  "img/",
			Force:   force,
		}
		lines := tpl.Gen(n)
		if err := passRate(tpl.OutDir, lines, lang, psm); err != nil {
			log.Fatal(err)
		}
	case "single": // 单个测试用例
		text, err := ocrText("train/0.tif", lang, psm)
		log.Printf("err:%v text=%s\n", err, text)
	default:
		log.Println("unknown mode:", mode)
	}
}

type template struct {
	idxFile string // 索引文件
	font    string // 使用的字体
	// 序列生成
	Seq    SeqGenerator // 序列生成器
	OutDir string       // 输出目录
	Force  bool         // 强制重写图片
}

// Gen 生成用例 (模型训练/识别率测试)
// n 随机用例个数
func (t *template) Gen(n int) []string {
	if oldItems := loadExistedCase(t.idxFile); len(oldItems) > 0 {
		if t.Force {
			t.writeImg(oldItems)
		}
		return oldItems
	}
	items := t.Seq.Gen(n)
	t.writeImg(items)
	_ = os.WriteFile(t.idxFile, []byte(strings.Join(items, "\n")), 0777)
	return items
}

func (t *template) writeImg(items []string) {
	if err := genImgByFont(items, t.OutDir, t.font); err != nil {
		panic(fmt.Sprintf("生成图片失败：%v", err))
	}
	log.Printf("图片写入成功: %d 条\n", len(items))
}
