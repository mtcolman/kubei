# KubeClarity

![](images/kubeclarity-logo.png)

KubeClarity is a tool for detection and management of Software Bill Of Materials (SBOM) and vulnerabilities of container images and filesystems. It scans both runtime K8s clusters and CI/CD pipelines for enhanced software supply chain security.

## SBOM & vulnerability detection challenges

* Effective vulnerability scanning requires an accurate Software Bill Of Materials (SBOM) detection:
  * Various programming languages and package managers
  * Various OS distributions
  * Package dependency information is usually stripped upon build
* Which one is the best scanner/SBOM analyzer?
* What should we scan: Git repos, builds, container images or runtime?
* Each scanner/analyzer has its own format - how to compare the results?
* How to manage the discovered SBOM and vulnerabilities?
* How are my applications affected by a newly discovered vulnerability?

## Solution

* Separate vulnerability scanning into 2 phases:
  * Content analysis to generate SBOM
  * Scan the SBOM for vulnerabilities
* Create a pluggable infrastructure to:
  * Run several content analyzers in parallel
  * Run several vulnerability scanners in parallel
* Scan and merge results between different CI stages using KubeClarity CLI
* Runtime K8s scan to detect vulnerabilities discovered post-deployment
* Group scanned resources (images/directories) under defined applications to navigate the object tree dependencies (applications, resources, packages, vulnerabilities)


![](images/dashboard-screen.png)

## Features

* Dashboard
  * Fixable vulnerabilities per severity
  * Top 5 vulnerable elements (applications, resources, packages)
  * New vulnerabilities trends
  * Package count per license type
  * Package count per programming language
  * General counters
* Applications
  * Automatic application detection in K8s runtime
  * Create/edit/delete applications
  * Per application, navigation to related:
    * Resources (images/directories)
    * Packages  
    * Vulnerabilities
    * Licenses in use by the resources
* Application Resources (images/directories)
  * Per resource, navigation to related:
    * Applications
    * Packages
    * Vulnerabilities
* Packages
    * Per package, navigation to related:
        * Applications
        * Linkable list of resources and the detecting SBOM analyzers
        * Vulnerabilities
* Vulnerabilities
    * Per vulnerability, navigation to related:
        * Applications
        * Resources
        * List of detecting scanners
* K8s Runtime scan
  * Automatic detection of target namespaces
  * Scan progress and result navigation per affected element (applications, resources, packages, vulnerabilities)
* CLI (CI/CD)
  * SBOM generation using multiple integrated content analyzers (Syft, cyclonedx-gomod)
  * SBOM/image/directory vulnerability scanning using multiple integrated scanners (Grype, Dependency-track)
  * Merging of SBOM and vulnerabilities across different CI/CD stages
  * Export results to KubeClarity backend
* API
  * The API for KubeClarity can be found [here](https://github.com/cisco-open/kubei/blob/master/api/swagger.yaml)
    
## High level architecture

![](images/architecture.png)


## Getting started

### Integration with SBOM generators and vulnerability scanners
KubeClarity content analyzer integrates with the following SBOM generators:
* [Syft](https://github.com/anchore/syft)
* [Cyclonedx-gomod](https://github.com/CycloneDX/cyclonedx-gomod)

KubeClarity vulnerability scanner integrates with the following scanners:
* [Grype](https://github.com/anchore/grype)
* [Dependency-Track](https://github.com/DependencyTrack/dependency-track)

The integrations with the SBOM generators can be found here [here](https://github.com/cisco-open/kubei/tree/master/shared/pkg/analyzer), and the integrations with the vulnerability scanners can be found here [here](https://github.com/cisco-open/kubei/tree/master/shared/pkg/scanner). 
To enable and configure the supported SBOM generators and vulnerability scanners, please check the "analyzer" and "scanner" config under the "vulnerability-scanner" section in Helm values.

Contributions of integrations with additional tools are more than welcome!

### Install KubeClarity in a K8s cluster using Helm:

1. Add Helm repo

   ```shell
   helm repo add kubeclarity https://cisco-open.github.io/kubei
   ```
2. Save KubeClarity default chart values
    ```shell
    helm show values kubeclarity/kubeclarity > values.yaml
    ```

3. Check the configuration in `values.yaml` and update the required values if needed

4. Deploy KubeClarity with Helm

   ```shell
   helm install --values values.yaml --create-namespace kubeclarity kubeclarity/kubeclarity -n kubeclarity
   ```

3. Port forward to KubeClarity UI:

   ```shell
   kubectl port-forward -n kubeclarity svc/kubeclarity-kubeclarity 9999:8080
   ```

4. Open KubeClarity UI in the browser: <http://localhost:9999/>

### Required K8s permissions

1. Read secrets in cluster scope. This is required for getting image pull secrets for scanning private image repositories.
2. Read config maps in cluster scope. This is required for getting the configured template of the scanner job.
3. List pods in cluster scope. This is required for calculating the target pods that need to be scanned.
4. List namespaces. This is required for fetching the target namespaces to scan in K8s runtime scan UI.
5. Create & delete jobs in cluster scope. This is required for managing the jobs that will scan the target pods in their namespaces.

### Build and run locally with demo data

1. Build UI & backend and start the backend locally (2 options):
    1. Using docker:
        1. Build UI and backend (the image tag is set using VERSION):
           ```shell
           VERSION=test make docker-backend
           ```
        2. Run the backend using demo data:
           ```shell
           docker run -p 8080:8080 -e FAKE_DATA=true -e ENABLE_DB_INFO_LOGS=true -e DATABASE_DRIVER=LOCAL ghcr.io/kubeclarity/kubeclarity:test run
           ```
    2. Local build:
        1. Build UI and backend
           ```shell
           make ui && make backend
           ```
        2. Copy the built site:
           ```shell
           cp -r ./ui/build ./site
           ```
        3. Run the backend locally using demo data:
           ```shell
           FAKE_RUNTIME_SCANNER=true DATABASE_DRIVER=LOCAL FAKE_DATA=true ENABLE_DB_INFO_LOGS=true ./backend/bin/backend run
           ```
    
2. Open KubeClarity UI in the browser: <http://localhost:8080/>

## CLI
KubeClarity includes a CLI that can be run locally and especially useful for CI/CD pipelines. 
It allows to analyze images and directories to generate SBOM, and scan it for vulnerabilities.
The results can be exported to KubeClarity backend.

### Binary Distribution

Download the release distribution for your OS from the [releases page](
https://github.com/cisco-open/kubei/releases)

Unpack the `kubeclarity-cli` binary, add it to your PATH, and you are good to go!

### Docker Image

A Docker image is available at `ghcr.io/cisco-open/kubeclarity-cli` with list of
available tags [here](https://github.com/cisco-open/kubei/pkgs/container/kubeclarity-cli/versions).

### Local Compilation
```
make cli
```

Copy `./cli/bin/cli` to your PATH under `kubeclarity-cli`.

### SBOM generation using multiple integrated content analyzers

```
# A list of the content analyzers to use can be configured using the ANALYZER_LIST env variable seperated by a space (e.g ANALYZER_LIST="syft gomod")
kubeclarity-cli analyze <image/directory name> --input-type <dir|file|image(default)> -o <output file or stdout>

# For example:
ANALYZER_LIST="syft" kubeclarity-cli analyze nginx:latest -o nginx.sbom
```

### Vulnerability scanning using multiple integrated scanners
```
# A list of the vulnerability scanners to use can be configured using the SCANNERS_LIST env variable seperated by a space (e.g SCANNERS_LIST="grype dependency-track")
kubeclarity-cli scan <image/sbom/directoty/file name> --input-type <sbom|dir|file|image(default)> -f <output file>

# For example:
SCANNERS_LIST="grype" kubeclarity-cli scan nginx.sbom --input-type sbom 
```


### Export results to KubeClarity backend
To export CLI results to the KubeClarity backend, need to use an application ID as defined by the KubeClarity backend.
The application ID can be found in the Applications screen in the UI or using the KubeClarity API.

#### Exporting SBOM:
```
# The SBOM can be exported to KubeClarity backend by setting the BACKEND_HOST env variable and the -e flag.
# Note: Until TLS is supported, BACKEND_DISABLE_TLS=true should be set.
BACKEND_HOST=<KubeClarity backend address> BACKEND_DISABLE_TLS=true kubeclarity-cli analyze <image> --application-id <application ID> -e -o <SBOM output file>

# For example:
BACKEND_HOST=localhost:9999 BACKEND_DISABLE_TLS=true kubeclarity-cli analyze nginx:latest --application-id 23452f9c-6e31-5845-bf53-6566b81a2906 -e -o nginx.sbom
```

#### Exporting vulnerability scan results:
```
# The vulnerability scan result can be exported to KubeClarity backend by setting the BACKEND_HOST env variable and the -e flag.
# Note: Until TLS is supported, BACKEND_DISABLE_TLS=true should be set.

BACKEND_HOST=<KubeClarity backend address> BACKEND_DISABLE_TLS=true kubeclarity-cli scan <image> --application-id <application ID> -e

# For example:
SCANNERS_LIST="grype" BACKEND_HOST=localhost:9999 BACKEND_DISABLE_TLS=true kubeclarity-cli scan nginx.sbom --input-type sbom  --application-id 23452f9c-6e31-5845-bf53-6566b81a2906 -e
```

### Merging of SBOM and vulnerabilities across different CI/CD stages
```
# Additional SBOM will be berged into the final results when '--merge-sbom' is defined during analysis. The input SBOM can be CycloneDX XML or CyclonDX json format.
# For example:
ANALYZER_LIST="syft" kubeclarity-cli analyze nginx:latest -o nginx.sbom --merge-sbom inputsbom.xml
```

## Private registries support for K8s runtime scan

### Amazon ECR

Create an [AWS IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console) with `AmazonEC2ContainerRegistryFullAccess` permissions.

Use the user credentials (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_DEFAULT_REGION`) to create the following secret:

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: ecr-sa
  namespace: kubeclarity
type: Opaque
data:
  AWS_ACCESS_KEY_ID: $(echo -n 'XXXX'| base64 -w0)
  AWS_SECRET_ACCESS_KEY: $(echo -n 'XXXX'| base64 -w0)
  AWS_DEFAULT_REGION: $(echo -n 'XXXX'| base64 -w0)
EOF
```

Note:
1. Secret name must be `ecr-sa`
2. Secret data keys must be set to `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_DEFAULT_REGION`

### Google GCR

Create a [Google service account](https://cloud.google.com/docs/authentication/getting-started#creating_a_service_account) with `Artifact Registry Reader` permissions.

Use the service account json file to create the following secret

```
kubectl -n kubeclarity create secret generic --from-file=sa.json gcr-sa
```

Note:
1. Secret name must be `gcr-sa`
1. `sa.json` must be the name of the service account json file when generating the secret
2. KubeClarity is using [application default credentials](https://developers.google.com/identity/protocols/application-default-credentials). These only work when running KubeClarity from GCP.

## Limitations

1. Supports Docker Image Manifest V2, Schema 2 (https://docs.docker.com/registry/spec/manifest-v2-2/). It will fail to scan earlier versions.

## Roadmap
* Integration with additional content analyzers (SBOM generators)
* Integration with additional vulnerability scanners
* CIS Docker benchmark in UI
* Image signing using [Cosign](https://github.com/sigstore/cosign)
* CI/CD metadata signing and attestation using [Cosign](https://github.com/sigstore/cosign) and [in-toto](https://github.com/in-toto/in-toto) (supply chain security)
* System settings and user management

## Contributing

Pull requests and bug reports are welcome.

For larger changes please create an Issue in GitHub first to discuss your
proposed changes and possible implications.

## License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
