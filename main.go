package main

import (
	"drawdemo/tools"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

//Readtxt 读取excel拷贝到txt的文件数据，数据以制表符(\t)分割
// func Readtxt(filename string,[][])

//DrawPanel 根据长宽画Panel
func DrawPanel(img *image.NRGBA, centerX, centerY, length, width int, color color.Color) {
	// 遍历画每个点
	for x := centerX - length/2; x <= centerX+length/2; x++ {
		for y := centerY - width/2; y <= centerY+width/2; y++ {
			img.Set(x, y, color)
			if x != centerX-length/2 && x != centerX+length/2 {
				y = centerY + width/2
				img.Set(x, y, color)
			}
		}
	}
}

//DrawPoint 根据长宽画Panel
func DrawPoint(img *image.NRGBA, centerX, centerY, size int, color color.Color) {
	// switch types {
	// case 0:
	for x := centerX - size; x <= centerX+size; x++ {
		for y := centerY - size; y <= centerY+size; y++ {
			img.Set(x, y, color)
		}
	}
	// case 1:
	// 	for x := centerX - size; x <= centerX+size; x++ {
	// 		for y := centerY - size; y <= centerY + size; y++ {
	// 			img.Set(x, y, color)
	// 		}
	// 	}
	// }

}

const (
	dx = 1850
	dy = 1500
)

func main() {
	var pi tools.ProductInfo
	var coater tools.PinInfo

	tools.Txt2LayoutInfo("5.76D.txt", &pi)
	tools.Txt2PinInfo("Coater_Pin.txt", &coater)
	fmt.Println(pi)
	fmt.Println(coater)

	imgfile, _ := os.Create("test.png")
	defer imgfile.Close()
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	fmt.Println("Draw a Picture!")
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 0})
		}
	}
	img.Set(925, 750, color.RGBA{255, 255, 0, 255})
	for _, v := range pi.CentPoints {
		DrawPanel(img, v.X, v.Y, pi.PnlLength, pi.Pnlwidth, color.RGBA{255, 0, 255, 255})
	}

	for i := 0; i < len(coater.PinPoints); i++ {
		DrawPoint(img, coater.PinPoints[i].X, coater.PinPoints[i].Y, 3, color.RGBA{255, 0, 0, 255})
	}

	// 以PNG格式保存文件
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
