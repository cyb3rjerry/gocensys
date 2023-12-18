# GoCensys

GoCensys is a simple wrapper in Go for the v2 API exposed by Censys

## Todo

### APIs
- [ ] /v2/hosts
  - [x] /search
  - [x] /aggregate
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
  - [x] /search
  - [x] /aggregate
- [x] /v2/metadata/hosts

### Utils
- [ ] getBaseUrl
- [ ] validateToken
- [x] makeRequest
- [x] parseResponse
- [x] baseResponse (struct)
