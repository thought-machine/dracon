package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "go_tool",
    hashes = [
        "acc512fbab4f716a8f97a8b3fbaa9ddd39606a28be6c2515ef7c6c6311acffde",  # linux-amd64
    ],
    version = "1.19.1",
)

go_module(
    name = "golint",
    binary = True,
    hashes = [
        "6a2982682499be8804589a0082514d650ca4019a88daeed951a215d56ff4f8b8",
    ],
    install = ["golint"],
    module = "golang.org/x/lint",
    version = "83fdc39ff7b56453e3793356bcff3070b9b96445",
    deps = [
        "//third_party/go:x_tools",
    ],
)
