package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}
	fmt.Println("Websites:", websites)
	fmt.Println("Google URL:", websites["Google"])
	websites["Blog"] = "https://xiaolongya.cn"
	fmt.Println("Websites after adding Blog:", websites)
	delete(websites, "Google")
	fmt.Println("Websites after deleting Google:", websites)
}
