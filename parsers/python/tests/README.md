# PyShExC tests

## Requirements
These tests require the [requests](http://docs.python-requests.org/en/master/ and [yadict-compare](https://github.com/hsolbrig/dict_compare) modules.  They can be installed by:
```bash
> pip install -r requirements.txt
```

## Summary

### test_basic_parser
```test_basic_parser.py``` creates a dynamic unit test by listing all of the ```.shex``` files in the [shexTest/schemas](https://github.com/shexSpec/shexTest/schemas) directory and then invoking unittest.  Currently all tests pass with the exceptions of:

* Unicode characters > 4 bytes.  The ANTLR4 lexer does not support this construct.
```text
1literalPattern_with_REGEXP_escapes.shex
1literalPattern_with_REGEXP_escapes_bare.shex
1literalPattern_with_UTF8_boundaries.shex
1refbnode_with_spanning_PN_CHARS_BASE1.shex
1val1STRING_LITERAL1_with_ECHAR_escapes.shex
1val1STRING_LITERAL1_with_UTF8_boundaries.shex
_all.shex
```

* A curious optimization done by the Javascript implementation of ShExC to ShExJ that simplifies certain constructs of the form "AND(AND(b, c), a)" to "AND(b, c, a)". We have chosen NOT do implement this optimization as:

    **(a)** it is inconsistent -- "AND(OR(b, c), a)" is (obviously) not simplified

    **(b)** the simplification is not reversable - there is no way to convert the resulting ShExJ back to the original ShExC

```text
FocusIRI2groupBnodeNested2groupIRIRef.shex
```

### test_shexr
```test_shexr.py``` traverses the ```.json``` files in the [shexTest/schemas](https://github.com/shexSpec/shexTest/schemas) directory and, using the [rdflib-jsonld](https://github.com/RDFLib/rdflib-jsonld) package, converts them to RDF.  It then compares the result with the corresponding ```.ttl`` directory.


_Failing Tests_
The following tests fail the conversion:

* [1literalPattern_with_all_punctuation.ttl](https://github.com/shexSpec/shexTest/blob/master/schemas/1literalPattern_wiht_all_punctuation.ttl)
* [1val1STRING_LITERAL1_with_ECHAR_escapes.ttl](https://github.com/shexSpec/shexTest/blob/master/schemas/1val1STRING_LITERAL1_with_ECHAR_escapes.ttl)
* [1val1STRING_LITERAL1_with_all_punctuation.ttl](https://github.com/shexSpec/shexTest/blob/master/schemas/1val1STRING_LITERAL1_with_all_punctuation.ttl)

All three fail the python rdflib parse call: ```Graph().parse(<url>, format="turtle")``` with the following stack trace fragment:
```
     . . .
  File "rdflib/plugins/parsers/notation3.py", line 1591, in strconst
    "bad escape")
  File "rdflib/plugins/parsers/notation3.py", line 1615, in BadSyntax
    raise BadSyntax(self._thisDoc, self.lines, argstr, i, msg)
  File "<string>", line None
rdflib.plugins.parsers.notation3.BadSyntax: <no detail available>
```

This is due to an error in rdflib 4.2.2, as reported at [Issue 732](https://github.com/RDFLib/rdflib/issues/732) . A temporary fix for this problem can be found at [https://github.com/hsolbrig/rdflib](https://github.com/hsolbrig/rdflib).

