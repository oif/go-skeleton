# go-skeleton
Go application skeleton

# Arch

```
├── CHANGELOG // Write changelog here, and also read version on build from here.
├── LICENSE
├── README.md  // Intro of your application
├── _build  // hack/build.sh will put artifacts in this directory
│   ├── skeleton-0.1.0-darwin-amd64
│   ├── skeleton-0.1.0-linux-amd64
│   └── skeleton-0.1.0-windows-amd64.exe
├── cmd // The entrence of your application
│   ├── cmd.go // Index commands 
│   ├── echo // One of the sub commands
│   │   ├── command.go
│   │   └── option.go
│   └── runtime.go // Setup runtime here
├── docs // Write all your docs, such as API, Design or any other docs
│   └── DESIGN
├── hack // hack scripts
│   └── build.sh
├── manifests // Manifests for application, e.g. Kubernetes deployment spec
│   └── MANIFEST
├── pkg // packages here, the mainline of all
│   ├── echo
│   │   └── echo.go
│   ├── meta // Application meta, will set variables property on build and startup
│   │   ├── app.go
│   │   └── version.go
│   └── utils // utils package
│       └── dumpstack
│           ├── dumpstack.go
│           ├── dumpstack_unix.go
│           └── dumpstack_windows.go
└── types // We construct types by Type and Version, which gain the abilitiy for open-source. And reduce import cycle
    └── core
        └── v1
            └── echo.go
```
