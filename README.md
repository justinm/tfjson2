tfjson2 (for exporting Terraform plans as JSON) & OPA (Open Policy Agent)
==========

Export your Terraform plans to JSON.

Running with Docker
-------------------

tfjson2 is also available via a pre-built Docker container.

```bash
cat $pathToPlan | docker run -i varuntomar2003/tfjson_with_open-policy-agent --stdin
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


tfjson2 --plan /tmp/terraform.plan

Formatted:

tfjson2 --plan /tmp/terraform.plan |jq .
```

Open Policy Agent:
-----

All the policies are copied under policies directory

To use OPA:

```/opa opa eval --data sample-policy.rego --input terraform.plan "data.terraform.analysis.authz"```

License
-------

This software uses compiles directly against the Terraform libraries. As such, this project will adopt the same licensing
as Terraform. Please see LICENSE for more information.
