#!/usr/bin/env python3
"""build-py-package.py

Build ethereumapis protobuf libraries as a standalone Python package using
setuptools. Includes gRPC libraries for service definitions.

In order to compile the proto files themselves, we use a virtualenv to install
buildtime dependencies and invoke protoc. Protobuf library dependencies
(googleapis, gogoproto) are pulled from Git during this process. All other
functions (packaging, metadata, etc.) happens outside the virtualenv and is
handled by setuptools.

Usage:
# compile the libraries into dist/ethereumapis/
scripts/build-py-package.py build

# build the Python wheel package for distribution
scripts/build-py-package.py bdist_wheel
"""

import os
import sys

_REPOSITORY_ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), os.pardir))
_README_FILENAME = os.path.join(_REPOSITORY_ROOT, "README.md")
_LICENSE_FILENAME = os.path.join(_REPOSITORY_ROOT, "LICENSE")

if __name__ == "__main__" and sys.argv[1] == "build_proto_venv":
    if sys.prefix == sys.base_prefix:
        print(
            "\N{no entry} Tried to call build_proto_venv outside a virtualenv. Exiting."
        )
        sys.exit(1)

    import git
    import grpc_tools.protoc
    import pkg_resources
    import re
    import shutil
    import tempfile
    import tqdm

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
        "ethereumapis/_internal/github/__init__.py": "",
        "ethereumapis/_internal/github/com/__init__.py": "",
        "ethereumapis/_internal/github/com/gogo/__init__.py": "",
        "ethereumapis/_internal/github/com/gogo/protobuf/__init__.py": "",
        "ethereumapis/_internal/github/com/gogo/protobuf/gogoproto/__init__.py": "",
        "ethereumapis/_internal/google/__init__.py": "",
        "ethereumapis/_internal/google/api/__init__.py": "",
    }

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

        def install_git_repo(remote, dest_dir, tag=None):
            print(
                "\N{inbox tray} Cloning Git repo {remote}{at_tag}to {dest_dir}".format(
                    remote=remote,
                    at_tag=" at tag {tag} ".format(tag=tag) if tag else " ",
                    dest_dir=dest_dir,
                )
            )

            _progress = tqdm.tqdm()

            class Progress(git.remote.RemoteProgress):
                _total = 0

                def update(self, op_code, cur_count, max_count=None, message=""):
                    if self._total != max_count:
                        self._total = max_count
                        _progress.reset(total=self._total)
                    _progress.n = cur_count
                    _progress.refresh()

            repo = git.Repo.clone_from(
                remote,
                tmpfile(dest_dir),
                branch=tag,
                single_branch=True,
                progress=Progress(),
            )

            _progress.close()

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
            tag="v1.3.1",
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
                tmpfile("include/gogo/github.com/gogo/protobuf/gogoproto/gogo.proto"),
                tmpfile("include/googleapis/google/api/annotations.proto"),
                tmpfile("include/googleapis/google/api/http.proto"),
            ],
        )

        print("\U0001F50E Searching for protobuf definition files")
        proto_files = list()
        grpc_proto_files = list()
        for api in _API_VERSIONS:
            pf, gpf = _discover_protos(
                tmpfile("include/ethereumapis/eth/{}".format(api)),
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

        for api in _API_VERSIONS:
            os.rename(
                tmpfile("ethereumapis/eth/{}".format(api)),
                tmpfile("ethereumapis/{}".format(api)),
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
    sys.exit(0)


import setuptools
import setuptools.command.build_py
import setuptools.command.install_lib


class ProtoBuilder(setuptools.Command):
    """Discover and build ethereumapis protocol buffer definitions."""

    description = "compiles protocol buffer definitions with gRPC support"
    user_options = []

    def initialize_options(self):
        pass

    def finalize_options(self):
        pass

    def run(self):
        print("\N{package} Creating a virtualenv")

        import subprocess
        import tempfile
        import venv

        _BUILDTIME_REQUIREMENTS = [
            "GitPython>=3.1.0",
            "grpcio-tools==1.34.0",
            "setuptools==51.0.0",
            "tqdm>=4.54.0",
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
                ret = subprocess.run(
                    [venv_python_bin, __file__, "build_proto_venv"], check=True
                )

            except subprocess.CalledProcessError as e:
                # capture the result and make sure the temp dir is deleted
                ret = subprocess.CompletedProcess(e.args, e.returncode)

        if ret.returncode != 0:
            sys.exit(ret.returncode)


class ProtoBuildHook(setuptools.command.build_py.build_py):
    """Build hook to ensure ProtoBuilder is run at build-time."""

    def run(self):
        self.run_command("build_proto")
        setuptools.command.build_py.build_py.run(self)


class InstallHook(setuptools.command.install_lib.install_lib):
    """Bypass running ProtoBuilder during the install flow."""

    def run(self):
        self.install()


setuptools.setup(
    name="ethereumapis",
    version="0.12.0",
    url="https://github.com/prysmaticlabs/ethereumapis",
    description="Prysm's service interface definitions for the Ethereum 2.0 API",
    author="Prysmatic Labs",
    author_email="contact@prysmaticlabs.com",
    long_description=open(_README_FILENAME, "r").read(),
    license="Apache License 2.0",
    platforms=["any"],
    classifiers=[
        "License :: OSI Approved :: Apache Software License",
    ],
    install_requires=[
        "grpcio>=1.34.0",
        "protobuf>=3.14.0",
    ],
    python_requires=">=3.5.0",
    packages=[
        "ethereumapis",
        "ethereumapis.v1",
        "ethereumapis.v1alpha1",
        "ethereumapis._internal",
        "ethereumapis._internal.google",
        "ethereumapis._internal.google.api",
        "ethereumapis._internal.github",
        "ethereumapis._internal.github.com",
        "ethereumapis._internal.github.com.gogo",
        "ethereumapis._internal.github.com.gogo.protobuf",
        "ethereumapis._internal.github.com.gogo.protobuf.gogoproto",
    ],
    package_dir={
        "": "dist",
    },
    cmdclass={
        "build_proto": ProtoBuilder,
        "build_py": ProtoBuildHook,
        "install_lib": InstallHook,
    },
)
