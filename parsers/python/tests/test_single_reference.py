import unittest

from jsonasobj import as_json

from pyshexc.parser_impl.generate_shexj import parse

shex = """<http://a.example/S0> @<http://a.example/S1> 
<http://a.example/S1> { <http://a.example/p1> . }"""

shexj = """{
   "type": "Schema",
   "shapes": [
      "http://a.example/S1",
      {
         "type": "Shape",
         "id": "http://a.example/S1",
         "expression": {
            "type": "TripleConstraint",
            "predicate": "http://a.example/p1"
         }
      }
   ]
}"""

class SingleReferenceTestCase(unittest.TestCase):
    """ Test to determine what this parser does with a single reference """
    def test_shex(self):
        schema = parse(shex)
        self.assertEqual(shexj, as_json(schema))


if __name__ == '__main__':
    unittest.main()
