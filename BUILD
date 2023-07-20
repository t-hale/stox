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
        "@com_github_polygon_io_client_go//rest",
        "@com_github_polygon_io_client_go//rest/models",
    ],
)

#genrule(
#    name = "goagen",
#    #    srcs = [
#    #        "//design:goafiles",
#    #    ],
#    outs = [
#        "gen/http/cli/server/cli.go",
#        "gen/http/openapi.json",
#        "gen/http/openapi.yaml",
#        "gen/http/openapi3.json",
#        "gen/http/openapi3.yaml",
#        "gen/http/stox/client/cli.go",
#        "gen/http/stox/client/client.go",
#        "gen/http/stox/client/encode_decode.go",
#        "gen/http/stox/client/paths.go",
#        "gen/http/stox/client/types.go",
#        "gen/http/stox/server/encode_decode.go",
#        "gen/http/stox/server/paths.go",
#        "gen/http/stox/server/server.go",
#        "gen/http/stox/server/types.go",
#        "gen/log/logger.go",
#        "gen/stox/client.go",
#        "gen/stox/endpoints.go",
#        "gen/stox/service.go",
#    ],
#    cmd = "goa gen github.com/t-hale/stox/design",
#)

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
        "@com_github_google_go_cmp//cmp",
    ],
)
