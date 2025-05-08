package groundtruth

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ForrestSu/health160/captcha"
	"github.com/ForrestSu/health160/sdk"
	"github.com/ForrestSu/health160/sdk/model"
)

type H160 struct {
	idxFile string // 索引文件
	outDir  string // 输出目录
}

func NewH160(idxFile string, outDir string) *H160 {
	return &H160{idxFile: idxFile, outDir: outDir}
}

func (t *H160) Gen(n int) []string {
	if oldItems := loadExistedCase(t.idxFile); len(oldItems) > 0 {
		return oldItems
	}
	err := writeImg(t.outDir, n)
	if err != nil {
		panic(err)
	}
	log.Printf("成功下载 %d 用例\n", n)
	// _ = os.WriteFile(t.idxFile, []byte(strings.Join(items, "\n")), 0777)
	return nil
}

func writeImg(outDir string, n int) error {
	if err := os.MkdirAll(outDir, 0777); err != nil {
		return err
	}
	outDir = strings.TrimSuffix(outDir, "/")
	for i := 0; i < n; i++ {
		data, err := parseCaptcha()
		if err != nil {
			return err
		}
		code := data.CaptchaCode()
		// 生成输出文件基础名
		baseName := fmt.Sprintf("%s/eng_%03d_%s", outDir, i, code)
		// 训练的文本文件
		trainTextFile := baseName + ".gt.txt"
		if err := os.WriteFile(trainTextFile, []byte(code), 0777); err != nil {
			return fmt.Errorf("写入训练文本失败: %v", err)
		}
		// move cap.png to outDir
		if err := os.Rename("cap.png", baseName+".png"); err != nil {
			return fmt.Errorf("移动文件失败: %v", err)
		}
		fmt.Printf(">> 第%d张 code=%s\n", i, code)
		// 写入原始图片
		os.WriteFile(baseName+"_ori.png", data.ImgBytes, 0777)
		// 打印告警日志
		if len(code) != 4 {
			log.Printf("Hint: 第%d行 code=%s 长度不为4\n", i, code)
		}
	}
	return nil
}

var decoder = captcha.NewTesseractDecoderV2("arial", 13)

// 解析-登录验证码
func parseCaptcha() (*model.TokenData, error) {
	data, err := sdk.ParseCaptcha()
	if err != nil {
		return nil, err
	}
	// 使用 decoder 识别验证码
	text, err := decoder.Decode(data.ImgBytes)
	if err != nil {
		return nil, fmt.Errorf("验证码识别错误! err=%v", err)
	}
	data.SetCaptCode(text)
	return data, nil
}
