package docs

import "embed"

//go:embed index.md getting_started.md assets/logo.png assets/favicon.png assets/preview.png _meta.md
var InitFiles embed.FS
