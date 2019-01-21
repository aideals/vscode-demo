package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//定义结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希
	PrevBlockHash []byte
	//当前区块哈希，这是为了方便加入的字段，正常区块中没有这个字段
	Hash []byte
	//梅克尔根  //占时不用管
	MerkelRoot []byte
	//时间戳，从1970.1.1到至今，描述一个数字
	TimeStamp uint64
	//难度值，可以推导出哈希值
	Bits uint64
	//随机数Nonce,挖矿要求的值
	Nonce uint64
	//数据
	Data []byte
}

//创建区块
func NewBlock(data string, PrevBlockHash []byte) *Block {

	block := Block{
		Version:       00,
		PrevBlockHash: PrevBlockHash,
		Hash:          nil,
		MerkelRoot:    nil,
		TimeStamp:     uint64(time.Now().Unix()),
		Bits:          0, //随便写一个数
		Nonce:         0, //随便写一个数

		Data: []byte(data),
	}

	//设置哈希值
	block.setHash()

	return &block
}

//生成哈希，将所有的数据拼接起来，做sha256处理
func (b *Block) setHash() {
	var blockInfo []byte

	blockInfo = append(blockInfo, unit2Bytes(b.Version)...)
	blockInfo = append(blockInfo, b.PrevBlockHash...)
	blockInfo = append(blockInfo, b.Hash...)

	blockInfo = append(blockInfo, b.MerkelRoot...)
	blockInfo = append(blockInfo, unit2Bytes(b.TimeStamp)...)
	blockInfo = append(blockInfo, unit2Bytes(b.Bits)...)
	blockInfo = append(blockInfo, unit2Bytes(b.Nonce)...)
}

//将数字转字节流
func unit2Bytes(num uint64) []byte {

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)

	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}
