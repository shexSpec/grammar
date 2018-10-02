import unittest

from tests.utils.simple_shex_test import SimpleShexTestCase

shexc = """IMPORT <1dot>

<http://a.example/S2> { <http://a.example/p2> . }
"""

shexj = """{
  "@context": "http://www.w3.org/ns/shex.jsonld",
  "type": "Schema",
  "imports": ["1dot"],
  "shapes": [
    {
      "id": "http://a.example/S2",
      "type": "Shape",
      "expression": {
        "type": "TripleConstraint",
        "predicate": "http://a.example/p2"
      }
    }
  ]
}
"""


class ImportTestCase(SimpleShexTestCase):

    def test_import(self):
        self.shex_test(shexc, shexj)


if __name__ == '__main__':
    unittest.main()
