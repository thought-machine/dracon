FROM //build/docker:dracon-base-go

COPY jira_c /consume
COPY config.yaml /config/config.yaml

ENV DRACON_JIRA_CONFIG_PATH="/config/config.yaml"

ENTRYPOINT ["/consume"]
