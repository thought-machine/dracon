#!/bin/sh -e
# Enricher
statik -src=enrichment_service/configs/sql/migrations -dest=pkg/enrichment/db -p migrations
# Dracon CLI
statik -src=resources/patches -dest=pkg/template -p patches
