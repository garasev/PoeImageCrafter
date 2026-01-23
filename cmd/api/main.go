package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/garasev/poe-item-generator/internal/generator"
	"github.com/garasev/poe-item-generator/internal/parser"
)

const (
	a = ``
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Ошибка при чтении из буфера обмена:", err)
		return
	}

	fmt.Println("Содержимое буфера обмена:")
	fmt.Println(text)
	item := parser.ParseText(text)

	fontBytes, err := os.ReadFile("../../src/fontin/FontinSans_Cyrillic_B_46b.otf")
	if err != nil {
		log.Fatal("Шрифт не найден:", err)
	}
	g, err := generator.NewGenerator(fontBytes)
	if err != nil {
		return
	}
	g.CreateItem(item)
	log.Println(item)
	// // Текст который нужно нарисовать
	// text := "Привет, мир! Это длинный текст для примера Привет, мир! Это длинный текст для примера"

	// // Загружаем шрифт
	// fontBytes, err := os.ReadFile("src/fontin/FontinSans_Cyrillic_R_46b.otf")
	// if err != nil {
	// 	log.Fatal("Шрифт не найден:", err)
	// }

	// fontFace, err := opentype.Parse(fontBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// face, err := opentype.NewFace(fontFace, &opentype.FaceOptions{
	// 	Size:    48, // Размер шрифта
	// 	DPI:     72,
	// 	Hinting: font.HintingFull,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 1. Создаем временный контекст для измерения текста
	// tempDC := gg.NewContext(1, 1)
	// tempDC.SetFontFace(face)

	// // 2. Измеряем размеры текста
	// textWidth, textHeight := tempDC.MeasureString(text)

	// // 3. Добавляем отступы
	// padding := 40.0
	// borderWidth := 5.0
	// totalPadding := padding*2 + borderWidth*2

	// // 4. Вычисляем итоговые размеры
	// imgWidth := int(textWidth + totalPadding)
	// imgHeight := int(float64(textHeight)*1.5 + totalPadding) // Увеличиваем высоту для лучшего вида

	// fmt.Printf("Текст: %s\n", text)
	// fmt.Printf("Размер текста: %.2f x %.2f\n", textWidth, textHeight)
	// fmt.Printf("Размер изображения: %d x %d\n", imgWidth, imgHeight)

	// // 5. Создаём основное изображение с вычисленными размерами
	// dc := gg.NewContext(imgWidth, imgHeight)
	// dc.SetFontFace(face)

	// // Фон
	// dc.SetColor(color.RGBA{240, 248, 255, 255}) // AliceBlue
	// dc.Clear()

	// // Рамка
	// dc.SetColor(color.RGBA{70, 130, 180, 255}) // SteelBlue
	// dc.SetLineWidth(borderWidth)
	// dc.DrawRectangle(borderWidth/2, borderWidth/2,
	// 	float64(imgWidth)-borderWidth,
	// 	float64(imgHeight)-borderWidth)
	// dc.Stroke()

	// // Текст (по центру)
	// textX := float64(imgWidth) / 2
	// textY := float64(imgHeight) / 2

	// dc.SetColor(color.RGBA{25, 25, 112, 255}) // MidnightBlue
	// dc.DrawStringAnchored(text, textX, textY, 0.5, 0.5)

	// // Сохраняем
	// dc.SavePNG("dynamic_size.png")
}
