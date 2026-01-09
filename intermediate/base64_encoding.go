package intermediate

import (
	"encoding/base64"
	"fmt"
)

func base64Ex() {

	data := []byte("He~lo, Base64 Encoding")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	decodedStr := string(decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodedStr))

	// URL safe
	decodedUrlSafe := base64.URLEncoding.EncodeToString(data)
	decodedUrlSafeStr := string(decodedUrlSafe)
	fmt.Println(string(decodedUrlSafeStr))

}
