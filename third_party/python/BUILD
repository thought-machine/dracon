package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = [
        "19d549606c2b0fa7355209436b0a32c333e81f68333b93c00a9282b4e1d1946f", # TM Mirror
        "eaa100d9fc6059b9caccdf61b73d6c48b7a7f4b13b416d0d94426e4e68780331", # Public
        "43bd26c3aafb57eeace6d63550c9a1b34a6a205712a5cc544bef3b8390550717", # GitHub Actions
    ],
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    outs = ["google/protobuf"],
    hashes = [
        "f958e27563875e4e4374cf4fce3c061f2f7742faa1139acf8def40da07e0ce4e", # TM Mirror
        "67aef484fdd95e53d4556f545b09dfa07bd5d1ae63edd79c904e245959312181", # Public
        "95dfa3e09ff6a00dd9a02a250d82255d43a2a2b48144c034ee313efd15fd73da", # GitHub Actions
    ],
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = [
        "c9230375a8c4ea835e383dea8911e6cd291a52eec59f7271bd862bfbe27f615c", # TM Mirror
        "7d32f137ecddaacb225e4654501b87121374a503d8b33ffabab02d6cd157c73d", # Public
        "bd3adba527cf620cda98f80c1f2ff4e0d43e03c5b8a980465dfb852e22954002", # GitHub Actions
    ],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = [
        "4c468c181cb3dd03c3eb368940cb008d6d45bad48ee085d3f9655a15dc6603f4", # TM Mirror
        "e21a52ba3ee5c0bb30f07d28b671bd08820ac992f5f1c78f6ec51b8d25faf9ba", # Public
        "a1f4385f169be6c75c37e2f656fecbc5a014dda92357cf8509a4a651d35ae141", # GitHub Actions
    ],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
)

pip_library(
    name = "requests",
    version = "2.23.0",
    hashes = [
        "a749c5c5d6f10ab3ac2a0e46b5459b829a93b0ef08dec412236d5abde75d94f7", # TM Mirror
        "9db0caaab267a16c47717326883fdc7ac92af9007c057e6dd5ff5ee4961124af", # Public
        "13308be8e327c460966a1b6f89ddf3ceccc70fe60923888ab6a35542cae495fe", # GitHub Actions
    ],
    deps = [
        ":urllib3",
        ":chardet",
        ":certifi",
        ":idna",
    ],
)

pip_library(
    name = "urllib3",
    hashes = [
        "dfadcc16d8b81f332c6c4e4d54fa85baec2a94d5a3aab593e705b0f7d30c4494", # TM Mirror
        "170b110226e286c02224e1b685d8aca4d81635d789df0db7167ebd03af008925", # Public
        "d9bb0ea061189e901e9919934c23723749a329d5aa5bb53397ed1af038e1e9c9", # GitHub Actions
    ],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = [
        "98e5784d3d46c6b61d9ecc2a2692ecf989eb0b5e72899ab52ad1143eed6dc532", # TM Mirror
        "b982a8eb40d4be0a5aab7b8b81169a953d43e4b2b95685f47f24effd8fbb417d", # Public
        "4be7beacd1d95e5dc42606dae7147c5d9b8804f0b54d4683cbf02cb2a3f8ff43", # GitHub Actions
    ],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = [
        "ac4bab4ffa21a25e617a083da655f4f3515616fcbc0623cecd30169d8ac17d51", # TM Mirror
        "d57c002a95bb1384e38740daba0bce14e31234dd7e18182978b7a7a3936d737f", # Public
        "be1489489a31e19f2665de886fd89f37cf4325b7740c9e5131afa69b054ab128", # GitHub Actions
    ],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = [
        "6777f28e7113d6e780e0bcdcd837cc97b2cd5bfedaf3b25b202e64ad1819d69d", # TM Mirror
        "f42d88cf9a14ea6b819e5f27906d98a165a9fb9608a9e4fff9328e93b5a7ce58", # Public
        "44ca85ef74af1d787ccbb9d94031748b2a3839a60e38b5f700d042bfcc0de283", # GitHub Actions
    ],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    licences = ["MIT"],
    deps = [":requests"],
    version = "1.1.3",
)
