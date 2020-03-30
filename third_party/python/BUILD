
pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = ["beea5295ad7f122708d25b76ef4d8520c2234c1d"],
    licences = ["Apache 2.0"],
    version = "1.27.2",
    visibility = ["PUBLIC"]
)

pip_library(
    name = "protobuf",
    outs = ["google/protobuf"],
    hashes=["eb641f83c54084dd1e5d02abbd6a88b905b4191a"],
    licences=["BSD 3-Clause"],
    visibility = ["PUBLIC"],
    version="3.11.3",
    deps = [":six"]
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = ["aca2ec09a5db08878b8b10ad2e278edf8e2af570"],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = ["1555c64851feea9ec2b55a8540d4161302c9d744"],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
)

pip_library(
    name = "requests",
    visibility = ["PUBLIC"],
    version = "2.23.0",
    hashes = ["c87bcd4b18ad82110fb615085f9bf7e79ce4f94a"],
    deps = [":urllib3",":chardet",":certifi",":idna"]
)

pip_library(
    name = "urllib3",
    hashes = ["ead88fe14444ae6b459bfd72d708b08bf063a801"],
    version = "1.25.8",
)
pip_library(
    name = "certifi",
    hashes = ["aac49ec341265281fad7ab9216150710d42a1750"],
    version = "2019.11.28",
)
pip_library(
    name = "chardet",
    hashes = ["2266861b95169c95b78f9c85b5d4465182a1b23a"],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = ["9c9e76222692feddd706449c59910c377ae8e589"],
    version = "2.9",
)
pip_library(
    name = "defectdojo_api",
    hashes = ["0d6908b24baf50e193fef2abdc79a777c4608e3d"],
    licences = ["MIT"],
    deps = [":requests"],
    version = "1.1.3",
)
