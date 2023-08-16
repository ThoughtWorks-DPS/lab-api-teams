#!/usr/bin/env bash
set -e

# Note: You will only be able to use this script when you have 3 AZs in your cluster.
#  Meaning you must have at least one node in each AZ (a, b, c).

export AWS_DEFAULT_REGION=$1

cat <<EOF > yugabyte/yugabyte-${AWS_DEFAULT_REGION}a.yaml
isMultiAz: True
AZ: ${AWS_DEFAULT_REGION}a
masterAddresses: "yb-master-0.yb-masters.yugabyte-one.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-two.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-three.svc.cluster.local:7100"
enableLoadBalancer: False
storage:
  master:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
  tserver:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
replicas:
  master: 1
  tserver: 1
  totalMasters: 3
gflags:
  master:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}a"
  tserver:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}a"

EOF

cat <<EOF > yugabyte/yugabyte-${AWS_DEFAULT_REGION}b.yaml
isMultiAz: True
AZ: ${AWS_DEFAULT_REGION}b
masterAddresses: "yb-master-0.yb-masters.yugabyte-one.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-two.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-three.svc.cluster.local:7100"
enableLoadBalancer: False
storage:
  master:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
  tserver:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
replicas:
  master: 1
  tserver: 1
  totalMasters: 3
gflags:
  master:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}b"
  tserver:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}b"
EOF

cat <<EOF > yugabyte/yugabyte-${AWS_DEFAULT_REGION}c.yaml
isMultiAz: True
AZ: ${AWS_DEFAULT_REGION}c
masterAddresses: "yb-master-0.yb-masters.yugabyte-one.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-two.svc.cluster.local:7100,yb-master-0.yb-masters.yugabyte-three.svc.cluster.local:7100"
enableLoadBalancer: False
storage:
  master:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
  tserver:
    storageClass: "prod-${AWS_DEFAULT_REGION}-ebs-storage-class"
replicas:
  master: 1
  tserver: 1
  totalMasters: 3
gflags:
  master:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}c"
  tserver:
    placement_cloud: "aws"
    placement_region: "${AWS_DEFAULT_REGION}"
    placement_zone: "${AWS_DEFAULT_REGION}c"
EOF

kubectl create namespace "yugabyte-one" --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace "yugabyte-two" --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace "yugabyte-three" --dry-run=client -o yaml | kubectl apply -f -
helm repo add yugabytedb https://charts.yugabyte.com

helm upgrade --install yugabyte-one yugabytedb/yugabyte \
  --version 2.19.0 \
  --namespace yugabyte-one \
  -f yugabyte/yugabyte-${AWS_DEFAULT_REGION}a.yaml --wait

 helm upgrade --install yugabyte-two yugabytedb/yugabyte \
  --version 2.19.0 \
  --namespace yugabyte-two \
  -f yugabyte/yugabyte-${AWS_DEFAULT_REGION}b.yaml --wait

 helm upgrade --install yugabyte-three yugabytedb/yugabyte \
  --version 2.19.0 \
  --namespace yugabyte-three \
  -f yugabyte/yugabyte-${AWS_DEFAULT_REGION}c.yaml --wait


