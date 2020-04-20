package overrides

import (
	"fmt"
	"reflect"

	"github.com/zachdeibert/canvas-sync/utils/apisync"
)

type overrideContext struct {
	parent *interface{}
	s0     *[]*apisync.Model
	s1     *[]apisync.MethodAPIPair
}

func createContext(models *[]*apisync.Model, methods *[]apisync.MethodAPIPair) *overrideContext {
	return &overrideContext {
		s0: models,
		s1: methods,
	}
}

func (o *overrideContext) addModels(obj *apisync.Model) *overrideContext {
	for _, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s0 = append(*o.s0, obj)
	return o
}

func (o *overrideContext) removeModels(obj *apisync.Model) *overrideContext {
	for i, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			*o.s0 = append((*o.s0)[:i], (*o.s0)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

func (o *overrideContext) model(search string) *overrideContextModels {
	for i, obj := range *o.s0 {
		if obj.Name == search {
			return &overrideContextModels {
				parent: o,
				s0: &(*o.s0)[i].Properties,
				p0: &(*o.s0)[i].Description,
				p1: &(*o.s0)[i].Name,
			}
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to find key '%v'", search))
}

type overrideContextModels struct {
	parent *overrideContext
	s0     *[]apisync.ModelProperty
	p0     *string
	p1     *string
}

func (o *overrideContextModels) done() *overrideContext {
	return o.parent
}

func (o *overrideContextModels) addProperties(obj apisync.ModelProperty) *overrideContextModels {
	for _, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s0 = append(*o.s0, obj)
	return o
}

func (o *overrideContextModels) removeProperties(obj apisync.ModelProperty) *overrideContextModels {
	for i, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			*o.s0 = append((*o.s0)[:i], (*o.s0)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

func (o *overrideContextModels) property(search string) *overrideContextModelsProperties {
	for i, obj := range *o.s0 {
		if obj.Name == search {
			return &overrideContextModelsProperties {
				parent: o,
				s0: &(*o.s0)[i].EnumValues,
				p0: &(*o.s0)[i].Type,
				p1: &(*o.s0)[i].Example,
				p2: &(*o.s0)[i].Description,
				p3: &(*o.s0)[i].Name,
			}
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to find key '%v'", search))
}

type overrideContextModelsProperties struct {
	parent *overrideContextModels
	s0     *[]string
	p0     *string
	p1     *string
	p2     *string
	p3     *string
}

func (o *overrideContextModelsProperties) done() *overrideContextModels {
	return o.parent
}

func (o *overrideContextModelsProperties) addEnumValues(obj string) *overrideContextModelsProperties {
	for _, e := range *o.s0 {
		if e == obj {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s0 = append(*o.s0, obj)
	return o
}

func (o *overrideContextModelsProperties) removeEnumValues(obj string) *overrideContextModelsProperties {
	for i, e := range *o.s0 {
		if e == obj {
			*o.s0 = append((*o.s0)[:i], (*o.s0)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

func (o *overrideContextModelsProperties) setType(old, new string) *overrideContextModelsProperties {
	if !(*o.p0 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p0))
	}
	*o.p0 = new
	return o
}

func (o *overrideContextModelsProperties) setExample(old, new string) *overrideContextModelsProperties {
	if !(*o.p1 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p1))
	}
	*o.p1 = new
	return o
}

func (o *overrideContextModelsProperties) setDescription(old, new string) *overrideContextModelsProperties {
	if !(*o.p2 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p2))
	}
	*o.p2 = new
	return o
}

func (o *overrideContextModelsProperties) setName(old, new string) *overrideContextModelsProperties {
	if !(*o.p3 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p3))
	}
	*o.p3 = new
	return o
}

func (o *overrideContextModels) setDescription(old, new string) *overrideContextModels {
	if !(*o.p0 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p0))
	}
	*o.p0 = new
	return o
}

func (o *overrideContextModels) setName(old, new string) *overrideContextModels {
	if !(*o.p1 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p1))
	}
	*o.p1 = new
	return o
}

func (o *overrideContext) addMethods(obj apisync.MethodAPIPair) *overrideContext {
	for _, e := range *o.s1 {
		if reflect.DeepEqual(e, obj) {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s1 = append(*o.s1, obj)
	return o
}

func (o *overrideContext) removeMethods(obj apisync.MethodAPIPair) *overrideContext {
	for i, e := range *o.s1 {
		if reflect.DeepEqual(e, obj) {
			*o.s1 = append((*o.s1)[:i], (*o.s1)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

type overrideContextMethods struct {
	parent *overrideContext
	s0     *[]apisync.MethodArgument
	p0     *string
	p1     *string
	p2     *string
	p3     *string
	p4     *string
}

func (o *overrideContextMethods) done() *overrideContext {
	return o.parent
}

func (o *overrideContextMethods) addMethodArguments(obj apisync.MethodArgument) *overrideContextMethods {
	for _, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s0 = append(*o.s0, obj)
	return o
}

func (o *overrideContextMethods) removeMethodArguments(obj apisync.MethodArgument) *overrideContextMethods {
	for i, e := range *o.s0 {
		if reflect.DeepEqual(e, obj) {
			*o.s0 = append((*o.s0)[:i], (*o.s0)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

func (o *overrideContextMethods) arg(search string) *overrideContextMethodsMethodArguments {
	for i, obj := range *o.s0 {
		if obj.Name == search {
			return &overrideContextMethodsMethodArguments {
				parent: o,
				s0: &(*o.s0)[i].EnumValues,
				p0: &(*o.s0)[i].Description,
				p1: &(*o.s0)[i].Type,
				p2: &(*o.s0)[i].Required,
				p3: &(*o.s0)[i].Name,
			}
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to find key '%v'", search))
}

type overrideContextMethodsMethodArguments struct {
	parent *overrideContextMethods
	s0     *[]string
	p0     *string
	p1     *string
	p2     *bool
	p3     *string
}

func (o *overrideContextMethodsMethodArguments) done() *overrideContextMethods {
	return o.parent
}

func (o *overrideContextMethodsMethodArguments) addEnumValues(obj string) *overrideContextMethodsMethodArguments {
	for _, e := range *o.s0 {
		if e == obj {
			panic(fmt.Errorf("Error applying override patch: unable to add '%v' to list due to it already existing", obj))
		}
	}
	*o.s0 = append(*o.s0, obj)
	return o
}

func (o *overrideContextMethodsMethodArguments) removeEnumValues(obj string) *overrideContextMethodsMethodArguments {
	for i, e := range *o.s0 {
		if e == obj {
			*o.s0 = append((*o.s0)[:i], (*o.s0)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%v' from list due to it not existing", obj))
}

func (o *overrideContextMethodsMethodArguments) setDescription(old, new string) *overrideContextMethodsMethodArguments {
	if !(*o.p0 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p0))
	}
	*o.p0 = new
	return o
}

func (o *overrideContextMethodsMethodArguments) setType(old, new string) *overrideContextMethodsMethodArguments {
	if !(*o.p1 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p1))
	}
	*o.p1 = new
	return o
}

func (o *overrideContextMethodsMethodArguments) setRequired(old, new bool) *overrideContextMethodsMethodArguments {
	if !(*o.p2 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p2))
	}
	*o.p2 = new
	return o
}

func (o *overrideContextMethodsMethodArguments) setName(old, new string) *overrideContextMethodsMethodArguments {
	if !(*o.p3 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p3))
	}
	*o.p3 = new
	return o
}

func (o *overrideContextMethods) setAPIName(old, new string) *overrideContextMethods {
	if !(*o.p0 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p0))
	}
	*o.p0 = new
	return o
}

func (o *overrideContextMethods) setMethodEndPoint(old, new string) *overrideContextMethods {
	if !(*o.p1 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p1))
	}
	*o.p1 = new
	return o
}

func (o *overrideContextMethods) setMethodReturnType(old, new string) *overrideContextMethods {
	if !(*o.p2 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p2))
	}
	*o.p2 = new
	return o
}

func (o *overrideContextMethods) setMethodDescription(old, new string) *overrideContextMethods {
	if !(*o.p3 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p3))
	}
	*o.p3 = new
	return o
}

func (o *overrideContextMethods) setMethodName(old, new string) *overrideContextMethods {
	if !(*o.p4 == old) {
		panic(fmt.Errorf("Error applying override patch: expected '%v' but got '%v'", old, *o.p4))
	}
	*o.p4 = new
	return o
}
