
import unittest
from contextlib import redirect_stdout
from io import StringIO

from ShExJSG import ShExJ
from dict_compare import compare_dicts, json_filtr
from jsonasobj import as_json, loads

from pyshexc.parser_impl.generate_shexj import parse

shexc = """
<J>
  ( ( /s......G....$/
      {
        <p> [0 1 2 3 5 6 7 8 9]+ ON SHAPE EXPRESSION
          (
            ( /sA..........$/ 
              {
                <p> [ 0 1 2 3 4 5 6 7 8 9 ] + 
              }
            )
            AND ( /s.B.........$/ 
                  {
                    <p> [ 0 1 3 4 5 6 7 8 9 ] + 
                  }
                )
          ) AND (
            /s..C........$/ 
            {
              <p> [ 0 1 2 4 5 6 7 8 9 ] + 
            }
          ) AND (
            /s...D.......$/ 
            { }
          );
        <p> [2] ON SHAPE EXPRESSION
          /s....E......$/ 
          {
          };
        <p> [3] ON SHAPE EXPRESSION
          /s.....F.....$/ 
          {
          };
        <p> [4]
      }
    )
    AND (
      /s.......H...$/
      {
        <p> [0 1 2 3 4 5 8 9]*
      }
    )
  ) AND (
    /s.........J.$/
    {
      <q> [5] ON SHAPE EXPRESSION
        /s........I..$/ 
        {
        };
      <q> [6]
    }
  )
"""

shexj = """
{
  "@context": "http://www.w3.org/ns/shex.jsonld",
  "shapes": [
    {
      "id": "http://localhost/shexSpec/shex.js/examples/inheritance/J",
      "shapeExprs": [
        {
          "shapeExprs": [
            {
              "shapeExprs": [
                {
                  "pattern": "s......G....$",
                  "type": "NodeConstraint"
                },
                {
                  "expression": {
                    "expressions": [
                      {
                        "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                        "valueExpr": {
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
                          ],
                          "type": "NodeConstraint"
                        },
                        "min": 1,
                        "max": -1,
                        "onShapeExpression": {
                          "shapeExprs": [
                            {
                              "shapeExprs": [
                                {
                                  "pattern": "sA..........$",
                                  "type": "NodeConstraint"
                                },
                                {
                                  "expression": {
                                    "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                                    "valueExpr": {
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
                                      ],
                                      "type": "NodeConstraint"
                                    },
                                    "min": 1,
                                    "max": -1,
                                    "type": "TripleConstraint"
                                  },
                                  "type": "Shape"
                                }
                              ],
                              "type": "ShapeAnd"
                            },
                            {
                              "shapeExprs": [
                                {
                                  "pattern": "s.B.........$",
                                  "type": "NodeConstraint"
                                },
                                {
                                  "expression": {
                                    "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                                    "valueExpr": {
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
                                      ],
                                      "type": "NodeConstraint"
                                    },
                                    "min": 1,
                                    "max": -1,
                                    "type": "TripleConstraint"
                                  },
                                  "type": "Shape"
                                }
                              ],
                              "type": "ShapeAnd"
                            },
                            {
                              "pattern": "s..C........$",
                              "type": "NodeConstraint"
                            },
                            {
                              "expression": {
                                "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                                "valueExpr": {
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
                                  ],
                                  "type": "NodeConstraint"
                                },
                                "min": 1,
                                "max": -1,
                                "type": "TripleConstraint"
                              },
                              "type": "Shape"
                            },
                            {
                              "pattern": "s...D.......$",
                              "type": "NodeConstraint"
                            },
                            {
                              "type": "Shape"
                            }
                          ],
                          "type": "ShapeAnd"
                        },
                        "type": "TripleConstraint"
                      },
                      {
                        "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                        "valueExpr": {
                          "values": [
                            {
                              "value": "2",
                              "type": "http://www.w3.org/2001/XMLSchema#integer"
                            }
                          ],
                          "type": "NodeConstraint"
                        },
                        "onShapeExpression": {
                          "shapeExprs": [
                            {
                              "pattern": "s....E......$",
                              "type": "NodeConstraint"
                            },
                            {
                              "type": "Shape"
                            }
                          ],
                          "type": "ShapeAnd"
                        },
                        "type": "TripleConstraint"
                      },
                      {
                        "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                        "valueExpr": {
                          "values": [
                            {
                              "value": "3",
                              "type": "http://www.w3.org/2001/XMLSchema#integer"
                            }
                          ],
                          "type": "NodeConstraint"
                        },
                        "onShapeExpression": {
                          "shapeExprs": [
                            {
                              "pattern": "s.....F.....$",
                              "type": "NodeConstraint"
                            },
                            {
                              "type": "Shape"
                            }
                          ],
                          "type": "ShapeAnd"
                        },
                        "type": "TripleConstraint"
                      },
                      {
                        "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                        "valueExpr": {
                          "values": [
                            {
                              "value": "4",
                              "type": "http://www.w3.org/2001/XMLSchema#integer"
                            }
                          ],
                          "type": "NodeConstraint"
                        },
                        "type": "TripleConstraint"
                      }
                    ],
                    "type": "EachOf"
                  },
                  "type": "Shape"
                }
              ],
              "type": "ShapeAnd"
            },
            {
              "shapeExprs": [
                {
                  "pattern": "s.......H...$",
                  "type": "NodeConstraint"
                },
                {
                  "expression": {
                    "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/p",
                    "valueExpr": {
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
                      ],
                      "type": "NodeConstraint"
                    },
                    "min": 0,
                    "max": -1,
                    "type": "TripleConstraint"
                  },
                  "type": "Shape"
                }
              ],
              "type": "ShapeAnd"
            }
          ],
          "type": "ShapeAnd"
        },
        {
          "shapeExprs": [
            {
              "pattern": "s.........J.$",
              "type": "NodeConstraint"
            },
            {
              "expression": {
                "expressions": [
                  {
                    "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/q",
                    "valueExpr": {
                      "values": [
                        {
                          "value": "5",
                          "type": "http://www.w3.org/2001/XMLSchema#integer"
                        }
                      ],
                      "type": "NodeConstraint"
                    },
                    "onShapeExpression": {
                      "shapeExprs": [
                        {
                          "pattern": "s........I..$",
                          "type": "NodeConstraint"
                        },
                        {
                          "type": "Shape"
                        }
                      ],
                      "type": "ShapeAnd"
                    },
                    "type": "TripleConstraint"
                  },
                  {
                    "predicate": "http://localhost/shexSpec/shex.js/examples/inheritance/q",
                    "valueExpr": {
                      "values": [
                        {
                          "value": "6",
                          "type": "http://www.w3.org/2001/XMLSchema#integer"
                        }
                      ],
                      "type": "NodeConstraint"
                    },
                    "type": "TripleConstraint"
                  }
                ],
                "type": "EachOf"
              },
              "type": "Shape"
            }
          ],
          "type": "ShapeAnd"
        }
      ],
      "type": "ShapeAnd"
    }
  ],
  "type": "Schema"
}"""


class OnShapeExpressionTestCase(unittest.TestCase):
    def test_erics_example(self):
        shex: ShExJ.Schema = \
            parse(shexc, default_base="http://localhost/shexSpec/shex.js/examples/inheritance/")
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

