tfjson2 - a tool for exporting Terraform plans as JSON
==========

Export your Terraform plans to JSON.

Running with Docker
-------------------

tfjson2 is also available via a pre-built Docker container.

```bash
cat $pathToPlan | docker run -i justoman05/tfjson2 --stdin
```


Installing
----------

```bash
go get github.com/justinm/tfjson2
```
 
 
Usage
-----

```bash
terraform plan -out=my.plan

tfjson2 --plan my.plan

Formatted:

tfjson2 --plan my.plan |jq .
```


License
-------

This software uses compiles directly against the Terraform libraries. As such, this project will adopt the same licensing
as Terraform. Please see LICENSE for more information.
