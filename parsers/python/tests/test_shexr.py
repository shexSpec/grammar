import re
import unittest

import json
from typing import Optional

import os

import rdflib
import requests

from rdflib import Graph, URIRef, Namespace
from rdflib.term import Identifier, BNode

from tests import schemas_base, RDFLIB_PARSING_ISSUE_FIXED
from tests.utils.build_test_harness import ValidationTestCase

# Comment this line out if you don't want context caching
from pyshexc.rdflib import contextcache

#
# Starting file name (with or without ".ttl" suffix)
START_AT = ""

# False if you want to start somewhere in the middle
SINGLE_FILE = bool(START_AT)

# True means use shex.jsonld from the tests directory. False means use the link in the file
# (Used for testing fixes to @context)
USE_LOCAL_CONTEXT = False

# Keep file prefixes to find errors
KEEP_PREFIXES = False

MF = Namespace("https://raw.githubusercontent.com/shexSpec/shexTest/master/validation/manifest")
SHX = Namespace("http://www.w3.org/ns/shex#")
fpath = os.path.abspath(os.path.split(__file__)[0])
SHT = Namespace("file://" + fpath + '/')


class ShexJToShexRTestCase(ValidationTestCase):
    pass

# ShexJToShexRTestCase.repo_url = "~/Development/git/shexSpec/shexTest/schemas"
ShexJToShexRTestCase.repo_url = schemas_base
ShexJToShexRTestCase.file_suffix = ".ttl"
ShexJToShexRTestCase.skip = {'coverage.ttl': 'Not ShEx',
                             'manifest.ttl': 'Not ShEx',
                             'meta.ttl': 'Not ShEx',
                             }
if not RDFLIB_PARSING_ISSUE_FIXED:
    issue_text = 'RDFLIB quote parsing issue not fixed'
    ShexJToShexRTestCase.skip['1literalPattern_with_all_punctuation.ttl'] = issue_text
    ShexJToShexRTestCase.skip['1val1STRING_LITERAL1_with_ECHAR_escapes.ttl'] = issue_text
    ShexJToShexRTestCase.skip['1val1STRING_LITERAL1_with_all_punctuation.ttl'] = issue_text

ShexJToShexRTestCase.start_at = START_AT if not START_AT or START_AT.endswith('.ttl') else START_AT + '.ttl'
ShexJToShexRTestCase.single_file = SINGLE_FILE


def strip_prefixes(txt: str) -> str:
    return txt if KEEP_PREFIXES else re.sub(r'@prefix .*', '',  txt, flags=re.MULTILINE).strip()


def complete_definition(subj: Identifier,
                        source_graph: Graph,
                        target_graph: Optional[Graph] = None) -> Graph:
    """
    Add a full definition for the supplied subject, following any object bnodes, to target_graph
    :param subj: URI or BNode for subject
    :param source_graph: Graph containing defininition
    :param target_graph: Graph to carry definition
    :return: target_graph
    """
    if not target_graph:
        target_graph = Graph()
        target_graph.bind("mf", MF)
        target_graph.bind("shx", SHX)
        target_graph.bind("sht", SHT)
    for p, o in source_graph.predicate_objects(subj):
        target_graph.add((subj, p, o))
        if isinstance(o, BNode):
            complete_definition(o, source_graph, target_graph)
    return target_graph

def bare_bnode(n: BNode, g: Graph) -> bool:
    """ Return True if n is a lone BNode"""
    return bool(g.subjects(None, n))

def compare_rdf(ttl_text: str, shex_ttl_url: str) -> bool:
    shex_json_url = shex_ttl_url.rsplit(".", 1)[0] + ".json"
    try:
        g1 = Graph().parse(data=ttl_text, format="turtle")
    except rdflib.plugins.parsers.notation3.BadSyntax as e:
        print(f"*****> RDFLIB Parsing Failure: {e}")
        return False
    if ':' in shex_json_url:
        resp = requests.get(shex_json_url)
        if not resp.ok:
            return False
        resp_json = resp.text
    else:
        with open(shex_json_url) as f:
            resp_json = f.read()

    if USE_LOCAL_CONTEXT:
        resp_obj = json.loads(resp_json)
        resp_obj['@context'] = "jsonld/shex.jsonld"
        g2 = Graph().parse(data=json.dumps(resp_obj), format="json-ld")
    else:
        g2 = Graph().parse(data=resp_json, format="json-ld")

    if g1.isomorphic(g2):
        return True

    g1_subjs = set([s for s in g1.subjects() if isinstance(s, URIRef) or bare_bnode(s, g1)])
    g2_subjs = set([s for s in g2.subjects() if isinstance(s, URIRef) or bare_bnode(s, g2)])
    for s in g1_subjs - g2_subjs:
        if not isinstance(s, BNode):
            print(f"Missed: {s}")

    for s in g2_subjs - g2_subjs:
        if not isinstance(s, BNode):
            print(f"Added: {s}")

    rval = True
    for s in g1_subjs.intersection(g2_subjs):
        s_in_g1 = complete_definition(s, g1)
        s_in_g2 = complete_definition(s, g2)
        if not s_in_g1.isomorphic(s_in_g2):
            print(f"{shex_json_url}")
            print(strip_prefixes(s_in_g1.serialize(format="turtle").decode()))
            print(f"{shex_ttl_url}")
            print(strip_prefixes(s_in_g2.serialize(format="turtle").decode()))
            rval = False

    return rval


def validate_file(download_url: str) -> bool:
    """
    Parse and validate the ShExC file in download_url
    :param download_url: ShExC file
    :return: True if success
    """
    if ':' in download_url:
        resp = requests.get(download_url)
        if resp.ok:
            return compare_rdf(resp.text, download_url)
        else:
            print("Error {}: {}".format(resp.status_code, resp.reason))
            return False
    else:
        with open(download_url) as f:
            return compare_rdf(f.read(), download_url)


ShexJToShexRTestCase.validation_function = validate_file
ShexJToShexRTestCase.build_test_harness()


if __name__ == '__main__':
    unittest.main()
