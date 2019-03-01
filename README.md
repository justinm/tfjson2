tfjson2 (export Terraform plan as JSON) & OPA (Open Policy Agent)
==========

Running with Docker
-------------------

tfjson2 with OPA is also available via a pre-built Docker container.

```bash
cat $pathToPlan | docker run -i cloudvar/tfjson_with_open-policy-agent:latest --stdin
```


Installing
----------

```bash
go get github.com/cloudvar/tfjson_with_open-policy-agent
```
 
 
Usage
-----

```bash
terraform plan -out=my.plan
```

```bash
tfjson2 --plan /tmp/terraform.plan
```

Run against the container
-----

```bash
docker run -it -v /tmp/terraform.plan:/tmp/terraform.plan cloudvar/tfjson_with_open-policy-agent:latest tfjson2 --plan /tmp/terraform.plan > /tmp/plan.json
```

Open Policy Agent
-----

All OPA policies reside under policies directory

To use OPA, from inside docker (cloudvar/tfjson_with_open-policy-agent:latest)

```bash
/opa eval --data sample-policy.rego --input terraform.plan "data.terraform.analysis.authz"
```

Run against the container

```bash
docker run -v /tmp/plan.json:/tmp/plan.json -it cloudvar/tfjson_with_open-policy-agent:latest /opa eval --data /opt/policies/sample-policy.rego --input /tmp/plan.json "data.terraform.analysis.authz"
```

License
-------

This software uses compiles directly against the Terraform libraries. As such, this project will adopt the same licensing
as Terraform. Please see LICENSE for more information.
