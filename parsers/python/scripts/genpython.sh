#!/usr/bin/env bash
# Note: need antlr4.5.1 or later to run this!
# see: http://www.antlr.org/download.html for details
wget https://raw.githubusercontent.com/shexSpec/grammar/master/ShExDoc.g4
antlr4 -Dlanguage=Python3 -package parser -o ../pyshexc/parser -no-listener -visitor ShExDoc.g4
rm ShExDoc.g4
wget https://raw.githubusercontent.com/hsolbrig/shexTest/master/doc/ShExJ.jsg -O ../pyshexc/shexj/ShExJ.jsg
generate_parser -e ../pyshexc/shexj/ShExJ.jsg -o ../pyshexc/shexj/ShExJ.py
