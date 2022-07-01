package main

import "fmt"




func main(){
 
  
  person := map[string]string{
    "name": "Jonh Doe",
    "address" : "Citayem",
  
  }
  
  fmt.Println(person["name"], person["address"])
  
  var movie map[string]string = make(map[string]string)
  movie["title"] = "Harry Potter And The Half Blood Prince"
  movie["author"] = "JK Rowling"
  movie["favorit"] = "Muwiw & Heker"
  
  delete(movie,"favorit")
  
  fmt.Println(movie)
  
  
  
  
  
  
  
  
}
