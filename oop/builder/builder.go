package builder

import "fmt"

type ResourcePoolConf struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfFunc func(option *ResourcePoolConfOption)

func NewResourcePoolConf(name string, opts ...ResourcePoolConfFunc) (*ResourcePoolConf, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	option := &ResourcePoolConfOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}
	for _, opt := range opts {
		opt(option)
	}

	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}
	return &ResourcePoolConf{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil

}
