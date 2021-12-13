package scripts

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
	"github.com/waigani/diffparser"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.DisableTimestamp = false
	customFormatter.DisableColors = false
	customFormatter.ForceColors = true
	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

var (
	resourceName = flag.String("resource", "", "the name of the terraform resource to diff")
	fileName     = flag.String("file_name", "", "the file to check diff")
)

type Resource struct {
	Name       string
	Arguments  map[string]interface{}
	Attributes map[string]interface{}
}

func TestConsistencyWithMD(t *testing.T) {
	flag.Parse()
	if len(*resourceName) == 0 {
		log.Errorf("The Resource Name is Empty")
		t.Fatal()
	}
	obj := alicloud.Provider().(*schema.Provider).ResourcesMap[*resourceName].Schema
	objSchema := make(map[string]interface{}, 0)
	objMd, err := parseResource(*resourceName)
	if err != nil {
		log.Error(err)
		t.Fatal()
	}
	mergeMaps(objSchema, objMd.Attributes, objMd.Arguments)
	if len(obj)+1 != len(objSchema) {
		log.Errorf("The Field Number of Schema is not consistent with the number in Document")
		t.Fatal()
	}
}

func TestFieldCompatibilityCheck(t *testing.T) {
	flag.Parse()
	if len(*fileName) == 0 {
		t.Fatal("The DiffFIle is empty")
	}
	byt, _ := ioutil.ReadFile(*fileName)
	diff, _ := diffparser.Parse(string(byt))
	res := false
	fileRegex := regexp.MustCompile("alicloud/resource[a-zA-Z_]*.go")
	fileTestRegex := regexp.MustCompile("alicloud/resource[a-zA-Z_]*_test.go")
	for _, file := range diff.Files {
		if fileRegex.MatchString(file.NewName) {
			if fileTestRegex.MatchString(file.NewName) {
				continue
			}
			for _, hunk := range file.Hunks {
				if hunk != nil {
					prev := ParseField(hunk.OrigRange, hunk.OrigRange.Length)
					current := ParseField(hunk.NewRange, hunk.NewRange.Length)
					res = CompatibilityRule(prev, current, file.NewName)
				}
			}
		}
	}
	if res {
		t.Fatal("Incompatible changes has occurred")
	}
}

func CompatibilityRule(prev, current map[string]map[string]interface{}, fileName string) (res bool) {
	for filedName, obj := range prev {
		// Optional -> Required
		_, exist1 := obj["Optional"]
		_, exist2 := current[filedName]["Required"]
		if exist1 && exist2 {
			res = true
			log.Errorf("The File: %v, Incompatible changes has occurred: Please Check Out The Field %v: Optional/Required.", fileName, filedName)
		}
		// Type changed
		_, exist1 = obj["Type"]
		_, exist2 = current[filedName]["Type"]
		if exist1 && exist2 {
			res = true
			log.Errorf("The File: %v, Incompatible changes has occurred: Please Check Out The Field %v Type.", fileName, filedName)
		}

		_, exist2 = current[filedName]["ForceNew"]
		if exist2 {
			res = true
			log.Errorf("The File: %v, Incompatible changes has occurred: Please Check Out The Field %v: ForceNew.", fileName, filedName)
		}

		// type string: valid values
		validateArrPrev, exist1 := obj["ValidateFuncString"]
		validateArrCurrent, exist2 := current[filedName]["ValidateFuncString"]
		if exist1 && exist2 && len(validateArrPrev.([]string)) > len(validateArrCurrent.([]string)) {
			res = true
			log.Errorf("The File: %v, Incompatible changes has occurred: Please Check Out The Field %v valid values. the prev valid values: %v, the current valid values: %v", fileName, filedName, validateArrPrev, validateArrCurrent)
		}

	}
	return
}

func ParseField(hunk diffparser.DiffRange, length int) map[string]map[string]interface{} {
	schemaRegex := regexp.MustCompile("^\\t*\"([a-zA-Z_]*)\"")
	typeRegex := regexp.MustCompile("^\\t*Type:\\s+schema.([a-zA-Z]*)")
	optionRegex := regexp.MustCompile("^\\t*Optional:\\s+([a-z]*),")
	forceNewRegex := regexp.MustCompile("^\\t*ForceNew:\\s+([a-z]*),")
	requiredRegex := regexp.MustCompile("^\\t*Required:\\s+([a-z]*),")
	validateStringRegex := regexp.MustCompile("^\\t*ValidateFunc: ?validation.StringInSlice\\(\\[\\]string\\{([a-z\\-A-Z_,\"\\s]*)")

	temp := map[string]interface{}{}
	schemaName := ""
	raw := make(map[string]map[string]interface{}, 0)
	for i := 0; i < length; i++ {
		currentLine := hunk.Lines[i]
		content := currentLine.Content
		fieldNameMatched := schemaRegex.FindAllStringSubmatch(content, -1)
		if fieldNameMatched != nil && fieldNameMatched[0] != nil {
			if len(schemaName) != 0 && schemaName != fieldNameMatched[0][1] {
				temp["Name"] = schemaName
				raw[schemaName] = temp
				temp = map[string]interface{}{}
			}
			schemaName = fieldNameMatched[0][1]
		}

		if !schemaRegex.MatchString(currentLine.Content) && currentLine.Mode == diffparser.UNCHANGED {
			continue
		}

		typeMatched := typeRegex.FindAllStringSubmatch(content, -1)
		typeValue := ""
		if typeMatched != nil && typeMatched[0] != nil {
			typeValue = typeMatched[0][1]
			temp["Type"] = typeValue
		}

		optionalMatched := optionRegex.FindAllStringSubmatch(content, -1)
		optionValue := ""
		if optionalMatched != nil && optionalMatched[0] != nil {
			optionValue = optionalMatched[0][1]
			op, _ := strconv.ParseBool(optionValue)
			temp["Optional"] = op
		}

		forceNewMatched := forceNewRegex.FindAllStringSubmatch(content, -1)
		forceNewValue := ""
		if forceNewMatched != nil && forceNewMatched[0] != nil {
			forceNewValue = forceNewMatched[0][1]
			fc, _ := strconv.ParseBool(forceNewValue)
			temp["ForceNew"] = fc
		}

		requiredMatched := requiredRegex.FindAllStringSubmatch(content, -1)
		requiredValue := ""
		if requiredMatched != nil && requiredMatched[0] != nil {
			requiredValue = requiredMatched[0][1]
			rq, _ := strconv.ParseBool(requiredValue)
			temp["Required"] = rq
		}

		validateStringMatched := validateStringRegex.FindAllStringSubmatch(content, -1)
		validateStringValue := ""
		if validateStringMatched != nil && validateStringMatched[0] != nil {
			validateStringValue = validateStringMatched[0][1]
			temp["ValidateFuncString"] = strings.Split(validateStringValue, ",")
		}

	}
	if _, exist := raw[schemaName]; !exist && len(temp) >= 1 {
		temp["Name"] = schemaName
		raw[schemaName] = temp
	}
	return raw
}

func parseResource(resourceName string) (*Resource, error) {
	splitRes := strings.Split(resourceName, "alicloud_")
	if len(splitRes) < 2 {
		log.Errorf("the resource name parsed failed")
		return nil, fmt.Errorf("the resource name parsed failed")
	}
	basePath := "../website/docs/r/"
	filePath := strings.Join([]string{basePath, splitRes[1], ".html.markdown"}, "")

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", filePath, err)
		return nil, err
	}
	defer file.Close()

	argsRegex := regexp.MustCompile("## Argument Reference")
	attribRegex := regexp.MustCompile("## Attributes Reference")
	secondLevelRegex := regexp.MustCompile("^\\#+")
	argumentsFieldRegex := regexp.MustCompile("^\\* `([a-zA-Z_0-9]*)`[ ]*-? ?(\\(.*\\)) ?(.*)")
	attributeFieldRegex := regexp.MustCompile("^\\* `([a-zA-Z_0-9]*)`[ ]*-?(.*)")

	name := filepath.Base(filePath)
	re := regexp.MustCompile("[a-zA-Z_]*")
	resourceName = "alicloud_" + re.FindString(name)
	result := &Resource{Name: resourceName, Arguments: map[string]interface{}{}, Attributes: map[string]interface{}{}}
	log.Infof("The ResourceName = %s\n", resourceName)

	scanner := bufio.NewScanner(file)
	argumentFlag := false
	attrFlag := false

	attrName := make([]string, 0)
	argumentName := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if argsRegex.MatchString(line) {
			argumentFlag = true
			continue
		}
		if attribRegex.MatchString(line) {
			attrFlag = true
			continue
		}
		if argumentFlag {
			if secondLevelRegex.MatchString(line) {
				argumentFlag = false
				continue
			}
			argumentsMatched := argumentsFieldRegex.FindAllStringSubmatch(line, 1)
			for _, argumentMatched := range argumentsMatched {
				Field := parseMatchLine(argumentMatched, true)
				if v, exist := Field["Name"]; exist {
					argumentName = append(argumentName, v.(string))
					result.Arguments[v.(string)] = Field
				}
			}
		}

		if attrFlag {
			if secondLevelRegex.MatchString(line) {
				attrFlag = false
				continue
			}
			attributesMatched := attributeFieldRegex.FindAllStringSubmatch(line, 1)
			for _, attributeParsed := range attributesMatched {
				Field := parseMatchLine(attributeParsed, false)
				if v, exist := Field["Name"]; exist {
					attrName = append(attrName, v.(string))
					result.Attributes[v.(string)] = Field
				}
			}
		}
	}
	log.Infof("the field name in arguments: %#v\n", argumentName)
	log.Infof("the field name in attrbuites: %#v\n", attrName)
	return result, nil
}

func parseMatchLine(words []string, argumentFlag bool) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	if argumentFlag && len(words) >= 4 {
		result["Name"] = words[1]
		result["Description"] = words[3]
		if strings.Contains(words[2], "Optional") {
			result["Optional"] = true
		}
		if strings.Contains(words[2], "Required") {
			result["Required"] = true
		}
		if strings.Contains(words[2], "ForceNew") {
			result["ForceNew"] = true
		}
		return result
	}
	if !argumentFlag && len(words) >= 3 {
		result["Name"] = words[1]
		result["Description"] = words[2]
		return result
	}
	return nil
}

func mergeMaps(Dst map[string]interface{}, maps ...map[string]interface{}) map[string]interface{} {
	for _, m := range maps {
		for k, v := range m {
			Dst[k] = v
		}
	}
	return Dst
}
