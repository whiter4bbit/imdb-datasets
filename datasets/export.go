package datasets

import (
	"os"
)

func Export(imdbPath, outputPath string, datasetsToExport []string) error {
	if err := ExportTitlesDataset(imdbPath+string(os.PathSeparator)+"movies.list.gz", outputPath+string(os.PathSeparator)+"titles"); err != nil {
		return err
	}
	return nil
}
