## Yugabyte deployment

This directory holds the output of templating the yugabyte helm chart values 
from the script `scripts/yugabyte-*.sh`

The `yugabyte-*.sh` scripts are used to generate the `yugabyte.yaml` file
which is then used to deploy yugabyte on the kubernetes cluster.

## Steps to deploy yugabyte on kubernetes manually

For example, if experimenting with our sandbox cluster:

1. Run the script `scripts/yugabyte-single-az.sh us-east-2 $NAMESPACE sandbox`
2. check for connectivity `kubectl exec --namespace $NAMESPACE -it yb-tserver-0 -- /home/yugabyte/bin/ysqlsh -h yb-tserver-0.yb-tservers.$NAMESPACE`
3. Optionally, create the gorm database `scripts/create-gorm-db.sh $NAMESPACE`

## Undeploying

Note: Pvcs will not be deleted by default when uninstalling
yugabyte with `helm uninstall yugabyte`.
To delete them, you can run the following command:

```bash
kubectl delete pvc -l 'app in (yb-master,yb-tserver)'
```