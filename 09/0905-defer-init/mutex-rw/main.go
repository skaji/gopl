package main

import "sync"

// Image is
type Image struct {
	name string
}

var (
	mu    sync.RWMutex
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
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	mu.Lock()
	if icons == nil { // XXX 再びチェックが必要
		loadIcons()
	}
	mu.Unlock()
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
