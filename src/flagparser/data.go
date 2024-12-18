package flagparser

type flagData struct {
	flagReady bool

	help       bool
	configFile string
}

func newFlagData() flagData {
	return flagData{
		flagReady:  false,
		help:       false,
		configFile: "",
	}
}

func (d *flagData) ready() {
	d.flagReady = true
}

func (d *flagData) isReady() bool {
	return d.flagReady
}

func (d *flagData) Help() bool {
	if !d.isReady() {
		panic("flag not ready")
	}

	return d.help
}

func (d *flagData) ConfigFile() string {
	if !d.isReady() {
		panic("flag not ready")
	}

	return d.configFile
}
