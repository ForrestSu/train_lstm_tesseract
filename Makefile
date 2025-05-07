.PHONY: clean
RCP = rsync -avh --compress --partial --progress
SOURCES := $(shell find * -type f -name "*.go")
ModelName = arial
InstallDir = /opt/homebrew/share/tessdata/

default: build

build: $(SOURCES) go.mod
	cd ground-truth && GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -v .

sync:
	$(RCP) ground-truth/ground-truth  hu40:~/tools/lstm/tutorial_tesseract/

# 开始训练模型
# 准备工作： best/eng.traineddata 拷贝到 ../tesseract/tessdata 目录下
#   DEBUG_INTERVAL=-1  # 每迭代一次，打印一次日志
#   PSM=13 # 单行文本
train:
	 TESSDATA_PREFIX=../tesseract/tessdata  make training MODEL_NAME=arial START_MODEL=eng TESSDATA=../tesseract/tessdata  MAX_ITERATIONS=10000

# 生成模型(best fast)
traineddata:
	make traineddata  MODEL_NAME=$(ModelName)

# 将模型拷贝回本地
copy_back_install:
	$(RCP) hu40:~/tools/lstm/tutorial_tesseract/tesstrain/data/$(ModelName).traineddata $(InstallDir)

## rcp hu40:~/tools/lstm/tutorial_tesseract/tesstrain/data/$(ModelName)/tessdata_fast/arial_0.000_30_1300.traineddata ./


# 绘图-训练迭代曲线图
plot:
	make plot MODEL_NAME=$(ModelName) START_MODEL=best_eng

clean:
	rm -rf ground-truth/ground-truth
