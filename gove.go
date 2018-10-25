package main

import (
	"errors"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

type SDL2Context struct {
	window       *sdl.Window
	font         *ttf.Font
	isSdl2Init   bool
	isTTF2Init   bool
	isWindowOk   bool
	isFontOk     bool
	isRendererOk bool
}

func prg_init() (SDL2Context, error) {
	ctx := SDL2Context{}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Could not initialize SDL2: %s\n", err.Error())
		return SDL2Context{}, errors.New("Counld not initialize SDL2")
	}
	ctx.isSdl2Init = true

	if err := ttf.Init(); err != nil {
		fmt.Printf("Could not initialize SDL2_TTF: %s\n", err.Error())
		return SDL2Context{}, errors.New("Counld not initialize SDL2_TTF")
	}
	ctx.isTTF2Init = true

	window, err := sdl.CreateWindow("Gove - 0.1.0", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Could not create the window: %s\n", err.Error())
		return SDL2Context{}, errors.New("Could not create the Window")
	}
	ctx.isWindowOk = true
	ctx.window = window

	font, err := ttf.OpenFont("./Data/Fonts/Terminus-ja.ttf", 16)
	if err != nil {
		fmt.Printf("Could not load the default font: %s\n", err.Error())
		return SDL2Context{}, errors.New("Could not load the default font")
	}
	ctx.isFontOk = true
	ctx.font = font

	return ctx, nil
}

func prg_exit(ctx SDL2Context) {
	if ctx.isWindowOk {
		ctx.window.Destroy()
	}
	if ctx.isFontOk {
		ctx.font.Close()
	}
	if ctx.isTTF2Init {
		ttf.Quit()
	}
	if ctx.isSdl2Init {
		sdl.Quit()
	}
}

func prg_main() int {
	ctx, err := prg_init()
	defer prg_exit(ctx)
	if err != nil {
		fmt.Printf("FATAL: %s\n", err.Error())
		return 1
	}

	// Get the window surface
	surface, err := ctx.window.GetSurface()
	if err != nil {
		fmt.Printf("FATAL: Could not get Window Surface: %s\n", err.Error())
		return 1
	}

	running := true
	for running {
		messageSolid, err := ctx.font.RenderUTF8Solid("Hello, World!", sdl.Color{0, 255, 255, 255})
		if err != nil {
			fmt.Printf("FATAL: Could not create font: %s\n", err.Error())
			return 1
		}
		err = messageSolid.Blit(nil, surface, nil)
		if err != nil {
			fmt.Printf("FATAL: messageSolid.Blit() failed: %s\n", err.Error())
			return 1
		}

		ctx.window.UpdateSurface()

		for evt := sdl.PollEvent(); evt != nil; evt = sdl.PollEvent() {
			switch evt.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}

	return 0
}

func main() {
	os.Exit(prg_main())
}
