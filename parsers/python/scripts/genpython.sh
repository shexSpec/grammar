#!/usr/bin/env bash
# Note: need antlr4.5.1 or later to run this!
# see: http://www.antlr.org/download.html for details
cp ../../../ShExDoc.g4 .
antlr4 -Dlanguage=Python3 -package parser -o ../pyshexc/parser -no-listener -visitor ShExDoc.g4
rm ShExDoc.g4

