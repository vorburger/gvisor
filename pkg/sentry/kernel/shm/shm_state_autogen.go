// automatically generated by stateify.

package shm

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (r *Registry) StateTypeName() string {
	return "pkg/sentry/kernel/shm.Registry"
}

func (r *Registry) StateFields() []string {
	return []string{
		"userNS",
		"shms",
		"keysToShms",
		"totalPages",
		"lastIDUsed",
	}
}

func (r *Registry) beforeSave() {}

// +checklocksignore
func (r *Registry) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.userNS)
	stateSinkObject.Save(1, &r.shms)
	stateSinkObject.Save(2, &r.keysToShms)
	stateSinkObject.Save(3, &r.totalPages)
	stateSinkObject.Save(4, &r.lastIDUsed)
}

func (r *Registry) afterLoad() {}

// +checklocksignore
func (r *Registry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.userNS)
	stateSourceObject.Load(1, &r.shms)
	stateSourceObject.Load(2, &r.keysToShms)
	stateSourceObject.Load(3, &r.totalPages)
	stateSourceObject.Load(4, &r.lastIDUsed)
}

func (s *Shm) StateTypeName() string {
	return "pkg/sentry/kernel/shm.Shm"
}

func (s *Shm) StateFields() []string {
	return []string{
		"ShmRefs",
		"mfp",
		"registry",
		"ID",
		"creator",
		"size",
		"effectiveSize",
		"fr",
		"key",
		"perms",
		"owner",
		"attachTime",
		"detachTime",
		"changeTime",
		"creatorPID",
		"lastAttachDetachPID",
		"pendingDestruction",
	}
}

func (s *Shm) beforeSave() {}

// +checklocksignore
func (s *Shm) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.ShmRefs)
	stateSinkObject.Save(1, &s.mfp)
	stateSinkObject.Save(2, &s.registry)
	stateSinkObject.Save(3, &s.ID)
	stateSinkObject.Save(4, &s.creator)
	stateSinkObject.Save(5, &s.size)
	stateSinkObject.Save(6, &s.effectiveSize)
	stateSinkObject.Save(7, &s.fr)
	stateSinkObject.Save(8, &s.key)
	stateSinkObject.Save(9, &s.perms)
	stateSinkObject.Save(10, &s.owner)
	stateSinkObject.Save(11, &s.attachTime)
	stateSinkObject.Save(12, &s.detachTime)
	stateSinkObject.Save(13, &s.changeTime)
	stateSinkObject.Save(14, &s.creatorPID)
	stateSinkObject.Save(15, &s.lastAttachDetachPID)
	stateSinkObject.Save(16, &s.pendingDestruction)
}

func (s *Shm) afterLoad() {}

// +checklocksignore
func (s *Shm) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.ShmRefs)
	stateSourceObject.Load(1, &s.mfp)
	stateSourceObject.Load(2, &s.registry)
	stateSourceObject.Load(3, &s.ID)
	stateSourceObject.Load(4, &s.creator)
	stateSourceObject.Load(5, &s.size)
	stateSourceObject.Load(6, &s.effectiveSize)
	stateSourceObject.Load(7, &s.fr)
	stateSourceObject.Load(8, &s.key)
	stateSourceObject.Load(9, &s.perms)
	stateSourceObject.Load(10, &s.owner)
	stateSourceObject.Load(11, &s.attachTime)
	stateSourceObject.Load(12, &s.detachTime)
	stateSourceObject.Load(13, &s.changeTime)
	stateSourceObject.Load(14, &s.creatorPID)
	stateSourceObject.Load(15, &s.lastAttachDetachPID)
	stateSourceObject.Load(16, &s.pendingDestruction)
}

func (r *ShmRefs) StateTypeName() string {
	return "pkg/sentry/kernel/shm.ShmRefs"
}

func (r *ShmRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *ShmRefs) beforeSave() {}

// +checklocksignore
func (r *ShmRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *ShmRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*Registry)(nil))
	state.Register((*Shm)(nil))
	state.Register((*ShmRefs)(nil))
}