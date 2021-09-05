package alicloud

import (
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type FoasService struct {
	client *connectivity.AliyunClient
}

func (s *FoasService) DescribeRealtimeComputeCluster(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	conn, err := s.client.NewFoasconsoleClient()
	if err != nil {
		return nil, WrapError(err)
	}
	action := "GetClusterDetails"
	request := map[string]interface{}{
		"RegionId":  s.client.RegionId,
		"clusterId": id,
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Details", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Details", response)
	}
	object = v.(map[string]interface{})
	return object, nil
}

func (s *FoasService) GetClusterDetails(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	conn, err := s.client.NewFoasconsoleClient()
	if err != nil {
		return nil, WrapError(err)
	}
	action := "GetClusterDetails"
	request := map[string]interface{}{
		"RegionId":  s.client.RegionId,
		"clusterId": id,
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Details", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Details", response)
	}
	object = v.(map[string]interface{})
	return object, nil
}
