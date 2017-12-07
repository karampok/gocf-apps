# gocf-apps


# Push just an app 
```
cf api api.bosh-lite.com --skip-ssl-validation && cf auth admin  <passwd>
cf create-org o && cf t -o o && cf create-space s && cf t -o o -s s
cf enable-feature-flag diego_docker
cf push test-app -o cloudfoundry/test-app -i 2
```
