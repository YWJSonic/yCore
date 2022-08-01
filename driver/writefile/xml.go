package writefile

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func ProtocolXml(path string, payload interface{}, mode os.FileMode) error {

	data, err := xml.Marshal(payload)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, data, mode); err != nil {
		return err
	}

	return nil
}
