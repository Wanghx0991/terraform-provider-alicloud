package scripts

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Resource struct {
	Name       string
	Arguments  map[string]interface{}
	Attributes map[string]interface{}
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
	fmt.Printf("The ResourceName = %s\n", resourceName)

	scanner := bufio.NewScanner(file)
	argumentFlag := false
	attrFlag := false

	for scanner.Scan() {
		line := scanner.Text()
		if argsRegex.MatchString(line) {
			argumentFlag = true
			continue
		}
		if attribRegex.MatchString(line) {
			attrFlag = true
			fmt.Printf("%s\n", line)
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
					fmt.Printf("field = %v\n", v)
					//result.Arguments = append(result.Arguments, Field)
					result.Arguments[v.(string)]=Field
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
					fmt.Printf("field = %v\n", Field["Name"])
					result.Attributes[v.(string)] = Field
					//result.Attributes = append(result.Attributes, Field)
				}
			}
		}
	}

	return result, nil
}

func parseMatchLine(words []string, argumentFlag bool) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	if argumentFlag && len(words) >= 4 {
		result["Name"] = string(words[1])
		result["Description"] = string(words[3])
		if strings.Contains(string(words[2]), "Optional") {
			result["Optional"] = true
		}
		if strings.Contains(string(words[2]), "Required") {
			result["Required"] = true
		}
		if strings.Contains(string(words[2]), "ForceNew") {
			result["ForceNew"] = true
		}
		return result
	}
	if !argumentFlag && len(words) >= 3 {
		result["Name"] = string(words[1])
		result["Description"] = string(words[2])
		return result
	}
	return nil
}
