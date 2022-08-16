package v1alpha1

import (
	"broadcast.logger/schemalogger/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *SchemaLogger) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.SchemaLogger)
	dst.Spec.Condition = "default-condition"
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = v1beta1.SchemaLoggerStatus(src.Status)
	return nil
}

func (dst *SchemaLogger) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.SchemaLogger)
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = SchemaLoggerStatus(src.Status)
	return nil
}
