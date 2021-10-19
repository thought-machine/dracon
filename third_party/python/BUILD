package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    licences = ["MIT"],
    version = "1.14.0",
)

pip_library(
    name = "requests",
    version = "2.23.0",
    deps = [
        ":certifi",
        ":chardet",
        ":idna",
        ":urllib3",
    ],
)

pip_library(
    name = "urllib3",
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    version = "3.0.4",
)

pip_library(
    name = "idna",
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    licences = ["MIT"],
    version = "1.1.3",
    deps = [":requests"],
)
