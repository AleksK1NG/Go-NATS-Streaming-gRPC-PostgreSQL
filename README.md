### Golang NATS gRPC Postgresql email microservice example 👋

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000

### Nats UI:

http://localhost:8222/

### Swagger UI:

https://localhost:5000/swagger/index.html

### MailHog UI:

http://localhost:8025/

For local development:
```
make cert // generates tls certificates
make migrate_up // run sql migrations
make swagger // generate swagger documentation
make local or develop // for run docker compose files
```