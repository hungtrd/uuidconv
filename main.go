package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	uuidStr := flag.String("uuid", "", "UUID dạng string cần chuyển đổi")
	decodeFlag := flag.Bool("decode", false, "Chuyển đổi từ base64 encoded sang UUID dạng string")
	flag.Parse()

	if *uuidStr == "" {
		log.Fatal("Bạn cần cung cấp một UUID dạng string để chuyển đổi.")
	}

	if *decodeFlag {
		decodedUUID, err := base64.StdEncoding.DecodeString(*uuidStr)
		if err != nil {
			log.Fatalf("Lỗi khi chuyển đổi base64: %v", err)
		}

		u, err := uuid.FromBytes(decodedUUID)
		if err != nil {
			log.Fatalf("Lỗi khi chuyển đổi UUID: %v", err)
		}

		fmt.Printf("Base64 encoded: %s\n", *uuidStr)
		fmt.Printf("UUID dạng string: %s\n", u.String())
	} else {
		u, err := uuid.Parse(*uuidStr)
		if err != nil {
			log.Fatalf("Lỗi khi chuyển đổi UUID: %v", err)
		}

		uuidBytes := u[:]
		base64Encoded := base64.StdEncoding.EncodeToString(uuidBytes)

		fmt.Printf("UUID dạng string: %s\n", *uuidStr)
		fmt.Printf("Base64 encoded: %s\n", base64Encoded)
	}
}
