all:
	cd 18.10 && vab build --arg KERNEL_VERSION=5.0.5 -p --ref docker.io/stellarproject/terra:18.10

FORCE:

containerd:
	cd containerd && vab build -p --ref docker.io/stellarproject/containerd:latest

extras:
	cd cni && vab build -p --ref docker.io/stellarproject/cni:latest
	cd node_exporter && vab build -p --ref docker.io/stellarproject/node_exporter:latest
	cd buildkit && vab build -p --ref docker.io/stellarproject/buildkit:latest

kernel:
	cd kernel && vab build -p --ref docker.io/stellarproject/kernel:5.0.5

base: FORCE
	cd base && vab build -c /tmp/terra -p --ref docker.io/stellarproject/ubuntu:18.10