package mgr

import (
	"context"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func setupCgroupCPUShare(ctx context.Context, meta *ContainerMeta, spec *SpecWrapper) error {
	s := spec.s
	if s.Linux.Resources.CPU == nil {
		s.Linux.Resources.CPU = &specs.LinuxCPU{}
	}
	cpu := s.Linux.Resources.CPU

	v := uint64(meta.HostConfig.CPUShares)
	cpu.Shares = &v
	return nil
}

func setupCgroupCPUSet(ctx context.Context, meta *ContainerMeta, spec *SpecWrapper) error {
	s := spec.s
	if s.Linux.Resources.CPU == nil {
		s.Linux.Resources.CPU = &specs.LinuxCPU{}
	}
	cpu := s.Linux.Resources.CPU
	cpu.Cpus = meta.HostConfig.CpusetCpus
	cpu.Mems = meta.HostConfig.CpusetMems
	return nil
}

func setupCgroupCPUPeriod(ctx context.Context, meta *ContainerMeta, spec *SpecWrapper) error {
	s := spec.s
	if s.Linux.Resources.CPU == nil {
		s.Linux.Resources.CPU = &specs.LinuxCPU{}
	}
	cpu := s.Linux.Resources.CPU
	period := uint64(meta.HostConfig.CPUPeriod)
	cpu.Period = &period
	return nil
}

func setupCgroupCPUQuota(ctx context.Context, meta *ContainerMeta, spec *SpecWrapper) error {
	s := spec.s
	if s.Linux.Resources.CPU == nil {
		s.Linux.Resources.CPU = &specs.LinuxCPU{}
	}
	cpu := s.Linux.Resources.CPU
	quota := meta.HostConfig.CPUQuota
	cpu.Quota = &quota
	return nil
}
