import unittest
from contextlib import redirect_stdout
from io import StringIO
from typing import Tuple, Union

from ShExJSG import ShExJ
from dict_compare import compare_dicts, json_filtr
from jsonasobj import loads, as_json

from pyshexc.ShExC import ShExC
from pyshexc.parser_impl.generate_shexj import parse


class SimpleShexTestCase(unittest.TestCase):

    @staticmethod
    def compare_shexj(shex: Union[ShExJ.Schema, str], shexj: Union[ShExJ.Schema, str]) -> Tuple[bool, StringIO]:
        d1 = loads(as_json(shex) if isinstance(shex, ShExJ.Schema) else shex)
        d2 = loads(as_json(shexj) if isinstance(shexj, ShExJ.Schema) else shexj)

        log = StringIO()
        with redirect_stdout(log):
            return compare_dicts(d1._as_dict, d2._as_dict, d1name="expected", d2name="actual  ", filtr=json_filtr), log

    def shexc_to_shexj(self, shexc, base=None) -> ShExJ.Schema:
        shex: ShExJ.Schema = parse(shexc, default_base=base)
        self.assertIsNotNone(shex, "Compile error")
        shex['@context'] = "http://www.w3.org/ns/shex.jsonld"
        return shex

    def shex_test(self, shexc, shexj, base=None):
        # ShExC --> ShExJ
        shex = self.shexc_to_shexj(shexc, base)
        rslt, log = self.compare_shexj(shex, shexj)
        msg = "ShExC to ShExJ Conversion Error"
        if not rslt:
            print(f"***** {msg} *****")
            print(log.getvalue())
            print("\n***** Actual ShExJ *****")
            print(as_json(shex))
        self.assertTrue(rslt, msg)

        # ShExJ --> ShExC
        shexc_rev = str(ShExC(shex, base))
        shex_rev = self.shexc_to_shexj(shexc_rev, base)
        rslt, log = self.compare_shexj(shex_rev, shexj)
        msg = "ShExJ to ShExC Conversion Error"
        print(shexc_rev)
        if not rslt:
            print(f"***** {msg} *****")
            print(log.getvalue())
            print("\n***** Actual ShExC *****")
            print(shexc_rev)
        self.assertTrue(rslt, msg)


if __name__ == '__main__':
    unittest.main()
