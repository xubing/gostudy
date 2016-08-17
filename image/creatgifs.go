package main

//TODO: 1 jpg图片色色彩现在压缩的太厉害。
//TODO: 2 图片大小的位置需要
import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	//	 "log"
	//   "os"
	//	 "path/filepath"
	//	 "io/ioutil"
	"flag"
	"io/ioutil"
	"path"
	"strings"

	"os"
	//	"log"
	"log"
)
svncd
// Converts the opened image to the required form of *Image.Paletted by
// drawing into a new *Image.Paletted instance
func convertImage(img image.Image) *image.Paletted {
	//	fmt.Println(img)
	new_img := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(new_img, img.Bounds(), img, image.ZP, draw.Src)
	return new_img
}

func ScanImageAtDir(dir string) (filenames []os.FileInfo) {
	acceptedTypes := map[string]bool{".png": true, ".jpg": true, ".jpeg": true}
	fileInfoArr, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	images := []os.FileInfo{}
	fmt.Println("scan dir at:", dir)
	for _, fileInfo := range fileInfoArr {
		fmt.Println(fileInfo.Name())
		fileExt := path.Ext(fileInfo.Name())
		fileExt = strings.ToLower(fileExt)
		if acceptedTypes[fileExt] {
			fmt.Println(fileInfo.Name(), fileInfo.Size(), "bytes")
			images = append(images, fileInfo)
		}

	}
	return images
}

func readImage(filename string) (image.Image, error) {

	reader, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("decoding image")
	var img image.Image
	img, _, err = image.Decode(reader)
	if err != nil {
		fmt.Println("image.Decode fail: ", err.Error())
		return nil, err
	}
	defer reader.Close()

	return img, nil
}

func main() {

	flag.Parse()

	imagedir := "/Users/Bing/Desktop/picdir3/"
	var images []os.FileInfo
	if images = ScanImageAtDir(imagedir); images == nil {
		fmt.Println("there is no images.\n")
		return
	}

	outfilename := "/Users/Bing/Desktop/output.gif"
	fmt.Println("create gif: ", outfilename)

	// Setting up the empty GIF struct
	var my_gif gif.GIF
	my_gif.LoopCount = 100

	// an accepted image file
	for _, file := range images {
		in_image, err := readImage(imagedir + file.Name())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("converting image")
		new_img := convertImage(in_image)
		fmt.Println("Adding image to gif")
		my_gif.Image = append(my_gif.Image, new_img)
		my_gif.Delay = append(my_gif.Delay, 50)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	out, err := os.Create(outfilename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("encode gif")
	err = gif.EncodeAll(out, &my_gif)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	out.Close()

}
