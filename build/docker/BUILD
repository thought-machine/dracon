subinclude("@pleasings//docker")

# Can be used to build Dracon if you don't want to set up a build environment
docker_image(
    name = "dracon-builder",
    dockerfile = "Dockerfile-dracon-builder",
    image = "dracon-builder-go",
)

docker_image(
    name = "dracon-base-go",
    dockerfile = "Dockerfile-base-go",
    image = "dracon-base-go",
    visibility = ["PUBLIC"],
)

docker_image(
    name = "bandit",
    dockerfile = "Dockerfile-bandit",
    image = "dracon-tool-bandit",
)

docker_image(
    name = "spotbugs",
    dockerfile = "Dockerfile-spotbugs",
    image = "dracon-tool-spotbugs",
)
