package(default_visibility = ["PUBLIC"])

# TODO(hjenkins): support ES 5 in ES consumer
go_module(
    name = "elastic_go-elasticsearch_v5",
    version = "v5.6.1",
    module = "github.com/elastic/go-elasticsearch/v5",
)

# TODO(hjenkins): support ES 6 in ES consumer
go_module(
    name = "elastic_go-elasticsearch_v6",
    module = "github.com/elastic/go-elasticsearch/v6",
    install = [
        ".",
        "esapi",
        "estransport",
        "internal/version",
    ],
    version = "v6.8.2",
)

go_module(
    name = "go-jira",
    licences = ["MIT"],
    module = "github.com/andygrunwald/go-jira",
    version = "v1.12.0",
    deps = [
        ":go-querystring",
        ":jwt-go",
        ":pkg_errors",
        ":structs",
        ":tgo",
    ],
)

go_module(
    name = "go-querystring",
    install = ["query"],
    licences = [
        "BSD-3-Clause",
    ],
    module = "github.com/google/go-querystring",
    version = "v1.0.0",
)

go_module(
    name = "jwt-go",
    licences = ["mit"],
    module = "github.com/dgrijalva/jwt-go",
    version = "v3.2.0+incompatible",
)

go_module(
    name = "structs",
    licences = ["mit"],
    module = "github.com/fatih/structs",
    version = "v1.1.0",
)

go_module(
    name = "tgo",
    install = [
        "tcontainer",
        "treflect",
    ],
    licences = ["BSD 3-Clause"],
    module = "github.com/trivago/tgo",
    version = "v1.0.7",
)

go_module(
    name = "elastic_go-elasticsearch_v7",
    install = [
        ".",
        "esapi",
        "estransport",
        "internal/version",
    ],
    module = "github.com/elastic/go-elasticsearch/v7",
    version = "v7.6.0",
)

go_module(
    name = "gogo_protobuf",
    install = [
        "gogoproto",
        "proto",
        "sortkeys",
    ],
    module = "github.com/gogo/protobuf",
    version = "v1.2.1",
)

go_module(
    name = "protobuf",
    download = ":protobuf_download",
    install = [
        "proto",
        "ptypes",
        "ptypes/any",
        "ptypes/duration",
        "ptypes/timestamp",
        "jsonpb",
    ],
    module = "github.com/golang/protobuf",
    strip = [
        "proto/proto3_proto",
        "conformance",
    ],
)

go_mod_download(
    name = "protobuf_download",
    module = "github.com/golang/protobuf",
    version = "v1.3.2",
)

go_module(
    name = "protoc-gen-go",
    binary = True,
    download = ":protobuf_download",
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    deps = [":protobuf"],
)

go_module(
    name = "google_uuid",
    module = "github.com/google/uuid",
    version = "v1.1.1",
)

go_module(
    name = "mitchellh_go-homedir",
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
)

go_module(
    name = "speps_go-hashids",
    module = "github.com/speps/go-hashids",
    version = "v2.0.0",
)

go_module(
    name = "spf13_cobra",
    licences = ["apache-2.0"],
    module = "github.com/spf13/cobra",
    version = "v0.0.5",
    deps = [
        ":spf13_pflag",
    ],
)

go_module(
    name = "spf13_viper",
    licences = ["mit"],
    module = "github.com/spf13/viper",
    version = "v1.7.1",
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
        ":subosit_gotenv",
        ":ini_v1",
    ],
)

go_module(
    name = "subosit_gotenv",
    module = "github.com/subosito/gotenv",
    version = "v1.2.0",
    licences = ["MIT"],
)

go_module(
    name = "ini_v1",
    module = "gopkg.in/ini.v1",
    version = "v1.62.0",
    licences = ["Apache-2.0"],
)

go_module(
    name = "stretchr_testify",
    install = [
        "assert",
        "mock",
        "require"
    ],
    module = "github.com/stretchr/testify",
    version = "v1.7.0",
    deps = [
        ":davecgh_go-spew",
        ":pmezard_go-difflib",
        ":stretchr_objx",
        ":yaml_v3",
    ],
)

go_module(
    name = "yaml_v3",
    licences = ["Apache-2.0"],
    module = "gopkg.in/yaml.v3",
    version = "496545a6307b2a7d7a710fd516e5e16e8ab62dbc",
)

go_module(
    name = "davecgh_go-spew",
    install = [
        "spew",
    ],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_module(
    name = "pmezard_go-difflib",
    install = [
        "difflib",
    ],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_module(
    name = "stretchr_objx",
    module = "github.com/stretchr/objx",
    version = "v0.1.1",
)

go_module(
    name = "fsnotify",
    module = "github.com/fsnotify/fsnotify",
    version = "v1.4.7",
    deps = [
        ":x_sys",
    ],
)

go_module(
    name = "x_sys",
    install = [
        "unix",
        "cpu",
    ],
    module = "golang.org/x/sys",
    version = "95b1ffbd15a57cc5abb3f04402b9e8ec0016a52c",
)

go_module(
    name = "hashicorp_hcl",
    install = [".", "...",],
    module = "github.com/hashicorp/hcl",
    version = "v1.0.0",
)

go_module(
    name = "magiconair_properties",
    module = "github.com/magiconair/properties",
    version = "v1.8.0",
)

go_module(
    name = "mitchellh_mapstructure",
    module = "github.com/mitchellh/mapstructure",
    version = "v1.1.2",
)

go_module(
    name = "pelletier_go-toml",
    module = "github.com/pelletier/go-toml",
    version = "v1.2.0",
)

go_module(
    name = "spf13_cast",
    module = "github.com/spf13/cast",
    version = "v1.3.0",
)

go_module(
    name = "spf13_jwalterweatherman",
    module = "github.com/spf13/jwalterweatherman",
    version = "v1.0.0",
)

go_module(
    name = "yaml.v2",
    module = "gopkg.in/yaml.v2",
    version = "v2.2.2",
)

go_module(
    name = "spf13_afero",
    install = [
        ".",
        "mem",
    ],
    module = "github.com/spf13/afero",
    version = "v1.1.2",
    deps = [
        ":x_text",
    ],
)

go_module(
    name = "x_text",
    install = [
        "encoding",
        "encoding/...",
        "transform",
        "unicode/...",
        "secure/bidirule",
    ],
    module = "golang.org/x/text",
    version = "v0.3.0",
)

go_module(
    name = "spf13_pflag",
    module = "github.com/spf13/pflag",
    version = "v1.0.3",
)

go_module(
    name = "pkg_errors",
    module = "github.com/pkg/errors",
    version = "v0.8.1",
)

go_module(
    name = "evanphx_json-patch",
    module = "github.com/evanphx/json-patch",
    version = "v4.5.0",
    deps = [
        ":pkg_errors",
    ],
)

go_module(
    name = "ghodss_yaml",
    module = "github.com/ghodss/yaml",
    version = "v1.0.0",
    deps = [
        ":yaml.v2",
    ],
)

go_module(
    name = "rakyll_statik",
    install = [
        "fs",
    ],
    module = "github.com/rakyll/statik",
    version = "v0.1.6",
)

KUBERNETES_VERSION = "1.13.12"

go_module(
    name = "apimachinery",
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
    module = "k8s.io/apimachinery",
    strip = [
        "pkg/util/proxy",
        "pkg/util/jsonmergepatch",
        "pkg/runtime/serializer/recognizer/testing",
    ],
    version = "kubernetes-%s" % KUBERNETES_VERSION,
    deps = [
        ":gogo_protobuf",
        ":google_gofuzz",
        ":inf.v0",
        ":klog",
        ":x_net",
    ],
)

go_module(
    name = "google_gofuzz",
    module = "github.com/google/gofuzz",
    version = "v1.0.0",
)

go_module(
    name = "inf.v0",
    module = "gopkg.in/inf.v0",
    version = "v0.9.1",
)

go_module(
    name = "x_net",
    install = ["..."],
    module = "golang.org/x/net",
    strip = [
        "http2/h2demo",
        "http2/h2i",
    ],
    version = "13f9640d40b9cc418fb53703dfbd177679788ceb",
    deps = [
        ":x_sys",
        ":x_text",
    ],
)

go_module(
    name = "klog",
    module = "k8s.io/klog",
    version = "v1.0.0",
)

go_module(
    name = "golang-migrate_migrate",
    install = [
        ".",
        "database/postgres",
        "source/go_bindata",
        "database",
        "internal/...",
        "source",
    ],
    module = "github.com/golang-migrate/migrate/v4",
    version = "v4.7.0",
    deps = [
        ":hashicorp_go-multierror",
        ":lib_pq",
    ],
)

go_module(
    name = "hashicorp_go-multierror",
    module = "github.com/hashicorp/go-multierror",
    version = "v1.0.0",
    deps = [
        ":hashicorp_errwrap",
    ],
)

go_module(
    name = "hashicorp_errwrap",
    module = "github.com/hashicorp/errwrap",
    version = "v1.0.0",
)

go_module(
    name = "lib_pq",
    install = [
        ".",
        "oid",
    ],
    module = "github.com/lib/pq",
    version = "v1.0.0",
)

go_module(
    name = "jmoiron_sqlx",
    install = [".", "..."],
    module = "github.com/jmoiron/sqlx",
    version = "v1.2.0",
)

go_mod_download(
    name = "mockgen_download",
    module = "github.com/golang/mock",
    version = "v1.4.4",
)

go_module(
    name = "mockgen",
    binary = True,
    download = ":mockgen_download",
    install = ["mockgen"],
    licences = ["apache-2.0"],
    module = "github.com/golang/mock",
    deps = [":x_tools"],
)

go_module(
    name = "mock",
    download = ":mockgen_download",
    install = ["..."],
    licences = ["apache-2.0"],
    module = "github.com/golang/mock",
    deps = [
        ":x_tools",
    ],
)

go_module(
    name = "h2non_parth",
    licences = ["MIT"],
    module = "github.com/h2non/parth",
    version = "v0.0.0-20190131123155-b4df798d6542",
)

go_module(
    name = "h2non_gentleman",
    install = ["..."],
    licences = ["MIT"],
    module = "gopkg.in/h2non/gentleman.v1",
    strip = ["_examples"],
    version = "v1.0.4",
    deps = [
        ":x_net",
    ],
)

go_module(
    name = "h2non_gock",
    licences = ["MIT"],
    module = "gopkg.in/h2non/gock.v1",
    version = "v1.0.16",
    deps = [
        ":h2non_gentleman",
        ":h2non_parth",
        ":x_net",
    ],
)

go_module(
    name = "x_tools",
    install = [
        "cover",
        "present",
        "go/packages",
        "go/internal/packagesdriver",
        "internal/gocommand",
        "internal/packagesinternal",
        "go/gcexportdata",
        "go/internal/gcimporter",
    ],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/tools",
    version = "5bcca83a78812bd91ce8cb29be4fc4521cdc8f6f",
    deps = [
        ":x_errors",
        ":x_mod",
    ],
)

go_module(
    name = "x_crypto",
    install = ["..."],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/crypto",
    strip = ["acme/autocert"],
    version = "123391ffb6de907695e1066dc40c1ff09322aeb6",
    deps = [":x_sys"],
)

go_module(
    name = "x_errors",
    install = [".", "..."],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/xerrors",
    version = "a5947ffaace3e882f334c1750858b4a6a7e52422",
    deps = [":x_sys"],
)

go_module(
    name = "x_mod",
    install = [
        "semver",
        "module",
    ],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/mod",
    version = "v0.4.2",
    deps = [":x_errors"],
)

go_module(
    name = "x_sync",
    install = ["..."],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/sync",
    version = "fd80eb99c8f653c847d294a001bdf2a3a6f768f5",
    deps = [
        ":x_net",
    ],
)
