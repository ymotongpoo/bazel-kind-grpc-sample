#!/bin/bash

echo "PROJECT_ID $(gcloud config list --format 'value(core.project)' 2>/dev/null)"
echo "REPO demo"
echo "K8S_CLUSTER_NAME $(kubectl config view --minify -o=jsonpath='{.contexts[0].context.cluster}')"
