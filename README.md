# # go-controller

Educational project of studying Golang and Kubernetes Controllers, written as part of a course Crash Course: Kubernetes controllers by fwdays
A CLI utility for copying secrets-store.csi.k8s.io generated secrets from nginx-ingress namespace to destination namespaces.

---

## Requirements

- Go 1.21+
- A running Kubernetes cluster
- A valid kubeconfig with access to your cluster

---

## Installation

```sh
git clone https://github.com/dvysh/go-controller
cd go-controller
go mod tidy
```

---