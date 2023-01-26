# spanner-bazel

## Steps to build

```
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -build_file_proto_mode=disable_global
bazel run //:gazelle
bazel build //:main
```