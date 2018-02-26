#!/usr/bin/env bash

# pip install git+git://github.com/c2nes/javalang

python parse.py ~/Projects/chargebee/chargebee-java 2> chargebee.json

swagger generate client -f chargebee.json \
    -A chargebee \
    --target ../
#    --skip-validation \
