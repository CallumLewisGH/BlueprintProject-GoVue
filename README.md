# BlueprintProject-GoVue

A template project for Golang and Vue integration. Enough architectural considerations
for most tasks - clone it, drop the two subfolders into a new project, and build the
actual product on top.

## What's here

- **[service-base](./service-base)** - Go backend: Gin + [Huma](https://huma.rocks/)
  (OpenAPI-documented routes), GORM/Postgres, Redis, JWT auth with Google OAuth (via
  [goth](https://github.com/markbates/goth)), roles/admin, a CQRS-ish command/query
  architecture, and a signed-URL upload pattern for Google Cloud Storage. Deployed via
  Google Cloud Buildpacks (`project.toml`).
- **[web-base](./web-base)** - Vue 3 + Vite frontend: JWT auth against the backend above,
  a request interceptor that attaches the bearer token and redirects to `/login` on 401,
  a small reusable UI kit, dark mode, profile editing with image upload, and a generated
  OpenAPI client (`npm run api` regenerates it from the backend's `/openapi.yaml`).

Each has its own README with setup instructions. Neither is meant to be used as-is - the
landing pages (`Home`/`About`/`Contact`/`Privacy`) carry obvious placeholder copy, ready
to be replaced with the real thing.
