package v1alpha1

import (
	"fmt"

	"broadcast.logger/schemalogger/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var broadcastlog = logf.Log.WithName("schemalogger-resource")

func (src *SchemaLogger) ConvertTo(dstRaw conversion.Hub) error {
	broadcastlog.Info("Conversion webhook: from v1alpha1 to v1beta1")
	dst := dstRaw.(*v1beta1.SchemaLogger)
	dst.Spec.Condition = "default-condition"
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = v1beta1.SchemaLoggerStatus(src.Status)
	fmt.Println(dst)
	return nil
}

func (dst *SchemaLogger) ConvertFrom(srcRaw conversion.Hub) error {
	broadcastlog.Info("Conversion webhook: from v1beta1 to v1alpha1")
	src := srcRaw.(*v1beta1.SchemaLogger)
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = SchemaLoggerStatus(src.Status)
	broadcastlog.Info("Converted v1beta1 to v1alpha1", dst)
	fmt.Println(dst)
	return nil
}
