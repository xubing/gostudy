package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

// Converts the opened image to the required form of *Image.Paletted by
// drawing into a new *Image.Paletted instance
func convertImage(img image.Image) *image.Paletted {
	//	fmt.Println(img)
	new_img := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(new_img, img.Bounds(), img, image.ZP, draw.Src)
	return new_img
}

func main() {
	fmt.Printf("Making a gif\n")
	dirname := "/Users/Bing/Desktop/picdir/" // + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading " + dirname)

	// Setting up the empty GIF struct
	var my_gif gif.GIF
	my_gif.LoopCount = 100

	// Map of accepted image types
	acceptedTypes := map[string]bool{".png": true, ".jpg": true, ".jpeg": true}

	// Iterating through each file and adding it to the GIF if it is
	// an accepted image file
	for _, file := range files {
		fmt.Println(filepath.Ext(file.Name()))
		if file.Mode().IsRegular() && acceptedTypes[filepath.Ext(file.Name())] {
			fmt.Println(file.Name(), file.Size(), "bytes")
			reader, err := os.Open(file.Name())
			fmt.Println("decoding image")
			in_image, _, err := image.Decode(reader)
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			fmt.Println("converting image")
			new_img := convertImage(in_image)
			fmt.Println("Adding image to gif")
			my_gif.Image = append(my_gif.Image, new_img)
			my_gif.Delay = append(my_gif.Delay, 100)
			if err != nil {
				log.Fatal(err)
			}

		}

	}

	out, err := os.Create("./finalGif.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = gif.EncodeAll(out, &my_gif)
	if err != nil {
		log.Fatal(err)
	}
	out.Close()
}
