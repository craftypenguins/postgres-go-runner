FROM scratch
MAINTAINER Nicole Macfarlane

COPY postgres-go-runner /postgres_go
COPY commands.sql /commands.sql
ENTRYPOINT ["/postgres-go-runner"]
