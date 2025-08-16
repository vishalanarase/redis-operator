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