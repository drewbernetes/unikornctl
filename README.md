# unikornctl

## Introduction
`unikornctl` is a simple utility for working with the Unikorn Kubernetes API, and which provides a marginal UX improvement over a bunch of curl commands wrapped in a shell script.

## Installation

Grab a binary from the [Releases](https://github.com/drewbernetes/unikornctl/releases) page.

It probably helps to familiarise yourself with the [Unikorn architecture and principles](https://unikorn-cloud.org/).

## Configuration

`unikornctl` can be configured via `~/.unikornctl.yaml`:

```yaml
username: "example@unikorn-cloud.org"
password: "hunter2"
url: "https://clusters.unikorn-cloud.org"
project: "abc123"
```

These options can also be set (or overriden) via the command line.

## Usage

### Get projects

```shell
% unikornctl get projects
Name: njones-demo	ID: abc123
Name: hotdog	Description: Management cluster ID: abc123
```

### List images

Get a list of compatible images:

```shell
% unikornctl get images
Name: unikorn-230408-ac77c987	UUID: ab59076c-142a-4da6-9149-e8426fb93aff	Created: 2023-04-08 09:34:16 +0000 UTC	Kubernetes version: v1.25.6	NVIDIA driver version: 525.85.05
Name: unikorn-230414-cc034e2d	UUID: e1ad0805-ecd9-4162-b243-e2bfd85e8aed	Created: 2023-04-14 12:47:16 +0000 UTC	Kubernetes version: v1.26.3	NVIDIA driver version: no_gpu
```

### List networks

Get external networks:

```shell
% unikornctl get networks
Name: Internet	ID: c9d130bc-301d-45c0-9328-a6964af65579
```

### List Unikorn control planes:

```shell
% unikornctl get controlplanes
Name: dibnah	Status: Provisioned	Version: 1.0.1
```

### List clusters

```shell
% unikornctl get clusters --controlplane default
Name: demo	Version: v1.26.1	Status: Provisioned
```

### Get versions

```shell
% unikornctl get versions
Cluster Bundles:
Name: kubernetes-cluster-1.3.1	Version: 1.3.1	EOL: 04 Dec 23 00:00 UTC
Name: kubernetes-cluster-1.4.1	Version: 1.4.1
Control Plane Bundles:
Name: control-plane-1.1.0	Version: 1.1.0	EOL: 04 Dec 23 00:00 UTC
Name: control-plane-1.2.0	Version: 1.2.0

```

### Create a control plane

```shell
% unikornctl create controlplane --name default --version 1.2.0
```

### Create a cluster

Creating a cluster requires a number of parameters to be defined, including workload pools plus associated flavours and whether to enable autoscaling.  For now this definition needs to be crafted as a blob of JSON, see the example in the [examples](https://github.com/drewbernetes/unikornctl/tree/main/examples) folder.  A cluster can be created once the associated Unikorn control plane is in status `Provisioned`:

```shell
% unikornctl get controlplane --name default
Name: default	Status: Provisioned	Version: 1.0.1
% unikornctl create cluster --name demo ./examples/cluster.json
% unikornctl get cluster --name demo
Cluster: demo, version: v1.26.1, status: Provisioning.
└── Pools:
    └── Name: worker	Flavor: g.2.standard	Image: unikorn-230408-3464bc03
```

### Update a cluster

Similar to `create`, an existing cluster can be updated with the `update` command and the (full) JSON definition:

```shell
% unikornctl update cluster --name demo --json ./examples/cluster.json
```

### Retrieve kubeconfig

Once the cluster's status is `Provisioned` you can retrieve its kubeconfig:

```shell
% unikornctl get clusters --name demo
Cluster: demo, version: v1.26.1, status: Provisioned.
└── Pools:
    └── Name: worker	Flavor: g.2.standard	Image: unikorn-230408-3464bc03
% unikornctl get kubeconfig --cluster demo > ~/kubeconfig-demo
% export KUBECONFIG=~/kubeconfig-demo
% kubectl version -o json | jq .serverVersion
{
  "major": "1",
  "minor": "26",
  "gitVersion": "v1.26.1",
  "gitCommit": "8f94681cd294aa8cfd3407b8191f6c70214973a4",
  "gitTreeState": "clean",
  "buildDate": "2023-01-18T15:51:25Z",
  "goVersion": "go1.19.5",
  "compiler": "gc",
  "platform": "linux/amd64"
}
```
### Testing Using Docker Image
Use the below snippet to create the image and then use the commands above to test things work as expected.

```shell
docker build -t unikornctl:v0.0.0 -f docker/Dockerfile .
docker run --name unikornctl -it --rm -v /home/$USER/.unikornctl.yaml:/home/unikornctl/.unikornctl.yaml unikornctl:v0.0.0 "unikornctl get images"
```
