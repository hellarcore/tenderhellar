# Docker

## Supported tags and respective `Dockerfile` links

DockerHub tags for official releases are [here](https://hub.docker.com/r/hellarpay/tenderhellar/tags). The "latest" tag will always point to the highest version number.

Official releases can be found [here](https://github.com/hellarcore/tenderhellar/releases).

The Dockerfile used for all builds can be found [here](https://github.com/hellarcore/tenderhellar/blob/master/DOCKER/Dockerfile).

Respective versioned files can be found at `https://raw.githubusercontent.com/tendermint/tendermint/vX.XX.XX/DOCKER/Dockerfile` (replace the Xs with the version number).

## How to use this image

### Start one instance of the Tendermint core with the `kvstore` app

A quick example of a built-in app and Tendermint core in one container.

```sh
mkdir --mode=0777 -p /tmp/tenderhellar
docker run -it --rm -v "/tmp/tenderhellar:/tenderhellar" hellarpay/tenderhellar 
```


## License

- Tendermint's license is [Apache 2.0](https://github.com/tendermint/tendermint/blob/master/LICENSE).

## Contributing

Contributions are most welcome! See the [contributing file](https://github.com/tendermint/tendermint/blob/master/CONTRIBUTING.md) for more information.
