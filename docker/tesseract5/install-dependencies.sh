#!/bin/bash

# os-release may be missing in container environment by default.
if [ -f "/etc/os-release" ]; then
    . /etc/os-release
elif [ -f "/etc/arch-release" ]; then
    export ID=arch
else
    echo "/etc/os-release missing."
    exit 1
fi

# define basic tools
basic_tools=(
    ascii
    tree
    unzip
    p7zip
    zsh
    fish
    git
    make
    man-db
    tmux  ## effective
    parallel
    stow
    tcpdump ## net tools
    traceroute
    net-tools
    ethtool
    bwm-ng
    mtr
    dsniff
    nmap
    htop ## top tools
    iotop
    iftop
    sysstat
    iperf ## perf
    iperf3
    strace
    lsof
    wget ## file
    file
    dos2unix
    rsync
    cloc
)

# various linux distributions
# (1) ubuntu 25.04
ubuntu_packages=(
    "${basic_tools[@]}"
    vim
    iproute2 # ss
    autoconf
    libtool
    curl
    gcc
    g++
    cmake
)

# tesseract depends
tesseract_deps=(
    make
    wget
    unzip
    python3
    python3-pip
    # tesseract
    tesseract-ocr
    tesseract-ocr-chi-sim
)

if [ "$ID" = "ubuntu" ] || [ "$ID" = "debian" ]; then
    apt-get install -y "${ubuntu_packages[@]}"
    apt-get install -y "${tesseract_deps[@]}"
else
    echo "Your system ($ID) is not supported by this script. Please install dependencies manually."
    exit 1
fi
