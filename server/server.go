package main

// we gotta import some packages that have useful functionality
import (
  // this contains everything to do with HTTP servers
  "net/http"

  // this is a formatting package which can also be used to log messages to the console
  "fmt"
)

// the main() function starts the program
func main() {

  // printing a log to the terminal
  fmt.Println("This writes to the terminal")

  // setting up a request handler
  http.HandleFunc("/", handler)

  // starting the server, running on port 8080
  http.ListenAndServe(":8080", nil)
}

// so we can count how many times the handler has been called, we're saving a counter variable to add to
var count = 0

// The handler is a function that's called every time you send the request
func handler(w http.ResponseWriter, r *http.Request) {

  // we're adding to our request count
  count = count + 1

  // preparing the text that we're going to return
  outputText := fmt.Sprintf("helloo this is a number - %d" + count
  
  // writing the text to the response
  w.Write([]byte(outputText))

  // the request has now finished
}