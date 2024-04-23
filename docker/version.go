package docker

import (
	"time"

	"sub-store-manager-cli/lib"
	"sub-store-manager-cli/vars"
)

type ReleaseInfo struct {
	Url         string    `json:"url"`
	AssetsUrl   string    `json:"assets_url"`
	UploadUrl   string    `json:"upload_url"`
	HtmlUrl     string    `json:"html_url"`
	Id          int       `json:"id"`
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"published_at"`
	Body        string    `json:"body"`
}

func (c *Container) getVersionInfos() (json []ReleaseInfo) {
	var repoURL string
	switch c.ContainerType {
	case vars.ContainerTypeFE:
		repoURL = "https://api.github.com/repos/sub-store-org/Sub-Store-Front-End/releases"
	case vars.ContainerTypeBE:
		repoURL = "https://api.github.com/repos/sub-store-org/Sub-Store/releases"
	}
	_, err := lib.HC.R().SetResult(&json).Get(repoURL)
	if err != nil {
		lib.PrintError("Failed to get versions info:", err)
	}
	return
}

func (c *Container) getVersionStrs() (v []string) {
	for _, info := range c.getVersionInfos() {
		v = append(v, info.TagName)
	}
	return
}

func (c *Container) SetLatestVersion() {
	versions := c.getVersionStrs()
	if len(versions) == 0 {
		lib.PrintError("no versions found", nil)
	}
	c.Version = versions[0]
}

func (c *Container) CheckVersionValid() bool {
	versions := c.getVersionStrs()
	for _, v := range versions {
		if v == c.Version {
			return true
		}
	}
	return false
}
