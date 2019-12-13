# Developing

We are using [Please](https://please.build) to build Dracon's binaries and docker containers. This abstracts away needing to understand the underling programming language or tooling.

## Running Dracon from Please

To run the Dracon command with the `--help` flag, we can:

```
plz run //cmd/dracon:dracon -- --help
```

This does everything a build system should, downloading dependencies, calling the compiler then finally calling the application with your arguments.

## Building with please in docker

If you don't want to install please in your development environment, there's the option of using our docker build container.

For example building all artefacts (all the binaries), you can use:

```bash
docker run --rm -it \
  --volume "${PWD}:/src" \
  --user $(id -u):$(id -g) \
  thoughtmachine/dracon-builder-go:53322138724d569c4ff037dc36443bb5e0107aecd85e39b094ed8689b6cbc9dc \
  plz build
```

Building the docker images is a little bit more tricky as you would need to mount the docker socket in to the container. At the moment we don't support doing this. Instead, you'll need to install please.
