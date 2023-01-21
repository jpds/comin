package config

import (
	"github.com/nlewo/comin/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	configPath := "./configuration.yaml"
	expected := types.Configuration{
		Hostname: "machine",
		StateDir: "/var/lib/comin",
		Poller: types.Poller{
			Period: 10,
		},
		Remotes: []types.Remote{
			types.Remote{
				Name: "origin",
				URL:  "https://framagit.org/owner/infra",
				Auth: types.Auth{
					AccessToken:     "my-secret",
					AccessTokenPath: "./secret",
				},
			},
			types.Remote{
				Name: "local",
				URL:  "/home/owner/git/infra",
				Auth: types.Auth{
					AccessToken:     "",
					AccessTokenPath: "",
				},
			},
		},
		Branches: types.Branches{
			Main: types.Branch{
				Name:      "main",
				Protected: true,
			},
			Testing: types.Branch{
				Name:      "testing-machine",
				Protected: false,
			},
		},
	}
	config, err := Read(configPath)
	assert.Nil(t, err)
	assert.Equal(t, expected, config)
}
