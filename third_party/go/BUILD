package(default_visibility = ["PUBLIC"])

# TODO(hjenkins): Support multiple versions of ES
# go_get(
#     name = "elastic_go-elasticsearch_v5",
#     revision = "v5.6.1",
#     get = "github.com/elastic/go-elasticsearch",
# )

go_get(
    name = "elastic_go-elasticsearch_v6",
    get = "github.com/elastic/go-elasticsearch",
    install = [
        "",
        "esapi",
        "estransport",
        "internal/version",
    ],
    revision = "v6.8.2",
)

# TODO(hjenkins): Support multiple versions of ES
# go_get(
#     name = "elastic_go-elasticsearch_v7",
#     revision = "v7.4.1",
#     get = "github.com/elastic/go-elasticsearch",
# )

go_get(
    name = "gogo_protobuf",
    get = "github.com/gogo/protobuf",
    install = [
        "gogoproto",
        "proto",
        "sortkeys",
    ],
    revision = "v1.2.1",
)

go_get(
    name = "golang_protobuf",
    get = "github.com/golang/protobuf",
    install = [
        "proto",
        "ptypes",
        "ptypes/any",
        "ptypes/duration",
        "ptypes/timestamp",
    ],
    revision = "v1.3.2",
)

go_get(
    name = "google_uuid",
    get = "github.com/google/uuid",
    revision = "v1.1.1",
)

go_get(
    name = "mitchellh_go-homedir",
    get = "github.com/mitchellh/go-homedir",
    revision = "v1.1.0",
)

go_get(
    name = "speps_go-hashids",
    get = "github.com/speps/go-hashids",
    revision = "v2.0.0",
)

go_get(
    name = "spf13_cobra",
    get = "github.com/spf13/cobra",
    licences = ["apache-2.0"],
    revision = "v0.0.5",
    deps = [
        ":spf13_pflag",
    ],
)

go_get(
    name = "spf13_viper",
    get = "github.com/spf13/viper",
    licences = ["mit"],
    revision = "v1.4.0",
    deps = [
        ":fsnotify",
        ":hashicorp_hcl",
        ":magiconair_properties",
        ":mitchellh_mapstructure",
        ":pelletier_go-toml",
        ":spf13_afero",
        ":spf13_cast",
        ":spf13_jwalterweatherman",
        ":spf13_pflag",
        ":yaml.v2",
    ],
)

go_get(
    name = "stretchr_testify",
    get = "github.com/stretchr/testify",
    install = [
        "",
        "assert",
    ],
    revision = "v1.2.2",
    strip = [
        "vendor",
    ],
    deps = [
        ":davecgh_go-spew",
        ":pmezard_go-difflib",
        ":stretchr_objx",
    ],
)

go_get(
    name = "davecgh_go-spew",
    get = "github.com/davecgh/go-spew",
    install = [
        "spew",
    ],
    revision = "v1.1.1",
)

go_get(
    name = "pmezard_go-difflib",
    get = "github.com/pmezard/go-difflib",
    install = [
        "difflib",
    ],
    revision = "v1.0.0",
)

go_get(
    name = "stretchr_objx",
    get = "github.com/stretchr/objx",
    revision = "v0.1.1",
)

go_get(
    name = "fsnotify",
    get = "github.com/fsnotify/fsnotify",
    revision = "v1.4.7",
    deps = [
        ":x_sys",
    ],
)

go_get(
    name = "x_sys",
    get = "golang.org/x/sys",
    install = [
        "unix",
        "cpu",
    ],
    revision = "95b1ffbd15a57cc5abb3f04402b9e8ec0016a52c",
)

go_get(
    name = "hashicorp_hcl",
    get = "github.com/hashicorp/hcl/...",
    revision = "v1.0.0",
)

go_get(
    name = "magiconair_properties",
    get = "github.com/magiconair/properties",
    revision = "v1.8.0",
)

go_get(
    name = "mitchellh_mapstructure",
    get = "github.com/mitchellh/mapstructure",
    revision = "v1.1.2",
)

go_get(
    name = "pelletier_go-toml",
    get = "github.com/pelletier/go-toml",
    revision = "v1.2.0",
)

go_get(
    name = "spf13_cast",
    get = "github.com/spf13/cast",
    revision = "v1.3.0",
)

go_get(
    name = "spf13_jwalterweatherman",
    get = "github.com/spf13/jwalterweatherman",
    revision = "v1.0.0",
)

go_get(
    name = "yaml.v2",
    get = "gopkg.in/yaml.v2",
    revision = "v2.2.2",
)

go_get(
    name = "spf13_afero",
    get = "github.com/spf13/afero",
    install = [
        "",
        "mem",
    ],
    revision = "v1.1.2",
    deps = [
        ":x_text",
    ],
)

go_get(
    name = "x_text",
    get = "golang.org/x/text",
    install = [
        "transform",
        "unicode/...",
        "secure/bidirule",
    ],
    revision = "v0.3.0",
)

go_get(
    name = "spf13_pflag",
    get = "github.com/spf13/pflag",
    revision = "v1.0.3",
)

go_get(
    name = "pkg_errors",
    get = "github.com/pkg/errors",
    revision = "v0.8.1",
)

go_get(
    name = "evanphx_json-patch",
    get = "github.com/evanphx/json-patch",
    revision = "v4.5.0",
    deps = [
        ":pkg_errors",
    ],
)

go_get(
    name = "ghodss_yaml",
    get = "github.com/ghodss/yaml",
    revision = "v1.0.0",
    deps = [
        ":yaml.v2",
    ],
)

go_get(
    name = "rakyll_statik",
    get = "github.com/rakyll/statik",
    install = [
        "fs",
    ],
    revision = "v0.1.6",
)

KUBERNETES_VERSION = "1.13.12"

go_get(
    name = "apimachinery",
    get = "k8s.io/apimachinery",
    install = [
        "pkg/apis/meta/v1",
        "pkg/runtime/schema",
        "pkg/api/resource",
        "pkg/conversion",
        "pkg/labels",
        "pkg/runtime",
        "pkg/util/intstr",
        "pkg/util/runtime",
        "pkg/util/sets",
        "pkg/util/validation",
        "pkg/util/json",
        "pkg/util/errors",
        "pkg/util/naming",
        "pkg/util/net",
        "pkg/util/validation/...",
        "pkg/fields",
        "pkg/conversion/...",
        "pkg/watch",
        "pkg/types",
        "pkg/selection",
        "third_party/forked/golang/reflect",
    ],
    repo = "github.com/kubernetes/apimachinery",
    revision = "kubernetes-%s" % KUBERNETES_VERSION,
    strip = [
        "pkg/util/proxy",
        "pkg/util/jsonmergepatch",
        "pkg/runtime/serializer/recognizer/testing",
    ],
    deps = [
        ":gogo_protobuf",
        ":google_gofuzz",
        ":inf.v0",
        ":klog",
        ":x_net",
    ],
)

go_get(
    name = "google_gofuzz",
    get = "github.com/google/gofuzz",
    revision = "v1.0.0",
)

go_get(
    name = "inf.v0",
    get = "gopkg.in/inf.v0",
    revision = "v0.9.1",
)

go_get(
    name = "x_net",
    get = "golang.org/x/net/...",
    revision = "13f9640d40b9cc418fb53703dfbd177679788ceb",
    strip = [
        "http2/h2demo",
        "http2/h2i",
    ],
    deps = [
        ":x_sys",
        ":x_text",
    ],
)

go_get(
    name = "klog",
    get = "k8s.io/klog",
    revision = "v1.0.0",
)

go_get(
    name = "golang-migrate_migrate",
    get = "github.com/golang-migrate/migrate",
    install = [
        "",
        "database/postgres",
        "source/go_bindata",
        "database",
        "internal/...",
        "source",
    ],
    revision = "v4.7.0",
    deps = [
        ":hashicorp_go-multierror",
        ":lib_pq",
    ],
)

go_get(
    name = "hashicorp_go-multierror",
    get = "github.com/hashicorp/go-multierror",
    revision = "v1.0.0",
    deps = [
        ":hashicorp_errwrap",
    ],
)

go_get(
    name = "hashicorp_errwrap",
    get = "github.com/hashicorp/errwrap",
    revision = "v1.0.0",
)

go_get(
    name = "lib_pq",
    get = "github.com/lib/pq/...",
    revision = "v1.0.0",
)

go_get(
    name = "jmoiron_sqlx",
    get = "github.com/jmoiron/sqlx/...",
    revision = "v1.2.0",
)
