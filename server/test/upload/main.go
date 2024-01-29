package main

import (
	"fmt"
	UploadFile "server/infrastructure/action"
	"server/infrastructure/logger"
)

func main() {
	// Search()
	action, err := UploadFile.NewFileClient()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	base64Image := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAzdJREFUeNrsWo1xrCAQNqmAEuwgdhBKsARTwbsS7IASLMF5FXCpgOvAvApIB8Zk8I1DFkRYEO/cmZ1JMhH359tvF6QoTjnllJ2knZRP2k9KH835btJR04cJwgVw/lvlpNW9O99oTnPl+Py7mJQ8ivOzszXw9/LenGcA3EtLcOS9cEKpMrp0bjDUegPwQnv0NicNsDdJbXjmUGioVZb1bDLH5ysANaNqnVlzA1WsPgKQp54IGg2BILnVeW/o622gsaUhqFLNE1m0NpkArtRQVru2TGic5ZENaoBASMU7yYQAJJXSCALMFWOqkoCc5zuREgXKr0vtPNuZg6BBq0vlfJNJF0piW5ep87YgoE2PdWawdw3CgMFNRCManvkYXmn2BieLaa3uCIcV+omT9ylTmXnd24RjoJZp9XQkoVryvKZTedDsQyjYzAW1VvtHlDoEwSzVeBlZVsnw2dJOZvmbqIfH6DBXg0+bAnBL4DyPNGO8a13NCzqxRUScMpu1Un7OoE5v2hBDEdf+WENADgF402r1T27seZjBBbMEPn3JI4Ctb1pAsA5MZvm3JQC3tfYRQZbt9hVpzRcDH2wKwGuiAFwjoI76tvOgMRJh94kxflch65Fin+srmOTbhY7zXZH+JAgrACgJpDugACsAaGcZyz21SDgLCKTaDz7L0BdrE+3feQD0BXbS2kSl0CJsiroC6UDUtmOLdZ9PBMKWxURrWcS91KivTwKd77H31NDlBMxP4r2n8aSA7ynMPBJUrtTg+AhEuwwMsA/HRLWPOSyso4F5vKgBsubiOPewr3aFlMvi3PJ/XJ3sVCstT3/e9jGzUmsKwzuFYyDYlh4K6bCIpCsEuaa2tXvF3q362SUZzUZ7Ohcist3PqzwgGFuXhNf6IOHiGOUWyeBBBbRzzNq4AeLE0c56CX254jxBzDo0Q5ACvizto2LFn1+bpLVoXTaQjItxZGXvIROWzg+yhwQvkhtGU+LARyNiQqJH3PfaLE2UnP/1JxAXHVTGMQ43aaSuAyamUoYLTzixiNvlUvFRSDD6Ban/yJND9EtLJj+UXov0QlXCiOXo/vvr8Kc6Et/DxlNOyV2+BBgAhmyMeZMFJvYAAAAASUVORK5CYII="
	fileName := "test.png"
	url, err := action.PostFile(&base64Image, fileName)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	fmt.Println(*url)
}
