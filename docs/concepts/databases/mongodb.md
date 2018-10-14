---
title: MongoDB
menu:
  docs_0.9.0-beta.0:
    identifier: mongodb-db
    name: MongoDB
    parent: databases
    weight: 20
menu_name: docs_0.9.0-beta.0
section_menu_id: concepts
---

> New to KubeDB? Please start [here](/docs/concepts/README.md).

# MongoDB

## What is MongoDB

`MongoDB` is a Kubernetes `Custom Resource Definitions` (CRD). It provides declarative configuration for [MongoDB](https://www.mongodb.com/) in a Kubernetes native way. You only need to describe the desired database configuration in a MongoDB object, and the KubeDB operator will create Kubernetes objects in the desired state for you.

## MongoDB Spec

As with all other Kubernetes objects, a MongoDB needs `apiVersion`, `kind`, and `metadata` fields. It also needs a `.spec` section. Below is an example MongoDB object.

```yaml
apiVersion: kubedb.com/v1alpha1
kind: MongoDB
metadata:
  name: mgo1
  namespace: demo
spec:
  version: "3.4"
  storage:
    storageClassName: "standard"
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 50Mi
  databaseSecret:
    secretName: mgo1-auth
  env:
    - name: MONGO_INITDB_DATABASE
      value: myDB
  nodeSelector:
    disktype: ssd
  init:
    scriptSource:
      gitRepo:
        repository: "https://github.com/kubedb/mongodb-init-scripts.git"
        directory: .
  backupSchedule:
    cronExpression: "@every 6h"
    storageSecretName: mg-snap-secret
    gcs:
      bucket: restic
      prefix: demo
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  doNotPause: true
  monitor:
    agent: prometheus.io/coreos-operator
    prometheus:
      namespace: demo
      labels:
        app: kubedb
      interval: 10s
  resources:
    requests:
      memory: "64Mi"
      cpu: "250m"
    limits:
      memory: "128Mi"
      cpu: "500m"
```

### spec.version

`spec.version` is a required field specifying the version of MongoDB database. Official [MongoDB docker images](https://hub.docker.com/r/library/mongo/tags/) will be used for the specific version.

### spec.storage

Since 0.8.0, `spec.storage` is a required field that specifies the StorageClass of PVCs dynamically allocated to store data for the database. This storage spec will be passed to the StatefulSet created by KubeDB operator to run database pods. You can specify any StorageClass available in your cluster with appropriate resource requests.

- `spec.storage.storageClassName` is the name of the StorageClass used to provision PVCs. PVCs don’t necessarily have to request a class. A PVC with its storageClassName set equal to "" is always interpreted to be requesting a PV with no class, so it can only be bound to PVs with no class (no annotation or one set equal to ""). A PVC with no storageClassName is not quite the same and is treated differently by the cluster depending on whether the DefaultStorageClass admission plugin is turned on.
- `spec.storage.accessModes` uses the same conventions as Kubernetes PVCs when requesting storage with specific access modes.
- `spec.storage.resources` can be used to request specific quantities of storage. This follows the same resource model used by PVCs.

To learn how to configure `spec.storage`, please visit the links below:

- https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims

### spec.databaseSecret

`spec.databaseSecret` is an optional field that points to a Secret used to hold credentials for `mongodb` superuser. If not set, KubeDB operator creates a new Secret `{mongodb-object-name}-auth` for storing the password for `mongodb` superuser for each MongoDB object. If you want to use an existing secret please specify that when creating the MongoDB object using `spec.databaseSecret.secretName`.

This secret contains a `user` key and a `password` key which contains the `username` and `password` respectively for `mongodb` superuser.

Example:

```console
$ kubectl create secret generic mgo1-auth -n demo --from-literal=user=jhon-doe --from-literal=password=6q8u_2jMOW-OOZXk
secret "mgo1-auth" created
```

```ini
apiVersion: v1
data:
  password: NnE4dV8yak1PVy1PT1pYaw==
  user: amhvbi1kb2U=
kind: Secret
metadata:
  ...
  name: mgo1-auth
  namespace: demo
  resourceVersion: "1702"
  ...
type: Opaque
```

### spec.env

`spec.env` is an optional field that specifies the environment variables to pass to the MongoDB docker image. To know about supported environment variables, please visit [here](https://hub.docker.com/r/_/mongo/).

Note that, Kubedb does not allow `MONGO_INITDB_ROOT_USERNAME` and `MONGO_INITDB_ROOT_PASSWORD` environment variables to set in `spec.env`. If you want to use custom superuser and password, please use `spec.databaseSecret` instead described earlier.

If you try to set `MONGO_INITDB_ROOT_USERNAME` or `MONGO_INITDB_ROOT_PASSWORD` environment variable in MongoDB crd, Kubed operator will reject the request with following error,

```ini
Error from server (Forbidden): error when creating "./mongodb.yaml": admission webhook "mongodb.validators.kubedb.com" denied the request: environment variable MONGO_INITDB_ROOT_USERNAME is forbidden to use in MongoDB spec
```

Also, note that Kubedb does not allow to update the environment variables as updating them does not have any effect once the database is created.  If you try to update environment variables, Kubedb operator will reject the request with following error,

```ini
Error from server (BadRequest): error when applying patch:
...
for: "./mongodb.yaml": admission webhook "mongodb.validators.kubedb.com" denied the request: precondition failed for:
...
spec:map[env:[map[name:MONGO_INITDB_DATABASE value:patched]]]].At least one of the following was changed:
    apiVersion
    kind
    name
    namespace
    spec.version
    spec.storage
    spec.databaseSecret
    spec.nodeSelector
    spec.init
    spec.env
```

### spec.nodeSelector

`spec.nodeSelector` is an optional field that specifies a map of key-value pairs. For the pod to be eligible to run on a node, the node must have each of the indicated key-value pairs as labels (it can have additional labels as well). To learn more, see [here](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector) .

### spec.init

`spec.init` is an optional section that can be used to initialize a newly created MongoDB database. MongoDB databases can be initialized in one of two ways:

#### Initialize via Script

To initialize a MongoDB database using a script (shell script, js script), set the `spec.init.scriptSource` section when creating a MongoDB object. It will execute files alphabetically with extensions `.sh`  and `.js` that are found in the repository. ScriptSource must have following information:

- [VolumeSource](https://kubernetes.io/docs/concepts/storage/volumes/#types-of-volumes): Where your script is loaded from.

Below is an example showing how a shell script from a git repository can be used to initialize a MongoDB database.

```yaml
apiVersion: kubedb.com/v1alpha1
kind: MongoDB
metadata:
  name: mgo1
spec:
  version: 3.4
  init:
    scriptSource:
      gitRepo:
        repository: "https://github.com/kubedb/mongodb-init-scripts.git"
        directory: .
```

In the above example, KubeDB operator will launch a Job to execute all js script of `mongodb-init-script` repo in alphabetical order once StatefulSet pods are running.

#### Initialize from Snapshots

To initialize from prior snapshots, set the `spec.init.snapshotSource` section when creating a MongoDB object. In this case, SnapshotSource must have following information:

- `name:` Name of the Snapshot

```yaml
apiVersion: kubedb.com/v1alpha1
kind: MongoDB
metadata:
  name: mgo1
spec:
  version: 3.4
  init:
    snapshotSource:
      name: "snapshot-xyz"
```

In the above example, MongoDB database will be initialized from Snapshot `snapshot-xyz` in `default` namespace. Here, KubeDB operator will launch a Job to initialize MongoDB once StatefulSet pods are running.

### spec.backupSchedule

KubeDB supports taking periodic snapshots for MongoDB database. This is an optional section in `.spec`. When `spec.backupSchedule` section is added, KubeDB operator immediately takes a backup to validate this information. After that, at each tick kubeDB operator creates a [Snapshot](/docs/concepts/snapshot.md) object. This triggers operator to create a Job to take backup. If used, set the various sub-fields accordingly.

- `spec.backupSchedule.cronExpression` is a required [cron expression](https://github.com/robfig/cron/blob/v2/doc.go#L26). This specifies the schedule for backup operations.
- `spec.backupSchedule.{storage}` is a required field that is used as the destination for storing snapshot data. KubeDB supports cloud storage providers like S3, GCS, Azure and OpenStack Swift. It also supports any locally mounted Kubernetes volumes, like NFS, Ceph, etc. Only one backend can be used at a time. To learn how to configure this, please visit [here](/docs/concepts/snapshot.md).
- `spec.backupSchedule.resources` is an optional field that can request compute resources required by Jobs used to take a snapshot or initialize databases from a snapshot.  To learn more, visit [here](http://kubernetes.io/docs/user-guide/compute-resources/).

### spec.doNotPause

`spec.doNotPause` is an optional field that tells KubeDB operator that if this MongoDB object is deleted, whether it should be reverted automatically. This should be set to `true` for production databases to avoid accidental deletion. If not set or set to false, deleting a MongoDB object put the database into a dormant state. THe StatefulSet for a DormantDatabase is deleted but the underlying PVCs are left intact. This allows users to resume the database later.

### spec.imagePullSecret

`KubeDB` provides the flexibility of deploying MongoDB database from a private Docker registry. To learn how to deploym MongoDB from a private registry, please visit [here](/docs/guides/mongodb/private-registry/using-private-registry.md).

### spec.monitor

MongoDB managed by KubeDB can be monitored with builtin-Prometheus and CoreOS-Prometheus operator out-of-the-box. To learn more,

- [Monitor MongoDB with builtin Prometheus](/docs/guides/mongodb/monitoring/using-builtin-prometheus.md)
- [Monitor MongoDB with CoreOS Prometheus operator](/docs/guides/mongodb/monitoring/using-coreos-prometheus-operator.md)

### spec.resources

`spec.resources` is an optional field. This can be used to request compute resources required by the database pods. To learn more, visit [here](http://kubernetes.io/docs/user-guide/compute-resources/).

## Next Steps

- Learn how to use KubeDB to run a MongoDB database [here](/docs/guides/mongodb/README.md).
- See the list of supported storage providers for snapshots [here](/docs/concepts/snapshot.md).
- Wondering what features are coming next? Please visit [here](/docs/roadmap.md).
- Want to hack on KubeDB? Check our [contribution guidelines](/docs/CONTRIBUTING.md).
