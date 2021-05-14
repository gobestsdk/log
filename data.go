package log

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

type Fields map[string]interface{}

func (f Fields) Marshal() (bs []byte, err error) {
	var header = make([]string, 0)
	var data = [][]string{}

	bf := bytes.NewBuffer(bs)
	table := tablewriter.NewWriter(bf)

	data_0 := []string{}
	for k, v := range f {
		header = append(header, k)
		data_0 = append(data_0, fmt.Sprintf("%v", v))
	}
	data = append(data, data_0)
	table.SetHeader(header)
	table.SetAutoFormatHeaders(false)
	for i := 0; i < len(data); i++ {
		table.Append(data[i])
	}
	table.Render()
	bs = bf.Bytes()
	return
}

type Struct []KV
type KV struct {
	K string
	V interface{}
}
