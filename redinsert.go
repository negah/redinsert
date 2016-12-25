package main

import(
  "fmt"
  "os"
  "bufio"
  "log"

  "github.com/garyburd/redigo/redis"
)

func main(){
  filename, prefix := os.Args[1], os.Args[2]

  c, err := redis.Dial("tcp", ":6379")

  if err != nil {
    log.Printf("redis connection error.")
  }

  file, err := os.Open(filename)

  if err != nil {
    log.Printf("redis connection error.")
  }

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    fmt.Println(scanner.Text())
    c.Do("SET",prefix + ":" + scanner.Text(), "1")
  }
}
