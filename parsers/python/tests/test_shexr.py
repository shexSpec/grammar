# Copyright (c) 2017, Mayo Clinic
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
# Redistributions of source code must retain the above copyright notice, this
#     list of conditions and the following disclaimer.
#
#     Redistributions in binary form must reproduce the above copyright notice,
#     this list of conditions and the following disclaimer in the documentation
#     and/or other materials provided with the distribution.
#
#     Neither the name of the Mayo Clinic nor the names of its contributors
#     may be used to endorse or promote products derived from this software
#     without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
# INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
# LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
# OF THE POSSIBILITY OF SUCH DAMAGE.
import unittest

import json
import requests
import sys

from rdflib import Graph
from rdflib.compare import to_isomorphic, graph_diff

from tests.build_test_harness import ValidationTestCase

class ShexJToShexRTestCase(ValidationTestCase):
    pass

ShexJToShexRTestCase.repo_url = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
ShexJToShexRTestCase.file_suffix = ".ttl"
ShexJToShexRTestCase.skip = ['coverage.ttl', 'manifest.ttl', 'meta.ttl']

# ShexJToShexRTestCase.start_at = "1literalPattern_with_all_punctionation.ttl"
# ShexJToShexRTestCase.single_file = True

# True means use shex.jsonld from the tests directory. False means use the link in the file
# (Used for testing fixes to @context)
USE_LOCAL_CONTEXT = False


def compare_rdf(shex_ttl_url: str) -> bool:
    shex_json_url = shex_ttl_url.rsplit(".", 1)[0] + ".json"
    g1 = Graph().parse(shex_ttl_url, format="turtle")
    if not USE_LOCAL_CONTEXT:
        g2 = Graph().parse(shex_json_url, format="json-ld")
    else:
        resp = requests.get(shex_json_url)
        if not resp.ok:
            print("Error {}: {}".format(resp.status_code, resp.reason), file=sys.stderr)
            return False
        resp_json = resp.json()
        resp_json['@context'] = "shex.jsonld"
        g2 = Graph().parse(data=json.dumps(resp_json), format="json-ld")
    if not (g1.isomorphic(g2)):
        g1_iso = to_isomorphic(g1)
        g2_iso = to_isomorphic(g2)
        _, in_first, in_second = graph_diff(g1_iso, g2_iso)
        print("Expect: ", file=sys.stderr)
        for l in sorted(in_first.serialize(format='nt').splitlines()):
            if l:
                print('\t' + l.decode('ascii'), file=sys.stderr)
        print("Got: ", file=sys.stderr)
        for l in sorted(in_second.serialize(format='nt').splitlines()):
            if l:
                print('\t' + l.decode('ascii'), file=sys.stderr)
        return False

    return True

ShexJToShexRTestCase.validation_function = compare_rdf
ShexJToShexRTestCase.build_test_harness()


if __name__ == '__main__':
    unittest.main()
