#!/usr/bin/env bash
# -*- coding: utf-8 -*-
#
# DESCRIPTION: example Example usage of trim, rtrim and ltrim utilities
#
# VERSION: 23.10.2
#
# OPTIONS:
#   -h, --help    Display this message.
#
# AUTHOR: Jordi Roca
# CREATED: 2023/10/02
#
# LICENSE: See LICENSE file.
#


chmod +x *.sh
./install.sh
./center_haikus.sh < sample.txt > haikus.txt

echo "Left trimming"
echo "============="
cat haikus.txt | ./ltrim | tr " " "_"
echo -e "\n\nRight trimming"
echo "=============="
cat haikus.txt | ./rtrim | tr " " "_"
echo -e "\n\nTrimming"
echo "========"
cat haikus.txt | ./trim | tr " " "_"
echo
