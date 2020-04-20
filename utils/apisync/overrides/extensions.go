package overrides

import (
	"fmt"

	"github.com/zachdeibert/canvas-sync/utils/apisync"
)

func (o *overrideContext) method(search string) *overrideContextMethods {
	for i, obj := range *o.s1 {
		if apisync.ToGoIdentifier(fmt.Sprintf("%s_%s", obj.APIName, obj.Method.Name), true) == search {
			return &overrideContextMethods{
				parent: o,
				s0:     &(*o.s1)[i].Method.Arguments,
				p0:     &(*o.s1)[i].APIName,
				p1:     &(*o.s1)[i].Method.EndPoint,
				p2:     &(*o.s1)[i].Method.ReturnType,
				p3:     &(*o.s1)[i].Method.Description,
				p4:     &(*o.s1)[i].Method.Name,
			}
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to find key '%v'", search))
}
