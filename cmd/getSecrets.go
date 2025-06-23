/*
Copyright © 2025 Dmytro Vyshniakov <vishdi@gmail.com>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve secrets from the specified namespace with the given label selector",
	Run: func(cmd *cobra.Command, args []string) {
		getSecrets()
	},
}

func getSecrets() ([]corev1.Secret, error) {
	// Retrieve secrets from the specified namespace with the given label selector
	secretsList, err := clientset.CoreV1().Secrets(srcNamespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		loggerStderr.Error().Err(err).Msg("Error retrieving secrets")

	}

	if len(secretsList.Items) == 0 {
		loggerStdout.Info().Str("namespace", targetNamespace).Msg("No secrets found in")
		return secretsList.Items, nil
	}

	loggerStdout.Info().Str("namespace", srcNamespace).Msg("List of secrets in")
	if labelSelector != "" {
		loggerStdout.Info().Str("label selector", labelSelector).Msg("With")
	}

	for _, secret := range secretsList.Items {
		loggerStdout.Info().Str("secret.Name", secret.Name).Msg("    ")
	}
	return secretsList.Items, nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
