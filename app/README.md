# gogitlog

It's like `git log` but terrible. Uses the `libgit2` bindings for Go to

## Build

```command
$ docker build -t gogitlog .
```

## Run

```command
$ docker run -v $(pwd)/../:/repo gogitlog
```

