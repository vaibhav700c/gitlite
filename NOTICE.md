# NOTICE

GitLite is a fork of [Gitea](https://github.com/go-gitea/gitea), forked at tag `v1.27.3`.
Gitea is MIT licensed; the upstream `LICENSE` file is preserved verbatim in this repository.

GitLite is not affiliated with, sponsored by, or endorsed by the Gitea project.
It exists only as the demo target for a documentation-sync hackathon submission
(Thally Sync Hackathon 2026, "Keep Product Knowledge Current" track).
The submitted work is the `gitlite-docs` documentation site and its Thally Track integration,
not this fork.

## Changes from upstream

Every change made to upstream Gitea source is listed here.

### Fork setup (`main`)

- `routers/api/v1/api.go`: swagger title and description say "GitLite API" instead of "Gitea API".
- `templates/swagger/v1_json.tmpl`, `templates/swagger/v1_openapi3_json.tmpl`: regenerated from the change above.
- `modules/setting/server.go`, `custom/conf/app.example.ini`: default `APP_NAME` is `GitLite`.
- `NOTICE.md`: this file.

No Gitea logos were replaced or reused in GitLite-branded material; the upstream assets remain in the tree untouched.
