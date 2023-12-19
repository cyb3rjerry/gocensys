# GoCensys

GoCensys is a simple wrapper in Go for the v2 API exposed by Censys

## Todo

### APIs
- [x] /v2/hosts
  - [x] /search
  - [x] /aggregate
  - [x] /{ip}
    - [x] /diff
    - [x] /names
    - [x] /comments
- [x] /v2/certificates
  - [x] /{fingerprint}
    - [x] /hosts
  - [x] /search
  - [x] /aggregate
- [x] /v2/metadata/hosts

### Utils
- [x] GetAccountInfo
- [x] makeRequest
- [x] parseResponse
- [x] baseResponse (struct)
