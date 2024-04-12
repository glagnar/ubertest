package featureimplementation1

import "github.com/glagnar/ubertest/feature"

type implementation struct{}

func NewFeature() feature.FeaturePlace {
	return &implementation{}
}

func (i *implementation) SayHello() {
	println("Hello from impl1")
}
