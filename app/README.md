# app

> xavier webapp

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

## Release

1. `go-bindata-assetfs -pkg cmd dist/...`
2. Copy the generated file into ../cmd
3. go build
