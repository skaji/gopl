package main

import "sync"

// Image is
type Image struct {
	name string
}

var (
	mu    sync.Mutex
	icons map[string]Image
)

func loadIcons() {
	icons = map[string]Image{
		"skaji": loadImage("skaji"),
		"you":   loadImage("you"),
		"he":    loadImage("he"),
	}
}

func loadImage(name string) Image {
	return Image{name: name}
}

// Icon is
func Icon(name string) Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

func main() {
	go func() {
		Icon("skaji")
	}()
	go func() {
		Icon("skaji")
	}()
}
