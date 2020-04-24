NS1 Terraform Provider
==================

> This project is in [active development](https://github.com/ns1/community/blob/master/project_status/ACTIVE_DEVELOPMENT.md).

- NS1 Website: https://www.ns1.com
- Terraform Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Contents
------
1. [Requirements](#requirements) - lists the requirements for building the provider
2. [Building The Provider](#building-the-provider) - lists the steps for building the provider
3. [Using The Provider](#using-the-provider) - details how to use the provider
4. [Developing The Provider](#developing-the-provider) - steps for contributing back to the provider
5. [Known Isssues/Roadmap](#known-issues) - check here for some of the improvements we are working on

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10+
-	[Go](https://golang.org/doc/install) 1.12+ (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-ns1`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers
$ cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-ns1
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-ns1
$ make build
```

Using The Provider
----------------------

Documentation and examples for NS1 `resources` and `data sources` is part of
this repository, and are published at
[www.terraform.io/docs/providers/ns1](https://www.terraform.io/docs/providers/ns1/index.html)
as part of the release process.

### NS1 Resources

1. [ApiKey](#apikey)
2. [Datafeed](#datafeed)
3. [Datasource](#datasource)
4. [MonitoringJob](#monitoringjob)
5. [NotifyList](#notifylist)
6. [Record](#record)
7. [Team](#team)
8. [User](#user)
9. [Zone](#zone)

#### ApiKey

[ApiKeys Api Doc](https://ns1.com/api#api-key)

[ApiKeys Resource Doc](/website/docs/r/apikey.html.markdown)

#### Datafeed

[Datafeed Api Doc](https://ns1.com/api#data-feeds)

[Datafeed Resource Doc](/website/docs/r/datafeed.html.markdown)

#### Datasource

[Datasource Api Doc](https://ns1.com/api#data-sources)

[Datasource Resource Doc](/website/docs/r/datasource.html.markdown)

#### MonitoringJob

[MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)

[MonitoringJob Resource Doc](/website/docs/r/monitoringjob.html.markdown)

#### NotifyList

[NotifyList Api Doc](https://ns1.com/api#notification-lists)

[NotifyList Resource Doc](/website/docs/r/notifylist.html.markdown)

#### Record

[Record Api Doc](https://ns1.com/api#records)

[Record Resource Doc](/website/docs/r/record.html.markdown)

#### Team

[Team Api Docs](https://ns1.com/api#team)

[Team Resource Doc](/website/docs/r/team.html.markdown)

#### User

[User Api Docs](https://ns1.com/api#user)

[User Resource Doc](/website/docs/r/user.html.markdown)

#### Zone

[Zone Api Docs](https://ns1.com/api#zones)

[Zone Resource Doc](/website/docs/r/zone.html.markdown)

### NS1 Data Sources

1. [DNSSEC](#dnssec)
2. [Zone](#zone-1)

#### DNSSEC

[DNSSEC Data Source Doc](/website/docs/d/dnssec.html.markdown)

#### Zone

[Zone Data Source Doc](/website/docs/d/zone.html.markdown)

Developing The Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine 
(version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH),
as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in 
the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-ns1
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Some helpful things for debugging:

* Set `TF_LOG=DEBUG` for verbose logging.
* Additionally set `NS1_DEBUG` environment variable to include details of the
  API requests in the logs.

Contributions
---

Pull Requests and issues are welcome. See the [NS1 Contribution Guidelines](https://github.com/ns1/community) for more information.

Known Issues/Roadmap
--------------------

* Currently, some arguments marked as required in resource documentation are
  de-facto optional. A resource will be created/updated without error, but
  in general will lead to a "dirty terraform" state, since the defaulted
  attributes on the returned state may not match the resource descriptions.
  We're working on making these either truly Required or truly Optional as
  appropriate.
* Currently, some resources do not return attributes for optional features that
  are unused. We are working on making the resource schemas fixed, with proper
  defaults returned for optional/unused features.
* We'll be adding a `record` data source ASAP, to cover simple read-only use
  cases
