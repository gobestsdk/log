package log

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

type LogRaw interface {
	TableMarshal() ([]byte, error)
	JsonMarshal() ([]byte, error)
}

type Fields map[string]interface{}

func (f Fields) TableMarshal() (bs []byte, err error) {
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

func (f Struct) TableMarshal() (bs []byte, err error) {
	var header = make([]string, 0)
	var data = [][]string{}

	bf := bytes.NewBuffer(bs)
	table := tablewriter.NewWriter(bf)

	data_0 := []string{}
	for i := 0; i < len(f); i++ {
		k := f[i].K
		v := f[i].V

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

type KV struct {
	K string
	V interface{}
}
