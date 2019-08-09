
def _collect_includes(gen_dir, srcs):
    """Build an include path mapping.

    It is important to not just collect unique dirnames, to also support
    proto files of the same name from different packages.

    The implementation below is similar to what bazel does in its
    ProtoCompileActionBuilder.java
    """
    includes = []
    for src in srcs:
        ref_path = src.path

        if ref_path.startswith(gen_dir):
            ref_path = ref_path[len(gen_dir):].lstrip("/")

        if src.owner.workspace_root:
            workspace_root = src.owner.workspace_root
            ref_path = ref_path[len(workspace_root):].lstrip("/")

        include = ref_path + "=" + src.path
        if include not in includes:
            includes.append(include)

    return includes

def _proto_descriptor_set_impl(ctx):
    """
    Build a descriptor set with imports and source info. Also see
    https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/grpc_json_transcoder_filter#how-to-generate-proto-descriptor-set.
    """
    proto = ctx.attr.proto[ProtoInfo]
    inputs = proto.direct_sources + ctx.files._well_known_protos + proto.transitive_sources.to_list()

    includes = _collect_includes(ctx.genfiles_dir.path, inputs)

    outfile = ctx.actions.declare_file("proto.pb")

    args = ctx.actions.args()
    args.add("--include_imports")
    args.add("--include_source_info")
    args.add("--descriptor_set_out=%s" % outfile.path)
    args.add_all(["-I%s" % include for include in includes])
    args.add_all([src.path for src in proto.direct_sources])

    ctx.actions.run(
        executable = ctx.executable._protoc,
        inputs = inputs,
        outputs = [outfile],
        arguments = [args],
    )

    files = depset([outfile])

    return [
        DefaultInfo(
            files = files,
        ),
    ]

proto_descriptor_set = rule(
    implementation = _proto_descriptor_set_impl,
    attrs = {
        "proto": attr.label(
            allow_rules = ["proto_library"],
            mandatory = True,
            providers = ["proto"],
        ),
        "_protoc": attr.label(
            default = "@com_google_protobuf//:protoc",
            executable = True,
            cfg = "host",
        ),
        "_well_known_protos": attr.label(
            default = "@com_google_protobuf//:well_known_protos",
            allow_files = True,
        ),
    },
)