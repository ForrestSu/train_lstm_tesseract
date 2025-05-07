## train_lstm_tesseract
training lstm new font tesseract5

## 前言

使用 tesseract5 训练 lstm 模型，文本所使用字体 arial。 发现个别识别率较低，如何提高识别率？

比如： `0 O` `1 I` `5 S`

## 1 搭建 train 环境 (ubuntu 25.04)

```bash
docker run -it --privileged --name tess -v ~/tools/lstm:/root/lstm sunquana/ubuntu:tesseract5 zsh
```

## 2 训练
```bash
make training MODEL_NAME=arial START_MODEL=best_eng TESSDATA=/usr/share/tesseract-ocr/5/tessdata  MAX_ITERATIONS=10000
```

## 3 识别率对比

- 使用默认的PSM = 13训练

| 模型    | 训练数据图片 | 测试数据图片 | psm | pass_rate |
|-------|--------|--------|-----|-----------|
| eng   | ---    | 200张   | 7   | 75%       |
| eng   | ---    | 200张   | 13  | 84%       |
| arial | 243张   | 200张   | 7   | 94.5%     |
| arial | 243张   | 200张   | 13  | 92%       |

### 3.1 优化方向
[x] 指定识别模式 --user-patterns。 
[ ] 提高训练的数据集 (200->1000)
    下载1000张真实数据，用于训练。 再下载1000张真实数据，用于测试准确率。

## 参考文档
- [OCR 100% accuracy of digital data](https://www.monperrus.net/martin/perfect-ocr-digital-data)
