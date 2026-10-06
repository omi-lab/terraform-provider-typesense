# Local Provider Test Directory

This directory is intended for manual Terraform testing against your locally built provider binary.

## 1) Build the provider

From the repository root:

```bash
go build -o bin/terraform-provider-typesense
```

## 2) Create a local Terraform CLI config

In this directory, create `terraform.rc` with the following content:

```hcl
provider_installation {
  dev_overrides {
    "ronati/typesense" = "/ABSOLUTE/PATH/TO/terraform-provider-typesense/bin"
  }
  direct {}
}
```

Use the absolute path to this repository's `bin` directory.

## 3) Create your variable file

```bash
cp terraform.tfvars.example terraform.tfvars
```

Then update values as needed.

## 4) Run Terraform with the local override

```bash
export TF_CLI_CONFIG_FILE="$(pwd)/terraform.rc"
terraform init
terraform plan
terraform apply
```

## 5) Cleanup

```bash
terraform destroy
```
