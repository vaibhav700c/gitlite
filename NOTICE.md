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

### Bearer authorization and pagination defaults (`feat/bearer-auth-deprecation`)

Upstream Gitea already accepts both `Authorization: token <sha>` and `Authorization: Bearer <sha>`.
GitLite does not add Bearer support; it makes Bearer the documented scheme and deprecates `token`.

- `modules/auth/httpauth/httpauth.go`, `httpauth_test.go`: record whether a token was sent with the `token` prefix.
- `services/auth/oauth2.go`: requests using the `token` prefix receive `Deprecation`, `Sunset`, and `Warning` headers and log a warning at most once per minute per user; they are rejected with `401` when `ALLOW_LEGACY_TOKEN_SCHEME` is `false`.
- `modules/setting/api.go`, `custom/conf/app.example.ini`: new `[api] ALLOW_LEGACY_TOKEN_SCHEME` (default `true`); `DEFAULT_PAGING_NUM` changed from `30` to `20` and `MAX_RESPONSE_ITEMS` from `50` to `25`.
- `services/context/api.go`: paginated list endpoints also send `X-Page`, `X-PerPage`, and `X-HasMore`. Upstream already sends `X-Total-Count` and `Link`.
- `routers/api/v1/api.go`, `templates/swagger/`: `AuthorizationHeaderToken` security definition describes the Bearer scheme.
- `tests/integration/api_token_scheme_test.go`: integration tests for the above.

No Gitea logos were replaced or reused in GitLite-branded material; the upstream assets remain in the tree untouched.
