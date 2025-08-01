#!/bin/sh
sed -re"s/^/ABC/" -e"s/ABC([0-9]{4})/ABD/g" -e"s/^(ABC).+$/\1/g"
