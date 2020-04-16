package tools

import (
	"encoding/csv"
	//"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//Point Glass上的点位,范围x-[0,1850] y-[0,1500] 单位mm
type Point struct {
	X int
	Y int
}

//ProductInfo 存储Lottype对应Panel layout信息
type ProductInfo struct {
	LotTpye    string
	PnlLength  int
	Pnlwidth   int
	time       string
	CentPoints map[string]Point
}

//PinInfo 存储各个设备Pin的坐标值
type PinInfo struct {
	PinName   string
	time      string
	PinPoints []Point
}

//Txt2LayoutInfo  txt文件提取玻璃排布信息
func Txt2LayoutInfo(f string, p *ProductInfo) {
	temp, err := os.Open(f)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer temp.Close()

	p.CentPoints = make(map[string]Point)
	cr := csv.NewReader(temp)
	cr.FieldsPerRecord = -1
	s, _ := cr.ReadAll()
	tempStr := strings.Replace(s[0][0], " ", "", -1)
	t := strings.Split(tempStr, "\t")
	p.LotTpye = t[0]
	l, _ := strconv.ParseFloat(t[1], 64)
	w, _ := strconv.ParseFloat(t[2], 64)
	p.PnlLength = int(l)
	p.Pnlwidth = int(w)
	p.time = time.Now().Format("2006/01/02")

	for i := 1; i < len(s); i++ {
		tempStr := strings.Replace(s[i][0], " ", "", -1)
		t = strings.Split(tempStr, "\t")
		var pp Point
		//pp.PanelID = t[0]
		x, _ := strconv.ParseFloat(t[1], 64)
		pp.Y = 750 - int(x)
		y, _ := strconv.ParseFloat(t[2], 64)
		pp.X = 925 - int(y)
		p.CentPoints[t[0]] = pp

	}
	if (p.CentPoints["AA"].X-p.PnlLength/2 < 20) || (p.CentPoints["AA"].Y-p.Pnlwidth/2 < 20) {
		p.PnlLength, p.Pnlwidth = p.Pnlwidth, p.PnlLength
		fmt.Println("Switch L & W!")
	}

}

//Txt2PointInfo  txt文件提取玻璃排布信息
func Txt2PointInfo(f string, p *ProductInfo) []Point {
	temp, err := os.Open(f)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	defer temp.Close()

	cr := csv.NewReader(temp)
	cr.FieldsPerRecord = -1
	s, _ := cr.ReadAll()

	res := make([]Point, 0, 0)

	for i := 0; i < len(s); i++ {
		tempStr := strings.Replace(s[i][0], " ", "", -1)
		t := strings.Split(tempStr, "\t")
		var pp Point
		x, _ := strconv.ParseFloat(t[1], 64)
		y, _ := strconv.ParseFloat(t[2], 64)

		pp.X = p.CentPoints[t[0]].X + int(x) - p.PnlLength/2
		pp.Y = p.CentPoints[t[0]].Y + int(y) - p.Pnlwidth/2

		res = append(res, pp)
	}

	return res
}

//Txt2PinInfo  txt文件提取玻璃排布信息
func Txt2PinInfo(f string, p *PinInfo) {
	temp, err := os.Open(f)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer temp.Close()

	cr := csv.NewReader(temp)
	cr.FieldsPerRecord = -1
	s, _ := cr.ReadAll()
	t := strings.Split(s[0][0], "\t")
	p.PinName = t[0]

	p.time = time.Now().Format("2006/01/02")

	for i := 1; i < len(s); i++ {
		t = strings.Split(s[i][0], "\t")
		var pp Point

		x, _ := strconv.ParseInt(t[0], 10, 0)
		pp.X = 925 + int(x)/1000
		y, _ := strconv.ParseInt(t[1], 10, 0)
		pp.Y = 750 - int(y)/1000
		p.PinPoints = append(p.PinPoints, pp)
	}

}

// //JSON2GlsInfo 读取Json文件到GlsInfo结构体
// func JSON2GlsInfo(f string) []GlsInfo {
// 	filePtr, err := os.Open(f)
// 	if err != nil {
// 		fmt.Println("Open file failed [Err:%s]", err.Error())
// 		return nil
// 	}
// 	defer filePtr.Close()

// 	var gls []GlsInfo

// 	// 创建json解码器
// 	decoder := json.NewDecoder(filePtr)
// 	err = decoder.Decode(&gls)
// 	if err != nil {
// 		fmt.Println("Decoder failed", err.Error())
// 		return nil

// 	}
// 	return gls

// }

// //GlsInfo2JSON 写入GlsInfo结构体到Json文件
// func GlsInfo2JSON(gls *[]GlsInfo) {
// 	// 创建文件
// 	filePtr, err := os.Create("glass_infos.json")
// 	if err != nil {
// 		fmt.Println("Create file failed", err.Error())
// 		return
// 	}
// 	defer filePtr.Close()

// 	// 创建Json编码器
// 	encoder := json.NewEncoder(filePtr)
// 	encoder.SetEscapeHTML(false)

// 	err = encoder.Encode(gls)
// 	if err != nil {
// 		fmt.Println("Encoder failed", err.Error())

// 	} else {
// 		fmt.Println("Encoder success")
// 	}
// }
