package utils

import (
	"context"
)

// Converter is the interface that wraps basic Convert method
type Converter interface {
	Convert() (map[int]string, error)
}

// Convert convert chinese from simplicity to traditional or simplicity to traditional
func Convert(c Converter) (map[int]string, error) {
	return c.Convert()
}

const (
	// batchWorker defines worker number
	batchWorker = 50
	// batchQueue defines queue number
	batchQueue = 100
)

// ConvertFunc defines convert func when doing batch
type ConvertFunc func(in string) (string, error)

type batchResult struct {
	batch int
	text  string
	err   error
}

// Batch defines bytes batch processing
func Batch(bytes [][]byte, processingFunc ConvertFunc) (map[int]string, error) {
	byteLen := len(bytes)
	if byteLen == 0 {
		return map[int]string{}, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	worker := make(chan struct{}, batchWorker)
	queue := make(chan *batchResult, batchQueue)

	done := make(chan interface{})
	results := make(map[int]string, byteLen)
	byteLen--
	var err error
	go func() {
		var count int
	LOOP:
		for q := range queue {
			results[q.batch] = q.text
			if err != nil || count == byteLen {
				err = q.err
				cancel()
				close(done)
				break LOOP
			}
			count++
		}
	}()

LOOP:
	for i, bs := range bytes {
		worker <- struct{}{}
		select {
		case <-ctx.Done():
			break LOOP
		default:
			go func(ctx context.Context, batch int, bs []byte, process ConvertFunc) {
				select {
				case <-ctx.Done():
					// if error occur do nothing
				default:
					queue <- func(ctx context.Context, batch int, bs []byte, process ConvertFunc) *batchResult {
						text, err := process(string(bs))
						return &batchResult{
							batch: batch,
							text:  text,
							err:   err,
						}
					}(ctx, batch, bs, process)
				}
				<-worker
			}(ctx, i, bs, processingFunc)
		}
	}
	<-done
	return results, err
}
