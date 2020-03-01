package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	url := `https://album.zhenai.com/u/1417321542`
	contents, err := Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", contents)
}
