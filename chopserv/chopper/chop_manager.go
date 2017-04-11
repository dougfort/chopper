package chopper

import "fmt"

// Chopper manages chopping and restoring URLs
type Chopper interface {

	// Chop 'shortens' a URL
	Chop(url string) (string, error)

	// Unchop restores a chopped URL
	UnChop(choppedURL string) (string, error)
}

type chopper struct {
	nextIndex int
	urls      map[string]string
}

// New returns an entity that implements the Chopper interface
func New() Chopper {
	return &chopper{urls: make(map[string]string)}
}

// Chop 'shortens' a URL
func (c *chopper) Chop(url string) (string, error) {
	// TODO: more sophisticated algorithm
	// TODO: persist
	key := fmt.Sprintf("%08d", c.nextIndex)
	c.urls[key] = url
	c.nextIndex++

	return key, nil
}

// Unchop restores a chopped URL
func (c *chopper) UnChop(choppedURL string) (string, error) {
	url, ok := c.urls[choppedURL]
	if !ok {
		return "", fmt.Errorf("unknown key '%s'", choppedURL)
	}
	delete(c.urls, choppedURL)
	return url, nil
}
