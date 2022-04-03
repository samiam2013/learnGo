#!/bin/bash

read -s -p "password for apt deps: " PW

echo "$PW" | sudo -S apt-get update &&
    sudo apt-get install \
    libc6-dev \
    libglu1-mesa-dev \
    libgl1-mesa-dev \
    libxcursor-dev \
    libxi-dev \
    libxinerama-dev \
    libxrandr-dev \
    libxxf86vm-dev \
    libasound2-dev \
    pkg-config \
    -y ;
