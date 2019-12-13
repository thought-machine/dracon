package version

// BuildVersion is the version of dracon. It's intended to be overriden using
// -X the linker flag.
var BuildVersion = "dev"

// TODO(hjenkins): Implement fetching the k8s and tekton version from the server
