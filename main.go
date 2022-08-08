package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Get cache
	c.Set("foo", "bbbbbar", cache.NoExpiration)
	c.SaveFile("cache")

	// Set cache
	c.LoadFile("cache")
	raw, found := c.Get("foo")
	if found {
		fmt.Println(raw.(string))
	}
}
