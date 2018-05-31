package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"fmt"
)

// Connecting
func main() {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		// handle err
	}
	defer client.Close()

	//_, err = client.Cmd("AUTH", "go-redis").Str()
	//if err != nil {
	//	fmt.Println("Auth error %s", err);
	//}

	_, err = client.Cmd("GET", "foo").Str()
	if err != nil {
		fmt.Println(`Error %s`, err);
	}

	err = client.Cmd("SET", "foo", "bar", "EX", 3600).Err
	if err != nil {
		fmt.Println(`Error %s`, err);
	}

	r := client.Cmd("MGET", "foo", "bar", "baz")
	if r.Err != nil {
		// handle error
	}

	// This:
	fmt.Println("MGET result:")
	l, _ := r.List()
	for _, elemStr := range l {
		fmt.Println(elemStr)
	}

	// is equivalent to this:
	elems, err := r.Array()
	for i := range elems {
		elemStr, _ := elems[i].Str()
		fmt.Println(elemStr)
	}

	client.PipeAppend("GET", "foo")
	client.PipeAppend("SET", "bar", "foo2")
	client.PipeAppend("DEL", "baz")

	// Read GET foo reply
	foo, err := client.PipeResp().Str()
	if err != nil {
		fmt.Printf("foo: %s", foo, "err: %s", err);
	}

	// Read SET bar foo reply
	if err := client.PipeResp().Err; err != nil {
		// handle err
	}

	// Read DEL baz reply
	if err := client.PipeResp().Err; err != nil {
		// handle err
	}

	// Flattening
	//
	// Radix will automatically flatten passed in maps and slices into the argument
	// list. For example, the following are all equivalent:
	//
	client.Cmd("HMSET", "myhash", "key1", "val1", "key2", "val2")
	client.Cmd("HMSET", "myhash", []string{"key1", "val1", "key2", "val2"})
	client.Cmd("HMSET", "myhash", map[string]string{
		"key1": "val1",
		"key2": "val2",
	})
	client.Cmd("HMSET", "myhash", [][]string{
		[]string{"key1", "val1"},
		[]string{"key2", "val2"},
	})

}
