import unittest

from tests.utils.simple_shex_test import SimpleShexTestCase

shexc = """<http://a.example/S1> {
   <http://a.example/p1> [@fr]
}
"""

shexj = """{
  "@context": "http://www.w3.org/ns/shex.jsonld",
  "type": "Schema",
  "shapes": [
    {
      "id": "http://a.example/S1",
      "type": "Shape",
      "expression": {
        "type": "TripleConstraint",
        "predicate": "http://a.example/p1",
        "valueExpr": {
          "type": "NodeConstraint",
          "values": [
            {
              "type": "Language",
              "languageTag": "fr"
            }
          ]
        }
      }
    }
  ]
}"""


class LanguageValuesetTestCase(SimpleShexTestCase):
    def test_language_valueset(self):
        self.shex_test(shexc, shexj)


if __name__ == '__main__':
    unittest.main()
