package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	uuidStr := flag.String("uuid", "", "UUID dạng string cần chuyển đổi")
	decodeFlag := flag.Bool("d", false, "Chuyển đổi từ base64 encoded sang UUID dạng string")
	b64Flag := flag.Bool("b64", false, "Chọn định dạng muốn copy là base64")
	flag.Parse()

	if *uuidStr == "" {
		u := uuid.New()
		encoded := base64.StdEncoding.EncodeToString(u[:])

		fmt.Println("-------------------------")
		fmt.Println("UUID string: ", u.String())
		fmt.Println("Base64 encoded: ", encoded)
		fmt.Println("-------------------------")
		if *b64Flag {
			if err := clipboard.WriteAll(encoded); err != nil {
				log.Printf("write to clipboard failed: %v", err)
			}
			fmt.Printf("Đã copy dạng base64(%s) vào clipboard!\n", encoded)
		} else {
			if err := clipboard.WriteAll(u.String()); err != nil {
				log.Printf("write to clipboard failed: %v", err)
			}
			fmt.Printf("Đã copy dạng string(%s) vào clipboard!\n", u.String())
		}

		return
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

		fmt.Println("-------------------------")
		fmt.Printf("Base64 encoded: %s\n", *uuidStr)
		fmt.Printf("UUID dạng string: %s\n", u.String())
		fmt.Println("-------------------------")
		if err := clipboard.WriteAll(u.String()); err != nil {
			log.Printf("write to clipboard failed: %v", err)
		}
		fmt.Printf("Đã copy dạng string(%s) vào clipboard!\n", u.String())
	} else {
		u, err := uuid.Parse(*uuidStr)
		if err != nil {
			log.Fatalf("Lỗi khi chuyển đổi UUID: %v", err)
		}

		uuidBytes := u[:]
		base64Encoded := base64.StdEncoding.EncodeToString(uuidBytes)

		fmt.Println("-------------------------")
		fmt.Printf("UUID dạng string: %s\n", *uuidStr)
		fmt.Printf("Base64 encoded: %s\n", base64Encoded)
		fmt.Println("-------------------------")
		if err := clipboard.WriteAll(base64Encoded); err != nil {
			log.Printf("write to clipboard failed: %v", err)
		}
		fmt.Printf("Đã copy dạng base64(%s) vào clipboard!\n", base64Encoded)
	}
}
