.PHONY: clean
SOURCES := $(shell find * -type f -name "*.go")

build: $(SOURCES) go.mod
	cd ground-truth && GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -v .

sync:
	rcp ground-truth/ground-truth  hu40:~/tools/lstm/tutorial_tesseract/

# 开始训练模型
train:
	make training MODEL_NAME=yahei START_MODEL=best_eng TESSDATA=/usr/share/tesseract-ocr/5/tessdata  MAX_ITERATIONS=10000

# 绘图-训练迭代曲线图
plot:
	make plot MODEL_NAME=yahei START_MODEL=best_eng

clean:
	rm -rf ground-truth/ground-truth