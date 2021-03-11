package main


import "C"
import (
    "encoding/binary"
    "fmt"
    "image"
    "image/color"
)
// SHORT = int16
// LONG = int32
// BYTE = uint8
// WORD = uint16
// DWORD = uint32
type SHORT int16
type LONG int32
type BYTE uint8
type WORD uint16
type DWORD uint32

type BITMAPFILEHEADER struct {
    bfType      uint16  // 2
    bfSize      uint32 // 4
    bfReserved1 uint16  // 2
    bfReserved2 uint16  // 2
    bfOffBits   uint32 // 4
}

type BITMAPINFOHEADER struct {
    biSize          uint32
    biWidth         int32
    biHeight        int32
    biPlanes        uint16
    biBitCount      uint16
    biCompression   uint32
    biSizeImage     uint32
    biXPelsPerMeter int32
    biYPelsPerMeter int32
    biClrUsed       uint32
    biClrImportant  uint32
}

type BmpReader struct {
    data []byte
    size int64
}
func (b BmpReader)Read(p []byte) (n int, err error) {
    pLen := len(p)
    copy(p, b.data)
    b.size -= int64(pLen)
    if b.size < 0 {
        return -1, nil
    }
    return pLen, nil
}

func bmpDecoder(input []byte) (image.Image, error) {
    
    bmfHeader := BITMAPFILEHEADER{}
    bmfHeader.bfType = uint16(WORD(binary.LittleEndian.Uint16(input[0:2])))
    if bmfHeader.bfType != 0x4D42 {
        fmt.Println("unsupport format")
        return nil,nil
    }
    bmfHeader.bfSize = uint32(DWORD(binary.LittleEndian.Uint32(input[2:6])))
    bmfHeader.bfReserved1 = uint16(WORD(binary.LittleEndian.Uint16(input[6:8])))
    bmfHeader.bfReserved2 = uint16(WORD(binary.LittleEndian.Uint16(input[8:10])))
    bmfHeader.bfOffBits = uint32(DWORD(binary.LittleEndian.Uint32(input[10:14])))
    
    bi := BITMAPINFOHEADER{}
    bi.biSize = uint32(DWORD(binary.LittleEndian.Uint32(input[14:18])))
    bi.biWidth = int32(LONG(binary.LittleEndian.Uint32(input[18:22])))
    bi.biHeight = int32(LONG(binary.LittleEndian.Uint32(input[22:26])))
    bi.biPlanes = uint16(WORD(binary.LittleEndian.Uint16(input[26:28])))
    bi.biBitCount = binary.LittleEndian.Uint16(input[28:30])
    bi.biCompression = binary.LittleEndian.Uint32(input[30:34])
    bi.biSizeImage = binary.LittleEndian.Uint32(input[34:38])
    bi.biXPelsPerMeter = int32(binary.LittleEndian.Uint32(input[38:42]))
    bi.biYPelsPerMeter = int32(binary.LittleEndian.Uint32(input[42:46]))
    bi.biClrUsed = binary.LittleEndian.Uint32(input[46:50])
    bi.biClrImportant = binary.LittleEndian.Uint32(input[50:54])

    fmt.Println(bmfHeader)
    fmt.Println(bi)
    maxX := int(bi.biWidth)
    maxY := int(bi.biHeight)
    rgba := image.NewRGBA(image.Rectangle{
        Min: image.Point{
            X: 0,
            Y: 0,
        },
        Max: image.Point{
            X: maxX,
            Y: maxY,
        },
    })
    // w -> y
    // h -> x
    base := bmfHeader.bfOffBits
    // 注意：由于位图信息头中的图像高度是正数，所以位图数据在文件中的排列顺序是从左下角到右上角，以行为主序排列的。
    // 从左下角到右上角,从左到右一行行扫描
    fmt.Println("maxX ", maxX,"maxY ", maxY)
    for y := maxY - 1; y >= 0;y -- {
        for x := 0;x < maxX ; x ++ {
            rgba.Set(x, y, color.NRGBA{
                R: input[base + 2],
                G: input[base + 1],
                B: input[base + 0],
                A: input[base + 3],
            })
            base += 4
        }
    }
    return rgba, nil
}
