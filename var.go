package main

import (
	"fmt"
 	"net/http"
 	"net/url"
 	"io/ioutil"
)
func main(){
	//resp, err := http.Get("http://google.com")
	//resp, err := http.Post("http://google.com", "1")
	data := make(map[string]string)

	data["username"] = "kevin"
	data["password"] = "abcd1234"

	host := "http://ptsv2.com/t/bsare-1533929861/post"

	body := curl(host, data)

	fmt.Printf("%s",body)

}


func curl(host string, data map[string]string) string{

	postdata := url.Values{}

	for k,v := range data{
		postdata.Add(k,v)
	}

	resp,err := http.PostForm(host ,postdata)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	


	if resp.StatusCode == http.StatusOK {
   	 	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
   	 	if err2 != nil {
   	 		fmt.Println(err2)
   	 	}
    	bodyString := string(bodyBytes)
    	return bodyString
	}
	return ""
}