
import unittest
from rdflib import Namespace

from ShExJSG import ShExJ

from pyshexc.parser_impl.generate_shexj import parse
from tests import git_branch

BASE = Namespace(f"https://raw.githubusercontent.com/shexSpec/shexTest/{git_branch}/validation/")
FOO = Namespace("/some/location/file/")
EX = Namespace("http://example.org/")


class TestIssue2(unittest.TestCase):
    def test_no_base(self):
        shex_str = '<S1> {<p1> [<o1>]}'
        shex: ShExJ.Schema = parse(shex_str)
        self.assertEqual("S1", str(shex.shapes[0].id))
        self.assertEqual("p1", str(shex.shapes[0].expression.predicate))
        self.assertEqual("o1", str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_default_base(self):
        shex_str = '<S1> {<p1> [<o1>]}'
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(BASE.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(BASE.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(BASE.o1), str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_explicit_base(self):
        shex_str = f'BASE <{str(FOO)}>\n<S1> {{<p1> [<o1>]}}'
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(FOO.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(FOO.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(FOO.o1), str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_explicit_uris(self):
        shex_str = f"""
BASE <{str(FOO)}>
PREFIX ex: <{EX}>

ex:S1 {{ex:p1 [ex:o1]}}"""
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(EX.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(EX.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(EX.o1), str(shex.shapes[0].expression.valueExpr.values[0]))


if __name__ == '__main__':
    unittest.main()
