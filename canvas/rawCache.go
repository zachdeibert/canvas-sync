package canvas

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

func (c *Canvas) getRawCacheLocation(u url.URL, accept string) string {
	if u.Host == fmt.Sprintf("%s.instructure.com", c.subdomain) {
		p := u.EscapedPath()
		ext := mimeToExt(accept)
		if len(u.RawQuery) > 0 {
			return path.Join(c.RawSaveFolder, fmt.Sprintf("%s%s", p, ext), fmt.Sprintf("%s%s", url.PathEscape(u.RawQuery), ext))
		}
		return path.Join(c.RawSaveFolder, fmt.Sprintf("%s%s", p, ext), fmt.Sprintf("default%s", ext))
	}
	return ""
}

func (c *Canvas) saveRequest(u url.URL, accept string, body []byte, res *http.Response) error {
	if fname := c.getRawCacheLocation(u, accept); len(fname) > 0 {
		dirname := path.Dir(fname)
		if err := os.MkdirAll(dirname, 0755); err != nil {
			return err
		}
		return ioutil.WriteFile(fname, body, 0644)
	}
	return nil
}
