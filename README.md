tfjson2 (export Terraform plans as JSON) & OPA (Open Policy Agent)
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

tfjson2 --plan /tmp/terraform.plan```

Inside docker
-----

```docker run -it -v /tmp/terraform.plan:/tmp/terraform.plan cloudvar/tfjson_with_open-policy-agent:latest --plan /tmp/terraform.plan```

Open Policy Agent
-----

All the OPA policies reside under policies directory

To use OPA

From inside docker (cloudvar/tfjson_with_open-policy-agent:latest)

```/opa opa eval --data sample-policy.rego --input terraform.plan "data.terraform.analysis.authz"```

License
-------

This software uses compiles directly against the Terraform libraries. As such, this project will adopt the same licensing
as Terraform. Please see LICENSE for more information.
