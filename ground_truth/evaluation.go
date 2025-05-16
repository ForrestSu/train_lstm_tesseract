package groundtruth

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// PassRate 通过率测试
func PassRate(dir string, lang string, psm string) error {
	imgFiles, err := travelImg(dir)
	if err != nil {
		return err
	}
	var pass int
	var total = len(imgFiles)
	fmt.Printf("加载到 %d 个图片文件\n", total)
	for k, fileName := range imgFiles {
		var truth = parseTruth(fileName)
		if len(truth) != 4 {
			return fmt.Errorf("fileName: %s 非法的文件名！k=%d", fileName, k)
		}
		gotText, err := OcrText(fileName, lang, psm)
		if err != nil {
			return err
		}
		if gotText == truth {
			pass++
		} else {
			fmt.Printf("FAIL: k=%d (expected=%s, got=%s)\n", k, truth, gotText)
		}
		var cnt = k + 1
		if (cnt)%50 == 0 {
			fmt.Printf("progress %.0f%% ", float32(cnt)/float32(total)*100)
			ratio := float32(pass) / float32(cnt)
			fmt.Printf("pass: %d/%d, pass rate %.2f%%\n", pass, cnt, ratio*100)
		}
	}
	fmt.Printf("pass=%d, total=%d, pass rate %.2f%%\n", pass, total, float32(pass)/float32(total)*100)
	return nil
}

func travelImg(dir string) ([]string, error) {
	var imgFiles []string
	// 递归遍历目录
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		extName := strings.ToLower(filepath.Ext(path))
		if slices.Contains([]string{".png", ".jpg", ".tif"}, extName) {
			imgFiles = append(imgFiles, path)
			// fmt.Println(">> " + path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return imgFiles, nil
}

// eng_000_P5SY.png
func parseTruth(fileName string) string {
	baseName := filepath.Base(fileName)
	end := strings.LastIndex(baseName, ".")
	if end < 0 {
		return ""
	}
	start := strings.LastIndex(baseName, "_")
	if start < 0 {
		return ""
	}
	return baseName[start+1 : end]
}

const limited = "tessedit_char_whitelist=" + allowChars

func OcrText(fileName string, lang string, psm string) (string, error) {
	// fmt.Println(">>> " + fileName)
	args := []string{fileName, "stdout", "-c", limited, "-l", lang, "--psm", psm}
	args = append(args, "--user-patterns", "my.patterns")
	args = append(args, "quiet")
	cmd := exec.Command("tesseract", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("调用cmd2失败! err=%v, msg=%s", err, stderr.String())
	}
	if stderr.Len() > 0 {
		return "", fmt.Errorf("执行返回 stderr=%s", stderr.String())
	}
	return strings.TrimSpace(stdout.String()), nil
}
