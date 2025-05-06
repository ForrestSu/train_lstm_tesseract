package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// 通过率测试
func passRate(dir string, items []string, lang string, psm string) error {
	var pass int
	var total = len(items)
	dir = strings.TrimSuffix(dir, "/")
	for k, one := range items {
		baseName := fmt.Sprintf("%s/eng_%03d_%s", dir, k, one)
		text, err := ocrText(baseName+".tif", lang, psm)
		if err != nil {
			return err
		}
		if text == one {
			pass++
		} else {
			fmt.Printf("failure! k=%d (expected=%s, got=%s)\n", k, one, text)
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

const limited = "tessedit_char_whitelist=" + allowChars

func ocrText(fileName string, lang string, psm string) (string, error) {
	// fmt.Println(">>> " + fileName)
	args := []string{fileName, "stdout", "-c", limited, "-l", lang, "--psm", psm, "quiet"}
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
