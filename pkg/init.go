package pkg

import (
	"fmt"
	"io"
	"strings"

	"github.com/appscode/go/types"
	"github.com/k8sdb/apimachinery/pkg/docker"
	"github.com/k8sdb/cli/pkg/kube"
	"github.com/k8sdb/cli/pkg/util"
	"github.com/spf13/cobra"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	init_long = templates.LongDesc(`
		Install or upgrade unified operator for kubedb databases.`)

	init_example = templates.Examples(`
		# Install latest released operator.
		kubedb init

		# Upgrade operator to use another version.
		kubedb init --version=0.2.0 --upgrade`)
)

func NewCmdInit(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Create or upgrade unified operator",
		Long:    init_long,
		Example: init_example,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(RunInit(cmd, out, errOut))
		},
	}

	util.AddInitFlags(cmd)
	return cmd
}

func RunInit(cmd *cobra.Command, out, errOut io.Writer) error {
	upgrade := cmdutil.GetFlagBool(cmd, "upgrade")
	namespace := cmdutil.GetFlagString(cmd, "operator-namespace")
	serviceAccount := cmdutil.GetFlagString(cmd, "operator-service-account")
	version := cmdutil.GetFlagString(cmd, "version")

	client, err := kube.NewKubeClient(cmd)
	if err != nil {
		return err
	}

	if upgrade {
		deployment, err := getOperatorDeployment(client, namespace)
		if err != nil {
			if kerr.IsNotFound(err) {
				message := fmt.Sprintf("Operator deployment \"%v\" not found.\n\n"+
					"Create operator using following commnad:\n"+
					"kubedb init --version=%v --operator-namespace=%v", docker.OperatorName, version, namespace)

				fmt.Fprintln(errOut, message)
				return nil
			}

			return err
		}

		containers := deployment.Spec.Template.Spec.Containers
		if len(containers) == 0 {
			fmt.Fprintln(errOut, fmt.Sprintf(`Invalid operator deployment "%v"`, docker.OperatorName))
			return nil
		}

		items := strings.Split(containers[0].Image, ":")

		image := items[0]
		tag := items[1]

		if image != docker.ImageOperator {
			fmt.Fprintln(errOut, fmt.Sprintf(`Operator image mismatch. Can't upgrade to version "%v"`, version))
			return nil
		}

		if tag == version {
			fmt.Fprintln(out, "Operator deployment is already using this version.")
			return nil
		}

		if err := docker.CheckDockerImageVersion(docker.ImageOperator, version); err != nil {
			fmt.Fprintln(errOut, fmt.Sprintf(`Operator image %v:%v not found.`, docker.ImageOperator, version))
			return nil
		}

		deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%v:%v", docker.ImageOperator, version)

		if err := updateOperatorDeployment(client, deployment); err != nil {
			return err
		}

		fmt.Fprintln(out, "Successfully upgraded operator deployment.")
	} else {
		if err := docker.CheckDockerImageVersion(docker.ImageOperator, version); err != nil {
			fmt.Fprintln(errOut, fmt.Sprintf(`Operator image %v:%v not found.`, docker.ImageOperator, version))
			return nil
		}

		if err := createOperatorDeployment(client, namespace, serviceAccount, version); err != nil {
			if kerr.IsAlreadyExists(err) {
				fmt.Fprintln(errOut, "Operator deployment already exists.")
			} else {
				return err
			}
		} else {
			fmt.Fprintln(out, "Successfully created operator deployment.")
		}

		if err := createOperatorService(client, namespace); err != nil {
			if kerr.IsAlreadyExists(err) {
				fmt.Fprintln(errOut, "Operator service already exists.")
			} else {
				return err
			}
		} else {
			fmt.Fprintln(out, "Successfully created operator service.")
		}
	}

	return nil
}

func getOperatorDeployment(client kubernetes.Interface, namespace string) (*extensions.Deployment, error) {
	return client.ExtensionsV1beta1().Deployments(namespace).Get(docker.OperatorName, metav1.GetOptions{})
}

var operatorLabel = map[string]string{
	"app": docker.OperatorName,
}

func createOperatorDeployment(client kubernetes.Interface, namespace, serviceAccount, version string) error {
	deployment := &extensions.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      docker.OperatorName,
			Namespace: namespace,
		},
		Spec: extensions.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: operatorLabel,
			},
			Replicas: types.Int32P(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: operatorLabel,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  docker.OperatorContainer,
							Image: fmt.Sprintf("%v:%v", docker.ImageOperator, version),
							Args: []string{
								"run",
								fmt.Sprintf("--address=:%v", docker.OperatorPortNumber),
								"--v=3",
							},
							Env: []apiv1.EnvVar{
								{
									Name: "OPERATOR_NAMESPACE",
									ValueFrom: &apiv1.EnvVarSource{
										FieldRef: &apiv1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.namespace",
										},
									},
								},
								{
									Name: "OPERATOR_SERVICE_ACCOUNT",
									ValueFrom: &apiv1.EnvVarSource{
										FieldRef: &apiv1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "spec.serviceAccountName",
										},
									},
								},
							},
							Ports: []apiv1.ContainerPort{
								{
									Name:          docker.OperatorPortName,
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: docker.OperatorPortNumber,
								},
							},
						},
					},
					ServiceAccountName: serviceAccount,
				},
			},
		},
	}

	_, err := client.ExtensionsV1beta1().Deployments(namespace).Create(deployment)
	return err
}

func createOperatorService(client kubernetes.Interface, namespace string) error {
	svc := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      docker.OperatorName,
			Namespace: namespace,
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeClusterIP,
			Ports: []apiv1.ServicePort{
				{
					Name:       docker.OperatorPortName,
					Port:       docker.OperatorPortNumber,
					Protocol:   apiv1.ProtocolTCP,
					TargetPort: intstr.FromString(docker.OperatorPortName),
				},
			},
			Selector: operatorLabel,
		},
	}

	_, err := client.CoreV1().Services(namespace).Create(svc)
	return err
}

func updateOperatorDeployment(client kubernetes.Interface, deployment *extensions.Deployment) error {
	_, err := client.ExtensionsV1beta1().Deployments(deployment.Namespace).Update(deployment)
	return err
}
