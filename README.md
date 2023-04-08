[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
# go-payment-service

> The project status is `WIP` (Work in progress) which means the author continously evaluate and improve the project.

Payment Service implementation using stripe as payment-gateway. The project using [typical-go](https://github.com/typical-go/typical-go) as its build-tool.


- Application
  - [ ] [Go-Standards](https://github.com/golang-standards/project-layout) Project Layout
  - [ ] Environment Variable Configuration
  - [ ] Health-Check and Debug API
  - [ ] Graceful Shutdown
- Layered architecture
  - [ ] Dependency Injection
  - [ ] Database Transaction
- HTTP Server
  - [ ] [Echo framework](https://echo.labstack.com/)
  - [ ] Server Side Caching
    - [ ] Cache but revalidate (Header `Cache-Control: no-cache`)
    - [ ] Set Expiration Time (Header `Cache-Control: max-age=120`)
    - [ ] Return 304 if not modified (Header `If-Modified-Since: Sat, 31 Oct 2020 10:28:02 GMT`)
  - [ ] Request ID in logger (Header `X-Request-Id: xxx`)
- RESTful
  - [ ] Create Resource (`POST` verb)
  - [ ] Update Resource (`PUT` verb)
  - [ ] Partially Update Resource (`PATCH` verb)
  - [ ] Find Resource (`GET` verb)
    - [ ] Offset Pagination (Query param `?limit=100&offset=0`)
    - [ ] Sorting (Query param `?sort=-title,created_at`)
    - [ ] Total count (Header `X-Total-Count: 99`)
  - [ ] Delete resource (`DELETE` verb, idempotent)
- Testing
  - [ ] Table Driven Test
