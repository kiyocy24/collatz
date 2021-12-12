package command

import (
	"collatz/collatz"
	"collatz/helper"
	"encoding/csv"
	"fmt"
	"os"
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

	var f *os.File
	if op.output != "" {
		f, err := os.OpenFile(op.output, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	for start := n; start <= op.end; start += Chunk {
		end := start + Chunk
		if end > op.end {
			end = op.end
		}
		result := collatz.Collatzs(start, end)
		var resultStr [][]string
		for _, v := range result {
			resultStr = append(resultStr, helper.SliceUint64ToString(v))
		}

		if op.output == "" {
			for _, v := range resultStr {
				fmt.Println(strings.Join(v, ","))
			}
		} else {
			err := output(f, resultStr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func output(f *os.File, records [][]string) error {
	w := csv.NewWriter(f)
	err := w.WriteAll(records)
	if err != nil {
		return err
	}
	w.Flush()

	err = w.Error()
	if err != nil {
		return err
	}

	return nil
}
