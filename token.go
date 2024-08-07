package idempotence

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

var tokenNotExist = fmt.Errorf("token not exist")
var once sync.Once
var tkService *tokenService

type TokenStore interface {
	Get(key string) (any, bool)
	Put(key string, val any)
	Delete(key string) error
}

// MapStore default impl map store
// key is token , value is struct{}{}
type MapStore struct {
	Store sync.Map
}

func (mt *MapStore) Get(key string) (any, bool) {
	v, ok := mt.Store.Load(key)
	if ok {
		return v.(string), true
	}
	return "", false
}

func (mt *MapStore) Put(key string, val any) {
	mt.Store.Store(key, val)
}

func (mt *MapStore) Delete(key string) error {
	_, exist := mt.Store.LoadAndDelete(key)
	if !exist {
		return tokenNotExist
	}

	return nil
}

type TokenBuilder interface {
	Build() string
}

type UUIDToken struct{}

func (UUIDToken) Build() string {
	return uuid.NewV1().String()
}

type tokenService struct {
	Store   TokenStore
	Builder TokenBuilder
}

func (ts *tokenService) GetToken() string {
	token := ts.Builder.Build()
	//token as key ,struct{}{} as value
	ts.Store.Put(token, struct{}{})
	return token
}

func (ts *tokenService) CheckToken(token string) error {
	return ts.Store.Delete(token)
}

func NewDefaultTokenService() *tokenService {
	once.Do(func() {
		tkService = &tokenService{
			Store:   &MapStore{Store: sync.Map{}},
			Builder: &UUIDToken{},
		}
	})
	return tkService
}

func NewTokenService(store TokenStore, builder TokenBuilder) *tokenService {
	once.Do(func() {
		tkService = &tokenService{
			Store:   store,
			Builder: builder,
		}
	})
	return tkService
}
