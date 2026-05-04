# env-demo

Two Gofr microservices to verify env-var wiring across services.

- **`service-a`** is the **caller**. It exposes `GET /call-b`.
- **`service-b`** is the **callee**. It exposes `GET /hello`.

`service-a` reads the URL of `service-b` from the **`SERVICE_B_URL`** environment variable. Change that env var and `service-a` calls a different `service-b` — no code change needed.
