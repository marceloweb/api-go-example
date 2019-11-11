package main

import (
  "fmt",
  "net/http"
)

const (
   maxPosts = 20
)

type Parameters struct {
   Name string `json:"name"`
   User string `json:"user"`
   Followers int `json:"followers"`
}

func getData(key string) {
   apiError := new(APIError)
   resp, err := s.sling.New().Get()
   return resp, err
}

func main() {
  return getData()
} 
