package cmd

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy secrets to a specified namespace",
	Run: func(cmd *cobra.Command, args []string) {
		copySecrets()
	},
}

func copySecrets() {
	secrets, err := getSecrets()

	if err != nil {
		loggerStderr.Error().Err(err).Msg("Error retrieving secrets")
		return
	}
	secretsCount := strconv.Itoa(len(secrets))
	loggerStdout.Info().Str("Count", secretsCount).Msg(("Copying secrets to destination namespace:"))
	for _, secret := range secrets {
		// Prepare new Secret object
		newSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secret.Name,
				Namespace: targetNamespace,
				Labels:    secret.Labels,
			},
			Data: secret.Data,
			Type: secret.Type,
		}

		// Create the Secret in the target namespace
		_, err := clientset.CoreV1().Secrets(targetNamespace).Create(context.TODO(), newSecret, metav1.CreateOptions{})
		if err != nil {
			// If already exists - skip and warn
			loggerStdout.Warn().Err(err).Str("namespace", targetNamespace).Msg("Warning:")
		} else {
			loggerStdout.Info().Str("namespace", targetNamespace).Str("secret.Name", secret.Name).Msg("Secret copied to:")
		}
	}
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
