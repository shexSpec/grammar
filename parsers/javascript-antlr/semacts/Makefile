

lib/ShExDocParser.js: ShExDoc.g4
	antlr4 -listener -o lib -Dlanguage=JavaScript ShExDoc.g4
	echo "exports.ShExDocListener = null;" > lib/ShExDocListener.js # empty out listener

clean:
	rm lib/*

demo: lib/ShExDocParser.js demo.shex
	node ./shexc-to-shexj < demo.shex
