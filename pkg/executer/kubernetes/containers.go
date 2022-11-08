package kubernetes

import (
	apiv1 "k8s.io/api/core/v1"
)

// Implement class of container queue
type container_queue []apiv1.Container

// Create container queue
func newContainerQueue() container_queue {
	que := make(container_queue, 0)
	return que
}

func (que *container_queue) push(name string, port int, image string, command []string, args []string) {
	port_queue := make([]apiv1.ContainerPort, 0)
	container_port := apiv1.ContainerPort{
		ContainerPort: int32(port),
	}
	port_queue = append(port_queue, container_port)

	container := apiv1.Container{
		Name:            name,
		Image:           image,
		Command:         command,
		Args:            args,
		Ports:           port_queue,
		ImagePullPolicy: "IfNotPresent",
	}
	*que = append(*que, container)
}

func (que *container_queue) pop() {
	if len(*que) > 0 {
		*que = (*que)[:len(*que)-1]
	}
}
