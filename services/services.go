package services

import (
	"fmt"
	"strings"
)

type Service struct {
	domain IDomain
}

func New(domain IDomain) *Service {
	return &Service{domain: domain}
}

func (s *Service) BeefSummaryService() map[string]int {
	input := s.domain.GetResponseBody()
	fmt.Println(string(input))
	ss := strings.Split(strings.ReplaceAll(string(input), "\n", ""), " ")
	types := make(map[string]int)
	for _, s := range ss {
		if len(s) == 0 {
			continue
		}
		s = Filter(s)
		types[s] += 1
	}

	return types
}

func Filter(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ",", "", -1)
	return strings.ToLower(s)
}
