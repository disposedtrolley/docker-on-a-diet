# Notes

## Demo

### 1 - Naive Dockerfile

- Total image size approaches 1GB.
- Takes around a minute to build from scratch.
- Most `RUN` steps in the Dockerfile are repeated when anything in the repository
  is changed.

### 2 - Order from least to most changing

- No difference in image size or initial build time.
- Incremental builds are _much_ faster as we can leverage the builder cache.

### 3 - Explicitly download your dependencies

- Dependencies likely aren't being modified as freqently as the source code.
- Downloading dependencies before the build/compilation step can help warm
  the cache.
- Especially important for ecosystems where a large number of dependencies
  is the norm, i.e. Node.

### 4 - Consolidate system package manager commands

- Each Docker command (i.e. `RUN`) spawns another layer.
- Once a layer is created, it can't be removed from the image. This means
  that if you create a file in one layer and delete it in a subsequent
  layer, you're not actually going to save any space.
- Running package manager commands together reduces layer overhead and
  allows you to "tidy up" before the layer is committed.
- For example, deleting the `apt` cache shaves about 40MB from the total
  image size.

### 5 - Multi-stage builds

- Separate the build steps into its own stage. No need to ship your dev
  dependencies!
- Instead of installing `libgit2-dev`, we can just install the compiled
  binary.
- 10x reduction in image size. ~800MB -> ~80MB.

