package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	_ "image/png"
	"log"
	"strconv"
)

func init() {
	InitBackground()
	InitImages()
	InitNumber2Image()
	InitText()
}

type Player struct {
	x int
	y int
}

type Game struct {
	gameMap [Number][Number]int
	mario   Player
	luigi   Player
	flower  int
}

func checkMarioOrLuigi(x int, y int, mario Player, luigi Player) bool {
	if mario.x == x && mario.y == y {
		return true
	}
	if luigi.x == x && luigi.y == y {
		return true
	}
	return false
}

func MoveRight(gameMap [Number][Number]int, player Player, flower int, mario Player, luigi Player) ([Number][Number]int, Player, int) {
	x, y := player.x, player.y
	if (gameMap[y][x+1] == 0 || gameMap[y][x+1] == 3) && (!checkMarioOrLuigi(x+1, y, mario, luigi)) {
		player.x += 1
	} else if gameMap[y][x+1] == 4 && gameMap[y][x+2] == 0 && (!checkMarioOrLuigi(x+2, y, mario, luigi)) {
		gameMap[y][x+1] = 0
		gameMap[y][x+2] = 4
		player.x += 1
	} else if gameMap[y][x+1] == 4 && gameMap[y][x+2] == 3 && (!checkMarioOrLuigi(x+2, y, mario, luigi)) {
		gameMap[y][x+1] = 0
		gameMap[y][x+2] = 5
		player.x += 1
		flower -= 1
	} else if gameMap[y][x+1] == 5 && gameMap[y][x+2] == 0 && (!checkMarioOrLuigi(x+2, y, mario, luigi)) {
		gameMap[y][x+1] = 3
		gameMap[y][x+2] = 4
		player.x += 1
		flower += 1
	} else if gameMap[y][x+1] == 5 && gameMap[y][x+2] == 3 && (!checkMarioOrLuigi(x+2, y, mario, luigi)) {
		gameMap[y][x+1] = 3
		gameMap[y][x+2] = 5
		player.x += 1
	}
	return gameMap, player, flower
}

func MoveLeft(gameMap [Number][Number]int, player Player, flower int, mario Player, luigi Player) ([Number][Number]int, Player, int) {
	x, y := player.x, player.y
	if (gameMap[y][x-1] == 0 || gameMap[y][x-1] == 3) && (!checkMarioOrLuigi(x-1, y, mario, luigi)) {
		player.x -= 1
	} else if gameMap[y][x-1] == 4 && gameMap[y][x-2] == 0 && (!checkMarioOrLuigi(x-2, y, mario, luigi)) {
		gameMap[y][x-1] = 0
		gameMap[y][x-2] = 4
		player.x -= 1
	} else if gameMap[y][x-1] == 4 && gameMap[y][x-2] == 3 && (!checkMarioOrLuigi(x-2, y, mario, luigi)) {
		gameMap[y][x-1] = 0
		gameMap[y][x-2] = 5
		player.x -= 1
		flower -= 1
	} else if gameMap[y][x-1] == 5 && gameMap[y][x-2] == 0 && (!checkMarioOrLuigi(x-2, y, mario, luigi)) {
		gameMap[y][x-1] = 3
		gameMap[y][x-2] = 4
		player.x -= 1
		flower += 1
	} else if gameMap[y][x-1] == 5 && gameMap[y][x-2] == 3 && (!checkMarioOrLuigi(x-2, y, mario, luigi)) {
		gameMap[y][x-1] = 3
		gameMap[y][x-2] = 5
		player.x -= 1
	}
	return gameMap, player, flower
}

func MoveDown(gameMap [Number][Number]int, player Player, flower int, mario Player, luigi Player) ([Number][Number]int, Player, int) {
	x, y := player.x, player.y
	if (gameMap[y+1][x] == 0 || gameMap[y+1][x] == 3) && (!checkMarioOrLuigi(x, y+1, mario, luigi)) {
		player.y += 1
	} else if gameMap[y+1][x] == 4 && gameMap[y+2][x] == 0 && (!checkMarioOrLuigi(x, y+2, mario, luigi)) {
		gameMap[y+1][x] = 0
		gameMap[y+2][x] = 4
		player.y += 1
	} else if gameMap[y+1][x] == 4 && gameMap[y+2][x] == 3 && (!checkMarioOrLuigi(x, y+2, mario, luigi)) {
		gameMap[y+1][x] = 0
		gameMap[y+2][x] = 5
		player.y += 1
		flower -= 1
	} else if gameMap[y+1][x] == 5 && gameMap[y+2][x] == 0 && (!checkMarioOrLuigi(x, y+2, mario, luigi)) {
		gameMap[y+1][x] = 3
		gameMap[y+2][x] = 4
		player.y += 1
		flower += 1
	} else if gameMap[y+1][x] == 5 && gameMap[y+2][x] == 3 && (!checkMarioOrLuigi(x, y+2, mario, luigi)) {
		gameMap[y+1][x] = 3
		gameMap[y+2][x] = 5
		player.y += 1
	}
	return gameMap, player, flower
}

func MoveUp(gameMap [Number][Number]int, player Player, flower int, mario Player, luigi Player) ([Number][Number]int, Player, int) {
	x, y := player.x, player.y
	if (gameMap[y-1][x] == 0 || gameMap[y-1][x] == 3) && (!checkMarioOrLuigi(x, y-1, mario, luigi)) {
		player.y -= 1
	} else if gameMap[y-1][x] == 4 && gameMap[y-2][x] == 0 && (!checkMarioOrLuigi(x, y-2, mario, luigi)) {
		gameMap[y-1][x] = 0
		gameMap[y-2][x] = 4
		player.y -= 1
	} else if gameMap[y-1][x] == 4 && gameMap[y-2][x] == 3 && (!checkMarioOrLuigi(x, y-2, mario, luigi)) {
		gameMap[y-1][x] = 0
		gameMap[y-2][x] = 5
		player.y -= 1
		flower -= 1
	} else if gameMap[y-1][x] == 5 && gameMap[y-2][x] == 0 && (!checkMarioOrLuigi(x, y-2, mario, luigi)) {
		gameMap[y-1][x] = 3
		gameMap[y-2][x] = 4
		player.y -= 1
		flower += 1
	} else if gameMap[y-1][x] == 5 && gameMap[y-2][x] == 3 && (!checkMarioOrLuigi(x, y-2, mario, luigi)) {
		gameMap[y-1][x] = 3
		gameMap[y-2][x] = 5
		player.y -= 1
	}
	return gameMap, player, flower
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		if Level <= MaxLevel {
			g.gameMap, g.flower, g.mario, g.luigi = InitLevel(Level)
		} else {
			g.gameMap, g.flower, g.mario, g.luigi = InitLevel(1)
			Level = 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.gameMap, g.mario, g.flower = MoveRight(g.gameMap, g.mario, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.gameMap, g.mario, g.flower = MoveLeft(g.gameMap, g.mario, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.gameMap, g.mario, g.flower = MoveUp(g.gameMap, g.mario, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.gameMap, g.mario, g.flower = MoveDown(g.gameMap, g.mario, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.gameMap, g.luigi, g.flower = MoveRight(g.gameMap, g.luigi, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.gameMap, g.luigi, g.flower = MoveLeft(g.gameMap, g.luigi, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.gameMap, g.luigi, g.flower = MoveUp(g.gameMap, g.luigi, g.flower, g.mario, g.luigi)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.gameMap, g.luigi, g.flower = MoveDown(g.gameMap, g.luigi, g.flower, g.mario, g.luigi)
	}

	if g.flower == 0 {
		Level += 1
		if Level <= MaxLevel {
			g.gameMap, g.flower, g.mario, g.luigi = InitLevel(Level)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if Level > MaxLevel {
		text.Draw(screen, "Finish!", MplusNormalFont, (WriteSide+Width)/2, Height/2, color.White)
	} else {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(BackgroundImage, nil)
		for x := 0; x < Number; x++ {
			for y := 0; y < Number; y++ {
				op = &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(CellSide*x+LineSide), float64(CellSide*y+LineSide))
				if g.gameMap[y][x] >= 3 {
					screen.DrawImage(NumberToImage[g.gameMap[y][x]], op)
				}
			}
		}
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(CellSide*g.mario.x+LineSide), float64(CellSide*g.mario.y+LineSide))
		screen.DrawImage(MarioImage, op)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(CellSide*g.luigi.x+LineSide), float64(CellSide*g.luigi.y+LineSide))
		screen.DrawImage(LuigiImage, op)

		//ebitenutil.DebugPrintAt(screen, strconv.Itoa(g.flower), Width+100, 100)
		text.Draw(screen, "level", MplusNormalFont, Width+10, 80, color.White)
		text.Draw(screen, strconv.Itoa(Level)+" / "+strconv.Itoa(MaxLevel), MplusNormalFont, Width+10, 130, color.White)

		text.Draw(screen, "flowers", MplusNormalFont, Width+10, 220, color.White)
		text.Draw(screen, strconv.Itoa(g.flower), MplusNormalFont, Width+10, 270, color.White)

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width + WriteSide, Height
}

func newGame() *Game {
	g := &Game{}
	g.gameMap, g.flower, g.mario, g.luigi = InitLevel(Level)
	return g
}

func main() {

	ebiten.SetWindowSize(Width*Scale+Scale*WriteSide, Height*Scale)
	ebiten.SetWindowTitle("双人推箱子")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}

}
