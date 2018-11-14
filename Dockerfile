FROM scratch
MAINTAINER Nicole Macfarlane

COPY postgres-go-runner /postgres-go-runner
COPY commands.sql /commands.sql
ENTRYPOINT ["/postgres-go-runner"]
