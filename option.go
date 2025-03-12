package insomnium

// Option is used to help create an *Insomnuim instance.
type Option func(i *Insomnium)

// WithUnmarshaler sets the unmarshaler usede by the *Insomnium instance.
func WithUnmarshaler(u Unmarshaler) Option {
	return func(i *Insomnium) {
		i.unmarshaler = u
	}
}

// WithFileReader sets the file reader used by the *Insomnium instance.
func WithFileReader(fr FileReader) Option {
	return func(i *Insomnium) {
		i.fileReader = fr
	}
}
