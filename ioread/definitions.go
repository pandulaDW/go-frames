package ioread

//commonOptions describes the read options common to all formats
type commonOptions struct {
	Path string
}

//CsvOptions describes the read options specific to only csv format
type CsvOptions struct {
	commonOptions
	Delimiter string
}
