#!/bin/bash

mkdir tools
cd zbar
autoreconf -vfi
./configure
make -j10
cp -r ./zbarimg/ ../tools/
