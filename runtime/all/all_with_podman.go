//go:build linux && cgo && podman
// +build linux,cgo,podman

package all

import (
	_ "github.com/srl-labs/containerlab/runtime/containerd"
	_ "github.com/srl-labs/containerlab/runtime/docker"
	_ "github.com/srl-labs/containerlab/runtime/ignite"
	_ "github.com/srl-labs/containerlab/runtime/podman"
)