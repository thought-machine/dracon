package(default_visibility = ["PUBLIC"])

# TODO(hjenkins): support ES 5 in ES consumer
go_module(
    name = "elastic_go-elasticsearch_v5",
    hashes = ["41840bb5dbdd99ead7037d2f55633e9616202d1e58696f0c1d6c8cc22a569a20"],
    module = "github.com/elastic/go-elasticsearch/v5",
    version = "v5.6.1",
)

# TODO(hjenkins): support ES 6 in ES consumer
go_module(
    name = "elastic_go-elasticsearch_v6",
    hashes = ["e4691676db41046d251321bd12b88d86ecdab869a17d96374fd4eb633dddc400"],
    install = [
        ".",
        "esapi",
        "estransport",
        "internal/version",
    ],
    module = "github.com/elastic/go-elasticsearch/v6",
    version = "v6.8.2",
)

go_module(
    name = "go-jira",
    hashes = ["64094bcb972b117352e093cf7d268c2b57809886c0a6c97459ba3cb1a7e5dc9e"],
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
    hashes = ["0e60ec823c1ce579da5afb9e6441d3acc3ecf7317a704f2d328c0e3fbc556e46"],
    install = ["query"],
    licences = [
        "BSD-3-Clause",
    ],
    module = "github.com/google/go-querystring",
    version = "v1.0.0",
)

go_module(
    name = "jwt-go",
    hashes = ["330efa092b713c652945c28336a8c826fdb29f27348cbbf602f5005d11dee012"],
    licences = ["mit"],
    module = "github.com/dgrijalva/jwt-go",
    version = "v3.2.0+incompatible",
)

go_module(
    name = "structs",
    hashes = ["5a9a1c9439eda501943950c180d7cb19b72c5ce432beda0dc730f4c4a803effa"],
    licences = ["mit"],
    module = "github.com/fatih/structs",
    version = "v1.1.0",
)

go_module(
    name = "tgo",
    hashes = ["3a2e7df432a4b103d9c77fd0ca4b74fb5b13766d4cafefc2f1efad3c1db1dbe1"],
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
    hashes = ["943a5660abb1d2d08b47e9f3951d86667ca47c6ab2e85fd983cdd69673ccb541"],
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
    hashes = ["1cadc8a7da408a99a92e888f97c951e9da3e257585eb01a80bda33661d5387cb"],
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
    hashes = ["a7248005f1f893498d86d4823f9ce1bd50bc40078fcbc2fffcdc7232ab835f2a"],
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
    hashes = ["a7248005f1f893498d86d4823f9ce1bd50bc40078fcbc2fffcdc7232ab835f2a"],
    module = "github.com/golang/protobuf",
    version = "v1.3.2",
)

go_module(
    name = "protoc-gen-go",
    binary = True,
    download = ":protobuf_download",
    hashes = ["9c0a6d6312dc36e23d8738d2d15d0dfa06cb996ab7f45e0d124cb6d65fbdb40f"],
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    deps = [":protobuf"],
)

go_module(
    name = "google_uuid",
    hashes = ["0bee0a1c7fa446e79d9e56ecfadcca396989fe30e14c63d738dfdb3aaff3db99"],
    module = "github.com/google/uuid",
    version = "v1.1.1",
)

go_module(
    name = "mitchellh_go-homedir",
    hashes = ["57ca4d6fa8b119796322cf95c4ced599a65a1bc709bbbaaff7c8687a61bb832e"],
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
)

go_module(
    name = "speps_go-hashids",
    hashes = ["13789763eecd7beb189e86cf4ae65ae4e84c940f6f84a7a5eeb1a5d8589685bf"],
    module = "github.com/speps/go-hashids",
    version = "v2.0.0",
)

go_module(
    name = "spf13_cobra",
    hashes = ["66f76e3bca88c69d154ce98f9bf3f91768b3513f5bd4cf2f02384728bfb81668"],
    licences = ["apache-2.0"],
    module = "github.com/spf13/cobra",
    version = "v0.0.5",
    deps = [
        ":spf13_pflag",
    ],
)

go_module(
    name = "spf13_viper",
    hashes = ["788eccdfe9d5bead6d43b23273d2f9acb2b0a200b66475e3403f38dcdaaee519"],
    licences = ["mit"],
    module = "github.com/spf13/viper",
    version = "v1.7.1",
    deps = [
        ":fsnotify",
        ":hashicorp_hcl",
        ":ini_v1",
        ":magiconair_properties",
        ":mitchellh_mapstructure",
        ":pelletier_go-toml",
        ":spf13_afero",
        ":spf13_cast",
        ":spf13_jwalterweatherman",
        ":spf13_pflag",
        ":subosit_gotenv",
        ":yaml.v2",
    ],
)

go_module(
    name = "subosit_gotenv",
    hashes = ["b8f373d56054a2f8724ebd24679041970003095a4fa16987d3ad53a2780c72bb"],
    licences = ["MIT"],
    module = "github.com/subosito/gotenv",
    version = "v1.2.0",
)

go_module(
    name = "ini_v1",
    hashes = ["175543eaa9b210000e71fb67bdc44b610439915e66e7ec9cd3ee1f7a90126639"],
    licences = ["Apache-2.0"],
    module = "gopkg.in/ini.v1",
    version = "v1.62.0",
)

go_module(
    name = "stretchr_testify",
    hashes = ["274065a0b7acec70df07119ba7a7f0f294c42df9f9555286147b1ea94cf4e3dd"],
    install = [
        "assert",
        "mock",
        "require",
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
    hashes = ["5471affa54995ff549f3552d0ff9010dcf41fc6c67a61d9c0754e9ec8034a634"],
    licences = ["Apache-2.0"],
    module = "gopkg.in/yaml.v3",
    version = "496545a6307b2a7d7a710fd516e5e16e8ab62dbc",
)

go_module(
    name = "davecgh_go-spew",
    hashes = ["41d8bec76512ded0c55eaeacfb5167214d602df18ea04c254c336bb9b18e5d05"],
    install = [
        "spew",
    ],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_module(
    name = "pmezard_go-difflib",
    hashes = ["29a36405fc7d4d8b004644927774412693295d9a9a49e5260123f74d07006344"],
    install = [
        "difflib",
    ],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_module(
    name = "stretchr_objx",
    hashes = ["08456456036b6e9e76d8fd0343f2d9e5f7bd5d4bdbee7d1b16e835c9ba7558d9"],
    module = "github.com/stretchr/objx",
    version = "v0.1.1",
)

go_module(
    name = "fsnotify",
    hashes = ["f45ce25fa7d31128ddd2c5d1627b6d91361d288d6602ee725fedb8e24b14b72e"],
    module = "github.com/fsnotify/fsnotify",
    version = "v1.4.7",
    deps = [
        ":x_sys",
    ],
)

go_module(
    name = "x_sys",
    hashes = ["e9ba045ac49a174a9bb1220fc50e89b46e26c8ea18d987248baec06478c9c7bf"],
    install = [
        "unix",
        "cpu",
    ],
    module = "golang.org/x/sys",
    version = "95b1ffbd15a57cc5abb3f04402b9e8ec0016a52c",
)

go_module(
    name = "hashicorp_hcl",
    hashes = ["84ed05e07626b15b91bd9c6b40815385efc5f176c0e608800795c7143dc25d41"],
    install = [
        ".",
        "...",
    ],
    module = "github.com/hashicorp/hcl",
    version = "v1.0.0",
)

go_module(
    name = "magiconair_properties",
    hashes = ["119de712236037975e29d5c8a3fc10c9a8fd3dfc35aaa9a4184cd023ac81e692"],
    module = "github.com/magiconair/properties",
    version = "v1.8.0",
)

go_module(
    name = "mitchellh_mapstructure",
    hashes = ["439cdf7b4de1becfc42b82a0339e90b47dffa778040d2873e4402d236a95a5f2"],
    module = "github.com/mitchellh/mapstructure",
    version = "v1.1.2",
)

go_module(
    name = "pelletier_go-toml",
    hashes = ["a40df99183325ad5201f79a502f94e14fad72cd21691924bce6da18f24ae6d16"],
    module = "github.com/pelletier/go-toml",
    version = "v1.2.0",
)

go_module(
    name = "spf13_cast",
    hashes = ["0506c99ac955149543f8e1b12b0880a543a16914e5a775caa180777dd57f8866"],
    module = "github.com/spf13/cast",
    version = "v1.3.0",
)

go_module(
    name = "spf13_jwalterweatherman",
    hashes = ["6f24d0ecfc60741aa131b3b1332efde2d24bb4daa43bc2d83b96dc9613914c00"],
    module = "github.com/spf13/jwalterweatherman",
    version = "v1.0.0",
)

go_module(
    name = "yaml.v2",
    hashes = ["f4de942d0effb706f8c541ffa4da0f976ae40b890c3fbaab4939f65a8f1405f7"],
    module = "gopkg.in/yaml.v2",
    version = "v2.2.2",
)

go_module(
    name = "spf13_afero",
    hashes = ["b1d49c29819db7fc8853894a61d11a165f5a0af774fccc1217d42efcf5535cf4"],
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
    hashes = ["2c8ecc0a55d3e39182d644b3161232950fe3070b0c4885d776fcb68c49f3c7ea"],
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
    hashes = ["9cf06648897420904b5319c76908d8f27b109da1d2a1040f038dde64e1dbca84"],
    module = "github.com/spf13/pflag",
    version = "v1.0.3",
)

go_module(
    name = "pkg_errors",
    hashes = ["f20d68bc6301675e4662897d0086ab7610062d494d0792f076f7e8eebe4aa8f5"],
    module = "github.com/pkg/errors",
    version = "v0.8.1",
)

go_module(
    name = "evanphx_json-patch",
    hashes = ["fbf04a7dfd189625bff40d232be9c757c7201ab9e20ec59047a164f08f6165f8"],
    module = "github.com/evanphx/json-patch",
    version = "v4.5.0",
    deps = [
        ":pkg_errors",
    ],
)

go_module(
    name = "ghodss_yaml",
    hashes = ["8a1a79d62640820df92e2c197fef387f7ec3595c761c022ec147c909a32bdb0a"],
    module = "github.com/ghodss/yaml",
    version = "v1.0.0",
    deps = [
        ":yaml.v2",
    ],
)

KUBERNETES_VERSION = "1.13.12"

go_module(
    name = "apimachinery",
    hashes = ["1a8f138862f20f444ac2b0bcb9fc0f8249cc1fec4a2cdefa0e025ecad0506065"],
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
    hashes = ["fb4f53aa4452d2b03490cb7b854c3e6fc78a60e1ee3c38ddc3ff6b467be0a1ac"],
    module = "github.com/google/gofuzz",
    version = "v1.0.0",
)

go_module(
    name = "inf.v0",
    hashes = ["4667437f9b9cf0308643b7c6505257857a403c7e0f0ca2a7ab219240234b88d7"],
    module = "gopkg.in/inf.v0",
    version = "v0.9.1",
)

go_module(
    name = "x_net",
    hashes = ["30ad515d77b7dee160c3cd64871edcb88205a4e04ef14825f335556cb021578d"],
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
    hashes = ["0e223a4730f3eaeaa878c9cd880fd44850fe11f799a0fd70394be179c5de5398"],
    module = "k8s.io/klog",
    version = "v1.0.0",
)

go_module(
    name = "golang-migrate_migrate",
    hashes = ["bb960ab46b4c6b1a6b31bd7b00594c329632b54e7cc6dfe3ba67ef31189f34aa"],
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
    hashes = ["e763ffaddb979585a6a3d9ddbd444c1e8b8940e958994d8b5aab42465bf04591"],
    module = "github.com/hashicorp/go-multierror",
    version = "v1.0.0",
    deps = [
        ":hashicorp_errwrap",
    ],
)

go_module(
    name = "hashicorp_errwrap",
    hashes = ["4e7e911378220893d62e7807728529a06fd6ab10adc4b447277151dd69dabe26"],
    module = "github.com/hashicorp/errwrap",
    version = "v1.0.0",
)

go_module(
    name = "lib_pq",
    hashes = ["3d0e8ad6b65db01924c9506a8ae966ed42c56d9ec1785e57fd13e5c30ef0a0e3"],
    install = [
        ".",
        "oid",
    ],
    module = "github.com/lib/pq",
    version = "v1.0.0",
)

go_module(
    name = "jmoiron_sqlx",
    hashes = ["9e933a4d192e53183e3930f29283e3929f2a60fed359cfdba8b3ec8f413de21a"],
    install = [
        ".",
        "...",
    ],
    module = "github.com/jmoiron/sqlx",
    version = "v1.2.0",
)

go_mod_download(
    name = "mockgen_download",
    hashes = ["b1d226ce7c379454d7cabc008d6dc4e15da1aa004bc34ff79dea01dec7dabba6"],
    module = "github.com/golang/mock",
    version = "v1.4.4",
)

go_module(
    name = "mockgen",
    binary = True,
    download = ":mockgen_download",
    hashes = ["b1d226ce7c379454d7cabc008d6dc4e15da1aa004bc34ff79dea01dec7dabba6"],
    install = ["mockgen"],
    licences = ["apache-2.0"],
    module = "github.com/golang/mock",
    deps = [":x_tools"],
)

go_module(
    name = "mock",
    download = ":mockgen_download",
    hashes = ["9c0a6d6312dc36e23d8738d2d15d0dfa06cb996ab7f45e0d124cb6d65fbdb40f"],
    install = ["..."],
    licences = ["apache-2.0"],
    module = "github.com/golang/mock",
    deps = [
        ":x_tools",
    ],
)

go_module(
    name = "h2non_parth",
    hashes = ["e2cee03c005ff4bbaf0aa5180e225fe335d2a6773b3e8ab635103cbdf42fafbe"],
    licences = ["MIT"],
    module = "github.com/h2non/parth",
    version = "v0.0.0-20190131123155-b4df798d6542",
)

go_module(
    name = "h2non_gentleman",
    hashes = ["cbea4830e829135e69d3db6e0a74684e26f1442007812c53c929712ad75f661e"],
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
    hashes = ["e5c3c1fd3a62093da07f38b5f0d952b50c769118f5aa5010e9ecb9f17e0bf014"],
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
    hashes = ["dff907b66b1f2cfd3536c92d71dfd5f28078026249f850214ee1a860519decc6"],
    install = [
        "cover",
        "present",
        "go/packages",
        "go/internal/packagesdriver",
        "internal/gocommand",
        "internal/packagesinternal",
        "go/gcexportdata",
        "go/internal/gcimporter",
        "go/ast/astutil",
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
    hashes = ["b4db6037c0ecc342a86703ddd215028f1574187d5aed688196e2e91c3c9523c0"],
    install = ["..."],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/crypto",
    strip = ["acme/autocert"],
    version = "123391ffb6de907695e1066dc40c1ff09322aeb6",
    deps = [":x_sys"],
)

go_module(
    name = "x_errors",
    hashes = ["10780b586a60adfdea2f4d76dbaf3025eb3a8165376153c7febfb36516d1ec72"],
    install = [
        ".",
        "...",
    ],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/xerrors",
    version = "a5947ffaace3e882f334c1750858b4a6a7e52422",
    deps = [":x_sys"],
)

go_module(
    name = "x_mod",
    hashes = ["2a937734dda11bfc94447ce43128e5ffafcf0479a4f4a4558aa17baf7335c2d7"],
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
    hashes = ["e3fa3cce6777b75261be26b66d4fcc8f19539b6f30407b64cf0679858f0892c8"],
    install = ["..."],
    licences = ["bsd-3-clause"],
    module = "golang.org/x/sync",
    version = "fd80eb99c8f653c847d294a001bdf2a3a6f768f5",
    deps = [
        ":x_net",
    ],
)
