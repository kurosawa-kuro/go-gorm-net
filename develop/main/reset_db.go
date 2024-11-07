package main

import "go-gorm-net/develop"

func main() {
	develop.CleanupDatabase()
	develop.SeedDatabase()
}
