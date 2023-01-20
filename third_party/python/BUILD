package(default_visibility = ["PUBLIC"])

pip_library(
    name = "grpc",
    package_name = "grpcio",
    hashes = [
        "6b6839032627170b468ca08e2d75d78a66cd4b9c83af0742f6d4c46f3a877d52",  # TM Mirror
        "d1660bcf7d90b1ebe054f2e1a66ce0d96c80f5b5d0ccb3e52ee80649a938a944",  # Public
        "e0fd27277382e821b221303c41ed616953d876836a21eeb3fe2f9c4b0d6dc75c",  # GitHub Actions
    ],
    licences = ["Apache 2.0"],
    version = "1.38.0",
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
        "57f6f22bde4e042978bcd50176fdb381d7c21a9efa4041202288d3737a0c6a54",  # TM Mirror
        "a951e6c5a85abc380f46c0cd8780a1a03581e82edd08b64defebe8f52ae6ab8d",  # Public
        "c4a1ddda2582eb4c681ab9fc6d7e6ac5a0d34999d335623ec282a062ea384060",  # GitHub Actions
    ],
    licences = ["MIT"],
    version = "65.6.3",
)

pip_library(
    name = "six",
    package_name = "six",
    licences = ["MIT"],
    version = "1.16.0",
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
