package argo

import (
	"strconv"

	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/kubernetes"
)

type Args []string

func NewArgs(image string, port int, target string) Args {
	var args Args
	args = append(args, "--host")
	args = append(args, image+"-svc."+kubernetes.GetNamespace()+".svc.cluster.local")
	args = append(args, "--port")
	args = append(args, strconv.Itoa(port))
	args = append(args, "--target")
	args = append(args, target)
	return args
}

func NewArgsWithInput(name string, image string, port int, target string) Args {
	args := NewArgs(image, port, target)
	args = append(args, "--ifile")
	args = append(args, "/tmp/"+name+"-art.json")
	return args
}

func NewArgsWithOutput(name string, image string, port int, target string) Args {
	args := NewArgs(image, port, target)
	args = append(args, "--ofile")
	args = append(args, "/tmp/"+name+"-art.json")
	return args
}

func NewArgsWithInputAndOutput(iname string, oname string, image string, port int, target string) Args {
	args := NewArgs(image, port, target)
	args = append(args, "--ifile")
	args = append(args, "/tmp/"+iname+"-art.json")
	args = append(args, "--ofile")
	args = append(args, "/tmp/"+oname+"-art.json")
	return args
}
