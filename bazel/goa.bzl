def _goa_main(ctx):
    gen_dir = ctx.actions.declare_directory("gen")
    args = ctx.actions.args()

    args.add("gen")
    args.add(ctx.attr.package)

    outputs = []
    if ctx.outputs.output_dir:
        args.add("--output", ctx.outputs.output_dir.path)
        outputs.append(ctx.outputs.output_dir)
    else:
        outputs.append(gen_dir)

    print(dir(ctx))
    ctx.actions.run(
        executable = "/Users/tyler/go/bin/goa",
        arguments = [args],
        outputs = [gen_dir],
        #        use_default_shell_env = True,
        env = {
            "GOCACHE": str(ctx.path(".")) + "/.gocache",
            #            "GO111MODULE": "on",
            #            "GOPATH": "/Users/tyler/go",
            #            "GOROOT": "/usr/local/go",
            #            "HOME": "/Users/tyler",
        },
    )

    return [
        DefaultInfo(
            files = depset([gen_dir]),
            #            runfiles = ctx.runfiles(files = ctx.files.data),
        ),
    ]

goa = rule(
    implementation = _goa_main,
    attrs = {
        "package": attr.string(),
        "output_dir": attr.output(),
        "data": attr.label_list(allow_files = True),
        "deps": attr.label_list(allow_files = True),
    },
)
