import os

# TODO: get git_branch into this URL
schemas_base = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
schemas_base = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..', '..', '..', 'shexTest', 'schemas'))
if not os.path.exists(schemas_base):
    schemas_base = "https://api.github.com/repos/shexSpec/shexTest/contents/schemas"
git_branch = 'master'

print(f"***** Schema test base is {schemas_base} *****")
print(f"***** git test branch: {git_branch}*****\n")
