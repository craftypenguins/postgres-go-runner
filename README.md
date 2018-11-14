# postgres-go-runner

Overwrite the mounted in commands.sql, the default is just a 'SELECT 1;' that will work for init containers for checking DB connectivity.

Connects to a postgresurl provided by environment variable POSTGRES_URL, populate it via secret in k8s or something similar
