package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = [
        "19d549606c2b0fa7355209436b0a32c333e81f68333b93c00a9282b4e1d1946f",  # TM Mirror
        "0d28cf192fcffdb370273a54cb1600df283bd4c56dff76f626849cc6f11ae278",  # Public
        "43bd26c3aafb57eeace6d63550c9a1b34a6a205712a5cc544bef3b8390550717",  # GitHub Actions
    ],
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    hashes = [],
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = [
        "c9230375a8c4ea835e383dea8911e6cd291a52eec59f7271bd862bfbe27f615c",  # TM Mirror
        "841b1413fc9f13ed481b7c2df0d3a8c58879d405aa2f289a7998e78119ba1b73",  # Public
        "bd3adba527cf620cda98f80c1f2ff4e0d43e03c5b8a980465dfb852e22954002",  # GitHub Actions
    ],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = [],
    licences = ["MIT"],
    version = "1.14.0",
)

pip_library(
    name = "requests",
    hashes = [],
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
    hashes = [],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = [],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = [],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = [],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    hashes = [],
    licences = ["MIT"],
    version = "1.1.3",
    deps = [":requests"],
)
