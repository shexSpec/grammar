import unittest

import requests
from typing import Optional
from dict_compare import compare_dicts, json_filtr
from jsonasobj import loads as jao_loads
from pyjsg.jsglib.jsg import loads as jsg_loads
from pyjsg.jsglib.logger import Logger
from pyshexc.parser_impl.generate_shexj import parse
from ShExJSG import ShExJ

from tests.build_test_harness import ValidationTestCase

#
# Starting file name (with or without ".shex" suffix)
START_AT = ""

# False if you want to start somewhere in the middle
SINGLE_FILE = bool(START_AT)

# Notes:
#   you can use shexj._as_json_dumps() to print all or part of a ShEx Schema
#   you can use "ctx.getText()" to get the span of any parser context

V2_1_IMPORT = "Uses ShEx 2.1 Import feature"
INCONSISTENT_STEM_WORK = "Untyped literal stems"
LONG_UNICODE_LITERALS = "ANTLR does not support unicode literals > 4 hex digits"

skip = {
    "1dotIMPORT1dot.shex": V2_1_IMPORT,
    "1valExprRef-IV1.shex": V2_1_IMPORT,
    "1valExprRefbnode-IV1.shex": V2_1_IMPORT,
    "2EachInclude1-IS2.shex": V2_1_IMPORT,
    "2RefS1-IS2.shex": V2_1_IMPORT,
    "2RefS1-Icirc.shex": V2_1_IMPORT,
    "2RefS2-IS1.shex": V2_1_IMPORT,
    "2RefS2-Icirc.shex": V2_1_IMPORT,
    "3circRefS1-IS2-IS3-IS3.shex": V2_1_IMPORT,
    "3circRefS1-IS2-IS3.shex": V2_1_IMPORT,
    "3circRefS1-IS23.shex": V2_1_IMPORT,
    "3circRefS1-Icirc.shex": V2_1_IMPORT,
    "3circRefS123-Icirc.shex": V2_1_IMPORT,
    "3circRefS2-IS3.shex": V2_1_IMPORT,
    "3circRefS2-Icirc.shex": V2_1_IMPORT,
    "3circRefS3-IS12.shex": V2_1_IMPORT,
    "3circRefS3-Icirc.shex": V2_1_IMPORT,
    "1val1STRING_LITERAL1_with_UTF8_boundaries.shex": LONG_UNICODE_LITERALS,
    "1refbnode_with_spanning_PN_CHARS_BASE1.shex": LONG_UNICODE_LITERALS,
    "1literalPattern_with_REGEXP_escapes_bare.shex": LONG_UNICODE_LITERALS,
    "1val1STRING_LITERAL1_with_ECHAR_escapes.shex": LONG_UNICODE_LITERALS,
    "1literalPattern_with_REGEXP_escapes.shex": LONG_UNICODE_LITERALS,
    "1literalPattern_with_UTF8_boundaries.shex": LONG_UNICODE_LITERALS,
}


class BasicParserTestCase(ValidationTestCase):
    pass


BasicParserTestCase.repo_url = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
# BasicParserTestCase.repo_url = "(path to git)/shexSpec/shexTest/schemas"
BasicParserTestCase.file_suffix = ".shex"
BasicParserTestCase.start_at = START_AT if not START_AT or START_AT.endswith('.shex') else START_AT + '.shex'
BasicParserTestCase.single_file = SINGLE_FILE

BasicParserTestCase.skip = list(skip.keys())


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
    if ':' in json_url:
        resp = requests.get(json_url)
        if not resp.ok:
            return False
        json_text = resp.text
    else:
        with open(json_url) as f:
            json_text = f.read()
    d1 = jao_loads(json_text)
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
    shexj['@context'] = "http://www.w3.org/ns/shex.jsonld"
    if shexj is None:
        return False
    shex_obj = jsg_loads(shexj._as_json, ShExJ)
    if not shex_obj._is_valid(logger):
        print("File: {} - ".format(input_fname))
        print(log.log)
        return False
    elif not compare_json(input_fname, shex_obj._as_json, log):
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
    if ':' in download_url:
        resp = requests.get(download_url)
        if resp.ok:
            return validate_shexc(resp.text, download_url)
        else:
            print("Error {}: {}".format(resp.status_code, resp.reason))
            return False
    else:
        with open(download_url) as f:
            return validate_shexc(f.read(), download_url)


BasicParserTestCase.validation_function = validate_file
BasicParserTestCase.build_test_harness()


if __name__ == '__main__':
    unittest.main()
