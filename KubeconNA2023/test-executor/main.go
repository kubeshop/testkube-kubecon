package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var name string

	name, ok := os.LookupEnv("NAME")
	if !ok {
		name = "Kubecon User"
	}

	text := fmt.Sprintf(`
  .////////////////////////////////////* 
(((///(((((((((((((((//(((((((((((((((/ 
(((///(((((((((((((/(((((/((((((((((((/ 
/((*//(((((((((((/(/((((/((/(((((((((// 
 (((//(((((((((((///((((/((/(((((((((/  
 /((///((((((((((((///((((((((((((((//  
 /((///((((((((((((((//(((((((((((((/.  
  (((//(((((((((((((((((((((((((((((/   
  *((///(((###((((((((((((((((###((//   
  .((///((((#((((((((((((((((((((((/    
   ((///((((((((((((((((((((((((((//    
    /(///((((((((((((((((((((((((//     
       (((((((((((((  ((((((((           
       ,((((((((((((   (((((((           
            (((((                  

	 
Thanks %s for trying TestKube Cloud!

Check if there are still some plushies 
available in Testkube booth!
	`, name)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(time.Second)
	}
}
