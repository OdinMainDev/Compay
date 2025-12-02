package utils

import (
	"fmt"
)

type ProgressBar struct {
	total     int
	current   int
	barWidth  int
	fillChar  rune
	emptyChar rune
}

func (p *ProgressBar) Init(total int) {
	p.total = total
	p.current = 0
	p.barWidth = 50
	p.fillChar = '#'
	p.emptyChar = ' '
	p.render()
}

// Step увеличивает прогресс на 1 шаг
func (p *ProgressBar) Step() {
	if p.current < p.total {
		p.current++
		p.render()
	}
}

// Finish завершает прогресс и отображает финальный вид
func (p *ProgressBar) Finish() {
	p.current = p.total
	p.render()
	fmt.Println() // перенос строки
}

// render отрисовывает текущий прогресс
func (p *ProgressBar) render() {
	progress := p.current * p.barWidth / p.total
	fmt.Printf("\r[%s%s] %3d%%",
		string(repeat(p.fillChar, progress)),
		string(repeat(p.emptyChar, p.barWidth-progress)),
		p.current*100/p.total)
}

// repeat — замена strings.Repeat без импорта
func repeat(char rune, count int) []rune {
	s := make([]rune, count)
	for i := range s {
		s[i] = char
	}
	return s
}

// main демонстрирует использование
func main() {
	var bar ProgressBar
	bar.Init(100)

	for i := 0; i < 100; i++ {
		// Здесь могла бы быть полезная работа
		bar.Step()
		// Имитация задержки
		// time.Sleep(50 * time.Millisecond) // можно раскомментировать
	}

	bar.Finish()
}
