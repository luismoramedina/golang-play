package main

import (
   "fmt"
   "encoding/json"
   )

func main() {
   encoded := "{"+
      `  "tracks": {`+
      `  "items": [ {`+
      `    "uri": "spotify:track:6e7TwT8TUE3BHN7aokBKfq"`+
      `  }]`+
      `}}`

   type Item struct {
      Uri string
   }
   type Items map[string][]Item
   type Tracks map[string]Items

   var tracks Tracks
   
   // Decode the json object
   err := json.Unmarshal([]byte(encoded), &tracks)
   if err != nil {
      panic(err)
   }
   fmt.Println(tracks["tracks"]["items"][0].Uri)
   

}