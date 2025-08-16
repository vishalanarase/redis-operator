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
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] api/v1/rediscluster_types.go
INFO[0000] api/v1/groupversion_info.go
INFO[0000] internal/controller/suite_test.go
INFO[0000] internal/controller/rediscluster_controller.go
INFO[0000] internal/controller/rediscluster_controller_test.go
INFO[0000] Update dependencies:
$ go mod tidy
INFO[0000] Running make:
$ make generate
mkdir -p /Users/vishal/workspace/vishalanarase/redis-operator/bin
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.1
go: downloading sigs.k8s.io/controller-tools v0.16.1
go: downloading golang.org/x/tools v0.24.0
go: downloading github.com/fatih/color v1.17.0
go: downloading golang.org/x/net v0.28.0
go: downloading golang.org/x/text v0.17.0
go: downloading golang.org/x/sys v0.23.0
go: downloading golang.org/x/mod v0.20.0
/Users/vishal/workspace/vishalanarase/redis-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```
</details>