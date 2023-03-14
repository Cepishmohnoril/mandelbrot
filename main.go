package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 800
	height = 800
	maxIt  = 128
)

var palette [maxIt]byte

type Game struct {
	offscreen    *ebiten.Image
	offscreenPix []byte
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Mandelbrot")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func init() {
}

func NewGame() *Game {
	g := &Game{
		offscreen:    ebiten.NewImage(width, height),
		offscreenPix: make([]byte, width*height*4),
	}

	g.calculateSet(-0.75, 0.25, 2)
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.offscreen, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (gm *Game) calculateSet(centerX, centerY, size float64) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			x := float64(w)*size/width - size/2 + centerX
			y := (height-float64(h))*size/height - size/2 + centerY
			c := complex(x, y)
			z := complex(0, 0)
			it := 0
			for ; it < maxIt; it++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					break
				}
			}
			r, g, b := color(it)
			p := 4 * (w + h*width)
			gm.offscreenPix[p] = r
			gm.offscreenPix[p+1] = g
			gm.offscreenPix[p+2] = b
			gm.offscreenPix[p+3] = 0xff
		}
	}

	gm.offscreen.WritePixels(gm.offscreenPix)
}

func color(it int) (r, g, b byte) {
	if it == maxIt {
		return 0xff, 0xff, 0xff
	}
	c := palette[it]
	return c, c, c
}
