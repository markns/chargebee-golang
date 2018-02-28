#!/usr/bin/env bash

set -ex

# pip install git+git://github.com/c2nes/javalang

#python parse.py ~/Projects/chargebee/chargebee-java 2> chargebee.json

swagger generate client -f chargebee.json \
    -A chargebee \
    --target ../
#    --skip-validation \


cp -r ../models ../client \
    ~/Projects/src/bitbucket.org/gridarrow/gridarrow-manager/vendor/github.com/markns/chargebee-golang


#first_name=John&last_name=Doe&email=john%40test.com&locale=fr-CA&billing_address%5Bfirst_name%5D=John&billing_address%5Blast_name%5D=Doe&billing_address%5Bline1%5D=PO+Box+9999&billing_address%5Bcity%5D=Walnut&billing_address%5Bstate%5D=California&billing_address%5Bzip%5D=91789&billing_address%5Bcountry%5D=US