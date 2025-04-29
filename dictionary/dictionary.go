package dictionary

import "untref-ayp2/guia-conjuntos-hashes-diccionarios/hashtable"

type Dictionary[K comparable, V any] struct {
	hash hashtable.HashTable[K, V]
}

// Implementar
