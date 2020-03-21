package registry

import (
	"context"
	"fmt"
	"sync"
)

/*
插件管理类
可以用一个大的map管理，key字符串，value是Registry接口对象

调用流程：
	先注册插件，才能初始化插件
*/

type PluginMgr struct {
	//map维护所有插件
	plugins map[string]Registry
	//因为是要进行写操作，所以加个锁
	lock sync.Mutex
}

//初始化map变量
var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

//对外开放插件注册方法
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

//内部函数调用插件注册方法
func (p *PluginMgr) registerPlugin(registry Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	//map里的key是registry.Name，里面放的是插件类型的名称，value是Registry整个结构体属性
	_, ok := p.plugins[registry.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exits")
		return
	}
	p.plugins[registry.Name()] = registry
	return
}

//插件初始化
func InitPlugin(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initPlugin(ctx, name, opts...)
}

func (p *PluginMgr) initPlugin(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	//先查看服务列表，服务是否存在，只有服务存在，才能进行初始化
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exits", name)
		return
	}

	//服务存在，返回值赋值
	plugin = registry
	err = plugin.Init(ctx, opts...)
	return
}
