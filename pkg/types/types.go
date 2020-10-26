/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package types

const (
	VDeviceAnnotation       = "dana.794/vcuda-device"
	VCoreAnnotation         = "dana.794/vcuda-core"
	VCoreLimitAnnotation    = "dana.794/vcuda-core-limit"
	VMemoryAnnotation       = "dana.794/vcuda-memory"
	PredicateTimeAnnotation = "dana.794/predicate-time"
	PredicateGPUIndexPrefix = "dana.794/predicate-gpu-idx-"
	GPUAssigned             = "dana.794/gpu-assigned"
	ClusterNameAnnotation   = "clusterName"

	VCUDA_MOUNTPOINT = "/etc/vcuda"

	/** 256MB */
	MemoryBlockSize = 268435456

	KubeletSocket                 = "kubelet.sock"
	VDeviceSocket                 = "vcuda.sock"
	CheckPointFileName            = "kubelet_internal_checkpoint"
	PreStartContainerCheckErrMsg  = "PreStartContainer check failed"
	PreStartContainerCheckErrType = "PreStartContainerCheckErr"
	UnexpectedAdmissionErrType    = "UnexpectedAdmissionError"
)

const (
	NvidiaCtlDevice    = "/run/nvidia/driver/dev/nvidiactl"
	NvidiaUVMDevice    = "/run/nvidia/driver/dev/nvidia-uvm"
	NvidiaFullpathRE   = `^/run/nvidia/driver/dev/nvidia([0-9]*)$`
	NvidiaDevicePrefix = "/run/nvidia/driver/dev/nvidia"
)

const (
	ManagerSocket = "/var/run/gpu-manager.sock"
)

const (
	CGROUP_BASE  = "/sys/fs/cgroup/memory"
	CGROUP_PROCS = "cgroup.procs"
)

type VCudaRequest struct {
	PodUID string
	//Deprecated
	ContainerName string
	//Deprecated
	Cores int64
	//Deprecated
	Memory int64
	Done   chan error
}

type PodDevicesEntry struct {
	PodUID        string
	ContainerName string
	ResourceName  string
	DeviceIDs     []string
	AllocResp     []byte
}

type Checkpoint struct {
	PodDeviceEntries  []PodDevicesEntry
	RegisteredDevices map[string][]string
}

type CheckpointData struct {
	Data *Checkpoint `json:"Data"`
}

var (
	DriverVersionMajor      int
	DriverVersionMinor      int
	DriverLibraryPath       string
	DriverOriginLibraryPath string
)

const (
	ContainerNameLabelKey = "io.kubernetes.container.name"
	PodNamespaceLabelKey  = "io.kubernetes.pod.namespace"
	PodNameLabelKey       = "io.kubernetes.pod.name"
	PodUIDLabelKey        = "io.kubernetes.pod.uid"
	PodCgroupNamePrefix   = "pod"
)
