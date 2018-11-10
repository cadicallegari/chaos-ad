# chaos-ad

named after the amazing album
https://en.wikipedia.org/wiki/Chaos_A.D.



## Run locally

```
make run
```

## Testing

```
make check-integration
```

## Releasing

```
make release version=<version>
```

It will create a git tag with the provided **<version>**
and build and publish a docker image.

## Vendoring

To update vendored dependencies run:

```
make vendor
```
