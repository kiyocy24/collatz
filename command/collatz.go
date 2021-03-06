package command

import (
	"collatz/collatz"
	"collatz/helper"
	"fmt"
	"log"
	"strings"
)

const (
	Chunk = 100000
)

type Options struct {
	output string
	end    uint64
}

type option func(*Options)

func Output(v string) option {
	return func(opt *Options) {
		opt.output = v
	}
}

func End(v uint64) option {
	return func(ops *Options) {
		ops.end = v
	}
}

func Collatz(n uint64, opts ...option) error {
	op := &Options{
		output: "",
		end:    0,
	}

	for _, opt := range opts {
		opt(op)
	}
	if op.end == 0 {
		op.end = n
	}

	var writer *helper.Writer
	var err error
	if op.output != "" {
		writer, err = helper.NewWriter(op.output)
		if err != nil {
			return err
		}
		defer writer.Close()
	}

	for start := n; start <= op.end; start += Chunk {
		end := start + Chunk - 1
		if end > op.end {
			end = op.end
		}
		result := collatz.Collatzs(start, end)
		var resultStr [][]string
		for _, v := range result {
			resultStr = append(resultStr, helper.SliceUint64ToString(v))
		}

		if writer != nil {
			err = writer.Write(resultStr)
			if err != nil {
				return err
			}
			log.Printf("%d / %d", end, op.end)
		} else {
			for _, v := range resultStr {
				fmt.Println(strings.Join(v, ","))
			}
		}
	}

	return nil
}
