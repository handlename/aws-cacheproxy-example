# aws-cacheproxy-example

An implementation example of a cache proxy using AWS Lambda and CloudFront.

## Prerequisites

- Go
- [aqua](https://aquaproj.github.io/)

## Setup

```console
aqua i
```

## Deployment

```console
cd terraform && make apply
```

If you want to deploy Lambda function only, run the following command:

```console
cd lambda && make deploy
```

## Author

@handlename

## License

MIT
