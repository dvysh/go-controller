/*
Copyright © 2025 Dmytro Vyshniakov <vishdi@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve secrets from the specified namespace with the given label selector",
	Run: func(cmd *cobra.Command, args []string) {
		getSecrets()
	},
}

func getSecrets() {
	// Retrieve secrets from the specified namespace with the given label selector
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		log.Fatalf("Error retrieving secrets: %v", err)
	}

	if len(secrets.Items) == 0 {
		fmt.Printf("No secrets found in namespace '%s' with label '%s'.\n", namespace, labelSelector)
		return
	}

	fmt.Printf("List of secrets in namespace '%s'", namespace)
	if labelSelector != "" {
		fmt.Printf(" with label '%s'", labelSelector)
	}
	fmt.Println(":")

	for _, secret := range secrets.Items {
		fmt.Println("- ", secret.Name)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
