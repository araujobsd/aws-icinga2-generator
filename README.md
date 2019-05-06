[![GoDoc](https://godoc.org/github.com/araujobsd/aws-icinga2-generator/plugins?status.svg)](https://godoc.org/github.com/araujobsd/aws-icinga2-generator/)
[![GitHub issues](https://img.shields.io/github/issues/araujobsd/aws-icinga2-generator.svg)](https://github.com/araujobsd/aws-icinga2-generator/issues)
[![GitHub forks](https://img.shields.io/github/forks/araujobsd/aws-icinga2-generator.svg)](https://github.com/araujobsd/aws-icinga2-generator/network)
[![Go Report Card](https://goreportcard.com/badge/github.com/araujobsd/aws-icinga2-generator)](https://goreportcard.com/report/github.com/araujobsd/aws-icinga2-generator)

aws-icinga2-generator
================
This software automatically scan all instances on your AWS account and auto-generate the node configuration file for Icinga2.

## Build instructions
1) `make build`

## How to use
First you need to setup the aws command line, you can get more information from Amazon website.
```
[root@taipei aws-icinga2-generator]# go run .
===> Instance: zone-c.cluster.blabla.com-i-000000000
===> Instance: zone-d.cluster.blabla.com-i-000000000
===> Instance: zone-e.cluster.blabla.com-i-000000000
===> Instance: zone-f.cluster.blabla.com-i-000000000
```

It will save all the configuration files into the output directory.

## Copy it to icinga2
Now you just need to copy it into your zones.d and restart icinga2.

## Copyright and licensing
Distributed under [2-Clause BSD License](https://github.com/araujobsd/aws-icinga2-generator/blob/master/LICENSE).
