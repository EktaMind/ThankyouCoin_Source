package launcher

import (
	"github.com/EktaMind/Thank_ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://f95e8f28f5fb0d61469ebd3d3cea8f5e3deee69d939723eedcde8c8b548a276139a60a3134e803cd5019d96116c7a59ec9c1652e81aba02db1a0c834611eac78@178.62.69.60:5050",
		"enode://ae6370f028f72161ba17449575078876351a82b2ab0e1a33f5a48d73bd5213646486f5ce3fe05f8c1dfdeefcfa72a85e76f8bc597134b5abc9c5e8c3f6e28482@167.99.199.7:5050",
		"enode://c2053f1b5d3ab0342a6b6a8ae88f62aa43dab3a4da5fbadb7318957931db2a4837b428585d7af79dd030103413012625b5371c51bbf0a4bdc53834bdc9ff97ff@142.93.33.192:5050",
		"enode://be401f4f653b57eef0ef2020e0c348ead56a6c6297d96ddef6613e0e9ae1bfc0188f867c8046fee5620540c773141df0ff4c7e995ba5ae0bbbf53bf2b6f6734c@159.65.52.165:5050",
		"enode://a6f0a6f09e3428b2ec2e4587006fc540581fcfa9fc8ddb84b7203d3e91b1928567f3031ac6ba53379a2f848e591fbe61ede03dfab25ea6f31ad8459aacf5755a@46.101.28.195:5050",
		"enode://cb13ff7e0e51ff55346c9c828921025c5f1f53694e6faec01132c4bb2f8da992def466ba2d6bc8b8ad39a0952741ce65f1029cb1e42d4dee5c3d07aba5518705@142.93.42.96:5050",
		"enode://8a86aaa0b2dd8132f18d6d6ab010c3d1814eebbee633a101ad10222040137335bba965d99d4fc7789a965d743babce14ec613905e7e26af7f89cec3d053aa74a@62.171.179.134:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
