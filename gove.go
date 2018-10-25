package main

import (
	"errors"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type SDL2Context struct {
	window         *sdl.Window
	renderer       *sdl.Renderer
	is_sdl2_init   bool
	is_window_ok   bool
	is_renderer_ok bool
}

func prg_init() (SDL2Context, error) {
	ctx := SDL2Context{}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Could not initialize SDL2: %s\n", err.Error())
		return SDL2Context{}, errors.New("Counld not initialize SDL2")
	}
	ctx.is_sdl2_init = true

	window, err := sdl.CreateWindow("Gove - 0.1.0", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Could not create the window: %s\n", err.Error())
		return SDL2Context{}, errors.New("Could not create the Window")
	}
	ctx.is_window_ok = true
	ctx.window = window

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		fmt.Printf("Could not initialize the Renderer. :%s\n", err.Error())
		return SDL2Context{}, errors.New("Could not initialize the Renderer")
	}
	ctx.is_renderer_ok = true
	ctx.renderer = renderer

	return ctx, nil
}

func prg_exit(ctx SDL2Context) {
	if ctx.is_renderer_ok {
		ctx.renderer.Destroy()
	}
	if ctx.is_window_ok {
		ctx.window.Destroy()
	}
	if ctx.is_sdl2_init {
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
