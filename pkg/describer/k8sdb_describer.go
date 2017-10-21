package describer

import (
	"fmt"
	"io"

	tapi "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/printers"
)

const statusUnknown = "Unknown"

func (d *humanReadableDescriber) describeElastic(item *tapi.Elasticsearch, describerSettings *printers.DescriberSettings) (string, error) {
	clientSet, err := d.ClientSet()
	if err != nil {
		return "", err
	}

	snapshots, err := d.extensionsClient.Snapshots(item.Namespace).List(
		metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(
				map[string]string{
					tapi.LabelDatabaseKind: item.ResourceKind(),
					tapi.LabelDatabaseName: item.Name,
				},
			).String(),
		},
	)
	if err != nil {
		return "", err
	}

	var events *kapi.EventList
	if describerSettings.ShowEvents {
		item.Kind = tapi.ResourceKindElasticsearch
		events, err = clientSet.Core().Events(item.Namespace).Search(scheme.Scheme, item)
		if err != nil {
			return "", err
		}
	}

	return tabbedString(func(out io.Writer) error {
		fmt.Fprintf(out, "Name:\t%s\n", item.Name)
		fmt.Fprintf(out, "Namespace:\t%s\n", item.Namespace)
		fmt.Fprintf(out, "CreationTimestamp:\t%s\n", timeToString(&item.CreationTimestamp))
		if item.Labels != nil {
			printLabelsMultiline(out, "Labels", item.Labels)
		}
		fmt.Fprintf(out, "Status:\t%s\n", string(item.Status.Phase))
		if len(item.Status.Reason) > 0 {
			fmt.Fprintf(out, "Reason:\t%s\n", item.Status.Reason)
		}
		fmt.Fprintf(out, "Replicas:\t%d  total\n", item.Spec.Replicas)
		if item.Annotations != nil {
			printLabelsMultiline(out, "Annotations", item.Annotations)
		}

		describeStorage(item.Spec.Storage, out)

		statefulSetName := fmt.Sprintf("%v-%v", item.Name, item.ResourceCode())

		d.describeStatefulSet(item.Namespace, statefulSetName, out)
		d.describeService(item.Namespace, item.Name, out)

		if item.Spec.Monitor != nil {
			describeMonitor(item.Spec.Monitor, out)
		}

		listSnapshots(snapshots, out)

		if events != nil {
			describeEvents(events, out)
		}

		return nil
	})
}

func (d *humanReadableDescriber) describePostgres(item *tapi.Postgres, describerSettings *printers.DescriberSettings) (string, error) {
	clientSet, err := d.ClientSet()
	if err != nil {
		return "", err
	}

	snapshots, err := d.extensionsClient.Snapshots(item.Namespace).List(
		metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(
				map[string]string{
					tapi.LabelDatabaseKind: item.ResourceKind(),
					tapi.LabelDatabaseName: item.Name,
				},
			).String(),
		},
	)
	if err != nil {
		return "", err
	}

	var events *kapi.EventList
	if describerSettings.ShowEvents {
		item.Kind = tapi.ResourceKindPostgres
		events, err = clientSet.Core().Events(item.Namespace).Search(scheme.Scheme, item)
		if err != nil {
			return "", err
		}
	}

	return tabbedString(func(out io.Writer) error {
		fmt.Fprintf(out, "Name:\t%s\n", item.Name)
		fmt.Fprintf(out, "Namespace:\t%s\n", item.Namespace)
		fmt.Fprintf(out, "StartTimestamp:\t%s\n", timeToString(&item.CreationTimestamp))
		if item.Labels != nil {
			printLabelsMultiline(out, "Labels", item.Labels)
		}
		fmt.Fprintf(out, "Status:\t%s\n", string(item.Status.Phase))
		if len(item.Status.Reason) > 0 {
			fmt.Fprintf(out, "Reason:\t%s\n", item.Status.Reason)
		}
		if item.Annotations != nil {
			printLabelsMultiline(out, "Annotations", item.Annotations)
		}

		describeStorage(item.Spec.Storage, out)

		statefulSetName := fmt.Sprintf("%v-%v", item.Name, item.ResourceCode())

		d.describeStatefulSet(item.Namespace, statefulSetName, out)
		d.describeService(item.Namespace, item.Name, out)
		if item.Spec.DatabaseSecret != nil {
			d.describeSecret(item.Namespace, item.Spec.DatabaseSecret.SecretName, "Database", out)
		}

		if item.Spec.Monitor != nil {
			describeMonitor(item.Spec.Monitor, out)
		}

		listSnapshots(snapshots, out)

		if events != nil {
			describeEvents(events, out)
		}

		return nil
	})
}

func (d *humanReadableDescriber) describeSnapshot(item *tapi.Snapshot, describerSettings *printers.DescriberSettings) (string, error) {
	clientSet, err := d.ClientSet()
	if err != nil {
		return "", err
	}

	var events *kapi.EventList
	if describerSettings.ShowEvents {
		item.Kind = tapi.ResourceKindSnapshot
		events, err = clientSet.Core().Events(item.Namespace).Search(scheme.Scheme, item)
		if err != nil {
			return "", err
		}
	}

	return tabbedString(func(out io.Writer) error {
		fmt.Fprintf(out, "Name:\t%s\n", item.Name)
		fmt.Fprintf(out, "Namespace:\t%s\n", item.Namespace)
		fmt.Fprintf(out, "CreationTimestamp:\t%s\n", timeToString(&item.CreationTimestamp))
		if item.Status.CompletionTime != nil {
			fmt.Fprintf(out, "CompletionTimestamp:\t%s\n", timeToString(item.Status.CompletionTime))
		}
		if item.Labels != nil {
			printLabelsMultiline(out, "Labels", item.Labels)
		}
		fmt.Fprintf(out, "Status:\t%s\n", string(item.Status.Phase))
		if len(item.Status.Reason) > 0 {
			fmt.Fprintf(out, "Reason:\t%s\n", item.Status.Reason)
		}
		if item.Annotations != nil {
			printLabelsMultiline(out, "Annotations", item.Annotations)
		}

		d.describeSecret(item.Namespace, item.Spec.StorageSecretName, "Storage", out)

		if events != nil {
			describeEvents(events, out)
		}

		return nil
	})
}

func (d *humanReadableDescriber) describeDormantDatabase(item *tapi.DormantDatabase, describerSettings *printers.DescriberSettings) (string, error) {
	clientSet, err := d.ClientSet()
	if err != nil {
		return "", err
	}

	snapshots, err := d.extensionsClient.Snapshots(item.Namespace).List(
		metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(
				map[string]string{
					tapi.LabelDatabaseKind: item.Labels[tapi.LabelDatabaseKind],
					tapi.LabelDatabaseName: item.Name,
				},
			).String(),
		},
	)
	if err != nil {
		return "", err
	}

	var events *kapi.EventList
	if describerSettings.ShowEvents {
		item.Kind = tapi.ResourceKindDormantDatabase
		events, err = clientSet.Core().Events(item.Namespace).Search(scheme.Scheme, item)
		if err != nil {
			return "", err
		}
	}

	return tabbedString(func(out io.Writer) error {
		fmt.Fprintf(out, "Name:\t%s\n", item.Name)
		fmt.Fprintf(out, "Namespace:\t%s\n", item.Namespace)
		fmt.Fprintf(out, "CreationTimestamp:\t%s\n", timeToString(&item.CreationTimestamp))
		if item.Status.PausingTime != nil {
			fmt.Fprintf(out, "PausedTimestamp:\t%s\n", timeToString(item.Status.PausingTime))
		}
		if item.Status.WipeOutTime != nil {
			fmt.Fprintf(out, "WipeOutTimestamp:\t%s\n", timeToString(item.Status.WipeOutTime))
		}
		if item.Labels != nil {
			printLabelsMultiline(out, "Labels", item.Labels)
		}
		fmt.Fprintf(out, "Status:\t%s\n", string(item.Status.Phase))
		if len(item.Status.Reason) > 0 {
			fmt.Fprintf(out, "Reason:\t%s\n", item.Status.Reason)
		}
		if item.Annotations != nil {
			printLabelsMultiline(out, "Annotations", item.Annotations)
		}

		describeOrigin(item.Spec.Origin, out)

		if item.Status.Phase != tapi.DormantDatabasePhaseWipedOut {
			listSnapshots(snapshots, out)
		}

		if events != nil {
			describeEvents(events, out)
		}

		return nil
	})
}

func describeStorage(pvcSpec *core.PersistentVolumeClaimSpec, out io.Writer) {
	if pvcSpec == nil {
		fmt.Fprint(out, "No volumes.\n")
		return
	}

	accessModes := getAccessModesAsString(pvcSpec.AccessModes)
	val, _ := pvcSpec.Resources.Requests[core.ResourceStorage]
	capacity := val.String()
	fmt.Fprint(out, "Volume:\n")
	if pvcSpec.StorageClassName != nil {
		fmt.Fprintf(out, "  StorageClass:\t%s\n", *pvcSpec.StorageClassName)
	}
	fmt.Fprintf(out, "  Capacity:\t%s\n", capacity)
	fmt.Fprintf(out, "  Access Modes:\t%s\n", accessModes)
}

func describeMonitor(monitor *tapi.MonitorSpec, out io.Writer) {
	if monitor == nil {
		return
	}

	fmt.Fprint(out, "\n")
	fmt.Fprint(out, "Monitoring System:\n")
	fmt.Fprintf(out, "  Agent:\t%s\n", monitor.Agent)
	if monitor.Prometheus != nil {
		prom := monitor.Prometheus
		fmt.Fprint(out, "  Prometheus:\n")
		fmt.Fprintf(out, "    Namespace:\t%s\n", prom.Namespace)
		if prom.Labels != nil {
			printLabelsMultiline(out, "    Labels", prom.Labels)
		}
		fmt.Fprintf(out, "    Interval:\t%s\n", prom.Interval)
	}
}

func listSnapshots(snapshotList *tapi.SnapshotList, out io.Writer) {
	fmt.Fprint(out, "\n")

	if len(snapshotList.Items) == 0 {
		fmt.Fprint(out, "No Snapshots.\n")
		return
	}

	fmt.Fprint(out, "Snapshots:\n")
	w := printers.GetNewTabWriter(out)

	fmt.Fprint(w, "  Name\tBucket\tStartTime\tCompletionTime\tPhase\n")
	fmt.Fprint(w, "  ----\t------\t---------\t--------------\t-----\n")
	for _, e := range snapshotList.Items {
		location, err := e.Spec.SnapshotStorageSpec.Location()
		if err != nil {
			location = "<invalid>"
		}
		fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%s\n",
			e.Name,
			location,
			timeToString(e.Status.StartTime),
			timeToString(e.Status.CompletionTime),
			e.Status.Phase,
		)
	}
	w.Flush()
}

func describeOrigin(origin tapi.Origin, out io.Writer) {
	fmt.Fprint(out, "\n")
	fmt.Fprint(out, "Origin:\n")
	fmt.Fprintf(out, "  Name:\t%s\n", origin.Name)
	fmt.Fprintf(out, "  Namespace:\t%s\n", origin.Namespace)
	if origin.Labels != nil {
		printLabelsMultiline(out, "  Labels", origin.Labels)
	}
	if origin.Annotations != nil {
		printLabelsMultiline(out, "  Annotations", origin.Annotations)
	}
}
