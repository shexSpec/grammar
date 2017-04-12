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
from typing import Optional
from dict_compare import compare_dicts, json_filtr
from jsonasobj import loads as jao_loads
from pyjsg.jsglib.jsg import loads as jsg_loads
from pyjsg.jsglib.logger import Logger
from pyshexc.parser_impl.generate_shexj import parse
from pyshexc.shexj import ShExJ

from tests.build_test_harness import ValidationTestCase

ValidationTestCase.repo_url = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
# ValidationTestCase.start_at = "1literalPattern_with_REGEXP_escapes.shex"
ValidationTestCase.skip = ['coverage.json', 'manifest.json']
ValidationTestCase.file_suffix = ".shex"


class MemLogger:
    def __init__(self, prefix: Optional[str] = None):
        self.prefix = prefix
        self.log = ""

    def write(self, txt):
        self.log += self.prefix + txt


def compare_json(shex_url: str, shex_json: str, log: Logger) -> bool:
    """
    Compare the JSON generated from shex_url to the JSON in the target directory
    :param shex_url: URL where we got the ShExC
    :param shex_json: ShExJ equivalent of ShExC
    :param log: Where comparison errors are recorded
    :return: True if they match, false otherwise.  If no match, the offending string is printed
    """
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
    return not all(ord(c) <= 0xFFFF for c in text)


def validate_shexc(shexc_str: str, input_fname: str) -> bool:
    """
    Validate json_str against ShEx Schema
    :param shexc_str: String to validate
    :param input_fname: Name of source file for error reporting
    :return: True if pass
    """
    if has_invalid_chars(shexc_str):
        print("ANTLR does not support unicode literals > 4 hex digits.")
        return False
    log = MemLogger('\t')
    logger = Logger(log)
    shexj = parse(shexc_str)
    if shexj is None:
        return False
    shex_obj = jsg_loads(shexj._as_json, ShExJ)
    if not shex_obj._is_valid(logger):
        print("File: {} - ".format(input_fname))
        print(log.log)
        return False
    elif not compare_json(input_fname, shex_obj._as_json, logger):
        print("File: {} - ".format(input_fname))
        print(log.log)
        return False
    return True


def validate_file(download_url: str) -> bool:
    """
    Parse and validate the ShExC file in download_url
    :param download_url: ShExC file
    :return: True if success
    """
    resp = requests.get(download_url)
    if resp.ok:
        return validate_shexc(resp.text, download_url)
    else:
        print("Error {}: {}".format(resp.status_code, resp.reason))
        return False

ValidationTestCase.validation_function = validate_file
ValidationTestCase.build_test_harness()

if __name__ == '__main__':
    unittest.main()
