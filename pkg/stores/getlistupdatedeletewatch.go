package stores

import (
	"github.com/obot-platform/kinm/pkg/strategy"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	_ rest.Getter             = (*GetListUpdateDeleteWatchStore)(nil)
	_ rest.Lister             = (*GetListUpdateDeleteWatchStore)(nil)
	_ rest.Updater            = (*GetListUpdateDeleteWatchStore)(nil)
	_ rest.Watcher            = (*GetListUpdateDeleteWatchStore)(nil)
	_ rest.RESTDeleteStrategy = (*GetListUpdateDeleteWatchStore)(nil)
	_ strategy.Base           = (*GetListUpdateDeleteWatchStore)(nil)
)

type GetListUpdateDeleteWatchStore struct {
	*strategy.SingularNameAdapter
	*strategy.GetAdapter
	*strategy.UpdateAdapter
	*strategy.ListAdapter
	*strategy.DeleteAdapter
	*strategy.WatchAdapter
	*strategy.DestroyAdapter
	*strategy.TableAdapter
}

func (g *GetListUpdateDeleteWatchStore) NamespaceScoped() bool {
	return g.WatchAdapter.NamespaceScoped()
}
