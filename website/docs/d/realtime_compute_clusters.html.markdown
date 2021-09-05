---
subcategory: "Realtime Compute"
layout: "alicloud"
page_title: "Alicloud: alicloud_realtime_compute_clusters"
sidebar_current: "docs-alicloud-datasource-realtime-compute-clusters"
description: |-
  Provides a list of Realtime Compute Clusters to the user.
---

# alicloud\_realtime\_compute\_clusters

This data source provides the Realtime Compute Clusters of the current Alibaba Cloud user.

-> **NOTE:** Available in v1.134.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_realtime_compute_clusters" "ids" {}
output "realtime_compute_cluster_id_1" {
  value = data.alicloud_realtime_compute_clusters.ids.clusters.0.id
}
            
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional, ForceNew) The name of the cluster.
* `enable_details` - (Optional) Default to `false`. Set it to `true` can output more details about resource attributes.
* `ids` - (Optional, ForceNew, Computed)  A list of Cluster IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `page_index` - (Optional, ForceNew) The page index.
* `page_size` - (Optional, ForceNew) The page size.
* `region` - (Optional, ForceNew) The region.
* `status` - (Optional, ForceNew) The status of the resource.

## Argument Reference

The following attributes are exported in addition to the arguments listed above:

* `clusters` - A list of Realtime Compute Clusters. Each element contains the following attributes:
	* `cluster_id` - The clusterId.
	* `create_time` - Cluster VPC creation time.
	* `description` - Cluster note.
	* `display_name` - The name of the cluster.
	* `gmt_modified` - Cluster VPC modification time.
	* `id` - The ID of the Cluster.
	* `instance_infos` - Instance information.
	* `is_mix_deploy` - Mixed deployment.
	* `operator` - The last operator of the cluster is UID.
	* `owner_id` - Cluster owner UID.
	* `status` - The status of the resource.
	* `storage_type` - Storage type.
	* `user_oss_info` - OSS in the Bucket.
	* `user_vpc_id` - The name of the VPC where the cluster resides.
	* `zone_id` - Available region ID.