
import unittest
from contextlib import redirect_stdout
from io import StringIO

from ShExJSG import ShExJ
from dict_compare import compare_dicts, json_filtr
from jsonasobj import as_json, loads

from pyshexc.parser_impl.generate_shexj import parse

shexc = """abstract <A>           /sA..........$/
                   and { <p> [0 1 2 3 4 5 6 7 8 9]+}
abstract <B> -<A>      /s.B.........$/
                   and { <p> [0 1   3 4 5 6 7 8 9]+}
         <C>           /s..C........$/
                   and { <p> [0 1 2   4 5 6 7 8 9]+}
abstract <D> -<B> -<C> /s...D.......$/
                   and { <p> [0 1 2 3   5 6 7 8 9]+}
abstract <E>           /s....E......$/
                   and &<D>
                       { <p> [    2              ] }
         <F>           /s.....F.....$/
                   and { <p> [      3            ] }
abstract <G>           /s......G....$/
                   and &<E> &<F>
                       { <p> [        4          ] }
abstract <H> -<G>      /s.......H...$/
                   and { <p> [0 1 2 3 4 5     8 9]*}
         <I>           /s........I..$/
                   and { <q> [          5        ] }
         <J> -<H>      /s.........J.$/
                   and &<I>
                       { <q> [            6      ] }
"""

shexj = """{
  "type": "Schema",
  "shapes": [
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/A",
      "type": "ShapeDecl",
      "abstract": true,
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "sA..........$"
          },
          {
            "type": "Shape",
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "0",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "1",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "2",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "3",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "4",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "5",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "6",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "7",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "8",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "9",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              },
              "min": 1,
              "max": -1
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/B",
      "type": "ShapeDecl",
      "abstract": true,
      "restricts": [
        "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/A"
      ],
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s.B.........$"
          },
          {
            "type": "Shape",
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "0",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "1",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "3",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "4",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "5",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "6",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "7",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "8",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "9",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              },
              "min": 1,
              "max": -1
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/C",
      "type": "ShapeAnd",
      "shapeExprs": [
        {
          "type": "NodeConstraint",
          "pattern": "s..C........$"
        },
        {
          "type": "Shape",
          "expression": {
            "type": "TripleConstraint",
            "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
            "valueExpr": {
              "type": "NodeConstraint",
              "values": [
                {
                  "value": "0",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "1",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "2",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "4",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "5",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "6",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "7",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "8",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                },
                {
                  "value": "9",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                }
              ]
            },
            "min": 1,
            "max": -1
          }
        }
      ]
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/D",
      "type": "ShapeDecl",
      "abstract": true,
      "restricts": [
        "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/B",
        "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/C"
      ],
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s...D.......$"
          },
          {
            "type": "Shape",
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "0",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "1",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "2",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "3",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "5",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "6",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "7",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "8",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "9",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              },
              "min": 1,
              "max": -1
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/E",
      "type": "ShapeDecl",
      "abstract": true,
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s....E......$"
          },
          {
            "type": "Shape",
            "extends": [
              "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/D"
            ],
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "2",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              }
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/F",
      "type": "ShapeAnd",
      "shapeExprs": [
        {
          "type": "NodeConstraint",
          "pattern": "s.....F.....$"
        },
        {
          "type": "Shape",
          "expression": {
            "type": "TripleConstraint",
            "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
            "valueExpr": {
              "type": "NodeConstraint",
              "values": [
                {
                  "value": "3",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                }
              ]
            }
          }
        }
      ]
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/G",
      "type": "ShapeDecl",
      "abstract": true,
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s......G....$"
          },
          {
            "type": "Shape",
            "extends": [
              "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/E",
              "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/F"
            ],
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "4",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              }
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/H",
      "type": "ShapeDecl",
      "abstract": true,
      "restricts": [
        "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/G"
      ],
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s.......H...$"
          },
          {
            "type": "Shape",
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "0",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "1",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "2",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "3",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "4",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "5",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "8",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  },
                  {
                    "value": "9",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              },
              "min": 0,
              "max": -1
            }
          }
        ]
      }
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/I",
      "type": "ShapeAnd",
      "shapeExprs": [
        {
          "type": "NodeConstraint",
          "pattern": "s........I..$"
        },
        {
          "type": "Shape",
          "expression": {
            "type": "TripleConstraint",
            "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
            "valueExpr": {
              "type": "NodeConstraint",
              "values": [
                {
                  "value": "5",
                  "type": "http://www.w3.org/2001/XMLSchema#integer"
                }
              ]
            }
          }
        }
      ]
    },
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/J",
      "type": "ShapeDecl",
      "restricts": [
        "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/H"
      ],
      "shapeExpr": {
        "type": "ShapeAnd",
        "shapeExprs": [
          {
            "type": "NodeConstraint",
            "pattern": "s.........J.$"
          },
          {
            "type": "Shape",
            "extends": [
              "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/I"
            ],
            "expression": {
              "type": "TripleConstraint",
              "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
              "valueExpr": {
                "type": "NodeConstraint",
                "values": [
                  {
                    "value": "6",
                    "type": "http://www.w3.org/2001/XMLSchema#integer"
                  }
                ]
              }
            }
          }
        ]
      }
    }
  ],
  "@context": "http://www.w3.org/ns/shex.jsonld"
}"""


class RestrictExtendTestCase(unittest.TestCase):
    def test_erics_example(self):
        shex: ShExJ.Schema = \
            parse(shexc, default_base="https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/")
        shex['@context'] = "http://www.w3.org/ns/shex.jsonld"
        d1 = loads(shexj)
        d2 = loads(as_json(shex))

        log = StringIO()
        with redirect_stdout(log):
            rval = compare_dicts(d1._as_dict, d2._as_dict, d1name="expected", d2name="actual  ", filtr=json_filtr)
        if not rval:
            print(log.getvalue())
            print(as_json(d2, indent="  "))
        self.assertTrue(rval)


if __name__ == '__main__':
    unittest.main()
