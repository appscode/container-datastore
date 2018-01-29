package cmds

import (
	"fmt"
	"io"

	"github.com/kubedb/apimachinery/apis/kubedb"
	apps "k8s.io/api/apps/v1beta1"
	batch "k8s.io/api/batch/v1"
	core "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	rbac "k8s.io/api/rbac/v1beta1"
	storage "k8s.io/api/storage/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const ServiceAccountName = operatorName

var policyRuleOperator = []rbac.PolicyRule{
	{
		APIGroups: []string{apiextensions.GroupName},
		Resources: []string{"customresourcedefinitions"},
		Verbs:     []string{"create", "delete", "get", "list"},
	},
	{
		APIGroups: []string{extensions.GroupName},
		Resources: []string{"thirdpartyresources"},
		Verbs:     []string{"create", "delete", "get", "list"},
	},
	{
		APIGroups: []string{rbac.GroupName},
		Resources: []string{"rolebindings", "roles"},
		Verbs:     []string{"create", "delete", "get", "patch"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"services"},
		Verbs:     []string{"create", "delete", "get", "patch"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"secrets", "serviceaccounts"},
		Verbs:     []string{"create", "delete", "get"},
	},
	{
		APIGroups: []string{apps.GroupName},
		Resources: []string{"deployments", "statefulsets"},
		Verbs:     []string{"create", "delete", "get", "patch", "update"},
	},
	{
		APIGroups: []string{batch.GroupName},
		Resources: []string{"jobs"},
		Verbs:     []string{"create", "delete", "get"},
	},
	{
		APIGroups: []string{storage.GroupName},
		Resources: []string{"storageclasses"},
		Verbs:     []string{"get"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"pods"},
		Verbs:     []string{"deletecollection", "get", "list", "patch", "watch"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"persistentvolumeclaims"},
		Verbs:     []string{"delete", "get", "list", "watch"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"configmaps"},
		Verbs:     []string{"create", "delete", "get", "update"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"events"},
		Verbs:     []string{"create"},
	},
	{
		APIGroups: []string{core.GroupName},
		Resources: []string{"nodes"},
		Verbs:     []string{"get", "list", "watch"},
	},
	{
		APIGroups: []string{kubedb.GroupName},
		Resources: []string{rbac.ResourceAll},
		Verbs:     []string{rbac.VerbAll},
	},
	{
		APIGroups: []string{"monitoring.coreos.com"},
		Resources: []string{"servicemonitors"},
		Verbs:     []string{"create", "delete", "get", "list", "update"},
	},
}

func EnsureRBACStuff(client kubernetes.Interface, namespace string, out io.Writer) error {
	name := ServiceAccountName
	// Ensure ClusterRoles for operator
	clusterRoleOperator, err := client.RbacV1beta1().ClusterRoles().Get(name, metav1.GetOptions{})
	if err != nil {
		if !kerr.IsNotFound(err) {
			return err
		}
		// Create new one
		role := &rbac.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name:   name,
				Labels: operatorLabel,
			},
			Rules: policyRuleOperator,
		}
		if _, err := client.RbacV1beta1().ClusterRoles().Create(role); err != nil {
			return err
		}
		fmt.Fprintln(out, "Successfully created cluster role.")
	} else {
		// Update existing one
		clusterRoleOperator.Rules = policyRuleOperator
		if _, err := client.RbacV1beta1().ClusterRoles().Update(clusterRoleOperator); err != nil {
			return err
		}
		fmt.Fprintln(out, "Successfully updated cluster role.")
	}

	// Ensure ServiceAccounts
	if _, err := client.CoreV1().ServiceAccounts(namespace).Get(name, metav1.GetOptions{}); err != nil {
		if !kerr.IsNotFound(err) {
			return err
		}
		sa := &core.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    operatorLabel,
			},
		}
		if _, err := client.CoreV1().ServiceAccounts(namespace).Create(sa); err != nil {
			return err
		}
		fmt.Fprintln(out, "Successfully created service account.")
	}

	var roleBindingRef = rbac.RoleRef{
		APIGroup: rbac.GroupName,
		Kind:     "ClusterRole",
		Name:     name,
	}
	var roleBindingSubjects = []rbac.Subject{
		{
			Kind:      rbac.ServiceAccountKind,
			Name:      name,
			Namespace: namespace,
		},
	}

	// Ensure ClusterRoleBindings
	roleBinding, err := client.RbacV1beta1().ClusterRoleBindings().Get(name, metav1.GetOptions{})
	if err != nil {
		if !kerr.IsNotFound(err) {
			return err
		}

		roleBinding := &rbac.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    operatorLabel,
			},
			RoleRef:  roleBindingRef,
			Subjects: roleBindingSubjects,
		}

		if _, err := client.RbacV1beta1().ClusterRoleBindings().Create(roleBinding); err != nil {
			return err
		}
		fmt.Fprintln(out, "Successfully created cluster role bindings.")
	} else {
		roleBinding.RoleRef = roleBindingRef
		roleBinding.Subjects = roleBindingSubjects
		if _, err := client.RbacV1beta1().ClusterRoleBindings().Update(roleBinding); err != nil {
			return err
		}
		fmt.Fprintln(out, "Successfully updated cluster role bindings.")
	}

	return nil
}
