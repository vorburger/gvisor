// automatically generated by stateify.

package ilist

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (l *List) StateTypeName() string {
	return "pkg/ilist.List"
}

func (l *List) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *List) beforeSave() {}

func (l *List) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *List) afterLoad() {}

func (l *List) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *Entry) StateTypeName() string {
	return "pkg/ilist.Entry"
}

func (e *Entry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *Entry) beforeSave() {}

func (e *Entry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *Entry) afterLoad() {}

func (e *Entry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*List)(nil))
	state.Register((*Entry)(nil))
}