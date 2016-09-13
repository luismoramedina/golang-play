package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
   resp, error := http.Get("https://api.gopro.com/v2/channels/feed/playlists/photo-of-the-day.json?platform=web&page=1&per_page=9")
   fmt.Println("error: ", error)
   if resp != nil {
      a, err := ioutil.ReadAll(resp.Body)
      fmt.Printf("%+q\n", a)
      fmt.Println("err reading: ", err)
   }
}