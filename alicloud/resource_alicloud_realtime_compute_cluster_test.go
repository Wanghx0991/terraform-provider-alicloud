package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudRealtimeComputeCluster_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_realtime_compute_cluster.default"
	ra := resourceAttrInit(resourceId, AlicloudRealtimeComputeClusterMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &FoasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRealtimeComputeCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%srealtimecomputecluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRealtimeComputeClusterBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"user_oss_info": "UserOssInfo",
					"description":   "Description",
					"zone_id":       "ZoneId",
					"user_vpc_id":   "UserVpcId",
					"display_name":  "DisplayName",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_oss_info": "UserOssInfo",
						"description":   "Description",
						"zone_id":       "ZoneId",
						"user_vpc_id":   "UserVpcId",
						"display_name":  "DisplayName",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"order_id", "user_vswitch"},
			},
		},
	})
}

var AlicloudRealtimeComputeClusterMap0 = map[string]string{
	"user_vswitch": NOSET,
	"order_id":     NOSET,
	"status":       CHECKSET,
}

func AlicloudRealtimeComputeClusterBasicDependence0(name string) string {
	return fmt.Sprintf(` 
variable "name" {
  default = "%s"
}
`, name)
}
