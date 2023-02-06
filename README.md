<h1 align="center">Welcome to SimpleStorage 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
</p>

> Simple storage application for storing string data

## Build

```sh
make       //run lint and build
```


## Run tests

```sh
make test
```

## Run app

```sh
make run   //build and run app and start memcache via docker-compose
make stop  //shutdown docker-compose memcache
```

## Run app with keys
```
Usage of ./bin/storage_app:
  -gp string
        gRPC port (default "50051")
  -ii
        Is use internal storage instead of memcached (default "false")
  -itm
        Is use memcache and third-party lib (default "false")
  -mu string
        Memcached connection URL (default "0.0.0.0:11211")

By default using my simple memchached connection
```

## Author

👤 **Starikov Mikhail**

* Github: [@4ynyky](https://github.com/4ynyky)
