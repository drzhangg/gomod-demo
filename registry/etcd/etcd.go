package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"gomod-demo/registry"
	"sync"
	"sync/atomic"
)

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service //map类型，value存放service的name和节点信息
}

type RegisterService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

type EtcdRegistry struct {
	options            *registry.Options
	client             *clientv3.Client
	serviceCh          chan *registry.Service
	value              atomic.Value
	lock               sync.Mutex
	registryServiceMap map[string]RegisterService
}

const (
	MaxServiceNum = 8
)

var (
	//初始化变量
	etcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaxServiceNum),
		registryServiceMap: make(map[string]RegisterService),
	}
)

//初始化etcd方法
func init() {

}

//循环操作
func (e *EtcdRegistry) run() {

}

// 
func (e *EtcdRegistry) syncServiceFromEtcd() {

}

//插件名字	实现接口
func (e *EtcdRegistry) Name() string {
	return "etcd"
}

//初始化	实现接口
func (e *EtcdRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	e.options = &registry.Options{}
	for _, opt := range opts {
		opt(e.options)
	}
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.options.Addrs,
		DialTimeout: e.options.Timeout,
	})
	if err != nil {
		err = fmt.Errorf("init etcd err:%v", err)
		return
	}
	return
}

//服务注册	实现接口
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {

}

//服务反注册	实现接口
func (e *EtcdRegistry) UnRegister(ctx context.Context, service *registry.Service) (err error) {

}

//服务发现	实现接口
func (e *EtcdRegistry) GetService(ctx context.Context, name string) (service *registry.Service, err error) {

}
