package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = ["a287cfbf4d674eb826e9686d0fc5def34bfadaf1", "eaa100d9fc6059b9caccdf61b73d6c48b7a7f4b13b416d0d94426e4e68780331"],
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    outs = ["google/protobuf"],
    hashes = ["6ad6b86dcf7d29e8868decb0269c145f003f04c9","67aef484fdd95e53d4556f545b09dfa07bd5d1ae63edd79c904e245959312181"],
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = ["96d9674fba0dd86980f1bd484b69de2dc2079f75", "7d32f137ecddaacb225e4654501b87121374a503d8b33ffabab02d6cd157c73d"],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = ["cf7af9d5871157d18d2f491202f9e0ed3854bfd4", "e21a52ba3ee5c0bb30f07d28b671bd08820ac992f5f1c78f6ec51b8d25faf9ba"],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
)

pip_library(
    name = "requests",
    version = "2.23.0",
    hashes = ["a864f264eb58c13e03408ad4becbe896ce9f510f", "9db0caaab267a16c47717326883fdc7ac92af9007c057e6dd5ff5ee4961124af"],
    deps = [
        ":urllib3",
        ":chardet",
        ":certifi",
        ":idna",
    ],
)

pip_library(
    name = "urllib3",
    hashes = ["a116a02e43b5ffe5d316c9e9ef2ffa946143030b", "170b110226e286c02224e1b685d8aca4d81635d789df0db7167ebd03af008925"],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = ["cf6e1706ad1119c9e474033c2302aa0f0af42865", "b982a8eb40d4be0a5aab7b8b81169a953d43e4b2b95685f47f24effd8fbb417d"],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = ["d6644e25f4f93987dc58b527f57832491d88eae1", "d57c002a95bb1384e38740daba0bce14e31234dd7e18182978b7a7a3936d737f"],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = ["ef35db113f6d552c3cf08cc8d5501e9871f1107c", "f42d88cf9a14ea6b819e5f27906d98a165a9fb9608a9e4fff9328e93b5a7ce58"],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    licences = ["MIT"],
    deps = [":requests"],
    version = "1.1.3",
)
