package(default_visibility = ["PUBLIC"])

GO_VERSION = "1.15.2"

remote_file(
    name = "go_download",
    url = "https://golang.org/dl/go%s.linux-amd64.tar.gz" % GO_VERSION,
    hashes = ["b49fda1ca29a1946d6bb2a5a6982cf07ccd2aba849289508ee0f9918f6bb4552"],
)

genrule(
    name = "goroot",
    srcs = [":go_download"],
    cmd = "tar -xzf $SRCS",
    outs = ["go/"],
)

genrule(
    name = "go",
    srcs = [":goroot"],
    cmd = """
absolute_repo_path=$(pwd | sed 's#plz-out/.*##g')
go_root=$(out_location :goroot)
absolute_go_root="${absolute_repo_path}${go_root}"
absolute_go_bin="${absolute_go_root}/bin/go"

cat <<EOF >> $OUTS
#!/bin/bash
set -Eeuxo pipefail
export GOROOT="${absolute_go_root}"
${absolute_go_bin} "\\\$@"
EOF
    """,
    outs = ["go.sh"],
    binary = True,
)
