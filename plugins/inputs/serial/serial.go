package serial
// serial.go

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
)

type Serial struct {
    Ok  bool            `toml:"ok"`
    Log telegraf.Logger `toml:"-"`
}

func (s *Serial) Description() string {
    return "a serial plugin"
}

func (s *Serial) SampleConfig() string {
    return `
  ## Indicate if everything is fine
  ok = true
`
}

// Init is for setup, and validating config.
func (s *Serial) Init() error {
    return nil
}

func (s *Serial) Gather(acc telegraf.Accumulator) error {
    if s.Ok {
        acc.AddFields("serial", map[string]interface{}{"value": "pretty good"}, nil)
    } else {
        acc.AddFields("serial", map[string]interface{}{"value": "not great"}, nil)
    }

    return nil
}

func init() {
    inputs.Add("serial", func() telegraf.Input { return &Serial{} })
}
