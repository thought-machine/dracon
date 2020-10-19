FROM //build/docker:dracon-base-go

RUN apk add --update --no-cache jq

COPY golang_gosec /parse

ENTRYPOINT ["/parse"]
