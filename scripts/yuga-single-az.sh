#!/usr/bin/env bash
set -e

# Note: You will only be able to use this script when you have 3 AZs in your cluster.
#  Meaning you must have at least one node in each AZ (a, b, c).

export AWS_DEFAULT_REGION=$1
export NAMESPACE=$2

cat <<EOF > yugabyte/yugabyte.yaml
enableLoadBalancer: False
storage:
  master:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
  tserver:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
replicas:
  master: 1
  tserver: 1
  totalMasters: 1
resource:
  master:
    requests:
      cpu: 2
      memory: 2Gi
    limits:
      cpu: 2
      memory: 2Gi
  tserver:
    requests:
      cpu: 2
      memory: 4Gi
    limits:
      cpu: 2
      memory: 4Gi

EOF

helm repo add yugabytedb https://charts.yugabyte.com

helm upgrade --install yugabyte yugabytedb/yugabyte \
  --version 2.19.0 \
  --namespace ${NAMESPACE} \
  -f yugabyte/yugabyte.yaml --wait

# gflags:
#   master:
#     placement_cloud: "aws"
#     placement_region: "${AWS_DEFAULT_REGION}"
#     placement_zone: "${AWS_DEFAULT_REGION}a"
#   tserver:
#     placement_cloud: "aws"
#     placement_region: "${AWS_DEFAULT_REGION}"
#     placement_zone: "${AWS_DEFAULT_REGION}a"

kubectl exec -n ${NAMESPACE} -it yb-tserver-0 -- ysqlsh  -c 'create database gorm;' 


