/*
Copyright © 2025 Dmytro Vyshniakov <vishdi@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var secretsList string
var srcNamespace string = "nginx-ingress"
var labelSelector string = "secrets-store.csi.k8s.io/managed=true"
var targetNamespace string = "default"

var clientset *kubernetes.Clientset

var rootCmd = &cobra.Command{
	Use:   "controller",
	Short: "CLI to retrieve secrets from Kubernetes",
	Long:  "CLI utility for copying secrets-store.csi.k8s.io generated secrets from nginx-ingress namespace",

	//Creating a new or continue using existing client for connection to Kubernetes client
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if clientset != nil {
			fmt.Println("Using existing kubernetes connection...")
			return nil
		}

		fmt.Println("Trying to get config from KUBECONFIG env variable...")
		kubeConfig := os.Getenv("KUBECONFIG")
		if kubeConfig == "" {
			fmt.Println("No KUBECONFIG env variable found, using default path ~/.kube/config")
			kubeConfig = clientcmd.RecommendedHomeFile
		}

		config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return fmt.Errorf("Error loading kubeconfig: %w", err)
		}
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			return fmt.Errorf("Error creating Kubernetes client: %w", err)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
