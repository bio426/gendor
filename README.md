## Build instructions

- docker build -t vugopo .
- docker run --rm -it -p 8080:8080 vugopo

## Deploy steps

- docker build -t us-central1-docker.pkg.dev/gendor-417403/app/server .
- docker push us-central1-docker.pkg.dev/gendor-417403/app/server
- implementar nueva revision en la instancia de cloud run con la nueva imagen

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
