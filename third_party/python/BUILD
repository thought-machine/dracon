
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
    visibility = ["PUBLIC"]
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = ["1555c64851feea9ec2b55a8540d4161302c9d744"],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
    visibility = ["PUBLIC"]
)
