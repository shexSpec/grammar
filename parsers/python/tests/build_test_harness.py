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
import sys

import requests


class ValidationTestCase(unittest.TestCase):
    longMessage = True

    repo_url = None                 # str
    file_suffix = None              # file suffix (e.g. ".shex")
    start_at = ""                   # Optional[str]
    skip = []                       # List[str]
    validation_function = None      # Callable[[str, str], bool]
    single_file = False             # bool

    @classmethod
    def make_test_function(cls, url):
        def test(self):
            self.assertTrue(cls.validation_function(url))
        return test

    @classmethod
    def build_test_harness(cls) -> None:
        started = not bool(cls.start_at)
        resp = requests.get(cls.repo_url)
        if resp.ok:
            for f in resp.json():
                fname = f['name']
                if fname.endswith(cls.file_suffix):
                    if fname not in cls.skip and (started or fname >= cls.start_at):
                        started = True
                        test_func = cls.make_test_function(f['download_url'])
                        setattr(cls, 'test_{0}'.format(fname.rsplit('.', 1)[0]), test_func)
                        if cls.single_file:
                            break
        else:
            print("Error {}: {}".format(resp.status_code, resp.reason), file=sys.stderr)

    def blank_test(self):
        self.assertTrue(True)
