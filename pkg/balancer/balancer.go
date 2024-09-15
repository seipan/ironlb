package balancer

import (
	"net/http"
)

type LoadBalancer struct {
	targets []string
	current int
}

func NewLoadBalancer(targets []string) *LoadBalancer {
	return &LoadBalancer{
		targets: targets,
		current: 0,
	}
}

func (lb *LoadBalancer) GetNextTarget() string {
	target := lb.targets[lb.current]
	lb.current = (lb.current + 1) % len(lb.targets)
	return target
}

func (lb *LoadBalancer) Start(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := lb.GetNextTarget()
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})
	return http.ListenAndServe(addr, nil)
}
