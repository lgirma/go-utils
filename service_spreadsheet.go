package utils

type SpreadsheetService interface {
	OpenFile(fileName string) (SpreadsheetInstance, error)
}

type SpreadsheetInstance interface {
	GetSheets() []string
	GetRowsIterator(sheetName string) (SpreadsheetRowsIterator, error)
	GetSheetCount() int
}

type SpreadsheetRowsIterator interface {
	GetCurrentRow() ([]string, error)
	GetNextRow() bool
	Close() error
}