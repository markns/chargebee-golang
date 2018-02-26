#!/usr/bin/env bash


python parse.py ~/Projects/chargebee/chargebee-java 2> chargebee.json

swagger generate client -f chargebee.json \
    -A chargebee \
    --target ../
#    --skip-validation \
