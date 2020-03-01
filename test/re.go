package main

import (
	"crawZhenai/model"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

const workRe = `<div class="m-btn purple" data-v-8b1eac0c>([^<]*)</div>`

func main() {
	profile := model.Profile{}
	contents, err := ioutil.ReadFile("F:\\testData\\reData.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(workRe)
	matches := re.FindAllSubmatch(contents, -1)
	if matches != nil {
		for i, n := range matches {
			data := string(n[1])
			switch i {

			case 0:
				profile.Marriage = data
			case 1:
				profile.Age, _ = strconv.Atoi(data)
			case 2:
				profile.Xinzuo = data
			case 3:
				profile.Height, _ = strconv.Atoi(data)
			case 4:
				profile.Weight, _ = strconv.Atoi(data)
			case 5:
				profile.Hokou = data
			case 6:
				profile.Income = data
			case 7:
				profile.Education = data
			default:
				fmt.Println("No conditions are met......")
			}

		}

	} else {
		fmt.Println("error occured...")
	}
	fmt.Println(profile.Marriage)
	fmt.Println(profile.Education)

}
