package sources

import "github.com/seekr-osint/seekr/api/functions"

func (sources Sources) Parse() (Sources, error) {
	newsources, err := functions.FullParseMapRet(sources, "url")
	return newsources, err
}

func (source Source) Parse() (Source, error) {
	return source, nil
}
