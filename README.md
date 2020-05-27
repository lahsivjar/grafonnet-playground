# Grafonnet Playground
Playground for grafanna with jsonnet. Allows users to write [grafonnet](https://github.com/grafana/grafonnet-lib) and see the rendered [grafana](https://grafana.com/) dashboard.

![Grafonnet playground example](assets/grafonnet-example.png)

## Usage

Run using dockerhub image (note the variables in the command)

```
docker run \
    -e "GRAFANA_URL=${grafana_url}" \
    -e "GRAFANA_API_KEY=${grafana_api_key}" \
    -e "GRAFONNET_PLAYGROUND_FOLDER_ID=${grafana_playground_folder_id}" \
    -e "GIN_MODE=release" \
    -p 8080:8080 \
    lahsivjar/grafonnet-playground
```

*Note*: For debugging remove the `GIN_MODE` environment variable from above command

## Development build
- Copy config file
```
make copy-config
```
- Update properties in the file `application.yml`
- Build code
```
make build-dev
```
- Execute go binary
```
./out/grafonnet-playground
```
