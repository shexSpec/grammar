
import unittest

from tests.utils.simple_shex_test import SimpleShexTestCase

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

# The ShExJ below is what the shex.js software emits -- it does no "and" optimization
# shexj = """
# {
#   "type": "Schema",
#   "shapes": [
#     {
#       "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/J",
#       "type": "ShapeAnd",
#       "shapeExprs": [
#         {
#           "type": "ShapeAnd",
#           "shapeExprs": [
#             {
#               "type": "ShapeAnd",
#               "shapeExprs": [
#                 {
#                   "type": "NodeConstraint",
#                   "pattern": "s......G....$"
#                 },
#                 {
#                   "type": "Shape",
#                   "expression": {
#                     "type": "EachOf",
#                     "expressions": [
#                       {
#                         "type": "TripleConstraint",
#                         "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                         "valueExpr": {
#                           "type": "NodeConstraint",
#                           "values": [
#                             {
#                               "value": "0",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "1",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "2",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "3",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "5",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "6",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "7",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "8",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             },
#                             {
#                               "value": "9",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             }
#                           ]
#                         },
#                         "min": 1,
#                         "max": -1,
#                         "onShapeExpression": {
#                           "type": "ShapeAnd",
#                           "shapeExprs": [
#                             {
#                               "type": "ShapeAnd",
#                               "shapeExprs": [
#                                 {
#                                   "type": "NodeConstraint",
#                                   "pattern": "sA..........$"
#                                 },
#                                 {
#                                   "type": "Shape",
#                                   "expression": {
#                                     "type": "TripleConstraint",
#                                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                                     "valueExpr": {
#                                       "type": "NodeConstraint",
#                                       "values": [
#                                         {
#                                           "value": "0",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "1",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "2",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "3",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "4",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "5",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "6",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "7",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "8",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "9",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         }
#                                       ]
#                                     },
#                                     "min": 1,
#                                     "max": -1
#                                   }
#                                 }
#                               ]
#                             },
#                             {
#                               "type": "ShapeAnd",
#                               "shapeExprs": [
#                                 {
#                                   "type": "NodeConstraint",
#                                   "pattern": "s.B.........$"
#                                 },
#                                 {
#                                   "type": "Shape",
#                                   "expression": {
#                                     "type": "TripleConstraint",
#                                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                                     "valueExpr": {
#                                       "type": "NodeConstraint",
#                                       "values": [
#                                         {
#                                           "value": "0",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "1",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "3",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "4",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "5",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "6",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "7",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "8",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "9",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         }
#                                       ]
#                                     },
#                                     "min": 1,
#                                     "max": -1
#                                   }
#                                 }
#                               ]
#                             },
#                             {
#                               "type": "ShapeAnd",
#                               "shapeExprs": [
#                                 {
#                                   "type": "NodeConstraint",
#                                   "pattern": "s..C........$"
#                                 },
#                                 {
#                                   "type": "Shape",
#                                   "expression": {
#                                     "type": "TripleConstraint",
#                                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                                     "valueExpr": {
#                                       "type": "NodeConstraint",
#                                       "values": [
#                                         {
#                                           "value": "0",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "1",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "2",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "4",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "5",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "6",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "7",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "8",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         },
#                                         {
#                                           "value": "9",
#                                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                                         }
#                                       ]
#                                     },
#                                     "min": 1,
#                                     "max": -1
#                                   }
#                                 }
#                               ]
#                             },
#                             {
#                               "type": "ShapeAnd",
#                               "shapeExprs": [
#                                 {
#                                   "type": "NodeConstraint",
#                                   "pattern": "s...D.......$"
#                                 },
#                                 {
#                                   "type": "Shape"
#                                 }
#                               ]
#                             }
#                           ]
#                         }
#                       },
#                       {
#                         "type": "TripleConstraint",
#                         "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                         "valueExpr": {
#                           "type": "NodeConstraint",
#                           "values": [
#                             {
#                               "value": "2",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             }
#                           ]
#                         },
#                         "onShapeExpression": {
#                           "type": "ShapeAnd",
#                           "shapeExprs": [
#                             {
#                               "type": "NodeConstraint",
#                               "pattern": "s....E......$"
#                             },
#                             {
#                               "type": "Shape"
#                             }
#                           ]
#                         }
#                       },
#                       {
#                         "type": "TripleConstraint",
#                         "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                         "valueExpr": {
#                           "type": "NodeConstraint",
#                           "values": [
#                             {
#                               "value": "3",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             }
#                           ]
#                         },
#                         "onShapeExpression": {
#                           "type": "ShapeAnd",
#                           "shapeExprs": [
#                             {
#                               "type": "NodeConstraint",
#                               "pattern": "s.....F.....$"
#                             },
#                             {
#                               "type": "Shape"
#                             }
#                           ]
#                         }
#                       },
#                       {
#                         "type": "TripleConstraint",
#                         "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                         "valueExpr": {
#                           "type": "NodeConstraint",
#                           "values": [
#                             {
#                               "value": "4",
#                               "type": "http://www.w3.org/2001/XMLSchema#integer"
#                             }
#                           ]
#                         }
#                       }
#                     ]
#                   }
#                 }
#               ]
#             },
#             {
#               "type": "ShapeAnd",
#               "shapeExprs": [
#                 {
#                   "type": "NodeConstraint",
#                   "pattern": "s.......H...$"
#                 },
#                 {
#                   "type": "Shape",
#                   "expression": {
#                     "type": "TripleConstraint",
#                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
#                     "valueExpr": {
#                       "type": "NodeConstraint",
#                       "values": [
#                         {
#                           "value": "0",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "1",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "2",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "3",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "4",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "5",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "8",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         },
#                         {
#                           "value": "9",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         }
#                       ]
#                     },
#                     "min": 0,
#                     "max": -1
#                   }
#                 }
#               ]
#             }
#           ]
#         },
#         {
#           "type": "ShapeAnd",
#           "shapeExprs": [
#             {
#               "type": "NodeConstraint",
#               "pattern": "s.........J.$"
#             },
#             {
#               "type": "Shape",
#               "expression": {
#                 "type": "EachOf",
#                 "expressions": [
#                   {
#                     "type": "TripleConstraint",
#                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
#                     "valueExpr": {
#                       "type": "NodeConstraint",
#                       "values": [
#                         {
#                           "value": "5",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         }
#                       ]
#                     },
#                     "onShapeExpression": {
#                       "type": "ShapeAnd",
#                       "shapeExprs": [
#                         {
#                           "type": "NodeConstraint",
#                           "pattern": "s........I..$"
#                         },
#                         {
#                           "type": "Shape"
#                         }
#                       ]
#                     }
#                   },
#                   {
#                     "type": "TripleConstraint",
#                     "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
#                     "valueExpr": {
#                       "type": "NodeConstraint",
#                       "values": [
#                         {
#                           "value": "6",
#                           "type": "http://www.w3.org/2001/XMLSchema#integer"
#                         }
#                       ]
#                     }
#                   }
#                 ]
#               }
#             }
#           ]
#         }
#       ]
#     }
#   ],
#   "@context": "http://www.w3.org/ns/shex.jsonld"
# }
# """

# This is the flattened equivalent of the above
shexj = """{
  "@context": "http://www.w3.org/ns/shex.jsonld",
  "shapes": [
    {
      "id": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/J",
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
                        "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                                    "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                                    "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                                "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                        "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                        "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                        "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                    "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/p",
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
                    "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
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
                    "predicate": "https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/q",
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
}
"""


class OnShapeExpressionTestCase(SimpleShexTestCase):
    @unittest.skipIf(True, "ON shapexpression not in 2.1 release")
    def test_erics_example(self):
        self.shex_test(shexc, shexj,
                       base="https://rawgit.com/shexSpec/shex.js/on-shape-expression/examples/inheritance/")


if __name__ == '__main__':
    unittest.main()
