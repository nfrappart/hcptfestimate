# hcptfestimate Project

## About the project

This project was born from the several dicussions between Hashicorp Ambassadors like myself and the company, about the lack of visibility and predictability for practitioners, users, customers, wanting to move to HCP Terraform (formerly known as Terraform Cloud). The idea is to bridge the gap with offial tooling to allow us to evaluate the cost of moving our existing workspaces to HCP.

Calculation allow only estimation of Standard Tier offer, since Hashicorp does not disclose the price for Plus Tier. As of writing, pricing model is RUM based, more details available on the company's page ([HCP Terraform Pricing model](https://www.hashicorp.com/products/terraform/pricing)).

### Design

This project is build as a single function binary. It's doing just one thing: count resources and calculate cost based on current Hashicorp pricing, based on a provided terraform state file.

### Status

The hcptfestimate project is at an embryonic stage. It may evolve into a more refined CLI if the community (or even just myself) find a need for it.

## Getting started

Retrieve the repository content:

```bash
# clone the repository
git clone git@github.com:nfrappart/hcptfestimate.git

```

The binary has to be build for your machine's architecture using golang tooling:

```bash
# build the binary with build command
cd hcptfestimate
go build hcptfestimate.go

```

Update your `PATH` to allow calling the binary from any path (use your prefered method: bashrc, zshrc, /usr/loca/bin...), and you're good to go.

