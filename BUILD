load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/t-thale/stox
gazelle(name = "gazelle")

go_library(
    name = "stox",
    srcs = ["stox.go"],
    importpath = "stox",
    visibility = ["//visibility:public"],
    deps = [
        "//gen/log",
        "//gen/stox",
        "//utils",
        "@com_github_alpacahq_alpaca_trade_api_go_v3//marketdata",
        "@com_github_polygon_io_client_go//rest",
        "@com_github_polygon_io_client_go//rest/models",
    ],
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
