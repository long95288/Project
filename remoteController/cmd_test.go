package main
/*
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
 */

import (
	"encoding/binary"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func BMPDecode(d []byte) {
	// 解析BMP格式图片

}

type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BitCount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImportant   uint32
}

func TestDecode2(t *testing.T) {
	// out, err := ioutil.ReadFile("hh.bmp")
	out, err := GetScreenCapture()
	if err != nil {
		fmt.Println(err)
		return
	}
	img,_ := bmpDecoder(out)
	f ,_ := os.Create("xxx.jpg")
	e := jpeg.Encode(f, img, nil)
	f2,_ := os.Create("hhh.png")
	r90 := r90d(img)
	e = png.Encode(f2, r90)
	
	fmt.Println(e)
}

func TestDeconde(t *testing.T) {
	file, err := os.Open("hh.bmp")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	//type拆成两个byte来读
	var headA, headB byte
	//Read第二个参数字节序一般windows/linux大部分都是LittleEndian,苹果系统用BigEndian
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	//文件大小
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	//预留字节
	var reservedA, reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	//偏移字节
	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reservedA, reservedB, offbits)

	infoHeader := new(BitmapInfoHeader)
	binary.Read(file, binary.LittleEndian, infoHeader)
	fmt.Println(infoHeader)
}

func TestGetScreenCapture(t *testing.T) {
	out, err := GetScreenCapture()
	fmt.Printf("out size %d err : %v \n", len(out), err)
	out, err = resizeImage(out)
	fmt.Printf(" resize %d err : %v \n", len(out), err)
}
