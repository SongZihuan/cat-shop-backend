package flagparser

var data flagData

func Help() bool {
	return data.Help()
}

func ConfigFile() string {
	return data.ConfigFile()
}
