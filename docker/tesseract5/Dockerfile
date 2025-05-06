FROM ubuntu:25.04 AS tesseract5
LABEL maintainer="ForrestSu <sunquana@gmail.com>"

ENV LANG=en_US.UTF-8 \
    LANGUAGE=en_US:en \
    LC_ALL=en_US.UTF-8 \
    TZ=Asia/Shanghai

ENV HOME=/root
WORKDIR /root

COPY sources.list $HOME/sources.list
## use tsinghua source
RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak && mv $HOME/sources.list /etc/apt/ \
    && apt-get update && apt-get -y upgrade \
    && apt-get install -y apt-utils build-essential

COPY install-dependencies.sh .
RUN ./install-dependencies.sh

COPY requirements.txt .
RUN pip install --break-system-packages -r ./requirements.txt

COPY fonts/* /usr/share/fonts/truetype/
COPY zshrc $HOME/.custom_zshrc
RUN cat ~/.custom_zshrc >> ~/.zshrc && rm ~/.custom_zshrc \
    && usermod -s /bin/zsh root \
    && ldconfig

CMD ["zsh"]
