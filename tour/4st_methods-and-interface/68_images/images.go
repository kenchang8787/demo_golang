/*
Image

image 套件定義了 Image 接口:

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

需要注意的是，Bounds 方法的返回值 Rectangle 實際上是一個 image.Rectangle，它聲明在 image 套件中。

另外就是 color.Color 和 color.Model 類型也是接口，
但是通常因為直接使用預定義的實現 image.RGBA 和 image.RGBAModel 而被忽視了，
這些接口和類型是由 image/color 套件定義的。
*/
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
