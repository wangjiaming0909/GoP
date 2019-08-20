#!/bin/bash

ProjectName=GoP

for d in $(go list ./... | grep -v -E ".+$ProjectName$" | grep -w 'tests');do
    go test $d
done


