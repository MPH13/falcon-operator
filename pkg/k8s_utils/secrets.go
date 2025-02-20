package k8s_utils

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeQuerySecretsMethod func(ctx context.Context) (*corev1.SecretList, error)

func QuerySecrets(namespace string, cli client.Client) KubeQuerySecretsMethod {
	return func(ctx context.Context) (*corev1.SecretList, error) {
		secrets := &corev1.SecretList{}
		err := cli.List(ctx, secrets, client.InNamespace(namespace))
		if err != nil {
			return nil, err
		}
		return secrets, nil
	}
}
