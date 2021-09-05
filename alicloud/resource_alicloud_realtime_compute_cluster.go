package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudRealtimeComputeCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudRealtimeComputeClusterCreate,
		Read:   resourceAlicloudRealtimeComputeClusterRead,
		Delete: resourceAlicloudRealtimeComputeClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_oss_info": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"order_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_vswitch": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudRealtimeComputeClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateCluster"
	request := make(map[string]interface{})
	conn, err := client.NewFoasconsoleClient()
	if err != nil {
		return WrapError(err)
	}
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	if v, ok := d.GetOk("display_name"); ok {
		request["displayName"] = v
	}
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("user_oss_info"); ok {
		request["userOssBucket"] = v
	}
	if v, ok := d.GetOk("user_vpc_id"); ok {
		request["userVpcId"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["zoneId"] = v
	}
	if v, ok := d.GetOk("order_id"); ok {
		request["orderId"] = v
	}
	if v, ok := d.GetOk("user_vswitch"); ok {
		request["userVSwitch"] = v
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-11-11"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_realtime_compute_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ClusterId"]))

	return resourceAlicloudRealtimeComputeClusterRead(d, meta)
}
func resourceAlicloudRealtimeComputeClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	foasService := FoasService{client}
	object, err := foasService.DescribeRealtimeComputeCluster(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_realtime_compute_cluster foasService.DescribeRealtimeComputeCluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	d.Set("description", object["Description"])
	d.Set("display_name", object["DisplayName"])
	d.Set("status", object["State"])
	d.Set("user_oss_info", object["UserOssInfo"])
	d.Set("user_vpc_id", object["UserVpcId"])
	d.Set("zone_id", object["ZoneId"])
	return nil
}
func resourceAlicloudRealtimeComputeClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "DestroyCluster"
	var response map[string]interface{}
	conn, err := client.NewFoasconsoleClient()
	if err != nil {
		return WrapError(err)
	}
	request := map[string]interface{}{
		"clusterId": d.Id(),
	}

	request["RegionId"] = client.RegionId
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2018-11-11"), StringPointer("AK"), request, nil, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}
	return nil
}
