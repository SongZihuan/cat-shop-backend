package config

type ConfigStruct struct {
	configReady   bool
	yamlHasParser bool

	Yaml YamlConfig
	File FileLocationConfig
}

func (c *ConfigStruct) init() error {
	err := c.File.init()
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigStruct) parser() ParserError {
	err := c.Yaml.parser()
	if err != nil {
		return err
	}

	c.yamlHasParser = true
	return nil
}

func (c *ConfigStruct) setDefault() {
	if !c.yamlHasParser {
		panic("yaml must parser first")
	}

	c.Yaml.setDefault()
}

func (c *ConfigStruct) check() (err ConfigError) {
	err = c.Yaml.check(&c.File)
	if err != nil && err.IsError() {
		return err
	}

	return nil
}

func (c *ConfigStruct) ready() (err ConfigError) {
	if c.configReady {
		return nil
	}

	initErr := c.init()
	if initErr != nil {
		return NewConfigError("init error: " + initErr.Error())
	}

	parserErr := c.parser()
	if parserErr != nil {
		return NewConfigError("parser error: " + parserErr.Error())
	} else if !c.yamlHasParser {
		return NewConfigError("parser error: unknown")
	}

	c.setDefault()
	err = c.check()
	if err != nil && err.IsError() {
		return err
	}

	c.configReady = true
	return nil
}
