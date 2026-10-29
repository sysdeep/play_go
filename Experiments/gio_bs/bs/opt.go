package bs

import "giobs/bs/theme"

type OptionsWidget interface {
	SetClass(className theme.ClassName)
}

type Opt func(w OptionsWidget)

func WithClass(className theme.ClassName) Opt {
	return func(w OptionsWidget) {
		w.SetClass(className)
	}
}
