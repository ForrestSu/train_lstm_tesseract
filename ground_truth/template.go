package groundtruth

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// NewRandom 随机用例生成
func NewRandom(idxFile string, font string, outDir string, force bool) *Template {
	seq := randomGen{}
	return &Template{idxFile: idxFile, font: font, Seq: seq, OutDir: outDir, Force: force}
}

// NewAmbiguous 易混淆用例生成
func NewAmbiguous(idxFile string, font string, outDir string, force bool) *Template {
	seq := ambiguousGen{}
	return &Template{idxFile: idxFile, font: font, Seq: seq, OutDir: outDir, Force: force}
}

type Template struct {
	idxFile string // 索引文件
	font    string // 使用的字体
	// 序列生成
	Seq    IGenerator // 序列生成器
	OutDir string     // 输出目录
	Force  bool       // 强制重写图片
}

// Gen 生成用例 (模型训练/识别率测试)
// n 随机用例个数
func (t *Template) Gen(n int) []string {
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

func (t *Template) writeImg(items []string) {
	if err := genImgByFont(items, t.OutDir, t.font); err != nil {
		panic(fmt.Sprintf("生成图片失败：%v", err))
	}
	log.Printf("图片写入成功: %d 条\n", len(items))
}

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
