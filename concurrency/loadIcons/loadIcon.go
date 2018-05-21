package loadIcons

import (
	"image"
	"sync"
)

var mu sync.RWMutex // 监控icons共享变量
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func Icon(name string) image.Image {
	mu.RLock() // 通过读锁实现对变量的并发访问
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	mu.Lock() // 通过互斥量确保变量只初始化一次，可以使用sync.Once简化操作
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}
