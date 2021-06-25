package main

import (
	"hash/fnv"
)

var paidTable map[uint32][]string

func hashFnv(tokens []string) uint32 {
	hash := fnv.New32a()

	for _, v := range tokens {
		hash.Write([]byte(v))
		hash.Write([]byte(":"))
	}

	return hash.Sum32()
}

func setupPaid(tokens []string) int8 {
	/*
		pick the incoming paid, url and volume name. Compute the hash corresponding to the new dataset
		lookup on the database table, and if there is a hit, read the entire row or any of the generic field which can hold
		the computed hash.
		Once you retrieve the hashVal in the datastore, check against the newly computed hash if they matches then proceed
		with the existing data present in the database
		If the retrieve hash is not matching with the computed one, continue to learn content bundle from the cserver
		and update the database tables

		Note: Consider coherency issues

		This can be a one time stuff where you are performing the setup request. What other cases we are considering the setup
		routine??
	*/
	lVal := hashFnv(tokens)
	if _, ok := paidTable[lVal]; ok {
		return -1
	}

	if paidTable == nil {
		paidTable = make(map[uint32][]string)
	}
	paidTable[lVal] = tokens

	return 0
}
