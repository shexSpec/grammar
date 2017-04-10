# PyShExC tests

## Requirements
These tests require the [requests](http://docs.python-requests.org/en/master/) and [yadict-compare](https://github.com/hsolbrig/dict_compare) modules.  They can be installed by:
```bash
> pip install -r requirements.txt
```

## Summary
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
kitchenSink.shex
```


* A curious optimization done by the Javascript implementation of ShExC to ShExJ that simplifies certain constructs of the form "AND(AND(b, c), a)" to "AND(b, c, a)". We have chosen NOT do implement this optimization as:

    **(a)** it is inconsistent -- "AND(OR(b, c), a)" is (obviously) not simplified

    **(b)** the simplification is not reversable - there is no way to convert the resulting ShExJ back to the original ShExC

```text
FocusIRI2groupBnodeNested2groupIRIRef.shex
```

