package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "go_tool",
    hashes = [
        "951a3c7c6ce4e56ad883f97d9db74d3d6d80d5fec77455c6ada6c1f7ac4776d2",  # linux-amd64
    ],
    version = "1.16.3",
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
