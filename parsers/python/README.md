# PyShExC - The Python ShExC Parser
This package converts the Shape Expression Compact (ShExC) into Python JSON Schema Binding (pyjsg) objects. The pyjsg object representation can be used to:

* Produce equivalent ShEx definitions in ShEx JSON (ShExJ) notation
* Produce equivalent ShEx definitions in ShEx RDF (ShExR) notation
* Implement a Python ShEx interpreter

## Installation
```bash
> pip install PyShExC
```

## Usage

### Command Line
```bash
> shexc_to_shexj -h
usage: shexc_to_shexj [-h] [-o OUTFILE] infile

positional arguments:
  infile                Input ShExC specification

optional arguments:
  -h, --help            show this help message and exit
  -o OUTFILE, --outfile OUTFILE
                        Output ShExJ file (Default: {infile}.json)
```
```bash
 > wget https://raw.githubusercontent.com/shexSpec/shexTest/master/schemas/FocusIRI2groupBnodeNested2groupIRIRef.shex -O test.shex
 > shexc_to_shexj test.shex
   Output written to test.json
 ```
**test.shex**
```
<http://a.example/S1> IRI /^https?:\/\// {
   <http://a.example/p1> <http://a.example/dt1> ;
   <http://a.example/p2> BNODE {
     <http://a.example/p3> LITERAL ;
     <http://a.example/p4> IRI /^https?:\/\// @<http://a.example/S1>?
   } AND CLOSED { <http://a.example/p3> . ; <http://a.example/p4> . }
}
```

**test.json**
```json
{
   "type": "Schema",
   "@context": "http://www.w3.org/ns/shex.jsonld",
   "shapes": [
      {
         "type": "ShapeAnd",
         "id": "http://a.example/S1",
         "shapeExprs": [
            {
               "type": "NodeConstraint",
               "nodeKind": "iri",
               "pattern": "^https?://"
            },
            {
               "type": "Shape",
               "expression": {
                  "type": "EachOf",
                  "expressions": [
                     {
                        "type": "TripleConstraint",
                        "predicate": "http://a.example/p1",
                        "valueExpr": {
                           "type": "NodeConstraint",
                           "datatype": "http://a.example/dt1"

                        }
                     },
                     {
                        "type": "TripleConstraint",
                        "predicate": "http://a.example/p2",
                        "valueExpr": {
                           "type": "ShapeAnd",
                           "shapeExprs": [
                              {
                                 "type": "ShapeAnd",
                                 "shapeExprs": [
                                    {
                                       "type": "NodeConstraint",
                                       "nodeKind": "bnode"
                                    },
                                    {
                                       "type": "Shape",
                                       "expression": {
                                          "type": "EachOf",
                                          "expressions": [
                                             {
                                                "type": "TripleConstraint",
                                                "predicate": "http://a.example/p3",
                                                "valueExpr": {
                                                   "type": "NodeConstraint",
                                                   "nodeKind": "literal"
                                                }
                                             },
                                             {
                                                "type": "TripleConstraint",
                                                "predicate": "http://a.example/p4",
                                                "valueExpr": {
                                                   "type": "ShapeAnd",
                                                   "shapeExprs": [
                                                      {
                                                         "type": "NodeConstraint",
                                                         "nodeKind": "iri",
                                                         "pattern": "^https?://"
                                                      },
                                                      "http://a.example/S1"
                                                   ]
                                                },
                                                "min": 0,
                                                "max": 1
                                             }
                                          ]
                                       }
                                    }
                                 ]
                              },
                              {
                                 "type": "Shape",
                                 "closed": "true",
                                 "expression": {
                                    "type": "EachOf",
                                    "expressions": [
                                       {
                                          "type": "TripleConstraint",
                                          "predicate": "http://a.example/p3"
                                       },
                                       {
                                          "type": "TripleConstraint",
                                          "predicate": "http://a.example/p4"
                                       }
                                    ]
                                 }
                              }
                           ]
                        }
                     }
                  ]
               }
            }
         ]
      }
   ]
}
```

