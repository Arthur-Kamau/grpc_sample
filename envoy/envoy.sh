## Start envoy
## in envoy  folder enter envoy folder
envoy -c envoy.yaml


# in project root
envoy -c envoy/envoy.yaml


## Start Envoy Admin
## * check admin port in envoy.yaml
docker run --name=proxy-with-admin -d \
    -p 9901:9901 \
    -p 10000:10000 \
    -v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
    envoyproxy/envoy:latest