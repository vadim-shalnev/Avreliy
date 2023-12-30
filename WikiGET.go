package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func WikiGET() string {
	resp, err := http.Get("https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D1%80%D1%82%D0%B0%D0%BB:%D0%A4%D0%B8%D0%BB%D0%BE%D1%81%D0%BE%D1%84%D0%B8%D1%8F")

	if err != nil {
		fmt.Println("Ошибка в GET запросе:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	// Извлечение и анализ HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при анализе HTML:", err)
		os.Exit(1)
	}

	// Извлечение заголовка статьи
	TargetClass := "ts-Цитата"
	text := ExtractTextByClass(doc, TargetClass)
	fmt.Print(TargetClass, ":", text)
	return text
}

func ExtractTextByClass(n *html.Node, targetClass string) string {
	if n.Type == html.ElementNode && HasClass(n, targetClass) {
		// Найден элемент с указанным классом
		return GetTextContent(n)
	}

	// Поиск в дочерних элементах
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := ExtractTextByClass(c, targetClass); result != "" {
			return result
		}
	}

	return ""
}

func HasClass(n *html.Node, targetClass string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, targetClass) {
			return true
		}
	}
	return false
}

// Получение текстового содержимого элемента
func GetTextContent(n *html.Node) string {
	var textContent string

	if n.Type == html.TextNode {
		// Текстовый узел
		textContent = n.Data
	} else {
		// Рекурсивный вызов для дочерних узлов
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			textContent += GetTextContent(c)
		}
	}

	return textContent
}
