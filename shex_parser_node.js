var FS = require('fs');
var PATH = require('path');

var shexFilename = null;
var dump = process.argv[4];
var Ds = {};
var m;

for (var i = 2; i < process.argv.length; ++i) {
    if ((m = process.argv[i].match("^-D([a-zA-Z0-9_]+)")))
        Ds[m[1]] = true;
    else if ((m = process.argv[i].match("^--dump")))
        dump = process.argv[++i];
    else if (shexFilename === null)
        shexFilename = process.argv[i];
    else
        console.log(process.argv[i]);
}

if (dump == "D") {
    console.log(JSON.stringify(Ds));
    return 0;
}

var shexParser = require("./shExDoc").parser;
var allEndpoints = [];
shexParser.yy = {
    log: function (s) { console.log(">>" + s + "<<"); },
    allEndpoints: allEndpoints,
    concomitants: [],
    medHistory: [],
    covariates: [],
    subtypes: {},
    name: null,
    type: null,
    file: shexFilename
};
var shexFile = FS.readFileSync(PATH.normalize(shexFilename), "utf8");
var shex = shexParser.parse(shexFile); // console.warn(shex);
