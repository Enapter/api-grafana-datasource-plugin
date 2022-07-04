package queryhandler

import "errors"

var (
	errEmptyQueryText                = errors.New("empty query text")
	ErrUnsupportedTimeseriesDataType = errors.New("unsupported timeseries data type")
)
