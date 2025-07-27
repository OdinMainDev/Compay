package cperrors

import "fmt"

var ErrorFileBadFileSize = fmt.Errorf("%s", "Error5.1")
var ErrorFileBadFileFormat = fmt.Errorf("%s", "Error5.2")
var ErrorFileClamavError = fmt.Errorf("%s", "Error5.3")
var ErrorFileInternalError = fmt.Errorf("%s", "Error5.4")
