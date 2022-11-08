package v1alpha1

import (
	"fmt"

	"broadcast.logger/schemalogger/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var broadcastlog = logf.Log.WithName("schemalogger-resource")

func (src *SchemaLogger) ConvertTo(dstRaw conversion.Hub) error {
	broadcastlog.Info("Conversion webhook: from v1alpha1 to v1")
	dst := dstRaw.(*v1.SchemaLogger)
	dst.Spec.Condition = "default-condition"
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = v1.SchemaLoggerStatus(src.Status)
	dst.Spec.AppName = src.Spec.AppName
	dst.Spec.Image = src.Spec.Image
	dst.Spec.Title = src.Spec.Title
	dst.Spec.CRDVersionTest1 = "default-tst"
	fmt.Println(dst)
	return nil
}

func (dst *SchemaLogger) ConvertFrom(srcRaw conversion.Hub) error {
	broadcastlog.Info("Conversion webhook: from v1 to v1alpha1")
	src := srcRaw.(*v1.SchemaLogger)
	dst.ObjectMeta = src.ObjectMeta
	dst.Status = SchemaLoggerStatus(src.Status)
	dst.Spec.AppName = src.Spec.AppName
	dst.Spec.Image = src.Spec.Image
	dst.Spec.Title = src.Spec.Title
	fmt.Println(dst)
	return nil
}
