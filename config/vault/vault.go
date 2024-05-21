package vault

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

const NamespaceCubbyhole = "cubbyhole"

var _ Vault = Client{}

// Client struct for managing vault requests.
type Client struct {
	client    *api.Client
	namespace string
}

// NewClient constructor.
func NewClient(namespace string) (*Client, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, ClientError{What: err.Error()}
	}

	if errs := validateClient(client); len(errs) > 0 {
		return nil, ValidationError{What: errs}
	}

	return &Client{
		client:    client,
		namespace: namespace,
	}, nil
}

// Pull data from vault via vault client.
func (c Client) Pull(key string) ([]byte, error) {
	data, err := c.read(key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PullRaw pull raw map data from vault via vault client.
func (c Client) PullRaw(key string) (map[string]any, error) {
	data, err := c.readRaw(key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Download data from vault via vault client to file via path.
func (c Client) Download(key, filepath string) error {
	data, err := c.read(key)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

// Upload data from file to the vault.
func (c Client) Upload(key, filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return c.Push(key, string(data))
}

// Push data to the vault via vault client.
func (c Client) Push(key string, value any) error {
	data := map[string]any{key: value}
	if _, err := c.client.Logical().Write(c.requestPath(key), data); err != nil {
		return WriteError{What: err.Error()}
	}

	return nil
}

// PushRaw data to the vault via vault client.
func (c Client) PushRaw(key string, values map[string]any) error {
	if _, err := c.client.Logical().Write(c.requestPath(key), values); err != nil {
		return WriteError{What: err.Error()}
	}

	return nil
}

// validateClient - validate client, checks token and address.
func validateClient(c *api.Client) []string {
	var (
		errorList []string
		secret    = os.Getenv("VAULT_SECRET")
	)

	if c.Address() == "" {
		errorList = append(errorList, "address must be in environment variable, like: VAULT_ADDR='...'")
	}

	if c.Token() == "" {
		if secret != "" {
			c.SetToken(secret)
			return errorList
		}

		errorList = append(errorList, "token must be in environment variable, like: VAULT_TOKEN='...'")
	}

	return errorList
}

func (c Client) read(key string) ([]byte, error) {
	secret, err := c.client.Logical().Read(c.requestPath(key))
	if err != nil {
		return nil, ReadError{What: err.Error()}
	}

	if secret == nil {
		return nil, ReadError{What: fmt.Sprintf("no data for key: '%s'", key)}
	}

	data, ok := secret.Data[key]
	if !ok {
		return nil, ReadError{What: fmt.Sprintf("no data for key: '%s'", key)}
	}

	value, ok := data.(string)
	if !ok {
		return nil, ReadError{What: fmt.Sprintf("convert secret to string")}
	}

	return []byte(value), nil
}

func (c Client) readRaw(key string) (map[string]any, error) {
	secret, err := c.client.Logical().Read(c.requestPath(key))
	if err != nil {
		return nil, ReadError{What: err.Error()}
	}

	return secret.Data, nil
}

func (c Client) requestPath(key string) string {
	return fmt.Sprintf("%s/%s", c.namespace, key)
}
