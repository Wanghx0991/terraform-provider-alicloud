---
subcategory: "Realtime Compute"
layout: "alicloud"
page_title: "Alicloud: alicloud_realtime_compute_cluster"
sidebar_current: "docs-alicloud-resource-realtime-compute-cluster"
description: |-
  Provides a Alicloud Realtime Compute Cluster resource.
---

# alicloud\_realtime\_compute\_cluster

Provides a Realtime Compute Cluster resource.

For information about Realtime Compute Cluster and how to use it, see [What is Cluster](https://help.aliyun.com/).

-> **NOTE:** Available in v1.134.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_realtime_compute_cluster" "example" {}

```

## Argument Reference

The following arguments are supported:

* `description` - (Optional, ForceNew) Cluster note.
* `display_name` - (Optional, ForceNew) The name of the cluster.
* `order_id` - (Optional) The order id.
* `user_oss_info` - (Optional, ForceNew) OSS in the Bucket.
* `user_vpc_id` - (Optional, ForceNew) The name of the VPC where the cluster resides.
* `user_vswitch` - (Optional) The user vswitch.
* `zone_id` - (Optional, ForceNew) Available region ID.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Cluster.
* `status` - The status of the resource.

## Import

Realtime Compute Cluster can be imported using the id, e.g.

```
$ terraform import alicloud_realtime_compute_cluster.example <id>
```