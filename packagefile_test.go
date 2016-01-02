package rpm

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadRPMFile(t *testing.T) {
	// get a directory full of rpms from RPM_DIR environment variable but allow
	// tests to suceed if no rpms are available (like on Travis CI).
	dir := os.Getenv("RPM_DIR")
	if dir != "" {

		// list RPM files
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatalf(err.Error())
		}

		valid := 0
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".rpm") {
				path := filepath.Join(dir, f.Name())
				rpm, err := OpenPackageFile(path)
				if err != nil {
					t.Errorf("Error loading RPM file %s: %s", f.Name(), err)
				} else {
					t.Logf("Loaded package: %v", rpm)
					valid++
				}
			}
		}

		if valid == 0 {
			t.Errorf("No RPM files found for testing with in %s", dir)
		} else {
			t.Logf("Validated %d RPM files", valid)
		}
	}
}
