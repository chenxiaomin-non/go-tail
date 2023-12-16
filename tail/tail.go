package tail

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Mode int

const (
	_ Mode = iota
	ReadFromHead
	ReadFromTail
)

type Observer struct {
	FilePath      string `yaml:"file_path"`
	CurrentOffset int64
	Mode          `yaml:"mode"`
	Interval      int `yaml:"interval"` // in millisecond
	Publisher     chan string
}

func NewObserver(filePath string) (*Observer, error) {
	// check file exists & readable
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return &Observer{
		FilePath:      filePath,
		CurrentOffset: 0,
		Mode:          ReadFromHead,
		Interval:      100,
	}, nil
}

func (o *Observer) SetReadFromTail() error {
	// get file size
	fp, err := os.Open(o.FilePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	stat, err := fp.Stat()
	if err != nil {
		return err
	}

	o.CurrentOffset = stat.Size()
	o.Mode = ReadFromTail
	return nil
}

func GetNewTailContent(fp *os.File, currentOffset int64) ([]byte, error) {
	// get file size
	stat, err := fp.Stat()
	if err != nil {
		return nil, err
	}

	// check file size
	if stat.Size() < currentOffset {
		return nil, nil
	}

	// read new content
	temp := make([]byte, stat.Size()-currentOffset)
	_, err = fp.ReadAt(temp, currentOffset)
	if err != io.EOF && err != nil {
		return nil, err
	}
	return temp, nil
}

// start observer, continuously read file content
// and publish to channel (which is a message queue
// for all observers, managed by ObsManager)
func (o *Observer) Start() error {
	go func() {
		fp, err := os.Open(o.FilePath)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		for {
			buf, err := GetNewTailContent(fp, o.CurrentOffset)
			if err != nil {
				panic(err)
			}

			len := len(buf)
			if len > 0 {
				o.Publisher <- bytes.NewBuffer(buf).String()
				fmt.Println("push new content to channel")
			}

			// update current offset
			o.CurrentOffset += int64(len)

			// sleep
			if o.Interval > 0 {
				time.Sleep(time.Millisecond * time.Duration(o.Interval))
			}
		}
	}()
	return nil
}
