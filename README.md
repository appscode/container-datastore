
[![Go Report Card](https://goreportcard.com/badge/github.com/kubedb/cli)](https://goreportcard.com/report/github.com/kubedb/cli)

[Website](https://kubedb.com) • [Slack](http://slack.kubernetes.io) • [Twitter](https://twitter.com/KubeDB)

# KubeDB
Running production quality database in Kubernetes can be tricky to say the least. In the early days of Kubernetes, replication controllers were used to run a single pod for a database. With the introduction of StatefulSet, it became easy to run a docker container for any database. But what about monitoring, taking periodic backups, restoring from backups or cloning from an existing database? KubeDB is a framework for writing operators for any database that support the following operational requirements:

 - Create a database declaratively using CRD.
 - Take one-off backups or periodical backups to various cloud stores, eg, S3, GCS, Azure etc.
 - Restore from backup or clone any database.
 - Native integration with Prometheus for monitoring via [CoreOS Prometheus Operator](https://github.com/coreos/prometheus-operator).
 - Apply deletion lock to avoid accidental deletion of database.
 - Keep track of deleted databases, cleanup prior snapshots with a single command.
 - Use cli to manage databases like kubectl for Kubernetes.

KubeDB is developed at [AppsCode](https://twitter.com/AppsCodeHQ) to run their SAAS platform on Kubernetes. Currently we include complete implementations for Postgres and Elasticsearch database.

## Supported Versions
Please pick a version of `kubedb` that matches your Kubernetes installation.

| KubeDB Version                                                      | Docs                                                       | Kubernetes Version |
|---------------------------------------------------------------------|------------------------------------------------------------|--------------------|
| [0.7.1](https://github.com/kubedb/cli/releases/tag/0.7.1) (uses CRD) | [User Guide](https://github.com/kubedb/cli/tree/0.7.1/docs) | 1.7.x+             |
| [0.6.0](https://github.com/kubedb/cli/releases/tag/0.6.0) (uses TPR) | [User Guide](https://github.com/kubedb/cli/tree/0.6.0/docs) | 1.5.x - 1.7.x      |

## Installation
To install `kubedb`, please follow the guide [here](/docs/install.md).

## Using KubeDB
Want to learn how to use `kubedb`? Please start [here](/docs/tutorials/README.md).

## Contribution guidelines
Want to help improve KubeDB? Please start [here](/CONTRIBUTING.md).

## Project Status
Wondering what features are coming next? Please visit [here](/ROADMAP.md).

---

**The KubeDB operator collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, talk to us on the [Kubernetes Slack team](http://slack.kubernetes.io/) channel `#kubedb`.

If you have found a bug with KubeDB or want to request for new features, please file issues under [kubedb/project](https://github.com/kubedb/project/issues/new) repository.
