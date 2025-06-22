# go-controller

A simple and flexible Go-based CLI tool for retrieving and optionally copying Kubernetes Secrets between namespaces. Built with [Cobra](https://github.com/spf13/cobra) and [client-go](https://github.com/kubernetes/client-go).

## Features

- ✅ Retrieve Kubernetes `Secrets` in a specific namespace
- ✅ Filter secrets by label selector
- ✅ Optionally copy found secrets into a target namespace
- ✅ Supports custom kubeconfig paths
- ✅ Simple and clear CLI interface

---

## Requirements

- Go 1.21+
- Access to a Kubernetes cluster
- Valid `kubeconfig` file

---

## Installation

```bash
git clone https://github.com/dvysh/go-controller.git
cd go-controller
go build -o controller

```

##  Usage

./controller get-secrets [flags]

| Flag                 | Description                                                 | Required | Default              |
| -------------------- | ----------------------------------------------------------- | -------- | -------------------- |
| `--kubeconfig`       | Path to kubeconfig file                                     | No       | `$HOME/.kube/config` |
| `--src-namespace`        | Source namespace to search secrets                          | No       | `nginx-ingress`            |
| `--label`            | Label selector for filtering secrets (example: `key=value`) | No       | *none*               |
| `--target-namespace` | Namespace where found secrets should be copied              | No       | `default`               |


## Examples

✔️ Get all secrets from default namespace:  
./controller get-secrets --namespace=default

✔️ Get secrets with label selector:  
./controller get-secrets --namespace=default --label=secrets-store.csi.k8s.io/managed=true

✔️ Get and copy secrets to another namespace:  
./controller get-secrets --namespace=default --label=secrets-store.csi.k8s.io/managed=true --target-namespace=prod


## License
MIT © 2025 Dmytro Vyshniakov vishdi@gmail.com
