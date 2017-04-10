# ShExC Python Parser Scripts
## genpython.sh
Generates the ANTLR4 lexer and parser from [ShExDoc.g4](../../../grammar/ShExDoc.g4), and the Python JSON Schema grammar binding [pyjsg](http://github.com/hsolbrig/pyjsg) for the target Python objects

### Requirements:
* [antlr4](http://www.antlr.org) version 4.5.1 or later. <br/> **NOTE:** The antl4 version used to generate the parser must match the version of the ```antlr4-python3-runtime``` module
* [python](http://python.org) version 3.5 or later
* [pyjsg](http://github.com/hsolbrig/pyjsg) version 0.1.2 or later

### Use
```bash
> cd scripts
> . genpython.sh
```

### shexc_to_shexj
Runtime script to convert ShExC (.shex) to ShExJ (.json)

### Use
See: [Main documentation page](../README.md) for details

