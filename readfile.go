package main

import (
   "fmt"
   "os"
   "bufio"
   )

func main() {

   file, err := os.Open("/home/luis/Dropbox/myinfo/google-music-like.txt")
   defer file.Close()
   fmt.Println(file)
   fmt.Println(err)

   scanner := bufio.NewScanner(file)
   var lines []string
   for scanner.Scan() {
      lines = append(lines, scanner.Text())
   }
   fmt.Println(len(lines))

   for _, line := range lines {
      fmt.Println(line)
   }

}