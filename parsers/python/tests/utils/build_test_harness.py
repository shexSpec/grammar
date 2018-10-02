import unittest
import sys
from typing import Tuple

import os
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
        for fname, fpath in \
                (cls.enumerate_http_files(cls.repo_url) if ':' in cls.repo_url else
                    cls.enumerate_directory(cls.repo_url)):
            if fname.endswith(cls.file_suffix):
                if fname not in cls.skip and (started or fname == cls.start_at):
                    started = True
                    test_func = cls.make_test_function(fpath)
                    setattr(cls, 'test_{0}'.format(fname.rsplit('.', 1)[0]), test_func)
                    if cls.single_file:
                        break

    @staticmethod
    def enumerate_http_files(url) -> Tuple[str, str]:
        resp = requests.get(url)
        if resp.ok:
            for f in resp.json():
                yield f['name'], f['download_url']
        else:
            print("Error {}: {}".format(resp.status_code, resp.reason), file=sys.stderr)

    @staticmethod
    def enumerate_directory(dir_) -> Tuple[str, str]:
        for fname in os.listdir(dir_):
            fpath = os.path.join(dir_, fname)
            if os.path.isfile(fpath):
                yield fname, fpath

    def blank_test(self):
        self.assertTrue(True)
