#!/usr/bin/env bash
set -e

export NAMESPACE=$1

if [ -z "$NAMESPACE" ]; then
  echo "Usage: $0 <namespace>"
  exit 1
fi

hostname="yb-tserver-0.yb-tservers.$NAMESPACE"
sleep_seconds=5
retries=24
attempt=1

check_connectivity() {
  echo "Checking DB Connectivity for $1"
  if kubectl exec --namespace "${NAMESPACE}" -it yb-tserver-0 -c yb-tserver -- ysqlsh -h "$1" -c '\q'; then
    echo "Connection successful"
    return 0
  else
    echo "Connection failed"
    return 1
  fi
}

while [ $attempt -le $retries ]; do
  if check_connectivity "$hostname"; then
    break
  fi
  echo "Attempt $attempt of $retries failed."
  if [ $attempt -eq $retries ]; then
    echo "All attempts failed. Exiting..."
    exit 1
  fi
  sleep $sleep_seconds
  ((attempt++))
done

result=$(kubectl exec -n "${NAMESPACE}" -it yb-tserver-0 -c yb-tserver -- ysqlsh -U yugabyte -h "$hostname" -tc "SELECT 1 FROM pg_database WHERE datname = 'gorm'")
exists=$(echo "$result" | tr -d '[:space:]')

if [[ "$exists" == "1" ]]; then
  echo "Database gorm already exists"
else
  echo "Creating database gorm"
  kubectl exec -n "${NAMESPACE}" -it yb-tserver-0 -c yb-tserver -- ysqlsh -U yugabyte -h "$hostname" -c "CREATE DATABASE gorm"
fi
