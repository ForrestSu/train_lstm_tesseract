package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func writeNCase(results []string) error {
	for k, one := range results {
		baseName := "img/" + strconv.Itoa(k) + "_" + one
		// 写入文本行到文件
		if err := os.WriteFile(baseName+".gt.txt", []byte(one), 0777); err != nil {
			return fmt.Errorf("写入训练文本失败: %w", err)
		}
		if err := text2Image(fontYahei, baseName+".gt.txt", baseName); err != nil {
			return err
		}
	}
	return nil
}

// 通过率测试
func passRate(lang string, psm string) error {
	var pass int
	var results []string
	var total = len(results)
	for k, one := range results {
		baseName := "img/" + strconv.Itoa(k) + "_" + one
		text, err := ocrText(baseName+".tif", lang, psm)
		if err != nil {
			return err
		}
		if text == one {
			pass++
		} else {
			fmt.Printf("case:%d, real=%s, got=%s\n", k, one, text)
		}
		if k%100 == 0 {
			fmt.Printf("progress %.0f%% ", float32(k)/float32(total)*100)
			ratio := float32(pass) / float32(k+1)
			fmt.Printf("pass: %d/%d, pass rate %.2f%%\n", pass, k+1, ratio*100)
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
		return "", fmt.Errorf("调用cmd2失败! err=%v", err)
	}
	if stderr.Len() > 0 {
		return "", fmt.Errorf("执行返回 stderr=%s", stderr.String())
	}
	return strings.TrimSpace(stdout.String()), nil
}
