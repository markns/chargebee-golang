#!/usr/bin/env bash


python parse.py 2> chargebee.json

swagger generate client -f chargebee.json \
    -A chargebee \
    --target ../
#    --skip-validation \
