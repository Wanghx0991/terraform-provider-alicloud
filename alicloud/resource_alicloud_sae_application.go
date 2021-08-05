package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAlicloudSaeApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudSaeApplicationCreate,
		Read:   resourceAlicloudSaeApplicationRead,
		Update: resourceAlicloudSaeApplicationUpdate,
		Delete: resourceAlicloudSaeApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"acr_assume_role_arn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"app_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_config": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_enable_application_scaling_rule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"batch_wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"change_order_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"command": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_args": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_map_mount_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{1000, 16000, 2000, 32000, 4000, 500, 8000}),
			},
			"custom_host_alias": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"edas_container_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_ahas": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_grey_tag_route": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"envs": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jar_start_args": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jar_start_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jdk": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"liveness": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{1024, 131072, 16384, 2048, 32768, 4096, 65536, 8192}),
			},
			"min_ready_instances": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mount_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mount_host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nas_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oss_ak_id": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"oss_ak_secret": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"oss_mount_descs": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"FatJar", "Image", "War"}, false),
			},
			"package_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"php_arms_config_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"php_config": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"php_config_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_stop": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"readiness": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"command", "initialDelaySeconds", "periodSeconds", "timeoutSeconds"}, false),
			},
			"replicas": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sls_configs": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"RUNNING", "STOPPED"}, false),
			},
			"termination_grace_period_seconds": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 60),
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tomcat_config": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"contextInputType", "contextPath", "httpPort", "maxThreads", "uriEncoding", "useBodyEncoding", "useDefaultConfig"}, false),
			},
			"update_strategy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"version_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"war_start_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_container": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudSaeApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateApplication"
	request := make(map[string]interface{})
	conn, err := client.NewServerlessClient()
	if err != nil {
		return WrapError(err)
	}
	if v, ok := d.GetOk("acr_assume_role_arn"); ok {
		request["AcrAssumeRoleArn"] = v
	}
	if v, ok := d.GetOk("app_description"); ok {
		request["AppDescription"] = v
	}
	request["AppName"] = d.Get("app_name")
	if v, ok := d.GetOkExists("auto_config"); ok {
		request["AutoConfig"] = v
	}
	if v, ok := d.GetOk("command"); ok {
		request["Command"] = v
	}
	if v, ok := d.GetOk("command_args"); ok {
		request["CommandArgs"] = v
	}
	if v, ok := d.GetOk("config_map_mount_desc"); ok {
		request["ConfigMapMountDesc"] = v
	}
	if v, ok := d.GetOk("cpu"); ok {
		request["Cpu"] = v
	}
	if v, ok := d.GetOk("custom_host_alias"); ok {
		request["CustomHostAlias"] = v
	}
	if v, ok := d.GetOkExists("deploy"); ok {
		request["Deploy"] = v
	}
	if v, ok := d.GetOk("edas_container_version"); ok {
		request["EdasContainerVersion"] = v
	}
	if v, ok := d.GetOk("envs"); ok {
		request["Envs"] = v
	}
	if v, ok := d.GetOk("image_url"); ok {
		request["ImageUrl"] = v
	}
	if v, ok := d.GetOk("jar_start_args"); ok {
		request["JarStartArgs"] = v
	}
	if v, ok := d.GetOk("jar_start_options"); ok {
		request["JarStartOptions"] = v
	}
	if v, ok := d.GetOk("jdk"); ok {
		request["Jdk"] = v
	}
	if v, ok := d.GetOk("liveness"); ok {
		request["Liveness"] = v
	}
	if v, ok := d.GetOk("memory"); ok {
		request["Memory"] = v
	}
	if v, ok := d.GetOk("mount_desc"); ok {
		request["MountDesc"] = v
	}
	if v, ok := d.GetOk("mount_host"); ok {
		request["MountHost"] = v
	}
	if v, ok := d.GetOk("namespace_id"); ok {
		request["NamespaceId"] = v
	}
	if v, ok := d.GetOk("nas_id"); ok {
		request["NasId"] = v
	}
	if v, ok := d.GetOk("oss_ak_id"); ok {
		request["OssAkId"] = v
	}
	if v, ok := d.GetOk("oss_ak_secret"); ok {
		request["OssAkSecret"] = v
	}
	if v, ok := d.GetOk("oss_mount_descs"); ok {
		request["OssMountDescs"] = v
	}
	request["PackageType"] = d.Get("package_type")
	if v, ok := d.GetOk("package_url"); ok {
		request["PackageUrl"] = v
	}
	if v, ok := d.GetOk("package_version"); ok {
		request["PackageVersion"] = v
	}
	if v, ok := d.GetOk("php_arms_config_location"); ok {
		request["PhpArmsConfigLocation"] = v
	}
	if v, ok := d.GetOk("php_config"); ok {
		request["PhpConfig"] = v
	}
	if v, ok := d.GetOk("php_config_location"); ok {
		request["PhpConfigLocation"] = v
	}
	if v, ok := d.GetOk("post_start"); ok {
		request["PostStart"] = v
	}
	if v, ok := d.GetOk("pre_stop"); ok {
		request["PreStop"] = v
	}
	if v, ok := d.GetOk("readiness"); ok {
		request["Readiness"] = v
	}
	request["Replicas"] = d.Get("replicas")
	if v, ok := d.GetOk("security_group_id"); ok {
		request["SecurityGroupId"] = v
	}
	if v, ok := d.GetOk("sls_configs"); ok {
		request["SlsConfigs"] = v
	}
	if v, ok := d.GetOk("termination_grace_period_seconds"); ok {
		request["TerminationGracePeriodSeconds"] = v
	}
	if v, ok := d.GetOk("timezone"); ok {
		request["Timezone"] = v
	}
	if v, ok := d.GetOk("tomcat_config"); ok {
		request["TomcatConfig"] = v
	}
	if v, ok := d.GetOk("war_start_options"); ok {
		request["WarStartOptions"] = v
	}
	if v, ok := d.GetOk("web_container"); ok {
		request["WebContainer"] = v
	}
	vswitchId := Trim(d.Get("vswitch_id").(string))
	if vswitchId != "" {
		vpcService := VpcService{client}
		vsw, err := vpcService.DescribeVSwitchWithTeadsl(vswitchId)
		if err != nil {
			return WrapError(err)
		}
		request["VpcId"] = vsw["VpcId"]
		request["VSwitchId"] = vswitchId
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-05-06"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sae_application", action, AlibabaCloudSdkGoERROR)
	}
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	responseData := response["Data"].(map[string]interface{})
	d.SetId(fmt.Sprint(responseData["AppId"]))

	return resourceAlicloudSaeApplicationUpdate(d, meta)
}
func resourceAlicloudSaeApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	saeService := SaeService{client}
	object, err := saeService.DescribeSaeApplication(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sae_application saeService.DescribeSaeApplication Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	d.Set("acr_assume_role_arn", object["AcrAssumeRoleArn"])
	d.Set("app_description", object["AppDescription"])
	d.Set("app_name", object["AppName"])
	d.Set("command", object["Command"])
	d.Set("command_args", object["CommandArgs"])
	d.Set("config_map_mount_desc", object["ConfigMapMountDesc"])
	if v, ok := object["Cpu"]; ok && fmt.Sprint(v) != "0" {
		d.Set("cpu", formatInt(v))
	}
	d.Set("custom_host_alias", object["CustomHostAlias"])
	d.Set("edas_container_version", object["EdasContainerVersion"])
	d.Set("envs", object["Envs"])
	d.Set("image_url", object["ImageUrl"])
	d.Set("jar_start_args", object["JarStartArgs"])
	d.Set("jar_start_options", object["JarStartOptions"])
	d.Set("jdk", object["Jdk"])
	d.Set("liveness", object["Liveness"])
	if v, ok := object["Memory"]; ok && fmt.Sprint(v) != "0" {
		d.Set("memory", formatInt(v))
	}
	if v, ok := object["MinReadyInstances"]; ok && fmt.Sprint(v) != "0" {
		d.Set("min_ready_instances", formatInt(v))
	}
	d.Set("mount_desc", object["MountDesc"])
	d.Set("mount_host", object["MountHost"])
	d.Set("namespace_id", object["NamespaceId"])
	d.Set("nas_id", object["NasId"])
	d.Set("oss_ak_id", object["OssAkId"])
	d.Set("oss_ak_secret", object["OssAkSecret"])
	d.Set("oss_mount_descs", object["OssMountDescs"])
	d.Set("package_type", object["PackageType"])
	d.Set("package_url", object["PackageUrl"])
	d.Set("package_version", object["PackageVersion"])
	d.Set("php_arms_config_location", object["PhpArmsConfigLocation"])
	d.Set("php_config", object["PhpConfig"])
	d.Set("php_config_location", object["PhpConfigLocation"])
	d.Set("post_start", object["PostStart"])
	d.Set("pre_stop", object["PreStop"])
	d.Set("readiness", object["Readiness"])
	if v, ok := object["Replicas"]; ok && fmt.Sprint(v) != "0" {
		d.Set("replicas", formatInt(v))
	}
	d.Set("security_group_id", object["SecurityGroupId"])
	d.Set("sls_configs", object["SlsConfigs"])
	if v, ok := object["TerminationGracePeriodSeconds"]; ok && fmt.Sprint(v) != "0" {
		d.Set("termination_grace_period_seconds", formatInt(v))
	}
	d.Set("timezone", object["Timezone"])
	d.Set("tomcat_config", object["TomcatConfig"])
	d.Set("vswitch_id", object["VSwitchId"])
	d.Set("war_start_options", object["WarStartOptions"])
	d.Set("web_container", object["WebContainer"])
	describeApplicationStatusObject, err := saeService.DescribeApplicationStatus(d.Id())
	if err != nil {
		return WrapError(err)
	}
	d.Set("status", describeApplicationStatusObject["CurrentStatus"])
	return nil
}
func resourceAlicloudSaeApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	saeService := SaeService{client}
	var response map[string]interface{}
	d.Partial(true)

	if !d.IsNewResource() && d.HasChange("security_group_id") {
		request := map[string]interface{}{
			"AppId": d.Id(),
		}
		if v, ok := d.GetOk("security_group_id"); ok {
			request["SecurityGroupId"] = v
		}
	}
	action := "UpdateAppSecurityGroup"
	conn, err := client.NewServerlessClient()
	if err != nil {
		return WrapError(err)
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2019-05-06"), StringPointer("AK"), request, nil, &util.RuntimeOptions{})
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
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	d.SetPartial("security_group_id")
	if !d.IsNewResource() && d.HasChange("cpu") {
		request := map[string]interface{}{
			"AppId": d.Id(),
		}
		if v, ok := d.GetOk("cpu"); ok {
			request["Cpu"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("memory") {
		if v, ok := d.GetOk("memory"); ok {
			request["Memory"] = v
		}
	}
	action := "RescaleApplicationVertically"
	conn, err := client.NewServerlessClient()
	if err != nil {
		return WrapError(err)
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-05-06"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	d.SetPartial("cpu")
	d.SetPartial("memory")
	if !d.IsNewResource() && d.HasChange("acr_assume_role_arn") {
		request := map[string]interface{}{
			"AppId": d.Id(),
		}
		if v, ok := d.GetOk("acr_assume_role_arn"); ok {
			request["AcrAssumeRoleArn"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("command") {
		if v, ok := d.GetOk("command"); ok {
			request["Command"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("command_args") {
		if v, ok := d.GetOk("command_args"); ok {
			request["CommandArgs"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("config_map_mount_desc") {
		if v, ok := d.GetOk("config_map_mount_desc"); ok {
			request["ConfigMapMountDesc"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("custom_host_alias") {
		if v, ok := d.GetOk("custom_host_alias"); ok {
			request["CustomHostAlias"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("edas_container_version") {
		if v, ok := d.GetOk("edas_container_version"); ok {
			request["EdasContainerVersion"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("envs") {
		if v, ok := d.GetOk("envs"); ok {
			request["Envs"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("image_url") {
		if v, ok := d.GetOk("image_url"); ok {
			request["ImageUrl"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("jar_start_args") {
		if v, ok := d.GetOk("jar_start_args"); ok {
			request["JarStartArgs"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("jar_start_options") {
		if v, ok := d.GetOk("jar_start_options"); ok {
			request["JarStartOptions"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("jdk") {
		if v, ok := d.GetOk("jdk"); ok {
			request["Jdk"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("liveness") {
		if v, ok := d.GetOk("liveness"); ok {
			request["Liveness"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("mount_desc") {
		if v, ok := d.GetOk("mount_desc"); ok {
			request["MountDesc"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("mount_host") {
		if v, ok := d.GetOk("mount_host"); ok {
			request["MountHost"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("nas_id") {
		if v, ok := d.GetOk("nas_id"); ok {
			request["NasId"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("oss_ak_id") {
		if v, ok := d.GetOk("oss_ak_id"); ok {
			request["OssAkId"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("oss_ak_secret") {
		if v, ok := d.GetOk("oss_ak_secret"); ok {
			request["OssAkSecret"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("oss_mount_descs") {
		if v, ok := d.GetOk("oss_mount_descs"); ok {
			request["OssMountDescs"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("package_url") {
		if v, ok := d.GetOk("package_url"); ok {
			request["PackageUrl"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("package_version") {
		if v, ok := d.GetOk("package_version"); ok {
			request["PackageVersion"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("php_arms_config_location") {
		if v, ok := d.GetOk("php_arms_config_location"); ok {
			request["PhpArmsConfigLocation"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("php_config") {
		if v, ok := d.GetOk("php_config"); ok {
			request["PhpConfig"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("php_config_location") {
		if v, ok := d.GetOk("php_config_location"); ok {
			request["PhpConfigLocation"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("post_start") {
		if v, ok := d.GetOk("post_start"); ok {
			request["PostStart"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("pre_stop") {
		if v, ok := d.GetOk("pre_stop"); ok {
			request["PreStop"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("readiness") {
		if v, ok := d.GetOk("readiness"); ok {
			request["Readiness"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("sls_configs") {
		if v, ok := d.GetOk("sls_configs"); ok {
			request["SlsConfigs"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("termination_grace_period_seconds") {
		if v, ok := d.GetOk("termination_grace_period_seconds"); ok {
			request["TerminationGracePeriodSeconds"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("timezone") {
		if v, ok := d.GetOk("timezone"); ok {
			request["Timezone"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("tomcat_config") {
		if v, ok := d.GetOk("tomcat_config"); ok {
			request["TomcatConfig"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("war_start_options") {
		if v, ok := d.GetOk("war_start_options"); ok {
			request["WarStartOptions"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("web_container") {
		if v, ok := d.GetOk("web_container"); ok {
			request["WebContainer"] = v
		}
	}
	if v, ok := d.GetOkExists("auto_enable_application_scaling_rule"); ok {
		request["AutoEnableApplicationScalingRule"] = v
	}
	if v, ok := d.GetOk("batch_wait_time"); ok {
		request["BatchWaitTime"] = v
	}
	if v, ok := d.GetOk("change_order_desc"); ok {
		request["ChangeOrderDesc"] = v
	}
	if v, ok := d.GetOk("enable_ahas"); ok {
		request["EnableAhas"] = v
	}
	if v, ok := d.GetOkExists("enable_grey_tag_route"); ok {
		request["EnableGreyTagRoute"] = v
	}
	if v, ok := d.GetOk("update_strategy"); ok {
		request["UpdateStrategy"] = v
	}
	action := "DeployApplication"
	conn, err := client.NewServerlessClient()
	if err != nil {
		return WrapError(err)
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-05-06"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	d.SetPartial("acr_assume_role_arn")
	d.SetPartial("command")
	d.SetPartial("command_args")
	d.SetPartial("config_map_mount_desc")
	d.SetPartial("custom_host_alias")
	d.SetPartial("edas_container_version")
	d.SetPartial("envs")
	d.SetPartial("image_url")
	d.SetPartial("jar_start_args")
	d.SetPartial("jar_start_options")
	d.SetPartial("jdk")
	d.SetPartial("liveness")
	d.SetPartial("mount_desc")
	d.SetPartial("mount_host")
	d.SetPartial("nas_id")
	d.SetPartial("oss_ak_id")
	d.SetPartial("oss_ak_secret")
	d.SetPartial("oss_mount_descs")
	d.SetPartial("package_url")
	d.SetPartial("package_version")
	d.SetPartial("php_arms_config_location")
	d.SetPartial("php_config")
	d.SetPartial("php_config_location")
	d.SetPartial("post_start")
	d.SetPartial("pre_stop")
	d.SetPartial("readiness")
	d.SetPartial("sls_configs")
	d.SetPartial("termination_grace_period_seconds")
	d.SetPartial("timezone")
	d.SetPartial("tomcat_config")
	d.SetPartial("war_start_options")
	d.SetPartial("web_container")
	if d.HasChange("status") {
		object, err := saeService.DescribeSaeApplication(d.Id())
		if err != nil {
			return WrapError(err)
		}
		target := d.Get("status").(string)
		if object[""].(string) != target {
			if target == "RUNNING" {
				request := map[string]interface{}{
					"AppId": d.Id(),
				}
				action := "StartApplication"
				conn, err := client.NewServerlessClient()
				if err != nil {
					return WrapError(err)
				}
				wait := incrementalWait(3*time.Second, 3*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2019-05-06"), StringPointer("AK"), request, nil, &util.RuntimeOptions{})
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
				if fmt.Sprint(response["Success"]) == "false" {
					return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
				}
			}
			if target == "STOPPED" {
				request := map[string]interface{}{
					"AppId": d.Id(),
				}
				action := "StopApplication"
				conn, err := client.NewServerlessClient()
				if err != nil {
					return WrapError(err)
				}
				wait := incrementalWait(3*time.Second, 3*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2019-05-06"), StringPointer("AK"), request, nil, &util.RuntimeOptions{})
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
				if fmt.Sprint(response["Success"]) == "false" {
					return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
				}
			}
			d.SetPartial("status")
		}
	}
	d.Partial(false)
	if d.HasChange("MinReadyInstances") {
		oldMinReadyInstances, newMinReadyInstances := d.GetChange("MinReadyInstances")

		if added.Len() > 0 {
			rescaleapplicationrequest := map[string]interface{}{
				"AppId": d.Id(),
			}
			rescaleapplicationrequest["Replicas"] = d.Get("replicas")
			rescaleapplicationrequest["MinReadyInstances"] = d.Get("min_ready_instances")
			if _, ok := d.GetOkExists("auto_enable_application_scaling_rule"); ok {
				rescaleapplicationrequest["AutoEnableApplicationScalingRule"] = d.Get("auto_enable_application_scaling_rule")
			}
			action := "RescaleApplication"
			conn, err := client.NewServerlessClient()
			if err != nil {
				return WrapError(err)
			}
			wait := incrementalWait(3*time.Second, 3*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2019-05-06"), StringPointer("AK"), rescaleapplicationrequest, nil, &util.RuntimeOptions{})
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, rescaleapplicationrequest)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
			if fmt.Sprint(response["Success"]) == "false" {
				return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
			}
			d.SetPartial("MinReadyInstances")
			d.SetPartial("MinReadyInstances")
			d.SetPartial("MinReadyInstances")
		}
	}
	return resourceAlicloudSaeApplicationRead(d, meta)
}
func resourceAlicloudSaeApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "DeleteApplication"
	var response map[string]interface{}
	conn, err := client.NewServerlessClient()
	if err != nil {
		return WrapError(err)
	}
	request := map[string]interface{}{
		"AppId": d.Id(),
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2019-05-06"), StringPointer("AK"), request, nil, &util.RuntimeOptions{})
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
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	return nil
}
