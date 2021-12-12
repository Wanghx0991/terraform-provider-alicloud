package scripts

import (
	"flag"
	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
	"testing"
)

var resourceName = flag.String("resource", "", "the name of the terraform resource to diff")

func TestConsistency(t *testing.T) {
	flag.Parse()
	if len(*resourceName) == 0 {
		log.Errorf("The Resource Name is Empty")
		t.Fatal()
	}
	obj := alicloud.Provider().(*schema.Provider).ResourcesMap[*resourceName].Schema
	objSchema := make(map[string]interface{},0)
	objMd,err := parseResource(*resourceName)
	if err != nil{
		log.Error(err)
		t.Fatal()
	}
	mergeMaps(objSchema,objMd.Attributes,objMd.Arguments)
	if len(obj)+1 != len(objSchema){
		log.Errorf("The Field Number of Schema is not consistent with the number in Document")
		t.Fatal()
	}
}

func mergeMaps(Dst map[string]interface{},maps ...map[string]interface{}) map[string]interface{} {
	for _, m := range maps {
		for k, v := range m {
			Dst[k] = v
		}
	}
	return Dst
}