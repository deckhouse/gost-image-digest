module github.com/deckhouse/gost-image-digest

go 1.20

require (
	github.com/google/go-containerregistry v0.16.1
	github.com/rs/zerolog v1.30.0
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
)

require (
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/docker/cli v24.0.0+incompatible // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.0+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	go.cypherpunks.ru/gogost/v5 v5.13.0
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace go.cypherpunks.ru/gogost/v5 v5.13.0 => github.com/flant/gogost/v5 v5.13.0
