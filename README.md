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
  secrets := secrets.NewDefaultSecrets()
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
  secrets := secrets.NewDefaultSecrets()
  secretMap, err := secrets.ReadAsMap("test-secret")
}
```

### Custom secrets location

Docker secrets are available to a container at `/run/secrets/`. Docker version 17.06 and above allows to configure the default secrets location. Secrets can be read from a custom location by supplying the location - 

```golang
package main

import (
  "github.com/dtmistry/secrets"
)

func main() {
  secrets, err  := secrets.NewSecrets("/custom/path")
  secretMap, err := secrets.ReadAsMap("test-secret")
}

```

### License

Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text.
