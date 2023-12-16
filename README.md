# GoCensys

GoCensys is a simple wrapper in Go for the v2 API exposed by Censys

## Todo

### APIs
- [ ] /v2/hosts
  - [ ] /search
  - [ ] /aggregate
  - [ ] /{ip}
    - [ ] /diff
    - [ ] /names
    - [ ] /comments
    - [ ] /tags
- [ ] /v2/certificates
  - [ ] /{fingerprint}
    - [ ] /hosts
    - [ ] /comments
  - [ ] /bulk
  - [ ] /search
  - [ ] /aggregate
- [ ] /v2/metadata/hosts

### Utils
- [ ] getBaseUrl
- [ ] validateToken
- [ ] makeRequest
- [ ] parseResponse
- [ ] baseResponse (struct)
