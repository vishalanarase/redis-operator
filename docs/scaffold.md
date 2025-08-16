# Redis Operator
> This document provides an overview of the Redis Operator, how it was scaffolded

## Scaffolding Process

### Init operator

```bash
operator-sdk init --domain devspace.com --repo github.com/vishalanarase/redis-operator
```

<details>
<summary>Output</summary>
```bash
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.19.0
INFO[0008] Update dependencies:
$ go mod tidy
go: downloading google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d
Next: define a resource with:
$ operator-sdk create api
```
</details>

### Create API

```bash
operator-sdk create api --group app --version v1 --kind RedisCluster --resource --controller
```

<details>
<summary>Output</summary>
```bash

```
</details>