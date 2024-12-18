package config

type ConfigStruct struct {
	configReady   bool
	yamlHasParser bool

	Yaml YamlConfig
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
	err = c.Yaml.check()
	if err != nil && err.IsError() {
		return err
	}

	return nil
}

func (c *ConfigStruct) ready() (err ConfigError) {
	if c.configReady {
		return nil
	}

	parserErr := c.parser()
	if parserErr != nil {
		return NewConfigError("parser yaml error: " + parserErr.Error())
	} else if !c.yamlHasParser {
		return NewConfigError("parser yaml error: unknown")
	}

	c.setDefault()
	err = c.check()
	if err != nil && err.IsError() {
		return err
	}

	c.configReady = true
	return nil
}
