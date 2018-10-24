package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Gove - 0.1.0", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true
	for running {
		for evt := sdl.PollEvent(); evt != nil; evt = sdl.PollEvent() {
			switch evt.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}
