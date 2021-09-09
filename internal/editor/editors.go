package editor

var editors = map[string][]string{
	"vim":     {"vim"},
	"nvim":    {"nvim"},
	"eclipse": {"eclipse --launcher.openFile"},
	// MS
	"atom":   {"atom"},
	"vscode": {"code"},
	//JetBrains
	"intellij-idea": {"idea64.exe", "idea", "idea.sh"},
	"webstorm":      {"webstorm.exe", "webstorm", "webstorm.sh"},
	"pycharm":       {"pycharm.exe", "pycharm", "pycharm.sh"},
	"goland":        {"golanf.exe", "goland", "goland.sh"},
	"phpstorm":      {"phpstorm.exe", "phpstorm", "phpstorm.sh"},
	"datagrip":      {"datagrip.exe", "datagrip", "datagrip.sh"},
	"rider":         {"rider.exe", "rider", "rider.sh"},
	"rubymine":      {"rubymine64.exe", "rubymine", "rubymine.sh"},
}
