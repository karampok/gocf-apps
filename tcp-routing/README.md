### Notes

Routing for TCP domains is layer 4 and protocol agnostic, so many features
 available to HTTP routing are not available for TCP routing. TCP domains are
 defined as being associated with the TCP Router Group. The TCP Router Group
 defines the range of ports available to be reserved with TCP Routes.
 Currently, only Shared Domains can be TCP.

### How to inspect the route status

```
cf api api.bosh-lite.com --skip-ssl-validation && cf auth admin <password> 
cf create-org o && cf t -o o && cf create-space s && cf t -o o -s s
cf create-shared-domain tcp.boh-lite.com --router-group default-tcp
cf domains
cf router-groups
cf routes
```

### Deploy with one command

```
GOARCH=amd64 GOOS=linux go build -ldflags "-X main.buildstamp=2017-11-23T10:38:33Z -X main.githash=bd1075f" -o ./bin/echoApp
cf push echoApp -d tcp.bosh-lite.com --random-route -c ./bin/echoApp -b https://github.com/cloudfoundry/binary-buildpack.git
```

### Deploy in two steps

Get the tcp ports 
```
cf oauth-token
cf curl /routing/v1/router_groups
```

Create the route
```
#cf create-route SPACE DOMAIN [--hostname HOSTNAME] [--path PATH]
cf create-route s tcp.bosh-lite.com --port 1051
```

Compile & push without route
```
GOARCH=amd64 GOOS=linux go build -ldflags "-X main.buildstamp=2017-11-23T10:38:33Z -X main.githash=bd1075f" -o ./bin/echoApp
cf push echoApp  -c ./bin/echoApp -b https://github.com/cloudfoundry/binary-buildpack.git --no-route
```

Map the route to the app
```
#cf map-route APP_NAME DOMAIN [--hostname HOSTNAME] [--path PATH]
cf map-route echoApp tcp.bosh-lite.com --port 1051 
```

### Using a docker image

```
cf push mqttbroker --docker-image toke/mosquitto -d tcp.bosh-lite.com --random-route
```


### Reservable ports

```
cf oauth-token
cf curl /routing/v1/router_groups
cf curl /routing/v1/router_groups/f7392031-a488-4890-8835-c4a038a3bded -X PUT -d '{"reservable_ports":"1024-1199"}'
 ```
