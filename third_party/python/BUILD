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
    hashes = [
        "f958e27563875e4e4374cf4fce3c061f2f7742faa1139acf8def40da07e0ce4e",  # TM Mirror
        "912fa826e4b9a16947853ca163c18ad530a42a159de357769052443898cae473",  # Public
        "95dfa3e09ff6a00dd9a02a250d82255d43a2a2b48144c034ee313efd15fd73da",  # GitHub Actions
    ],
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
    # outs = ["six.py"],
    hashes = [
        "4c468c181cb3dd03c3eb368940cb008d6d45bad48ee085d3f9655a15dc6603f4",  # TM Mirror
        "73031963564bba0bd0bbae94736110f52dc927f6f09c0627edf6f98a033ce669",  # Public
        "a1f4385f169be6c75c37e2f656fecbc5a014dda92357cf8509a4a651d35ae141",  # GitHub Actions
    ],
    licences = ["MIT"],
    version = "1.14.0",
)

pip_library(
    name = "requests",
    hashes = [
        "a749c5c5d6f10ab3ac2a0e46b5459b829a93b0ef08dec412236d5abde75d94f7",  # TM Mirror
        "66fa77b628df4eb564b5bb06117e0c8f756acc2865115d6ec89f05b31d20f471",  # Public
        "13308be8e327c460966a1b6f89ddf3ceccc70fe60923888ab6a35542cae495fe",  # GitHub Actions
    ],
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
    hashes = [
        "dfadcc16d8b81f332c6c4e4d54fa85baec2a94d5a3aab593e705b0f7d30c4494",  # TM Mirror
        "37b6e0d114199eb1edbde76b07e13efaa41331f4c79690d08f38a95368bbf0d2",  # Public
        "d9bb0ea061189e901e9919934c23723749a329d5aa5bb53397ed1af038e1e9c9",  # GitHub Actions
    ],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = [
        "98e5784d3d46c6b61d9ecc2a2692ecf989eb0b5e72899ab52ad1143eed6dc532",  # TM Mirror
        "aebcce6969f6c05658f124b4d13e0144acacd9cd1d258279fb634a96c10052a4",  # Public
        "4be7beacd1d95e5dc42606dae7147c5d9b8804f0b54d4683cbf02cb2a3f8ff43",  # GitHub Actions
    ],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = [
        "ac4bab4ffa21a25e617a083da655f4f3515616fcbc0623cecd30169d8ac17d51",  # TM Mirror
        "e9c3ba2c73032fb5f59330d2e1514fa6a440effa6b5ebace917884b7875fdd72",  # Public
        "be1489489a31e19f2665de886fd89f37cf4325b7740c9e5131afa69b054ab128",  # GitHub Actions
    ],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = [
        "6777f28e7113d6e780e0bcdcd837cc97b2cd5bfedaf3b25b202e64ad1819d69d",  # TM Mirror
        "9313e4342b41fa3eb7fa1cdc53392b392aa59b266e4a04edfc4a0edaaefd484d",  # Public
        "44ca85ef74af1d787ccbb9d94031748b2a3839a60e38b5f700d042bfcc0de283",  # GitHub Actions
    ],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    licences = ["MIT"],
    version = "1.1.3",
    deps = [":requests"],
)
