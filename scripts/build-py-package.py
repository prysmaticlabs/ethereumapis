#!/usr/bin/env python3

import sys

MIN_PYTHON_VERSION = (3, 5, 0)
if sys.version_info < MIN_PYTHON_VERSION:
    print("\N{no entry} Python %s.%s.%s. or later is required." % MIN_PYTHON_VERSION)
    sys.exit(1)

import os
import tempfile

if sys.prefix == sys.base_prefix:
    print("\N{package} Creating a virtualenv")

    import subprocess
    import venv

    _BUILDTIME_REQUIREMENTS = [
        "grpcio-tools==1.34.0",
        "pygit2==1.4.0",
        "setuptools==51.0.0",
    ]

    with tempfile.TemporaryDirectory() as venv_dir:
        """
        Create a new virtualenv to package build requirements.
        """
        venv.create(venv_dir, with_pip=True)
        venv_bin_dir = os.path.join(venv_dir, "bin")
        venv_pip_bin = os.path.join(venv_bin_dir, "pip")
        venv_python_bin = os.path.join(venv_bin_dir, "python")

        print("\N{hammer and wrench}  Installing buildtime dependencies")

        try:
            # Begin by upgrading pip
            ret = subprocess.run(
                [venv_pip_bin, "install", "--upgrade", "pip"], check=True
            )

            # Now install buildtime dependencies
            ret = subprocess.run(
                [venv_pip_bin, "install", *_BUILDTIME_REQUIREMENTS], check=True
            )

            # Finally, re-invoke the script with our venv-enabled python
            ret = subprocess.run([venv_python_bin, __file__, *sys.argv[1:]], check=True)

        except subprocess.CalledProcessError:
            pass  # clean up nicely

    sys.exit(ret.returncode)

import grpc_tools.protoc
import pkg_resources
import pygit2
import re
import setuptools
import setuptools.command.build_py
import shutil


_REPOSITORY_ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), os.pardir))
_README_FILENAME = os.path.join(_REPOSITORY_ROOT, "README.md")
_DIST_ROOT = os.path.join(_REPOSITORY_ROOT, "dist")

_API_VERSIONS = ["v1", "v1alpha1"]
_INIT_FILES = {
    "ethereumapis/__init__.py": """
import sys
import ethereumapis
import ethereumapis._internal.github
import ethereumapis._internal.google.api

sys.modules["eth"] = ethereumapis
sys.modules["github"] = ethereumapis._internal.github
sys.modules["google.api"] = ethereumapis._internal.google.api

__all__ = ["v1", "v1alpha1"]
""",
    "ethereumapis/v1alpha1/__init__.py": """
from ethereumapis.v1alpha1 import (
    attestation_pb2,
    beacon_block_pb2,
    beacon_chain_pb2,
    beacon_chain_pb2_grpc,
    node_pb2,
    node_pb2_grpc,
    validator_pb2,
    validator_pb2_grpc,
)
""",
    "ethereumapis/v1/__init__.py": """
from ethereumapis.v1 import (
    attestation_pb2,
    beacon_block_pb2,
    beacon_chain_service_pb2,
    beacon_chain_service_pb2_grpc,
    beacon_debug_service_pb2,
    beacon_debug_service_pb2_grpc,
    beacon_state_pb2,
    node_pb2,
    node_pb2_grpc,
    validator_pb2,
    validator_service_pb2,
    validator_service_pb2_grpc,
)
""",
    "ethereumapis/_internal/__init__.py": "",
}


class ProtoBuilder(setuptools.Command):
    """Discover and build ethereumapis protocol buffer definitions."""

    description = "compiles protocol buffer definitions with gRPC support"
    user_options = []

    def initialize_options(self):
        pass

    def finalize_options(self):
        pass

    @staticmethod
    def _discover_protos(dirname, excluded=[]):
        proto_files = list()
        grpc_proto_files = list()

        service_pattern = re.compile("^service .+ {")

        for root, dirs, files in os.walk(dirname, topdown=False):
            for name in files:
                if name.endswith(".proto") and name not in excluded:
                    proto_filename = os.path.join(root, name)
                    with open(proto_filename, "r") as proto_file:
                        service_match = False
                        for line in proto_file.readlines():
                            if service_pattern.match(line):
                                service_match = True
                                break

                        if service_match:
                            grpc_proto_files.append(proto_filename)
                        else:
                            proto_files.append(proto_filename)

        return (proto_files, grpc_proto_files)

    def run(self):
        with tempfile.TemporaryDirectory() as temp_dir:

            def tmpfile(subpath):
                return os.path.join(temp_dir, *subpath.split("/"))

            def install_local_repo(local_dir, dest_dir):
                print(
                    "\U0001F500 Copying local directory {local_dir} to {dest_dir}".format(
                        local_dir=local_dir, dest_dir=dest_dir
                    )
                )
                shutil.copytree(local_dir, tmpfile(dest_dir))

            def install_git_repo(remote, dest_dir):
                print(
                    "\N{inbox tray} Cloning Git repo {remote} to {dest_dir}".format(
                        remote=remote, dest_dir=dest_dir
                    )
                )
                pygit2.clone_repository(remote, tmpfile(dest_dir))

            def compile_proto(include_paths, out_dir, proto_srcs, grpc=False):
                print(
                    "\N{building construction}  Building {num_protos} Python proto libraries{grpc_option}into {out_dir}".format(
                        num_protos=len(proto_srcs),
                        grpc_option=" with gRPC support " if grpc else " ",
                        out_dir=out_dir,
                    )
                )

                _include_paths = [
                    tmpfile("include"),
                    pkg_resources.resource_filename("grpc_tools", "_proto"),
                ]
                _include_paths.extend(map(tmpfile, include_paths))

                args = list()
                for path in _include_paths:
                    args.append("--proto_path={}".format(path))

                if grpc:
                    args.append("--grpc_python_out={}".format(tmpfile(out_dir)))
                args.append("--python_out={}".format(tmpfile(out_dir)))

                args.extend(proto_srcs)
                return grpc_tools.protoc.main(args)

            install_local_repo(_REPOSITORY_ROOT, "include/ethereumapis")
            install_git_repo(
                "https://github.com/gogo/protobuf.git",
                "include/gogo/github.com/gogo/protobuf",
            )
            install_git_repo(
                "https://github.com/googleapis/googleapis.git",
                "include/googleapis",
            )

            os.makedirs(tmpfile("ethereumapis/_internal"))
            compile_proto(
                ["include/gogo", "include/googleapis"],
                "ethereumapis/_internal",
                [
                    tmpfile(
                        "include/gogo/github.com/gogo/protobuf/gogoproto/gogo.proto"
                    ),
                    tmpfile("include/googleapis/google/api/annotations.proto"),
                    tmpfile("include/googleapis/google/api/http.proto"),
                ],
            )

            print("\U0001F50E Searching for protobuf definition files")
            proto_files = list()
            grpc_proto_files = list()
            for api_version in _API_VERSIONS:
                pf, gpf = self._discover_protos(
                    tmpfile("include/ethereumapis/eth/{}".format(api_version)),
                    excluded=["swagger.proto"],
                )
                proto_files.extend(pf)
                grpc_proto_files.extend(gpf)

            compile_proto(
                [
                    "include/ethereumapis",
                    "include/gogo",
                    "include/googleapis",
                ],
                "ethereumapis",
                proto_files,
            )

            compile_proto(
                [
                    "include/ethereumapis",
                    "include/gogo",
                    "include/googleapis",
                ],
                "ethereumapis",
                grpc_proto_files,
                grpc=True,
            )

            os.rename(tmpfile("ethereumapis/eth/v1"), tmpfile("ethereumapis/v1"))
            os.rename(
                tmpfile("ethereumapis/eth/v1alpha1"),
                tmpfile("ethereumapis/v1alpha1"),
            )
            os.rmdir(tmpfile("ethereumapis/eth"))

            print("\N{adhesive bandage} Applying __init__.py patches")
            for file, contents in _INIT_FILES.items():
                with open(tmpfile(file), "w") as f:
                    f.write(contents)

            print("\N{incoming envelope} Copying ethereumapis module to dist/")
            os.makedirs(_DIST_ROOT, exist_ok=True)
            shutil.copytree(
                tmpfile("ethereumapis"),
                os.path.join(_DIST_ROOT, "ethereumapis"),
                dirs_exist_ok=True,
            )

            print("\N{check mark}  Done building ethereumapis")


class ProtoBuildHook(setuptools.command.build_py.build_py):
    """Build hook to ensure ProtoBuilder is run at build-time."""

    def run(self):
        self.run_command("build_proto")
        setuptools.command.build_py.build_py.run(self)


setuptools.setup(
    name="ethereumapis",
    version="0.0.1",
    url="https://github.com/prysmaticlabs/ethereumapis",
    description="Prysm's service interface definitions for the Ethereum 2.0 API",
    long_description=open(_README_FILENAME, "r").read(),
    packages=[
        "ethereumapis",
        "ethereumapis.v1",
        "ethereumapis.v1alpha1",
        "ethereumapis._internal",
    ],
    package_dir={
        "": "dist",
    },
    cmdclass={
        "build_proto": ProtoBuilder,
        "build_py": ProtoBuildHook,
    },
)
