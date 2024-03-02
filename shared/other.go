package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Check if the given path exists
// USED INTERNALLY
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Fetch JSON to a struct
// USED INTERNALLY
func FetchJson(url string, v any) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code not 200 (got %s)", strconv.Itoa(resp.StatusCode))
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// Download a file to the given filepath
// USED INTERNALLY
func Download(url string, filepath string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Dbg(exp interface{}) string {
	return fmt.Sprintf("%+v", exp)
}
