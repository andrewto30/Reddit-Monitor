package main

import (
	"fmt"
)

func main() {
	// subreddit := "mechmarket"
	// limit := 1
	// url := fmt.Sprintf("https://www.reddit.com/r/%s/new.json?limit=%d", subreddit, limit)

	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// fmt.Println(string(body))
	childData, err := FetchPosts("mechmart", 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, child := range childData {
		fmt.Printf("Subreddit: %s\n", child.Subreddit)
		fmt.Printf("Selftext: %s\n", child.Selftext)
		// ... (print other fields)
		fmt.Println("---------------------------")
	}

}
