package main

import (
	"fmt"

	backblaze "../backblaze"
)

func main() {
	fmt.Println("asd")
	filename := "img_071.mp4"

	id := "003a6915e1161c70000000001"
	key := "K003qLU0S1LvKIn//Tn3ltNPLNLz4x4"
	auth := &backblaze.UploadAuth{}

	// b2_authorize_account
	b2, err := NewB2(Credentials{
		KeyID:          id,
		ApplicationKey: key,
	})

	if err != nil {
		panic(err)
	}

	bucket, err := b2.Bucket("jonesphoto")
	if err != nil {
		panic(err)
	}
	auth, fileID, err := bucket.GetPartUploadURL(filename, "b2/x-auto")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", auth)
	fmt.Println(fileID)

}
