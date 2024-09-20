package zap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type PrettyJSONEncoder struct {
	zapcore.Encoder
}

func (enc *PrettyJSONEncoder) Clone() zapcore.Encoder { //nolint:ireturn // тут сигнатуру не поменять
	return enc.Encoder.Clone()
}

func (enc *PrettyJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	encodedEntry, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, fmt.Errorf("error encode entry: %w", err)
	}

	if entry.Stack == "" {
		return encodedEntry, nil
	}

	stack := strings.ReplaceAll(entry.Stack, "\n\t", " ")

	parts := strings.Split(stack, "\n")

	var stackRecords []StackRecord

	for index := range parts {
		part := parts[index]

		recParts := strings.Split(part, " ")

		if len(recParts) > 1 {
			stackRecords = append(stackRecords, StackRecord{
				Name: recParts[0],
				Path: recParts[1],
			})
		}
	}

	jsonStack, err := json.Marshal(stackRecords)
	if err != nil {
		return nil, fmt.Errorf("error marshal stack trace: %w", err)
	}

	newb := buffer.NewPool().Get()

	tempBytes := bytes.ReplaceAll(encodedEntry.Bytes(), []byte("\\n"), []byte("\n"))
	tempBytes = bytes.ReplaceAll(tempBytes, []byte("\\t"), []byte("\t"))
	newb.Write(bytes.ReplaceAll(tempBytes, []byte(`"`+entry.Stack+`"`), jsonStack))

	return newb, nil
}

type StackRecord struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
