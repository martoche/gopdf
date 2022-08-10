package gopdf

import (
	"compress/zlib"
	"fmt"
	"io"
)

type RawObj struct {
	Data        []byte
	BoundingBox [4]float64
	getRoot     func() *GoPdf
}

func (i *RawObj) init(getRoot func() *GoPdf) {
	i.getRoot = getRoot
}

func (i *RawObj) getType() string {
	return "XObject"
}

func (i *RawObj) write(w io.Writer, objId int) error {
	buf := GetBuffer()
	defer PutBuffer(buf)

	compressed, err := zlib.NewWriterLevel(buf, i.getRoot().compressLevel)
	if err != nil {
		return err
	}

	_, err = compressed.Write(i.Data)
	if err != nil {
		return err
	}

	err = compressed.Close()
	if err != nil {
		return err
	}

	content := "<<\n"
	content += fmt.Sprintf("\t/Type /%s\n", i.getType())
	content += "\t/Subtype /Form\n"
	content += "\t/Matrix [1 0 0 1 0 0]\n"
	content += fmt.Sprintf("\t/BBox [%.3f %.3f %.3f %.3f]\n", i.BoundingBox[0], i.BoundingBox[1], i.BoundingBox[2], i.BoundingBox[3])
	content += fmt.Sprintf("\t/Length %d\n", buf.Len())
	content += "\t/Filter /FlateDecode\n" // must come after /Length
	content += ">>\n"
	content += "stream\n"
	_, err = io.WriteString(w, content)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, "endstream\n")
	if err != nil {
		return err
	}

	return nil
}
