# VP K8s Controller

A little Helm chart for deploying the Ververica Platform Kubernetes Controller.

## Installing the Chart

| Parameter                    | Description                                          | Default                                            |
|------------------------------|------------------------------------------------------|----------------------------------------------------|
| `rbac.enabled`               | Whether or not to create RBAC resources.             | `true`                                             |
| `rbacProxy.enabled`          | Whether or not to create an RBAC proxy over `https`. |  `true`                                            |
| `rbacProxy.logLevel`         | Log level for the proxy.                             | `10`                                               |
| `rbacProxy.imageRepository`  |                                                      | `gcr.io/kubebuilder/kube-rbac-proxy`               |
| `rbacProxy.imageTag`         |                                                      | `v0.4.0`                                           |
| `rbacProxy.imagePullPolicy`  |                                                      | `IfNotPresent`                                     |
| `rbacProxy.port`             |                                                      | `8443`                                             |
| `controller.imageRepository` | Image repository for the Manager                     | `fintechstudios/ververica-platform-k8s-controller` |
| `controller.imageTag`        |                                                      | `v0.3.1`                                           |
| `controller.imagePullPolicy` |                                                      | `IfNotPresent`                                     |
| `controller.metricsHost`     | Host for the metrics reporter.                       | `127.0.0.1`                                        |
| `controller.metricsPort`     | Port for the metrics reporter.                       | `8080`                                             |
| `controller.vpApiUrl`        | URL for the Ververica Platform API.                  | `http://ververica-platform-appmanager/api`         |