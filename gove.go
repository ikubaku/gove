package main

import (
	"errors"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type SDL2Context struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	isSdl2Init   bool
	isWindowOk   bool
	isRendererOk bool
}

func prg_init() (SDL2Context, error) {
	ctx := SDL2Context{}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Could not initialize SDL2: %s\n", err.Error())
		return SDL2Context{}, errors.New("Counld not initialize SDL2")
	}
	ctx.isSdl2Init = true

	window, err := sdl.CreateWindow("Gove - 0.1.0", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Could not create the window: %s\n", err.Error())
		return SDL2Context{}, errors.New("Could not create the Window")
	}
	ctx.isWindowOk = true
	ctx.window = window

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		fmt.Printf("Could not initialize the Renderer. :%s\n", err.Error())
		return SDL2Context{}, errors.New("Could not initialize the Renderer")
	}
	ctx.isRendererOk = true
	ctx.renderer = renderer

	return ctx, nil
}

func prg_exit(ctx SDL2Context) {
	if ctx.isRendererOk {
		ctx.renderer.Destroy()
	}
	if ctx.isWindowOk {
		ctx.window.Destroy()
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

	// Clear screen
	ctx.renderer.SetDrawColor(0, 0, 0, 255)
	ctx.renderer.Clear()
	ctx.renderer.Present()

	running := true
	for running {
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
