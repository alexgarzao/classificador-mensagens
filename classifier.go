package main

type Classifier struct {
}

func New() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Types(text string) []string {
	return []string{}
}
