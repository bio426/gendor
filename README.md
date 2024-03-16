## Build instructions

- docker build -t vugopo .
- docker run --rm -it -p 8080:8080 vugopo

## Example env

```bash
#!/bin/bash

declare -A envs

# Tern
envs["TERN_CONFIG"]="tern.conf"
envs["TERN_MIGRATIONS"]="migrations/"

# Postgres
envs["PG_HOST"]="localhost"
envs["PG_PORT"]="5432"
envs["PG_USER"]="postgres"
envs["PG_PASSWORD"]="password"
envs["PG_DATABASE"]="gendor"

for env in "${!envs[@]}"; do
  export ${env}=${envs[${env}]}
done

echo "${#envs[@]} variables setted"

```
