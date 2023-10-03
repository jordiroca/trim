#!/usr/bin/env bash
# -*- coding: utf-8 -*-
#
# DESCRIPTION: install Install trim, ltrim and rtrim utilities
#
# VERSION: 23.10.2
#
# OPTIONS:
#   -h, --help    Display this message.
#
# AUTHOR: Jordi Roca
# CREATED: 2023/10/02
#
# GITHUB: https://github.com/jordiroca/
# WEBSITE: https://jordiroca.com
#
# LICENSE: See LICENSE file.
#

echo "Compiling trim, ltrim and rtrim utilities"
go build trim.go
if [ "$?" -eq 0 ]; then
    cp trim ltrim
    cp trim rtrim
    echo "OS may ask for sudo password to copy files to /usr/local/bin/"
    sudo cp -fv trim ltrim rtrim /usr/local/bin/
    echo
fi
