package groundtruth

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 使用字体生成文件 (dir不含尾部斜杠)
func genImgByFont(lines []string, outDir string, font string) error {
	if err := os.MkdirAll(outDir, 0777); err != nil {
		return err
	}
	outDir = strings.TrimSuffix(outDir, "/")
	for lineNum, line := range lines {
		// 生成输出文件基础名
		baseName := fmt.Sprintf("%s/eng_%03d_%s", outDir, lineNum, line)
		// 训练的文本文件
		trainTextFile := baseName + ".gt.txt"
		// 写入文本行到文件
		if err := os.WriteFile(trainTextFile, []byte(line), 0644); err != nil {
			return fmt.Errorf("写入训练文本失败: %w", err)
		}
		if err := text2Image(font, trainTextFile, baseName); err != nil {
			return err
		}
	}
	return nil
}

// 注意点：
// 1 fonts_dir 列出所有字体才有效。text2image --fonts_dir=xxx --list_available_fonts
func text2Image(font string, trainTextFile string, outputBaseDir string) error {
	// fmt.Println(">>> " + trainTextFile)
	args := []string{
		"--font=" + font,
		fmt.Sprintf("--text=%s", trainTextFile),
		fmt.Sprintf("--outputbase=%s", outputBaseDir),
		// "--fonts_dir=", // --fonts_dir /usr/share/fonts
		"--max_pages=1",
		"--strip_unrenderable_words",
		"--xsize=240",
		"--ysize=80",
		"--resolution=300",   // 图片分辨率DPI (默认300像素/英寸)
		"--margin=5",         // 图像边缘的圆角
		"--ptsize=12",        // 打印文本的大小
		"--leading=12",       // 行间间距 (以像素为单位) (type:int 默认值:12)
		"--box_padding=0",    // 生成的边界框周围的填充 (type:int 默认值:0)
		"--char_spacing=1.0", // 字符间距，单位：ems (type:double 默认值:0)
		"--exposure=0",       // 复印机的曝光级别 (type:int 默认值:0)
		"--blur=true",        // 模糊图像 (type:bool 默认值:true)
		// "--white_noise=false",   // 添加高斯噪声 (type:bool 默认值:true)
		// "--smooth_noise=false",  // 平滑噪声 (type:bool 默认值:true)
		"--degrade_image=false", // 使用斑点噪声、膨胀/侵蚀和旋转来降低渲染图像的质量
		"--rotate_image=false",  // 以随机方式旋转图像
		"--unicharset_file=eng.unicharset",
	}
	cmd := exec.Command("text2image", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("调用cmd1失败! err=%v, msg=%s", err, stderr.String())
	}
	if stderr.Len() > 0 {
		fmt.Printf(">> %s", stderr.String())
	}
	return nil
}
