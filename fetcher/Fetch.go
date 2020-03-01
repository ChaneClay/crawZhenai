package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

//
//func Fetch(url string)([]byte, error)  {
//	resp, err := http.Get(url)
//	if err!= nil{
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK{
//		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
//	}
//
//	return ioutil.ReadAll(resp.Body)		// 返回网页源码
//
//}

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
