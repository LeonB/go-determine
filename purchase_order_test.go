package determine_test

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/omniboost/go-determine"
)

func TestPurchaseOrder(t *testing.T) {
	files, err := filepath.Glob("po_*.xml")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		s := determine.BPACK{}
		err = xml.Unmarshal(b, &s)
		if err != nil {
			log.Fatal(err)
		}

		b, err = xml.MarshalIndent(s, "", "  ")
		log.Println(string(b))
	}
}
