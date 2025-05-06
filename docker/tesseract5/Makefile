.PHONY: clean

# 默认
default: build

build:
	docker build --target tesseract5 -t sunquana/ubuntu:tesseract5 .

run:
	docker run -it --privileged  -v ~/tools/lstm:/root/lstm  --name tess sunquana/ubuntu:tesseract5 zsh
