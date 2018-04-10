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
from typing import Optional, Union, List

from antlr4 import CommonTokenStream
from antlr4 import FileStream, InputStream
from antlr4.error.ErrorListener import ErrorListener
from rdflib import Graph
from rdflib.plugin import plugins as rdflib_plugins, Serializer as rdflib_Serializer
from rdflib.util import SUFFIX_FORMAT_MAP

from pyshexc.parser_impl.shex_doc_parser import ShexDocParser
from pyshexc.parser.ShExDocLexer import ShExDocLexer
from pyshexc.parser.ShExDocParser import ShExDocParser
from ShExJSG.ShExJ import Schema


class ParseErrorListener(ErrorListener):
    """ Record errors and text """

    def __init__(self):
        self.n_errors = 0
        self.errors = []

    def syntaxError(self, recognizer, offending_symbol, line, column, msg, e):
        self.n_errors += 1
        self.errors.append("line " + str(line) + ":" + str(column) + " " + msg)


def do_parse(infilename: str, jsonfilename: Optional[str], rdffilename: Optional[str], rdffmt: str,
             context: Optional[str] = None) -> bool:
    """
    Parse the jsg in infilename and save the results in outfilename
    :param infilename: name of the file containing the ShExC
    :param jsonfilename: target ShExJ equivalent
    :param rdffilename: target ShExR equivalent
    :param rdffmt: target RDF format
    :param context: @context to use for rdf generation. If None use what is in the file
    :return: true if success
    """
    shexj = parse(FileStream(infilename, encoding="utf-8"))
    if shexj is not None:
        shexj['@context'] = context if context else "http://www.w3.org/ns/shex.jsonld"
        if jsonfilename:
            with open(jsonfilename, 'w') as outfile:
                outfile.write(shexj._as_json_dumps())
        if rdffilename:
            g = Graph().parse(data=shexj._as_json, format="json-ld")
            g.serialize(open(rdffilename, "wb"), format=rdffmt)
        return True
    return False


def parse(input_: Union[str, FileStream], default_base: Optional[str]=None) -> Optional[Schema]:
    """
    Parse the text in infile and return the resulting schema
    :param input_: text or input stream to parse
    :param default_base_: base URI for relative URI's in schema
    :return: ShExJ Schema object.  None if error.
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
    parser = ShexDocParser(default_base=default_base)
    parser.visit(parse_tree)

    return parser.context.schema


def rdf_suffix(fmt: str) -> str:
    """ Map the RDF format to the approproate suffix """
    for k, v in SUFFIX_FORMAT_MAP.items():
        if fmt == v:
            return k
    return 'rdf'


def genargs() -> ArgumentParser:
    """
    Create a command line parser
    :return: parser
    """
    parser = ArgumentParser()
    parser.add_argument("infile", help="Input ShExC specification")
    parser.add_argument("-nj", "--nojson", help="Do not produce json output", action="store_true")
    parser.add_argument("-nr", "--nordf", help="Do not produce rdf output", action="store_true")
    parser.add_argument("-j", "--jsonfile", help="Output ShExJ file (Default: {infile}.json)")
    parser.add_argument("-r", "--rdffile", help="Output ShExR file (Default: {infile}.{fmt suffix})")
    parser.add_argument("--context", help="Alternative @context")
    parser.add_argument("-f", "--format",
                        choices=list(set(x.name for x in rdflib_plugins(None, rdflib_Serializer)
                                         if '/' not in str(x.name))),
                        help="Output format (Default: turtle)", default="turtle")
    return parser


def generate(argv: List[str]) -> bool:
    """ 
    Transform ShExC to ShExJ
    :param argv: Command line arguments
    :return: True if successful
    """
    opts = genargs().parse_args(argv)
    filebase = os.path.dirname(opts.infile) + str(os.path.basename(opts.infile).rsplit('.', 1)[0])
    if opts.nojson:
        opts.jsonfile = None
    elif not opts.jsonfile:
        opts.jsonfile = filebase + ".json"
    if opts.nordf:
        opts.rdffile = None
    elif not opts.rdffile:
        opts.rdffile = filebase + "." + rdf_suffix(opts.format)
    if do_parse(opts.infile, opts.jsonfile, opts.rdffile, opts.format, opts.context):
        if opts.jsonfile:
            print("JSON output written to {}".format(opts.jsonfile))
        if opts.rdffile:
            print("{} output written to {}".format(opts.format, opts.rdffile))
        return True
    else:
        print("Conversion failed")
        return False
