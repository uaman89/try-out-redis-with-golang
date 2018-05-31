package main

import (
	"log"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	ret, err := c.Do("SELECT","0")
	if (err != nil){
		log.Fatalf("Error on SELECT:: %v\n", err)
	} else {
		fmt.Printf("Select db 0: %s\n", ret)
	}

	ret, _ = c.Do("SET","food", "banana")
	if (err != nil){
		log.Fatalf("Error on SET:: %v\n", err)
	} else {
		fmt.Printf("SET: %s\n", ret)
	}

	ret, _ = c.Do("MSET","black", "batman", "red", "flash", "green", "lantern")
	if (err != nil){
		log.Fatalf("Error on MSET:: %v\n", err)
	} else {
		fmt.Printf("MSET:  %s\n", ret)
	}


	values, err := redis.Strings(c.Do("MGET", "black", "red", "green", "food"))
	if err != nil {
		// handle error
	} else{
		for _, v := range values {
			fmt.Println(v)
		}
	}

	ret, _ = c.Do("SET","delete-me", "1")
	ret, err = c.Do("DEL","delete-me")
	if err != nil {
		log.Fatalf("Error on del:: %v\n", err)
	} else {
		fmt.Printf("del %v\n", ret)
	}

}