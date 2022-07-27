package gopdf

import (
	"fmt"
	"io"
)

type cacheContentRaw struct {
	data      []byte
	transform [6]float64
}

func (c *cacheContentRaw) write(w io.Writer, protection *PDFProtection) error {
	_, err := fmt.Fprintf(w, "q %.2f %.2f %.2f %.2f %.2f %.2f cm\n",
		c.transform[0], c.transform[1], c.transform[2],
		c.transform[3], c.transform[4], c.transform[5])
	if err != nil {
		return err
	}

	_, err = w.Write(c.data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, "\nQ\n")
	if err != nil {
		return err
	}

	return nil
}
