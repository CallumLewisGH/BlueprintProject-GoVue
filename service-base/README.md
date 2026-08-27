# service-base

A barebones Go backend template: users, JWT auth (Google OAuth via [goth](https://github.com/markbates/goth)),
roles/admin, signed-URL file uploads (GCS), Postgres/GORM, Redis, and a CQRS-ish
command/query architecture. Meant to be copied as the starting point for a new project,
not used as-is.

## Stack

- Gin + [Huma](https://huma.rocks/) (OpenAPI-documented routes, served at `/docs`)
- GORM + Postgres
- Redis (rate limiting)
- JWT bearer auth, OAuth login (Google) via goth
- Google Cloud Storage signed-upload-URL pattern for file uploads
- Deployed via Google Cloud Buildpacks (see `project.toml`)

## Getting started

1. Copy `.dev.env.example` / `.prod.env.example` / `.test.env.example` to `.dev.env` /
   `.prod.env` / `.test.env` and fill in real values (DB connection string, `JWT_SECRET`,
   OAuth credentials, etc). These are gitignored — never commit real secrets.
2. `just dev` — starts local Postgres + Redis via docker-compose and runs the server with
   hot reload (air).
3. `just test` — runs the integration + unit test suite (uses testcontainers, needs Docker).

## Adding a new domain model

See the "Adding new stuff" notes at the bottom of `cmd/app/main.go` — model → DTO →
converter → validation → repository → specification → requests → queries/commands → routes.
