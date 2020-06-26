package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = ["a287cfbf4d674eb826e9686d0fc5def34bfadaf1"],
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    outs = ["google/protobuf"],
    hashes = ["6ad6b86dcf7d29e8868decb0269c145f003f04c9"],
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = ["96d9674fba0dd86980f1bd484b69de2dc2079f75"],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = ["cf7af9d5871157d18d2f491202f9e0ed3854bfd4"],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
)

pip_library(
    name = "requests",
    version = "2.23.0",
    hashes = ["a864f264eb58c13e03408ad4becbe896ce9f510f"],
    deps = [
        ":urllib3",
        ":chardet",
        ":certifi",
        ":idna",
    ],
)

pip_library(
    name = "urllib3",
    hashes = ["a116a02e43b5ffe5d316c9e9ef2ffa946143030b"],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = ["cf6e1706ad1119c9e474033c2302aa0f0af42865"],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = ["d6644e25f4f93987dc58b527f57832491d88eae1"],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = ["ef35db113f6d552c3cf08cc8d5501e9871f1107c"],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    hashes = ["df99c8fee0c6e6538beaa2dbde73d16516480bfd"],
    licences = ["MIT"],
    deps = [":requests"],
    version = "1.1.3",
)
