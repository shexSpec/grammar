import os

# TODO: get git_branch into this URL
USE_LOCAL_FILES = True
schemas_base = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..', '..', '..', 'shexTest', 'schemas'))
if not os.path.exists(schemas_base) or not USE_LOCAL_FILES:
    schemas_base = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
git_branch = 'master'

print(f"***** Schema test base is {schemas_base} *****")
print(f"***** git test branch: {git_branch}*****\n")

# Settings for rdflib parsing issue

#         See line 1578 in notation3.py:
#                 k = 'abfrtvn\\"\''.find(ch)
#                 if k >= 0:
#                     uch = '\a\b\f\r\t\v\n\\"\''[k]
RDFLIB_PARSING_ISSUE_FIXED = False
