# gpm-runtime

process manager in container using go language to manage process with solving zombie process.

## install

```shell
v1:
  go install github.com/jaronnie/gpm-runtime@v1.0.0
  
v2:
  go install github.com/jaronnie/gpm-runtime/v2@v2.0.0  
```

## Usage(v2)

You can use `gpm-runtime run -- npm run docs:dev` instead of `npm run docs:dev`

It just recycles zombie process.

use `gpm-runtime -h` to get more info.

`gpm-runtime run --loglevel debug -p 2 -- npm run docs:dev`