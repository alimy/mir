#!/usr/bin/env bash

GOFILES=$(find . -type f)

for filename in $GOFILES
do
  mv ${filename} ${filename}.tmpl
done