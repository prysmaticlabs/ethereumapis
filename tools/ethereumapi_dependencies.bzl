"""Load dependencies needed to compile the ethereumapi package as a 3rd-party."""

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

def _maybe(repo_rule, name, **kwargs):
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)

def ethereumapi_deps():
    _maybe(
        go_repository,
        name = "grpc_ecosystem_grpc_gateway",
        commit = "50c55a9810a974dc5a9e7dd1a5c0d295d525f283",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
    )
    _maybe(
        go_repository,
        name = "com_github_bazelbuild_buildtools",
        commit = "eb1a85ca787f0f5f94ba66f41ee66fdfd4c49b70",
        importpath = "github.com/bazelbuild/buildtools",
    )
    _maybe(
        go_repository,
        name = "com_github_ghodss_yaml",
        commit = "25d852aebe32c875e9c044af3eef9c7dc6bc777f",
        importpath = "github.com/ghodss/yaml",
    )
    _maybe(
        go_repository,
        name = "in_gopkg_yaml_v2",
        commit = "51d6538a90f86fe93ac480b35f37b2be17fef232",
        importpath = "gopkg.in/yaml.v2",
    )
    _maybe(
        go_repository,
        name = "com_github_ferranbt_fastssz",
        commit = "06015a5d84f9e4eefe2c21377ca678fa8f1a1b09",
        importpath = "github.com/ferranbt/fastssz",
    )
    _maybe(
        go_repository,
        name = "com_github_prysmaticlabs_go_bitfield",
        commit = "62c2aee7166951c456888f92237aee4303ba1b9d",
        importpath = "github.com/prysmaticlabs/go-bitfield",
    )