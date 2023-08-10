package main

import (
	"bufio"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

var CellSide int = 45
var LineSide int = 3

const Number int = 20

var Scale int = 1
var WriteSide = 200

var Level = 1
var MaxLevel = 5

//var StartGame = true

var ImageFile string = "images"
var MarioPath string = "mario.jpg"
var LuigiPath string = "luigi.jpg"
var FlowerPath string = "flower.jpg"

var Width int
var Height int

var BackgroundImage *ebiten.Image

func InitBackground() {
	Height = Number * CellSide
	Width = Number * CellSide

	backgroundImage := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			if x%CellSide < LineSide || x%CellSide >= CellSide-LineSide || y%CellSide < LineSide || y%CellSide >= CellSide-LineSide {
				backgroundImage.Set(x, y, color.RGBA{0, 0, 0, 0})
			} else {
				backgroundImage.Set(x, y, color.RGBA{255, 255, 255, 0})
			}
		}
	}
	BackgroundImage = ebiten.NewImageFromImage(backgroundImage)
}

func InitFullColor(R uint8, G uint8, B uint8, A uint8) *ebiten.Image {
	img := image.NewRGBA(image.Rect(0, 0, CellSide-2*LineSide, CellSide-2*LineSide))
	for x := 0; x < CellSide-2*LineSide; x++ {
		for y := 0; y < CellSide-2*LineSide; y++ {
			img.Set(x, y, color.RGBA{R, G, B, A})
		}
	}
	return ebiten.NewImageFromImage(img)
}

var MarioImage *ebiten.Image
var LuigiImage *ebiten.Image
var FlowerImage *ebiten.Image
var BoxImage *ebiten.Image
var BoxFlowerImage *ebiten.Image
var WallImage *ebiten.Image

func GetImageFromFile(path string) *ebiten.Image {
	//img, _, _ := ebitenutil.NewImageFromFile(ImageFile + "/" + path)
	f, _ := os.Open(ImageFile + "/" + path)
	img, _, _ := image.Decode(f)
	img = resize.Resize(uint(CellSide-2*LineSide), uint(CellSide-2*LineSide), img, resize.NearestNeighbor)
	ebitenImg := ebiten.NewImageFromImage(img)

	return ebitenImg
}

func InitImages() {
	MarioImage = GetImageFromFile(MarioPath)

	LuigiImage = GetImageFromFile(LuigiPath)
	FlowerImage = GetImageFromFile(FlowerPath)
	BoxImage = InitFullColor(244, 164, 96, 255)

	BoxFlowerImage = InitFullColor(255, 192, 203, 255)
	WallImage = InitFullColor(127, 127, 127, 255)

}

var NumberToImage map[int]*ebiten.Image

func InitNumber2Image() {
	NumberToImage = make(map[int]*ebiten.Image)
	NumberToImage[1] = MarioImage
	NumberToImage[2] = LuigiImage
	NumberToImage[3] = FlowerImage
	NumberToImage[4] = BoxImage
	NumberToImage[5] = BoxFlowerImage
	NumberToImage[6] = WallImage
}

func InitLevel(level int) ([Number][Number]int, int, Player, Player) {
	var mario Player
	var luigi Player
	var gameMap [Number][Number]int
	var flower int = 0
	path := "levels/" + strconv.Itoa(level) + ".txt"
	file, _ := os.Open(path)
	defer file.Close()
	reader := bufio.NewReader(file) // 读取文本数据
	for y := 0; y < Number; y++ {
		str, _ := reader.ReadString('\n')
		lst := strings.Split(str[:len(str)-2], " ")
		for x := 0; x < Number; x++ {
			i, _ := strconv.Atoi(lst[x])
			gameMap[y][x] = i
			if gameMap[y][x] == 3 {
				flower += 1
			} else if gameMap[y][x] == 1 {
				mario.y = y
				mario.x = x
				gameMap[y][x] = 0
			} else if gameMap[y][x] == 2 {
				gameMap[y][x] = 0
				luigi.y = y
				luigi.x = x
			}
		}
	}
	return gameMap, flower, mario, luigi
}

var MplusNormalFont font.Face

func InitText() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	MplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
}
