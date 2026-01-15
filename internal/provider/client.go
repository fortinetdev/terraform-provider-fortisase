package provider

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/auth"
	forticlient "github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"golang.org/x/time/rate"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
}

// FortiClient contains the basic FortiSASE SDK connection information to FortiSASE
// It can be used to as a client of FortiSASE for the plugin
// Now FortiClient contains two kinds of clients:
// Client is for FortiGate
type FortiClient struct {
	// to sdk client
	Client *forticlient.FortiSDKClient
	// to lock the resource
	ResourceLocks map[string]*sync.Mutex
	// mutex to protect ResourceLocks map from concurrent access
	resourceLocksMutex sync.RWMutex
}

func (f *FortiClient) GetResourceLock(name string) *sync.Mutex {
	// First, try to read with read lock (fast path)
	f.resourceLocksMutex.RLock()
	if lock, ok := f.ResourceLocks[name]; ok {
		f.resourceLocksMutex.RUnlock()
		return lock
	}
	f.resourceLocksMutex.RUnlock()

	// Lock for writing (slow path)
	f.resourceLocksMutex.Lock()
	defer f.resourceLocksMutex.Unlock()

	// Double-check after acquiring write lock (another goroutine might have created it)
	if lock, ok := f.ResourceLocks[name]; ok {
		return lock
	}

	// Create new mutex if it doesn't exist
	f.ResourceLocks[name] = &sync.Mutex{}
	return f.ResourceLocks[name]
}

type RateLimitedTransport struct {
	Transport http.RoundTripper
	Limiter   *rate.Limiter
}

func (r *RateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	err := r.Limiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	return r.Transport.RoundTrip(req)
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	err := createFortiSASEClient(&fClient, c)
	if err != nil {
		return nil, fmt.Errorf("Error create fortisase client: %v", err)
	}

	return &fClient, nil
}

func createFortiSASEClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Username, c.Password, c.AccessToken, c.RefreshToken)

	if auth.Username == "" {
		_, err := auth.GetEnvUsername()
		if err != nil {
			return fmt.Errorf("Error reading Username")
		}
	}

	if auth.Password == "" {
		_, err := auth.GetEnvPassword()
		if err != nil {
			return fmt.Errorf("Error reading Password")
		}
	}

	if auth.AccessToken == "" {
		_, err := auth.GetEnvAccessToken()
		if err != nil {
			return fmt.Errorf("Error reading AccessToken")
		}
	}

	if auth.RefreshToken == "" {
		_, err := auth.GetEnvRefreshToken()
		if err != nil {
			return fmt.Errorf("Error reading RefreshToken")
		}
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}
	limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 1)
	rateLimitedTransport := &RateLimitedTransport{
		Transport: tr,
		Limiter:   limiter,
	}
	client := &http.Client{
		Transport: rateLimitedTransport,
		Timeout:   time.Second * 250,
	}

	fc, err := forticlient.NewClient(auth, client)

	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	err = fc.CheckUP()
	if err != nil {
		return err
	}

	fClient.Client = fc
	// initialize the resource locks
	fClient.ResourceLocks = make(map[string]*sync.Mutex)
	return nil
}
