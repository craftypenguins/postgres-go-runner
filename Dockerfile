FROM scratch
MAINTAINER Nicole Macfarlane

COPY postgres_go /postgres_go
COPY commands.sql /commands.sql
ENTRYPOINT ["/postgres_go"]
