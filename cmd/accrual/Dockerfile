FROM --platform=linux/amd64 ubuntu:latest
COPY ./cmd/accrual/accrual_linux_amd64 /usr/local/bin
EXPOSE 8081
ENV RUN_ADDRESS=0.0.0.0:8081 \
    DATABASE_URI="postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable"
CMD ["sh", "-c", "sleep 4 && /usr/local/bin/accrual_linux_amd64"]
