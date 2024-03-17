#!/bin/bash

if [ $# -lt 1 ]
then
    echo "$0 <projectname>" 1>&2
    exit 1
fi

find . -type f -exec sed -i "s/gintemplate/$1/g" {} \;


