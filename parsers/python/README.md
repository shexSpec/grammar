# PyShExC - The Python ShExC Parser
This package converts the Shape Expression Compact (ShExC) into Python JSON Schema Binding (pyjsg) objects. The pyjsg object representation can be used to:

* Produce equivalent ShEx definitions in ShEx JSON (ShExJ) notation
* Produce equivalent ShEx definitions in ShEx RDF (ShExR) notation
* Implement a Python ShEx interpreter

## History
* 0.4.0 -- Almost ready for 2.1 release.  9 test errors remain, all involving unicode
* 0.4.1 -- Added contextcache. See test_shexr for how to use
* 0.4.2 -- Tweaks in 4 and 8 byte unicode parsing
* 0.5.0 -- Passes all 2.1 tests
* 0.5.1 -- Fix requirements
* 0.5.2 -- Add URL fetch ability and fix UTF8-BOM issue
* 0.5.3 -- Factor out fetch ability for re-use in other packages
* 0.5.4 -- Fix to 0.5.2 -- low probability type matches have to be ignored

## Installation
```bash
> pip install PyShExC
```

## Usage

### Command Line
```text
usage: shexc_to_shexj [-h] [-nj] [-nr] [-j JSONFILE] [-r RDFFILE]
                      [--context CONTEXT]
                      [-f {trix,trig,ttl,n3,turtle,ntriples,xml,json-ld,nt11,pretty-xml,nquads,nt}]
                      infile

positional arguments:
  infile                Input ShExC specification

optional arguments:
  -h, --help            show this help message and exit
  -nj, --nojson         Do not produce json output
  -nr, --nordf          Do not produce rdf output
  -j JSONFILE, --jsonfile JSONFILE
                        Output ShExJ file (Default: {infile}.json)
  -r RDFFILE, --rdffile RDFFILE
                        Output ShExR file (Default: {infile}.{fmt suffix})
  --context CONTEXT     Alternative @context for json to RDF
  -f {trix,trig,ttl,n3,turtle,ntriples,xml,json-ld,nt11,pretty-xml,nquads,nt}, --format {trix,trig,ttl,n3,turtle,ntriples,xml,json-ld,nt11,pretty-xml,nquads,nt}
                        Output format (Default: turtle)
```
```bash
 > wget https://raw.githubusercontent.com/shexSpec/shexTest/master/schemas/FocusIRI2groupBnodeNested2groupIRIRef.shex -O test.shex
 > shexc_to_shexj test.shex  --context shex.jsonld
  JSON output written to test.json
  turtle output written to test.ttl
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

**test.ttl**
```rdf
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix shex: <http://www.w3.org/ns/shex#> .
@prefix xml: <http://www.w3.org/XML/1998/namespace> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<http://a.example/S1> a shex:ShapeAnd ;
    shex:shapeExprs ( [ a shex:NodeConstraint ;
                shex:nodeKind shex:iri ;
                shex:pattern "^https?://" ] [ a shex:Shape ;
                shex:expression [ a shex:EachOf ;
                        shex:expressions ( [ a shex:TripleConstraint ;
                                    shex:predicate <http://a.example/p1> ;
                                    shex:valueExpr [ a shex:NodeConstraint ;
                                            shex:datatype <http://a.example/dt1> ] ] [ a shex:TripleConstraint ;
                                    shex:predicate <http://a.example/p2> ;
                                    shex:valueExpr [ a shex:ShapeAnd ;
                                            shex:shapeExprs ( [ a shex:ShapeAnd ;
                                                        shex:shapeExprs ( [ a shex:NodeConstraint ;
                                                                    shex:nodeKind shex:bnode ] [ a shex:Shape ;
                                                                    shex:expression [ a shex:EachOf ;
                                                                            shex:expressions ( [ a shex:TripleConstraint ;

        shex:predicate <http://a.example/p3> ;

        shex:valueExpr [ a shex:NodeConstraint ;

                shex:nodeKind shex:literal ] ] [ a shex:TripleConstraint ;

        shex:max 1 ;

        shex:min 0 ;

        shex:predicate <http://a.example/p4> ;

        shex:valueExpr [ a shex:ShapeAnd ;

                shex:shapeExprs ( [ a shex:NodeConstraint ;

                            shex:nodeKind shex:iri ;

                            shex:pattern "^https?://" ] <http://a.example/S1> ) ] ] ) ] ] ) ] [ a shex:Shape ;
                                                        shex:closed true ;
                                                        shex:expression [ a shex:EachOf ;
                                                                shex:expressions ( [ a shex:TripleConstraint ;
                                                                            shex:predicate <http://a.example/p3> ] [ a shex:TripleConstraint ;
                                                                            shex:predicate <http://a.example/p4> ] ) ] ] ) ] ] ) ] ] ) .

[] a shex:Schema ;
    shex:shapes <http://a.example/S1> .
```