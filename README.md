# Template for micro-service in Go #

### Description ###
Micro-service template to use with my [DevOpsDemo](https://github.com/butzist/DevOpsDemo)

### Features ###
* Build in multi-stage Docker container
* Configuration via environment variables
* Logging in JSON
* Health-check endpoint
* Prometheus metrics
* Unit tests with xunit-compatible output
* API/integration tests with docker-compose

### Links ###
* https://docs.microsoft.com/en-us/azure/devops/pipelines/languages/docker?view=azure-devops

### TODO ###
* Application container is built twice (once for test and once for buildAndPush) on Azure Pipelines - avoid this redundancy
