package flagparser

const MinWaitSec = 0
const MaxWaitSec = 60

type flagData struct {
	flagReady bool

	help       bool
	configFile string
	wait       uint
}

func newFlagData() flagData {
	return flagData{
		flagReady:  false,
		help:       false,
		configFile: "",
		wait:       0,
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

func (d *flagData) Wait() uint {
	if !d.isReady() {
		panic("flag not ready")
	}

	if d.wait > MaxWaitSec {
		return MaxWaitSec
	} else if d.wait < MinWaitSec {
		return MinWaitSec
	} else {
		return d.wait
	}
}
