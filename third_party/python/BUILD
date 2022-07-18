package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = [
        "f81a854eeda013553881fedf6395a2e253d2f269a79e9318efb5d22d09246668",  # TM Mirror
        "0d28cf192fcffdb370273a54cb1600df283bd4c56dff76f626849cc6f11ae278",  # Public
        "caa92e3af01edfb5c9bd050071d3ba6a99791bebfa3414ffda8680fe801cfc3c",  # GitHub Actions
    ],
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
    hashes = [
        "b3ce5a39751dd052c771fb6b840bf5c4eaeaa8356ed8549718eaab6acf06c0eb",  # TM Mirror
        "841b1413fc9f13ed481b7c2df0d3a8c58879d405aa2f289a7998e78119ba1b73",  # Public
        "ba8b2677634bce32cc743242ae5a2596c2f5c82ff1826b40ce0093a0f8d0b3c5",  # GitHub Actions
        "331ca5f64f49cba507b2dd66b0d327acf6dffa5acae53bad1e949e5e80b3fe5f",
    ],
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
