package gopdf

import (
	"fmt"
	"io"
)

type cacheContentRaw struct {
	index     int
	transform [6]float64
}

func (c *cacheContentRaw) write(w io.Writer, protection *PDFProtection) error {
	// the name of the object is "/Ixx" where xx == index+1
	// (see procset_obj.go)

	_, err := fmt.Fprintf(w, "q %.2f %.2f %.2f %.2f %.2f %.2f cm /I%d Do Q\n",
		c.transform[0], c.transform[1], c.transform[2],
		c.transform[3], c.transform[4], c.transform[5],
		c.index+1)
	if err != nil {
		return err
	}

	return nil
}
