package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	maxInt := int(^uint(0) >> 1)
	for i := range data {
		data[i] = rand.Intn(maxInt) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkMaxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var workers sync.WaitGroup
	workers.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// Include elements left over when the length is not divisible by CHUNKS.
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer workers.Done()
			if len(chunk) == 0 {
				// An existing value keeps empty chunks from overriding negative maxima.
				chunkMaxima[index] = data[0]
				return
			}
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	workers.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println()

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	fmt.Println()
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
