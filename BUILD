load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/t-hale/stox
gazelle(name = "gazelle")

go_library(
    name = "stox",
    srcs = ["stox.go"],
    importpath = "github.com/t-hale/stox",
    visibility = ["//visibility:public"],
    deps = [
        "//errors",
        "//gen/log",
        "//gen/stox",
        "//utils",
        "@com_github_alpacahq_alpaca_trade_api_go_v3//marketdata",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_polygon_io_client_go//rest",
        "@com_github_polygon_io_client_go//rest/models",
    ],
)

genrule(
    name = "goagen",
    srcs = [
        "//design:goafiles",
    ],
    outs = [
        "out.txt",
    ],
    cmd = "goa gen github.com/t-hale/stox/design > $@",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)

go_test(
    name = "stox_test",
    srcs = ["stox_test.go"],
    embed = [":stox"],
    deps = [
        "//gen/stox",
        "//utils",
        "@com_github_alpacahq_alpaca_trade_api_go_v3//marketdata",
        "@com_github_google_go_cmp//cmp",
        "@com_github_google_go_cmp//cmp/cmpopts",
        "@org_golang_google_protobuf//proto",
    ],
)
