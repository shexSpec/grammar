import unittest

from pyshexc.ShExC import ShExC
from tests.utils.simple_shex_test import SimpleShexTestCase

shexc = """abstract <A>      /sA.........$/
              and { <p> [0 1 2 3 4 5 6 7 8 9]+}
         <B> -<A> /s.B........$/
              and { <p> [0 1   3 4 5 6 7 8 9]+}
"""


class ShExCTestCase(SimpleShexTestCase):
    def test_shexc(self):
        base = 'http://example.org/'
        shex = self.shexc_to_shexj(shexc, base)

        # ShExJ --> ShExC
        shexc_rev = str(ShExC(shex, base))
        print(shexc_rev)
        shex_rev = self.shexc_to_shexj(shexc_rev, base)
        rslt, log = self.compare_shexj(shex, shex_rev)
        if not rslt:
            print(log.getvalue())
        self.assertTrue(rslt)


if __name__ == '__main__':
    unittest.main()
