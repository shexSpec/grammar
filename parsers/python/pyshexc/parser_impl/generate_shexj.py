# Copyright (c) 2017, Mayo Clinic
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
# Redistributions of source code must retain the above copyright notice, this
#     list of conditions and the following disclaimer.
#
#     Redistributions in binary form must reproduce the above copyright notice,
#     this list of conditions and the following disclaimer in the documentation
#     and/or other materials provided with the distribution.
#
#     Neither the name of the Mayo Clinic nor the names of its contributors
#     may be used to endorse or promote products derived from this software
#     without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
# INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
# LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
# OF THE POSSIBILITY OF SUCH DAMAGE.
# Copyright (c) 2017, Mayo Clinic
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
# Redistributions of source code must retain the above copyright notice, this
#     list of conditions and the following disclaimer.
#
#     Redistributions in binary form must reproduce the above copyright notice,
#     this list of conditions and the following disclaimer in the documentation
#     and/or other materials provided with the distribution.
#
#     Neither the name of the Mayo Clinic nor the names of its contributors
#     may be used to endorse or promote products derived from this software
#     without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
# INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
# LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
# OF THE POSSIBILITY OF SUCH DAMAGE.
import os
import sys
from argparse import ArgumentParser

from antlr4 import CommonTokenStream
from antlr4 import FileStream, InputStream
from antlr4.error.ErrorListener import ErrorListener
from pyshexc.parser.ShExDocLexer import ShExDocLexer
from pyshexc.parser.ShExDocParser import ShExDocParser
from pyjsg.jsglib.jsg import JSGObject
from typing import Optional, Union, List

from pyshexc.parser_impl.shex_doc_parser import ShexDocParser


class ParseErrorListener(ErrorListener):
    """ Record errors and text """

    def __init__(self):
        self.n_errors = 0
        self.errors = []

    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        self.n_errors += 1
        self.errors.append("line " + str(line) + ":" + str(column) + " " + msg)


def do_parse(infilename: str, outfilename: str) -> bool:
    """
    Parse the jsg in infilename and save the results in outfilename
    :param infilename: name of the file containing the ShExC
    :param outfilename: target ShExJ equivalent
    :return: true if success
    """
    shexj = parse(FileStream(infilename, encoding="utf-8"))
    if shexj is not None:
        with open(outfilename, 'w') as outfile:
            outfile.write(shexj._as_json_dumps())
        return True
    return False


def parse(input_: Union[str, FileStream]) -> Optional[JSGObject]:
    """
    Parse the text in infile and save the results in outfile
    :param input_: text or input stream to parse
    :return: ShExJ text if successful.  None if error.
    """

    # Step 1: Tokenize the input stream
    error_listener = ParseErrorListener()
    if not isinstance(input_, FileStream):
        input_ = InputStream(input_)
    lexer = ShExDocLexer(input_)
    lexer.addErrorListener(error_listener)
    tokens = CommonTokenStream(lexer)
    tokens.fill()
    if error_listener.n_errors:         # Lexer prints errors directly
        return None

    # Step 2: Generate the parse tree
    parser = ShExDocParser(tokens)
    parser.addErrorListener(error_listener)
    parse_tree = parser.shExDoc()
    if error_listener.n_errors:
        print('\n'.join(error_listener.errors), file=sys.stderr)
        return None

    # Step 3: Transform the results the results
    parser = ShexDocParser()
    parser.visit(parse_tree)

    return parser.context.schema


def genargs() -> ArgumentParser:
    """
    Create a command line parser
    :return: parser
    """
    parser = ArgumentParser()
    parser.add_argument("infile", help="Input ShExC specification")
    parser.add_argument("-o", "--outfile", help="Output ShExJ file (Default: {infile}.json)")
    return parser


def generate(argv: List[str]) -> bool:
    """ 
    Transform ShExC to ShExJ
    :param argv: Command line arguments
    :return: True if successful
    """
    opts = genargs().parse_args(argv)
    if not opts.outfile:
        opts.outfile = os.path.dirname(opts.infile) + str(os.path.basename(opts.infile).rsplit('.', 1)[0]) + ".json"
    if do_parse(opts.infile, opts.outfile):
        print("Output written to {}".format(opts.outfile))
        return True
    else:
        print("Conversion failed")
        return False
