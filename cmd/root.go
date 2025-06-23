/*
Copyright © 2025 Dmytro Vyshniakov <vishdi@gmail.com>
*/
package cmd

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var secretsList string
var srcNamespace string = "nginx-ingress"
var labelSelector string = "secrets-store.csi.k8s.io/managed=true"
var targetNamespace string = "default"

var outputStdout = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
var loggerStdout = zerolog.New(outputStdout).With().Timestamp().Logger()
var outputStderr = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339, NoColor: false}
var loggerStderr = zerolog.New(outputStderr).With().Timestamp().Logger()

var clientset *kubernetes.Clientset

var rootCmd = &cobra.Command{
	Use:   "controller",
	Short: "CLI to retrieve secrets from Kubernetes",
	Long:  "CLI utility for copying secrets-store.csi.k8s.io generated secrets from nginx-ingress namespace",

	//Creating a new or continue using existing client for connection to Kubernetes client
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if clientset != nil {
			loggerStdout.Info().Msg("Using existing kubernetes connection...")
		}

		loggerStdout.Debug().Msg("Trying to get config from KUBECONFIG env variable...")
		kubeConfig := os.Getenv("KUBECONFIG")
		if kubeConfig == "" {
			loggerStdout.Debug().Msg("No KUBECONFIG env variable found, using default path ~/.kube/config")
			kubeConfig = clientcmd.RecommendedHomeFile
		}

		config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			loggerStderr.Error().Err(err).Msg("Error loading kubeconfig")
		}
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			loggerStderr.Error().Err(err).Msg("Error creating Kubernetes client")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		loggerStderr.Error().Err(err).Msg("Error occured:")
		os.Exit(1)
	}
}
