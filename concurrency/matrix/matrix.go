package matrix

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

func GenerateRandomMatrixSequential(row, column, min, max int) [][]int {
	result := make([][]int, 0, column)
	for columnIndex := 0; columnIndex < column; columnIndex++ {
		result = append(result, make([]int, row))
		for rowIndex := 0; rowIndex < row; rowIndex++ {
			result[columnIndex][rowIndex] = rand.IntN(max-min) + min
		}
	}

	return result
}

func GenerateRandomMatrixConcurrent1(row, column, min, max int) [][]int {
	result := make([][]int, 0, column)
	var wg sync.WaitGroup

	for columnIndex := 0; columnIndex < column; columnIndex++ {
		result = append(result, make([]int, row))
		wg.Add(1)
		go func(columnIndex int) {
			defer wg.Done()
			for rowIndex := 0; rowIndex < row; rowIndex++ {
				result[columnIndex][rowIndex] = rand.IntN(max-min) + min
			}
		}(columnIndex)
	}

	wg.Wait()

	return result
}

func GenerateRandomMatrixConcurrent2(row, column, min, max int) [][]int {
	result := make([][]int, 0, column)
	var wg sync.WaitGroup

	routineCount := runtime.NumCPU()
	jobs := make(chan int, column) // each job represent a column

	worker := func() {
		for job := range jobs {
			for rowIndex := 0; rowIndex < row; rowIndex++ {
				result[job][rowIndex] = rand.IntN(max-min) + min
			}
		}
		wg.Done()
	}

	for i := 0; i < routineCount; i++ {
		wg.Add(1)
		go worker()
	}

	for columnIndex := 0; columnIndex < column; columnIndex++ {
		result = append(result, make([]int, row))
		jobs <- columnIndex
	}
	close(jobs)

	wg.Wait()
	return result
}

func GenerateRandomMatrixConcurrent3(row, column, min, max int) [][]int {
	result := make([][]int, 0, column)
	var wg sync.WaitGroup
	locker := semaphore.NewWeighted(int64(runtime.NumCPU()))

	for columnIndex := 0; columnIndex < column; columnIndex++ {
		result = append(result, make([]int, row))
		wg.Add(1)
		locker.Acquire(context.TODO(), 1)
		go func(columnIndex int) {
			defer wg.Done()
			defer locker.Release(1)
			for rowIndex := 0; rowIndex < row; rowIndex++ {
				result[columnIndex][rowIndex] = rand.IntN(max-min) + min
			}
		}(columnIndex)
	}

	wg.Wait()
	return result
}

func GenerateRandomMatrixConcurrent4(row, column, min, max int) [][]int {
	var (
		routineCount = runtime.NumCPU()
		wg           sync.WaitGroup
		bunch        = column / routineCount
		matrix       = make([][]int, column)
	)

	if bunch == 0 {
		bunch = 1
	}

	for start := 0; start < column; start += bunch {
		end := start + bunch
		if end >= column {
			end = column - 1
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			defer func() {
				_ = recover()
			}()

			for columnIndex := start; columnIndex <= end; columnIndex++ {
				matrix[columnIndex] = make([]int, row)

				for rowIndex := 0; rowIndex < row; rowIndex++ {
					matrix[columnIndex][rowIndex] = rand.IntN(max-min) + min
				}
			}
		}(start, end)
	}

	wg.Wait()
	return matrix
}

func Addition(matrix1, matrix2 [][]int) ([][]int, error) {
	column1Count := len(matrix1)
	column2Count := len(matrix2)
	if column1Count != column2Count {
		return nil, errors.New("invalid matrixes")
	}

	result := make([][]int, 0, column1Count)

	for columnIndex := 0; columnIndex < column1Count; columnIndex++ {
		rowMatrix1 := matrix1[columnIndex]
		rowMatrix2 := matrix2[columnIndex]
		row1Count := len(rowMatrix1)
		row2Count := len(rowMatrix2)
		if row1Count != row2Count {
			return nil, errors.New("invalid matrixes")
		}

		addition := make([]int, 0, len(rowMatrix1))
		for rowIndex := 0; rowIndex < len(rowMatrix1); rowIndex++ {
			addition = append(addition, rowMatrix1[rowIndex]+rowMatrix2[rowIndex])
		}

		result = append(result, addition)
	}

	return result, nil
}

func MultiplicationSequential(matrix1, matrix2 [][]int) ([][]int, error) {
	if matrix1 == nil || matrix2 == nil {
		return nil, errors.New("invalid")
	}

	column1Count := len(matrix1)
	row1Count := len(matrix1[0])

	column2Count := len(matrix2)
	row2Count := len(matrix2[0])

	if row1Count != column2Count {
		return nil, errors.New("invalid")
	}

	result := make([][]int, 0, column1Count)
	for columnIndex := 0; columnIndex < column1Count; columnIndex++ {
		result = append(result, make([]int, row2Count))
	}

	for columnIndex := 0; columnIndex < column1Count; columnIndex++ {
		for rowIndex := 0; rowIndex < row2Count; rowIndex++ {
			resultNumber := 0
			for index := 0; index < column1Count; index++ {
				resultNumber += matrix1[rowIndex][index] * matrix2[index][columnIndex]
			}
			result[columnIndex][rowIndex] = resultNumber
		}
	}

	return result, nil
}

func MultiplicateConcurrent() {}

func Print(matrix [][]int) {
	for indexColumn := 0; indexColumn < len(matrix); indexColumn++ {
		row := matrix[indexColumn]
		for indexRow := 0; indexRow < len(row); indexRow++ {
			fmt.Printf("%d ", row[indexRow])
		}
		fmt.Println()
	}
}
