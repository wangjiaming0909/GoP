#!/bin/bash

ProjectName=GoP


for d in $(go list ./... | grep -v -w 'tests' | grep -v -E ".+$ProjectName$"); do
    target=$(basename $d)
    echo 'building ' $target
    go build -o $target/$target $d && mv ./$target/$target $GOPATH/bin
done


echo 'building ' $ProjectName
go build -o $ProjectName . && mv ./$ProjectName $GOPATH/bin
