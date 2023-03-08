package csv2xlsx

type Options struct {
	InputFilePath      string
	OutputFilePath     string
	OutputFilePassword string
	Comma              string
}

func (opts *Options) Ajust() {
	if len(opts.OutputFilePath) == 0 {
		opts.OutputFilePath = opts.InputFilePath + ".xlsx"
	}
}
