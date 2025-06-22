/*
Copyright © 2025 Dmytro Vyshniakov <vishdi@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

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
		log.Fatalf("Error retrieving secrets: %v", err)
	}

	if len(secretsList.Items) == 0 {
		fmt.Printf("No secrets found in namespace '%s' with label '%s'.\n", srcNamespace, labelSelector)
		return secretsList.Items, nil
	}

	fmt.Printf("List of secrets in namespace '%s'", srcNamespace)
	if labelSelector != "" {
		fmt.Printf(" with label '%s'", labelSelector)
	}
	fmt.Println(":")

	for _, secret := range secretsList.Items {
		fmt.Println("- ", secret.Name)
	}
	return secretsList.Items, nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
