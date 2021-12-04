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

	var result [][]uint64
	if op.end == 0 {
		result = [][]uint64{collatz.Collatz(n)}
	} else {
		result = collatz.Collatzs(n, op.end)
	}

	var resultStr [][]string
	for _, v := range result {
		resultStr = append(resultStr, helper.SliceUint64ToString(v))
	}

	if op.output == "" {
		for _, v := range resultStr {
			fmt.Println(strings.Join(v, ","))
		}
	} else {
		f, err := os.Create(op.output)
		if err != nil {
			return err
		}
		defer f.Close()

		w := csv.NewWriter(f)
		err = w.WriteAll(resultStr)
		if err != nil {
			return err
		}
		w.Flush()

		err = w.Error()
		if err != nil {
			return err
		}
	}

	return nil
}
