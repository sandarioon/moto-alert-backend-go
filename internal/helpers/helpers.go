package helpers

import "fmt"

func CreateGeom(longitude, latitude float32) string {
	// SRID=4326;POINT(20.562519027 54.733065521)

	return fmt.Sprintf("SRID=4326;POINT(%f %f)", longitude, latitude)
}
