package scripts

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/waigani/diffparser"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var fileName = flag.String("file_name", "", "the file to check diff")
func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.DisableTimestamp = false
	customFormatter.DisableColors = false
	customFormatter.ForceColors = true
	log.SetFormatter(customFormatter)
	//log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestFieldCompatibilityCheck(t *testing.T) {
	flag.Parse()
	if len(*fileName) == 0 {
		t.Fatal("The DiffFIle is empty")
	}
	byt, _ := ioutil.ReadFile(*fileName)
	diff, _ := diffparser.Parse(string(byt))
	res := false
	for _, file := range diff.Files {
		fmt.Printf("FileName = %s\n", file.NewName)
		for _, hunk := range file.Hunks {
			if hunk != nil {
				prev,current := parseHunk(*hunk)
				res = CompatibilityRule(prev,current)
			}
		}
	}
	if res{
		t.Fatal("Incompatible changes has occurred")
	}
}

func CompatibilityRule(prev,current map[string]map[string]interface{}) (res bool)  {
	for filedName, obj := range prev{
		// Optional -> Required
		_,exist1 := obj["Optional"]
		_,exist2:=current[filedName]["Required"]
		if exist1&&exist2{
			res = true
			log.Errorf("Incompatible changes has occurred: Please Check Out The Field %v: Optional/Required.",filedName)
		}
		// Type changed
		_,exist1 = obj["Type"]
		_,exist2 =current[filedName]["Type"]
		if exist1&&exist2{
			res = true
			log.Errorf("Incompatible changes has occurred: Please Check Out The Field %v Type.",filedName)
		}

		_,exist2 =current[filedName]["ForceNew"]
		if exist2{
			res = true
			log.Errorf("Incompatible changes has occurred: Please Check Out The Field %v: ForceNew.",filedName)
		}

		// type string: valid values
		validateArrPrev,exist1 :=obj["ValidateFuncString"]
		validateArrCurrent,exist2 :=current[filedName]["ValidateFuncString"]
		if exist1&&exist2&&len(validateArrPrev.([]string))>len(validateArrCurrent.([]string)){
			res = true
			log.Errorf("Incompatible changes has occurred: Please Check Out The Field %v valid values. the prev valid values: %v, the current valid values: %v",filedName,validateArrPrev,validateArrCurrent)
		}


	}
	return
}

func parseHunk(hunk diffparser.DiffHunk) (map[string]map[string]interface{}, map[string]map[string]interface{}) {
	prev := parseField( hunk.OrigRange, hunk.OrigRange.Length)
	current := parseField( hunk.NewRange, hunk.NewRange.Length)
	return prev,current
}

func parseField(hunk diffparser.DiffRange,length int) map[string]map[string]interface{} {
	schemaRegex := regexp.MustCompile("^\\t*\"([a-zA-Z_]*)\"")
	typeRegex := regexp.MustCompile("^\\t*Type:\\s+schema.([a-zA-Z]*)")
	optionRegex := regexp.MustCompile("^\\t*Optional:\\s+([a-z]*),")
	forceNewRegex := regexp.MustCompile("^\\t*ForceNew:\\s+([a-z]*),")
	requiredRegex := regexp.MustCompile("^\\t*Required:\\s+([a-z]*),")
	validateStringRegex := regexp.MustCompile("^\\t*ValidateFunc: ?validation.StringInSlice\\(\\[\\]string\\{([a-z\\-A-Z_,\"\\s]*)")

	temp := map[string]interface{}{}
	schemaName := ""
	raw := make(map[string]map[string]interface{},0)
	for i := 0; i < length; i++ {
		currentLine := hunk.Lines[i]
		content := currentLine.Content
		fieldNameMatched := schemaRegex.FindAllStringSubmatch(content, -1)
		if fieldNameMatched != nil && fieldNameMatched[0] != nil {
			if len(schemaName)!=0 && schemaName != fieldNameMatched[0][1]{
				temp["Name"] = schemaName
				raw[schemaName] = temp
				//fmt.Printf("schemaName = %v, typeValue = %v, optionValue = %v,forceNewValue = %v,requiredValue = %v\n",temp["Name"],temp["Type"],temp["Optional"],temp["ForceNew"],temp["Required"])
				temp = map[string]interface{}{}
			}
			schemaName = fieldNameMatched[0][1]
		}

		if !schemaRegex.MatchString(currentLine.Content) && currentLine.Mode == diffparser.UNCHANGED{
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
			op,_ := strconv.ParseBool(optionValue)
			temp["Optional"]=op
		}

		forceNewMatched := forceNewRegex.FindAllStringSubmatch(content, -1)
		forceNewValue := ""
		if forceNewMatched != nil && forceNewMatched[0] != nil {
			forceNewValue = forceNewMatched[0][1]
			fc,_ := strconv.ParseBool(forceNewValue)
			temp["ForceNew"]=fc
		}

		requiredMatched := requiredRegex.FindAllStringSubmatch(content, -1)
		requiredValue := ""
		if requiredMatched != nil && requiredMatched[0] != nil {
			requiredValue = requiredMatched[0][1]
			rq,_ := strconv.ParseBool(requiredValue)
			temp["Required"]=rq
		}


		validateStringMatched := validateStringRegex.FindAllStringSubmatch(content, -1)
		validateStringValue := ""
		if validateStringMatched != nil && validateStringMatched[0] != nil {
			validateStringValue = validateStringMatched[0][1]
			temp["ValidateFuncString"] = strings.Split(validateStringValue,",")
		}

	}
	if _,exist := raw[schemaName];!exist&&len(temp)>=1{
		temp["Name"] = schemaName
		raw[schemaName] = temp
	}
	return raw
}
