package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

var (
	reWholeLine = regexp.MustCompile(`(?im)^\s+\/\/.*$`)
	reEndOfLine = regexp.MustCompile(`(?im)\/\/[^"\[\]]+$`)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("expecting <source url> <patch>")
	}

	sourceFile := os.Args[1]
	patchFile := os.Args[2]

	log.Println("Downloading source")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(sourceFile)
	if err != nil {
		log.Fatalf("Failed to download source file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download source file: %v", resp.Status)
	}

	originalSrc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	log.Println("Reading patch file")
	patchSrc, err := ioutil.ReadFile(patchFile)
	if err != nil {
		log.Fatalf("Failed to read patch file: %v", err)
	}

	patchSrc = reWholeLine.ReplaceAll(patchSrc, []byte{})
	patchSrc = reEndOfLine.ReplaceAll(patchSrc, []byte{})

	log.Println("Decoding patch")
	patch, err := jsonpatch.DecodePatch(patchSrc)
	if err != nil {
		log.Fatalf("Failed to decode patch: %v", err)
	}

	log.Println("Applying patch")
	patched, err := patch.Apply(originalSrc)
	if err != nil {
		log.Fatalf("Failed to apply patch: %v", err)
	}

	log.Println("Saving swagger.json")
	if err := ioutil.WriteFile("swagger.json", patched, 0644); err != nil {
		log.Fatalf("Failed to save patched swagger: %v", err)
	}
}
