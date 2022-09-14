package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func main() {
	strHash("Str test")
	fileHash()
}

func strHash(s string) {
	hMd5 := md5.Sum([]byte(s))
	hSha1 := sha1.Sum([]byte(s))
	hSha2 := sha256.Sum256([]byte(s))

	log.Printf("   MD5: %x\n", hMd5)
	log.Printf("  SHA1: %x\n", hSha1)
	log.Printf("SHA256: %x\n", hSha2)
}

func fileHash() {
	file, err := os.Open("invaliddata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	buf := make([]byte, 30*1024)
	hSha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := hSha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum := hSha256.Sum(nil)
	log.Printf("%x\n", sum)
}
