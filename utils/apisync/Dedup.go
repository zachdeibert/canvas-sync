package apisync

import "fmt"

// MethodAPIPair is a method paired to an API name
type MethodAPIPair struct {
	Method  *Method
	APIName string
}

// DedupModels deduplicates models
func DedupModels(models *[]*Model) {
	cache := map[string]*Model{}
	n := 0
	for _, model := range *models {
		other, ok := cache[model.Name]
		if !ok {
			cache[model.Name] = model
			(*models)[n] = model
			n++
		} else {
			for _, p := range model.Properties {
				found := false
				for _, p2 := range other.Properties {
					if p.Name == p2.Name {
						found = true
						break
					}
				}
				if !found {
					other.Properties = append(other.Properties, p)
				}
			}
		}
	}
	*models = (*models)[:n]
}

// DedupMethods deduplicates methods
func DedupMethods(methods *[]MethodAPIPair) {
	cache := map[string]MethodAPIPair{}
	nextID := map[string]int{}
	n := 0
	for _, method := range *methods {
		qualName := fmt.Sprintf("%s_%s", method.APIName, method.Method.Name)
		other, ok := cache[qualName]
		if !ok {
			cache[qualName] = method
			nextID[qualName] = 2
			(*methods)[n] = method
			n++
		} else if other.Method.ReturnType != method.Method.ReturnType || other.Method.EndPoint != method.Method.EndPoint {
			id := nextID[qualName]
			method.Method.Name = fmt.Sprintf("%s%d", method.Method.Name, id)
			qualName = fmt.Sprintf("%s%d", qualName, id)
			cache[qualName] = method
			(*methods)[n] = method
			n++
		} else {
			for _, a := range method.Method.Arguments {
				found := false
				for _, a2 := range other.Method.Arguments {
					if a.Name == a2.Name {
						found = true
						break
					}
				}
				if !found {
					other.Method.Arguments = append(other.Method.Arguments, a)
				}
			}
		}
	}
	*methods = (*methods)[:n]
}
