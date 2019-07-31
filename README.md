# condo3
[![CircleCI](https://circleci.com/gh/sue445/condo3/tree/master.svg?style=svg&circle-token=a9a9488053fc489f6cff7edfec8fe1d67d9da069)](https://circleci.com/gh/sue445/condo3/tree/master)

## Develop
### Setup
```bash
cp secrets.yaml.example secrets.yaml
vi secrets.yaml
```

Mac

```bash
brew cask install google-cloud-sdk
gcloud components install app-engine-go cloud-datastore-emulator
```

* https://cloud.google.com/appengine/docs/standard/go/download
* https://cloud.google.com/sdk/docs/

### Run local
```bash
make && ./bin/condo3
```
