# condo3
Feed generator for [connpass](https://connpass.com/) and [Doorkeeper](https://www.doorkeeper.jp/) group

https://condo3.appspot.com/

[![CircleCI](https://circleci.com/gh/sue445/condo3/tree/master.svg?style=svg&circle-token=a9a9488053fc489f6cff7edfec8fe1d67d9da069)](https://circleci.com/gh/sue445/condo3/tree/master)
[![Maintainability](https://api.codeclimate.com/v1/badges/e0d43c1cc012319a621c/maintainability)](https://codeclimate.com/github/sue445/condo3/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e0d43c1cc012319a621c/test_coverage)](https://codeclimate.com/github/sue445/condo3/test_coverage)

## Available APIs
* `https://condo3.appspot.com/api/connpass/{group}.{format}`
* `https://condo3.appspot.com/api/doorkeeper/{group}.{format}`

### Note
* `format` supports either `ics` or `atom`
  * e.g. https://condo3.appspot.com/api/connpass/gocon.ics

## Develop
### Setup
```bash
cp .envrc.example .envrc
vi .envrc
```

Mac

```bash
brew cask install google-cloud-sdk
gcloud components install app-engine-go
```

* https://cloud.google.com/appengine/docs/standard/go/download
* https://cloud.google.com/sdk/docs/

### Run local
```bash
npm run build
make && ./bin/condo3
```

### Encrypt credentials with KMS
```bash
vi credential.txt
gcloud kms encrypt --location=global --keyring condo3 --key app --plaintext-file credential.txt --ciphertext-file credential.enc
cat credential.enc | base64 > encrypted_credential_base64.txt
```

Write to [app.yaml](app.yaml)
