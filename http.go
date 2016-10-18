package main

import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"
import "encoding/json"

func main() {

   var myurl *url.URL
   myurl, err := url.Parse("https://api.spotify.com/v1/search")
   fmt.Println(err)
   parameters := url.Values{}
   parameters.Add("type", "track")
   parameters.Add("q", "She Him Why Do You Let Me Stay Here")
   myurl.RawQuery = parameters.Encode()

   resp, error := http.Get(myurl.String())
   fmt.Println("error: ", error)
   if resp != nil {
      a, err := ioutil.ReadAll(resp.Body)
      fmt.Printf("%+q\n", a)
      fmt.Println("err reading: ", err)
   }
}