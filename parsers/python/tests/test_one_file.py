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
import unittest

import requests
from dict_compare import compare_dicts, json_filtr
from jsonasobj import loads as jao_loads
from pyjsg.jsglib.jsg import loads as jsg_loads
from pyjsg.jsglib.logger import Logger
from typing import Optional

from pyshexc.parser_impl import parse
from pyshexc.shexj import ShExJ
from tests.memlogger import MemLogger

# This module is used for one-time tests on individual failing files or snippets of text.
# To test a file:  Enter the file URL below
# To test text:  Paste it in as "text"

FILE_NAME="https://raw.githubusercontent.com/shexSpec/shexTest/master/schemas/1val1STRING_LITERAL1_with_all_punctuation.shex"
FILE_NAME = None
text = r"""<http://a.example/S1> BNODE {
   <http://a.example/p2> BNODE {
     <http://a.example/p3> LITERAL
   } AND CLOSED { }
}
"""


def compare_json(shex_url: str, shex_json: str, log: Logger) -> bool:
    """ Compare the json equivalent of shex_fname to shex_json """
    json_url = shex_url.rsplit(".", 1)[0] + ".json"
    resp = requests.get(json_url)
    if not resp.ok:
        return False
    d1 = jao_loads(resp.text)
    d2 = jao_loads(shex_json)
    if not compare_dicts(d1._as_dict, d2._as_dict, d1name="expected", d2name="actual  ", file=log, filtr=json_filtr):
        print(d2._as_json_dumps())
        return False
    return True


def has_invalid_chars(text: str) -> bool:
    """ The ANTLR4 parser does not deal with utf characters > 4 digits.
        See: http://stackoverflow.com/questions/35938284/how-do-i-specify-a-unicode-literal-that-requires-more-than-four-hex-digits-in-an#35939479
        Also look at getCharValueFromCharInGrammarLiteral routine in tool/src/org/antlr4/v4/misc/CharSupport.java
    """
    return not all(ord(c) < 0xFFFF for c in text)


def validate_shexc(shexc_str: str, input_fname: str) -> bool:
    """
    Validate json_str against ShEx Schema
    :param json_str: String to validate
    :param input_fname: Name of source file for error reporting
    :return: True if pass
    """
    if has_invalid_chars(shexc_str):
        print("ANTLR does not support unicode literals > 4 hex digits.")
        return False
    log = MemLogger('\t')
    logger = Logger(log)
    shexj = parse(shexc_str, input_fname)
    if shexj is None:
        return False
    shex_obj = jsg_loads(shexj, ShExJ)
    if not shex_obj._is_valid(logger):
        print("File: {} - ".format(input_fname))
        print(log.log)
        return False
    elif not input_fname:
        print(shexj)
    elif input_fname and not compare_json(input_fname, shex_obj._as_json, logger):
        print("File: {} - ".format(input_fname))
        print(log.log)
        return False
    return True


def validate_file(download_url: Optional[str] = None, inline_text: Optional[str] = None) -> bool:
    if download_url:
        fname = download_url.rsplit('/', 1)[1]
        print("Testing {}".format(download_url))
        resp = requests.get(download_url)
        if resp.ok:
            inline_text = resp.text
        else:
            print("Error {}: {}".format(resp.status_code, resp.reason))
            return False
    else:
        download_url = ""
        inline_text = text

    return validate_shexc(inline_text, download_url)


class ShExCOneFileTestCase(unittest.TestCase):

    def test_one_file(self):
        self.assertTrue(validate_file(FILE_NAME, text))

if __name__ == '__main__':
    unittest.main()
