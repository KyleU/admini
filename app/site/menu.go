// Content managed by Project Forge, see [projectforge.md] for details.
package site

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/lib/menu"
	"github.com/kyleu/admini/app/lib/user"
)

const (
	keyInstall     = "install"
	keyDownload    = "download"
	keyQuickStart  = "quickstart"
	keyContrib     = "contributing"
	keyCustomizing = "customizing"
	keyTech        = "technology"
)

func Menu(as *app.State, _ *user.Profile, _ user.Accounts) menu.Items {
	return menu.Items{
		{Key: keyInstall, Title: "Install", Icon: "code", Route: "/" + keyInstall},
		{Key: keyDownload, Title: "Download", Icon: "download", Route: "/" + keyDownload},
		{Key: keyQuickStart, Title: "Quick Start", Icon: "bolt", Route: "/" + keyQuickStart},
		{Key: keyCustomizing, Title: "Customizing", Icon: "code", Route: "/" + keyCustomizing},
		{Key: keyContrib, Title: "Contributing", Icon: "cog", Route: "/" + keyContrib},
		{Key: keyTech, Title: "Technology", Icon: "cog", Route: "/" + keyTech},
	}
}
