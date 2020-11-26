subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "elasticsearch_filebeat",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/elasticsearch_filebeat/types:elasticsearch-filebeat-issue",
    ],
)

go_test(
    name = "elasticsearch_filebeat_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/elasticsearch_filebeat/types:elasticsearch-filebeat-issue",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-elasticsearchfilebeat",
    srcs = [
        ":elasticsearch_filebeat",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-elasticsearch-filebeat",
)
