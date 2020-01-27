package k8s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/sanjid133/vault-kube/config"
	"github.com/sanjid133/vault-kube/util"
	"io"
)

type Client struct {
	vClient *api.Client
	jwt     string
	role    string
	path    string
}

func New(vc *api.Client, cfg *config.Config) (*Client, error) {
	jwt, err := util.ReadData(cfg.ServiceAccountTokenPath)
	if err != nil {
		return nil, err
	}
	return &Client{
		vClient: vc,
		jwt:     jwt,
		role:    cfg.Role,
		path:    "kubernetes",
	}, nil

}

func (c *Client) Login() (string, error) {
	path := fmt.Sprintf("/v1/auth/%s/login", c.path)
	req := c.vClient.NewRequest("POST", path)
	payload := map[string]interface{}{
		"jwt":  c.jwt,
		"role": c.role,
	}
	if err := req.SetJSONBody(payload); err != nil {
		return "", err
	}

	resp, err := c.vClient.RawRequest(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var b bytes.Buffer
		if _, err := io.Copy(&b, resp.Body); err != nil {
			return "", fmt.Errorf("failed to copy response body: %s", err)
		}
		return "", fmt.Errorf("failed to get successful response: %#v, %s",
			resp, b.String())
	}
	var s struct {
		Auth struct {
			ClientToken    string `json:"client_token"`
			ClientAccessor string `json:"accessor"`
		} `json:"auth"`
	}

	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return "", err
	}
	return s.Auth.ClientToken, nil
}
