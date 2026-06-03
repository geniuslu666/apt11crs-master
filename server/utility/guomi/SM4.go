package guomi

import (
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tjfoc/gmsm/sm4"
)

type Sm4Secret struct {
	SM4Key []byte
	SM4Iv  []byte
}

func NewSM4(Key []byte, iv []byte) *Sm4Secret {
	return &Sm4Secret{
		SM4Key: Key,
		SM4Iv:  iv,
	}
}

func (Sm4Secrets *Sm4Secret) SM4En(plaintext []byte) ([]byte, error) {
	var (
		ecbMsgByte []byte
		err        error
	)
	if err = sm4.SetIV(Sm4Secrets.SM4Iv); err != nil {
		goto ERR
	}
	if ecbMsgByte, err = sm4.Sm4Cbc(Sm4Secrets.SM4Key, plaintext, true); err != nil {
		goto ERR
	}
	ecbMsgByte = HexEncode(string(ecbMsgByte))
	return ecbMsgByte, err
ERR:
	return ecbMsgByte, gerror.New("加密失败")
}

func (Sm4Secrets *Sm4Secret) SM4De(str []byte) ([]byte, error) {
	var (
		ecbMsg []byte
		err    error
	)
	if ecbMsg, err = HexDecode(string(str)); err != nil {
		return ecbMsg, err
	}
	if err = sm4.SetIV(Sm4Secrets.SM4Iv); err != nil {
		goto ERR
	}
	if ecbMsg, err = sm4.Sm4Cbc(Sm4Secrets.SM4Key, ecbMsg, false); err != nil {
		goto ERR
	}
	return ecbMsg, err
ERR:
	return ecbMsg, err
}

// HexDecode 16进制解码
func HexDecode(s string) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst, []byte(s))        //进制转换, src->dst
	if err != nil {
		g.Dump(err)
		err = gerror.New("解析失败")
		return nil, err
	}
	return dst[:n], err //返回0:n的数据.
}

// HexEncode 字符串转为16进制
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.EncodedLen
	n := hex.Encode(dst, []byte(s))             //字节流转化成16进制
	return dst[:n]
}
