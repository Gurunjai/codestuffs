package main

import (
	"fmt"
	"testing"
)

func TestHashFnv(t *testing.T) {
	data := [][]string{
		{
			"providerabc::assetabc",
			"http://127.0.0.1/xyzwe/provider,asset_vs_asset.m3u",
			"IPCDN",
		},
		{
			"providerblah::assetblah",
			"http://ipcdn/abcdef/xyz_test/123,987.m3u",
			"IPCDN",
		},
		{
			"providerxyz::assetxyz::index",
			"http://hlscdn.com/xyzqaesda/hksjksdfj@$_vs_sampletest.m3u8",
			"hls",
		},
		{
			"providerabc::assetabc",
			"http://127.0.0.1/xyzwe/provider,asset_vs_asset.m3u",
			"IPCDN",
		},
		{
			"providerabc::assetabc",
			"http://127.0.0.1/xyzwe/provider,asset_vs_asset2provider.m3u",
			"IPCDN",
		},
		{
			"providerblah::testasset",
			"http://ipcdn/abcdef/xyz_test/123,987.m3u",
			"IPCDN",
		},
		{
			"providerblah::assetblah",
			"http://ipcdn/abcdef/xyz_test/123,987.m3u",
			"hls",
		},
	}

	outVal := []int8{0, 0, 0, -1, 0, 0, 0}

	for i, v := range data {
		fmt.Printf("Data At Index(%v): %v\n\t hash: %v, Type: %[3]T\n", i+1, v, hashFnv(v))
		got := setupPaid(v)
		if got != outVal[i] {
			t.Errorf("Mismatch from setup process:\n\tGot: %v\n\tWant: %v\n", got, outVal[i])
		}
	}
}
