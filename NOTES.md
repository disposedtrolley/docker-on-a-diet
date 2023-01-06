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

