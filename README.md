# Grafonnet Playground
Playground for grafanna with jsonnet. Allows users to write [grafonnet](https://github.com/grafana/grafonnet-lib) and see the rendered [grafana](https://grafana.com/) dashboard.

## Usage

Run using dockerhub image (note the variables in the command)

```
docker run \
    -e "GRAFANA_URL=${grafana_url}" \
    -e "GRAFANA_API_KEY=${grafana_api_key}" \
    -e "GRAFONNET_PLAYGROUND_FOLDER_ID=${grafana_playground_folder_id}" \
    -p 8080:8080 \
    lahsivjar/grafonnet-playground
```

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
