package db0101

import (
	"bytes"
)

type KV struct {
	mem map[string][]byte // Go doesnt allow []byte as key so string are used
}

// Disc IO will be added later and those will generate error,i mean this one do generate error but it normal ones so we just dont count it enough to stop the program instead we use ok or updated etc

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // initialize empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	prev, exist := kv.mem[string(key)]
	kv.mem[string(key)] = val
	updated = !bytes.Equal(prev, val) || !exist
	return
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {

	_, exist := kv.mem[string(key)] // in go we cant declare variable we wont be using hence i needed to remove prev

	deleted = exist // it it exist then it will be deleted else it wont be deleted

	delete(kv.mem, string(key))

	return
}
