package main

// Image is
type Image struct {
	name string
}

var icons = make(map[string]Image)

func loadImage(name string) Image {
	return Image{name: name}
}

// Icon is
func Icon(name string) Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadImage(name)
		icons[name] = icon
	}
	return icon
}

func main() {
	go func() {
		Icon("skaji")
	}()
	go func() {
		Icon("skaji")
	}()
}
