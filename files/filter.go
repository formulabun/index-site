package files

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"go.formulabun.club/metadatadb"
)

func filterFiles(search string, files []metadatadb.File) []metadatadb.File {
	if search == "" {
		return files
	}

	searchTerms := []string{}
	for _, s := range strings.Split(strings.ToLower(strings.Trim(search, " ")), " ") {
		if s != "" {
			searchTerms = append(searchTerms, s)
		}
	}

	i := 0
	res := make([]metadatadb.File, len(files))
	for _, f := range files {
		for _, term := range searchTerms {
			if fuzzy.Match(term, strings.ToLower(f.Filename)) {
				res[i] = f
				i += 1
				break
			}
		}
	}

	return res[:i]
}
