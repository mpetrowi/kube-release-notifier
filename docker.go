package main

import (
    "fmt"
    "log"

    "github.com/google/go-containerregistry/pkg/name"
    "github.com/google/go-containerregistry/pkg/v1/remote"
)

func containerLabel(repoImage string) {
    imageRef := repoImage

    // Parse the image reference
    ref, err := name.ParseReference(imageRef)
    if err != nil {
        log.Fatalf("Error parsing image reference: %v", err)
    }

    // Fetch the remote image descriptor
    img, err := remote.Image(ref)
    if err != nil {
        log.Fatalf("Error fetching remote image: %v", err)
    }

    // Get the image config
    configFile, err := img.ConfigFile()
    if err != nil {
        log.Fatalf("Error getting image config: %v", err)
    }

    // Print image labels
    fmt.Println("Image Labels:")
    for key, value := range configFile.Config.Labels {
        fmt.Printf("%s: %s\n", key, value)
    }
}
