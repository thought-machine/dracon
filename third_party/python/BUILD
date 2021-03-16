package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = [
        "41bb2e981313b63ba572b8526230fabe396dba16ee30421f0a82461a2848dc6f", # TM Mirror
        "eaa100d9fc6059b9caccdf61b73d6c48b7a7f4b13b416d0d94426e4e68780331", # Public
        "858b9adc149550f79c74d97679a8ba5c52d7e40ec3defc570d66b597f05c1025", # GitHub Actions
    ],
    licences = ["Apache 2.0"],
    version = "1.27.2",
)

pip_library(
    name = "protobuf",
    outs = ["google/protobuf"],
    hashes = [
        "091f8543fa7332c700717efdb0f48b2c0cd683eaf4f9b8816df628d2f4b75a54", # TM Mirror
        "67aef484fdd95e53d4556f545b09dfa07bd5d1ae63edd79c904e245959312181", # Public
        "c6ff9fe51713459233c6f8a96dd9d5a3814993d781aca12cbf556a1196896b22", # GitHub Actions
    ],
    licences = ["BSD 3-Clause"],
    version = "3.11.3",
    deps = [":six"],
)

pip_library(
    name = "pkg_resources",
    package_name = "setuptools",
    hashes = [
        "270c08b9eabd7da9a4e75b544bbb05d4486aad67d3cad85ea2a46e8fa3d073e6", # TM Mirror
        "7d32f137ecddaacb225e4654501b87121374a503d8b33ffabab02d6cd157c73d", # Public
        "b867f0d461cb73a7d186683f8cdc773436520f88dea7251b20d794ea7f1c0727", # GitHub Actions
    ],
    licences = ["MIT"],
    version = "40.4.3",
)

pip_library(
    name = "six",
    package_name = "six",
    hashes = [
        "5e440c3c7dac784bfac48371c27c401a971514ee1ad8a989737db3a86b4202fd", # TM Mirror
        "e21a52ba3ee5c0bb30f07d28b671bd08820ac992f5f1c78f6ec51b8d25faf9ba", # Public
        "f33731e506b8c71ce050b2d1b5ca9f134371f59984f7fe552644a2055c6844bb", # GitHub Actions
    ],
    licences = ["MIT"],
    version = "1.14.0",
    outs = ["six.py"],
)

pip_library(
    name = "requests",
    version = "2.23.0",
    hashes = [
        "f6f7ddc63707be172af55bd804bd4e14cd1c67fbabb31ea17487d0b296dfe16c", # TM Mirror
        "9db0caaab267a16c47717326883fdc7ac92af9007c057e6dd5ff5ee4961124af", # Public
        "a627c90c486037f41eea41e329fbc7b72593faaa2d32fa591f9704b98a794f6b", # GitHub Actions
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
        "af6463d60692e0ed2af7092cedc3710927d6ddd8cfe7ae6dd5f52af5461050c7", # TM Mirror
        "170b110226e286c02224e1b685d8aca4d81635d789df0db7167ebd03af008925", # Public
        "a80b3c6f8b0f2d2ddb7cb2e1aa6d5984afe68b61b2f88d51c43ed913508cf081", # GitHub Actions
    ],
    version = "1.25.8",
)

pip_library(
    name = "certifi",
    hashes = [
        "df101fcf5e34c9090f5d0e18a17b1da14da1395fbffa693728db27a0a24c1b43", # TM Mirror
        "b982a8eb40d4be0a5aab7b8b81169a953d43e4b2b95685f47f24effd8fbb417d", # Public
        "12a0dfb553b6e365587537ff182cf8f159c3568dd54f05d9b86b01fa34adf2ad", # GitHub Actions
    ],
    version = "2019.11.28",
)

pip_library(
    name = "chardet",
    hashes = [
        "dfe2662abaee8d78363f61eb797f6e87e96566bfe6c53bdb98047fc30f7f0e2e", # TM Mirror
        "d57c002a95bb1384e38740daba0bce14e31234dd7e18182978b7a7a3936d737f", # Public
        "a642d84f029f2a5274d7931dd7a28635a6f6b6bf9507cd0060659b92c695ea23", # GitHub Actions
    ],
    version = "3.0.4",
)

pip_library(
    name = "idna",
    hashes = [
        "05c25be658247c9769823346378f719d2cb7d347662e8f6dee9e407450970afc", # TM Mirror
        "f42d88cf9a14ea6b819e5f27906d98a165a9fb9608a9e4fff9328e93b5a7ce58", # Public
        "a77f5960880c307d197dab8a0630b54edbd30a1ff30085ae70481112370a18b5", # GitHub Actions
    ],
    version = "2.9",
)

pip_library(
    name = "defectdojo_api",
    licences = ["MIT"],
    deps = [":requests"],
    version = "1.1.3",
)
