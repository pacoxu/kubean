//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerInfo) DeepCopyInto(out *DockerInfo) {
	*out = *in
	if in.VersionRange != nil {
		in, out := &in.VersionRange, &out.VersionRange
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerInfo.
func (in *DockerInfo) DeepCopy() *DockerInfo {
	if in == nil {
		return nil
	}
	out := new(DockerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerInfoStatus) DeepCopyInto(out *DockerInfoStatus) {
	*out = *in
	if in.VersionRange != nil {
		in, out := &in.VersionRange, &out.VersionRange
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerInfoStatus.
func (in *DockerInfoStatus) DeepCopy() *DockerInfoStatus {
	if in == nil {
		return nil
	}
	out := new(DockerInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostsMap) DeepCopyInto(out *HostsMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostsMap.
func (in *HostsMap) DeepCopy() *HostsMap {
	if in == nil {
		return nil
	}
	out := new(HostsMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalAvailable) DeepCopyInto(out *LocalAvailable) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]*SoftwareInfoStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SoftwareInfoStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = make([]*DockerInfoStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DockerInfoStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalAvailable.
func (in *LocalAvailable) DeepCopy() *LocalAvailable {
	if in == nil {
		return nil
	}
	out := new(LocalAvailable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalService) DeepCopyInto(out *LocalService) {
	*out = *in
	if in.ImageRepo != nil {
		in, out := &in.ImageRepo, &out.ImageRepo
		*out = make(map[ImageRepoType]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.YumRepos != nil {
		in, out := &in.YumRepos, &out.YumRepos
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.HostsMap != nil {
		in, out := &in.HostsMap, &out.HostsMap
		*out = make([]*HostsMap, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HostsMap)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalService.
func (in *LocalService) DeepCopy() *LocalService {
	if in == nil {
		return nil
	}
	out := new(LocalService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Manifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestList) DeepCopyInto(out *ManifestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Manifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestList.
func (in *ManifestList) DeepCopy() *ManifestList {
	if in == nil {
		return nil
	}
	out := new(ManifestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManifestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareInfo) DeepCopyInto(out *SoftwareInfo) {
	*out = *in
	if in.VersionRange != nil {
		in, out := &in.VersionRange, &out.VersionRange
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareInfo.
func (in *SoftwareInfo) DeepCopy() *SoftwareInfo {
	if in == nil {
		return nil
	}
	out := new(SoftwareInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareInfoStatus) DeepCopyInto(out *SoftwareInfoStatus) {
	*out = *in
	if in.VersionRange != nil {
		in, out := &in.VersionRange, &out.VersionRange
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareInfoStatus.
func (in *SoftwareInfoStatus) DeepCopy() *SoftwareInfoStatus {
	if in == nil {
		return nil
	}
	out := new(SoftwareInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	in.LocalService.DeepCopyInto(&out.LocalService)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]*SoftwareInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SoftwareInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = make([]*DockerInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DockerInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	in.LocalAvailable.DeepCopyInto(&out.LocalAvailable)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}
