package Models

import "fmt"

type Message struct{
	K string `form:"k"`
	V string `form:"v"`

}

type KV struct{
	kv map[string]string
}

func NewKV()*KV{
	return &KV{
		kv:make(map[string]string),
	}
}

func (kv *KV)GetValue(k Message,v *Message) error {
	var ok bool
	v.K=k.K
	v.V,ok=kv.kv[k.K]
	if !ok {
		return fmt.Errorf("unknown key %v",k.K)
	}
	return nil
}

func (kv *KV)SetValue(k Message,v *Message) error {
	kv.kv[k.K]=k.V
	v.K=k.K
	v.V=k.V
	return nil
}