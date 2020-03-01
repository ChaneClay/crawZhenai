package parser

import (
	"crawZhenai/engine"
	"crawZhenai/model"
	"fmt"
	"regexp"
)

const dataRe = `<div class="m-btn purple" data-v-8b1eac0c>([^<]*)</div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	re := regexp.MustCompile(dataRe)
	matches := re.FindAllSubmatch(contents, -1)
	//fmt.Println(len(matches))
	if matches != nil {
		for i, n := range matches {
			//fmt.Printf("%d   %s", i, string(n[1]))
			//str := string(n[1])
			//fmt.Printf("%d--%s\n", i, str)
			data := string(n[1])
			switch i {
			case 0:
				profile.Marriage = data
			case 1:
				profile.Age = data
			case 2:
				profile.Xinzuo = data
			case 3:
				profile.Height = data
			case 4:
				profile.Weight = data
			case 5:
				profile.Hokou = data
			case 6:
				profile.Income = data
			case 7:
				profile.Occupation = data
			case 8:
				profile.Education = data
			default:
				fmt.Printf("%s index: %d, data: %s\n", "No conditions are met  ", i, data)
			}

		}

	} else {
		fmt.Println("error occured...")
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	//fmt.Println("...................................................")

	return result

}
