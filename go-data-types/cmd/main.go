package main

import (
	aggregatedatatypes "go-data-types/pkg/aggregate-datatypes"
	basicdatatypes "go-data-types/pkg/basic-datatypes"
)

func main() {
	basicdatatypes.DisplayNumbers()
	basicdatatypes.DisplayStrings()
	basicdatatypes.DisplayBoolean()
	aggregatedatatypes.DisplayArrays()
	aggregatedatatypes.DisplayStructs()
}
