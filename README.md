# docker-secrets

Simple package to read [Docker Secrets](https://docs.docker.com/engine/swarm/secrets/) from a Swarm Cluster

## Usage

### Read

A simple Docker secret like - 

```bash
$ echo "test-secret" | docker secret create test-secret -
```

can be read using the Read method

```golang
package main

import (
  "github.com/dtmistry/secrets"
)

func main() {
  secrets := NewDefaultSecrets()
  secret, err := secrets.Read("test-secret")
}
```

### ReadAsMap

A secret created using a file like -

```properties
db-user=secret-user
db-pass=secret-pass
api-key=super-secret-apikey
```

can be read using the ReadAsMap method

```golang
package main

import (
  "github.com/dtmistry/secrets"
)

func main() {
  secrets := NewDefaultSecrets()
  secretMap, err := secrets.ReadAsMap("test-secret")
}
```




