package pico8

import (
	"io"
)

type writer interface {
	io.StringWriter
	io.Writer
}

func WriteBytesAsText(bytes []byte, writer writer) {

	for i := 0; i < len(bytes); i += 1 {
		switch bytes[i] {
		case 0:
			if i < len(bytes)-1 {
				nextCh := bytes[i+1]
				if nextCh >= 48 && nextCh <= 57 {
					writer.WriteString("\\000")
				} else {
					writer.WriteString("\\0")
				}
			} else {
				writer.WriteString("\\0")
			}
		case 10:
			writer.WriteString("\\n")
		case 13:
			writer.WriteString("\\r")
		case 92:
			writer.WriteString("\\\\")
		case 34:
			writer.WriteString("\\\"")
		default:
			writer.Write(bytes[i : i+1])
		}
	}

}
