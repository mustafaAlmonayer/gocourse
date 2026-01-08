package intermediate

import (
	"fmt"
	"net/url"
)

func urlParsing() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://mustafa:password@example.com:8080/path?query=param#fargment"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	pass, isSet := parsedUrl.User.Password()
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Username:", parsedUrl.User.Username())
	fmt.Println("Password is set:", isSet)
	fmt.Println("Password:", pass)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Query:", parsedUrl.Query())
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	// Buildin URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	query := baseURL.Query()
	query.Set("name", "John")
	baseURL.RawQuery = query.Encode()
	fmt.Println(baseURL.String())

}
