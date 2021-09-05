package alicloud

import (
	"fmt"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceAlicloudRealtimeComputeClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudRealtimeComputeClustersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DESTROYED", "DESTROYING", "EXPANDING", "MAINTAINING", "REDUCING", "STARTING", "UPGRADING"}, false),
			},
			"page_index": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"page_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gmt_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_infos": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mix_deploy": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_oss_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_vpc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enable_details": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourceAlicloudRealtimeComputeClustersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "ListCluster"
	request := make(map[string]interface{})
	if v, ok := d.GetOk("display_name"); ok {
		request["displayName"] = v
	}
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("status"); ok {
		request["state"] = v
	}
	if v, ok := d.GetOk("page_index"); ok {
		request["pageIndex"] = v
	}
	if v, ok := d.GetOk("page_size"); ok {
		request["pageSize"] = v
	}
	if v, ok := d.GetOk("region"); ok {
		request["region"] = v
	}
	var objects []map[string]interface{}

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			if vv == nil {
				continue
			}
			idsMap[vv.(string)] = vv.(string)
		}
	}
	var response map[string]interface{}
	conn, err := client.NewFoasconsoleClient()
	if err != nil {
		return WrapError(err)
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2018-11-11"), StringPointer("AK"), request, nil, &runtime)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)
	if err != nil {
		return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_realtime_compute_clusters", action, AlibabaCloudSdkGoERROR)
	}
	resp, err := jsonpath.Get("$.Clusters.Cluster", response)
	if err != nil {
		return WrapErrorf(err, FailedGetAttributeMsg, action, "$.Clusters.Cluster", response)
	}
	result, _ := resp.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		if len(idsMap) > 0 {
			if _, ok := idsMap[fmt.Sprint(item["ClusterId"])]; !ok {
				continue
			}
		}
		objects = append(objects, item)
	}
	ids := make([]string, 0)
	s := make([]map[string]interface{}, 0)
	for _, object := range objects {
		mapping := map[string]interface{}{
			"id":           fmt.Sprint(object["ClusterId"]),
			"cluster_id":   fmt.Sprint(object["ClusterId"]),
			"create_time":  fmt.Sprint(object["GmtCreate"]),
			"description":  object["Description"],
			"display_name": object["DisplayName"],
			"gmt_modified": fmt.Sprint(object["GmtModified"]),
			"operator":     object["Operator"],
			"owner_id":     object["OwnerId"],
			"status":       object["State"],
			"zone_id":      object["ZoneId"],
		}
		ids = append(ids, fmt.Sprint(mapping["id"]))
		if detailedEnabled := d.Get("enable_details"); !detailedEnabled.(bool) {
			s = append(s, mapping)
			continue
		}
		id := fmt.Sprint(object["ClusterId"])
		foasService := FoasService{client}
		getResp, err := foasService.DescribeRealtimeComputeCluster(id)
		if err != nil {
			return WrapError(err)
		}
		mapping["instance_infos"] = getResp["InstanceInfos"]
		mapping["is_mix_deploy"] = getResp["IsMixDeploy"]
		mapping["storage_type"] = getResp["StorageType"]
		mapping["user_oss_info"] = getResp["UserOssInfo"]
		mapping["user_vpc_id"] = getResp["UserVpcId"]
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}

	if err := d.Set("clusters", s); err != nil {
		return WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}

	return nil
}
