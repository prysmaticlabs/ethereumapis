load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    _go_repository = "go_repository",
)

def _maybe(repo_rule, name, **kwargs):
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)

def go_repository(name, **kwargs):
    _maybe(_go_repository, name, **kwargs)
