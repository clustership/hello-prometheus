# TODO

See: https://access.redhat.com/documentation/en-us/openshift_container_platform/4.5/html/monitoring/monitoring-your-own-services

[] Deploy the application
[] Integrate service monitor in Prometheus stack
[] Create AlertRule using PrometheusRule

# What's missing

route configuration (oc expose)
blackbox operator
Configuration blackbox to query the route



# BUILD

```bash
podman build -t quay.io/xymox/hello-prometheus:$(cat VERSION) .
podman tag quay.io/xymox/hello-prometheus:$(cat VERSION) quay.io/xymox/hello-prometheus:latest
podman push quay.io/xymox/hello-prometheus:$(cat VERSION)
```
