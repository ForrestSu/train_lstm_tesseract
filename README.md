# train_lstm_tesseract
training lstm new font tesseract5

# 1 搭建 train 环境 (ubuntu 25.04)
```bash
docker run -it --privileged --name tess -v ~/tools/lstm:/root/lstm sunquana/ubuntu:tesseract5 zsh
```

# 训练
```bash
make training MODEL_NAME=arial START_MODEL=best_eng TESSDATA=/usr/share/tesseract-ocr/5/tessdata  MAX_ITERATIONS=10000
```

## 识别率对比
| 模型    | 训练数据图片 | 测试数据图片 | pass_rate |
|-------|--------|--------|-----------|
| eng   | ---    | 200张   | 76.50%    |
| arial | 243张   | 200张   | 92.0%     |


```bash
> ➜ ground-truth git:(main) ./ground-truth -mode pass -lang arial
2025/05/06 22:33:09 加载已有的测试用例: 200 条
progress 0% pass: 1/1, pass rate 100.00%
case:9, real=OJ00, got=0J00
case:16, real=2AOD, got=2A0D
progress 25% pass: 49/51, pass rate 96.08%
case:53, real=O033, got=0033
case:60, real=3OAU, got=30AU
case:87, real=LO76, got=L076
progress 50% pass: 96/101, pass rate 95.05%
case:128, real=4ZI5, got=4Z15
case:133, real=J0H7, got=JOH7
case:135, real=72YO, got=72Y0
case:150, real=B0E7, got=BOE7
progress 75% pass: 142/151, pass rate 94.04%
case:165, real=UL0P, got=ULOP
case:166, real=U6O4, got=U604
case:171, real=0O0L, got=000L
case:172, real=IZE7, got=1ZE7
case:176, real=O4QM, got=04QM
case:177, real=D1UW, got=DIUW
case:188, real=14O6, got=1406
pass=184, total=200, pass rate 92.00%
```