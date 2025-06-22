package cmd

import (
	"context"
	"fmt"

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
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Copying %d secrets to destination namespace:\n", len(secrets))
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
			fmt.Printf("Secret '%s' already exists in target namespace '%s' or error occurred: %v\n", secret.Name, targetNamespace, err)
		} else {
			fmt.Printf("Secret '%s' copied to namespace '%s'\n", secret.Name, targetNamespace)
		}
	}
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
