package util

import (
	"testing"
)

func TestCache(t *testing.T) {
	// Создаем новый кэш
	cache := New()

	// Проверяем, что кэш пустой
	if len(cache.items) != 0 {
		t.Errorf("New cache is not empty")
	}

	// Добавляем элемент в кэш
	cache.Set("key", "value")

	// Проверяем, что элемент добавлен корректно
	if len(cache.items) != 1 {
		t.Errorf("Failed to set value in cache")
	}

	// Получаем элемент из кэша
	item, found := cache.Get("key")

	// Проверяем, что элемент найден
	if !found {
		t.Errorf("Item not found in cache")
	}

	// Проверяем, что значение элемента корректно
	if item != "value" {
		t.Errorf("Incorrect value retrieved from cache")
	}

	// Удаляем элемент из кэша
	err := cache.Delete("key")

	// Проверяем, что элемент успешно удален
	if err != nil {
		t.Errorf("Failed to delete item from cache")
	}

	// Проверяем, что кэш пустой после удаления элемента
	if len(cache.items) != 0 {
		t.Errorf("Cache is not empty after deleting item")
	}
}
