package mason

// Ctx (Context) model
type Ctx struct {
	Options map[string]string
}

// Option adds command line options to context
func (ctx *Ctx) Option(opts map[string]string) {
	if ctx.Options == nil {
		ctx.Options = make(map[string]string)
	}

	for key, value := range opts {
		ctx.Options[key] = value
	}
}

// GetOptions returns options map
func (ctx *Ctx) GetOptions() map[string]string {
	return ctx.Options
}

// GetOption return specified option
func (ctx *Ctx) GetOption(key string) string {
	return ctx.Options[key]
}
