package pkpass

import (
	"archive/zip"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Pass 代表一个 Apple Wallet 通行证
type Pass struct {
	Data       map[string]interface{}
	Images     map[string][]byte
	LocaleData map[string]map[string]interface{}
	Signer     *Signer
}

// Signer 包含签名所需的证书和私钥
type Signer struct {
	WWDRCert   *x509.Certificate
	PassCert   *x509.Certificate
	PrivateKey *rsa.PrivateKey
}

// NewPass 创建一个新的通行证
func NewPass() *Pass {
	return &Pass{
		Data:       make(map[string]interface{}),
		Images:     make(map[string][]byte),
		LocaleData: make(map[string]map[string]interface{}),
	}
}

// SetField 设置通行证的字段
func (p *Pass) SetField(key string, value interface{}) {
	p.Data[key] = value
}

// AddImage 添加图像资源
func (p *Pass) AddImage(name string, data []byte) {
	p.Images[name] = data
}

// AddLocalization 添加本地化数据
func (p *Pass) AddLocalization(locale string, data map[string]interface{}) {
	if p.LocaleData[locale] == nil {
		p.LocaleData[locale] = make(map[string]interface{})
	}
	for k, v := range data {
		p.LocaleData[locale][k] = v
	}
}

// SetSigner 设置签名者
func (p *Pass) SetSigner(signer *Signer) {
	p.Signer = signer
}

// CreateManifest 创建清单文件
func (p *Pass) CreateManifest() (map[string]string, error) {
	manifest := make(map[string]string)

	// 添加 pass.json 到清单
	passJSON, err := json.MarshalIndent(p.Data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化通行证数据失败: %v", err)
	}

	hash := sha1.Sum(passJSON)
	manifest["pass.json"] = base64.StdEncoding.EncodeToString(hash[:])

	// 添加图像到清单
	for name, data := range p.Images {
		hash := sha1.Sum(data)
		manifest[name] = base64.StdEncoding.EncodeToString(hash[:])
	}

	// 添加本地化数据到清单
	for _, data := range p.LocaleData {
		for name, content := range data {
			contentJSON, err := json.MarshalIndent(content, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("序列化本地化数据失败: %v", err)
			}

			hash := sha1.Sum(contentJSON)
			manifest[name] = base64.StdEncoding.EncodeToString(hash[:])
		}
	}

	return manifest, nil
}

// CreateSignature 创建签名
func (p *Pass) CreateSignature(manifest map[string]string) ([]byte, error) {
	if p.Signer == nil {
		return nil, fmt.Errorf("未设置签名者")
	}

	// 序列化清单
	manifestJSON, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化清单失败: %v", err)
	}

	// 计算清单的哈希值
	hashed := sha1.Sum(manifestJSON)

	// 使用私钥签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, p.Signer.PrivateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("签名失败: %v", err)
	}

	// 创建 CMS 签名
	signedData, err := createCMS(signature, p.Signer.PassCert, p.Signer.WWDRCert)
	if err != nil {
		return nil, fmt.Errorf("创建 CMS 签名失败: %v", err)
	}

	return signedData, nil
}

// Save 保存通行证到文件
func (p *Pass) Save(filename string) error {
	// 创建临时目录
	tempDir, err := ioutil.TempDir("", "pkpass")
	if err != nil {
		return fmt.Errorf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 保存 pass.json
	passJSONPath := filepath.Join(tempDir, "pass.json")
	passJSON, err := json.MarshalIndent(p.Data, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化通行证数据失败: %v", err)
	}
	if err := ioutil.WriteFile(passJSONPath, passJSON, 0644); err != nil {
		return fmt.Errorf("保存 pass.json 失败: %v", err)
	}

	// 保存图像
	for name, data := range p.Images {
		imagePath := filepath.Join(tempDir, name)
		if err := ioutil.WriteFile(imagePath, data, 0644); err != nil {
			return fmt.Errorf("保存图像 %s 失败: %v", name, err)
		}
	}

	// 保存本地化数据
	for locale, data := range p.LocaleData {
		localeDir := filepath.Join(tempDir, locale)
		if err := os.MkdirAll(localeDir, 0755); err != nil {
			return fmt.Errorf("创建本地化目录失败: %v", err)
		}

		for name, content := range data {
			contentJSON, err := json.MarshalIndent(content, "", "  ")
			if err != nil {
				return fmt.Errorf("序列化本地化数据失败: %v", err)
			}

			contentPath := filepath.Join(localeDir, name)
			if err := ioutil.WriteFile(contentPath, contentJSON, 0644); err != nil {
				return fmt.Errorf("保存本地化数据失败: %v", err)
			}
		}
	}

	// 创建清单
	manifest, err := p.CreateManifest()
	if err != nil {
		return fmt.Errorf("创建清单失败: %v", err)
	}

	// 保存清单
	manifestJSON, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化清单失败: %v", err)
	}

	manifestPath := filepath.Join(tempDir, "manifest.json")
	if err := ioutil.WriteFile(manifestPath, manifestJSON, 0644); err != nil {
		return fmt.Errorf("保存清单失败: %v", err)
	}

	// 创建签名
	signature, err := p.CreateSignature(manifest)
	if err != nil {
		return fmt.Errorf("创建签名失败: %v", err)
	}

	// 保存签名
	signaturePath := filepath.Join(tempDir, "signature")
	if err := ioutil.WriteFile(signaturePath, signature, 0644); err != nil {
		return fmt.Errorf("保存签名失败: %v", err)
	}

	// 创建 ZIP 文件
	if err := createZipFile(tempDir, filename); err != nil {
		return fmt.Errorf("创建 ZIP 文件失败: %v", err)
	}

	return nil
}

// createZipFile 创建 ZIP 文件
func createZipFile(sourceDir, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	zipWriter := zip.NewWriter(output)
	defer zipWriter.Close()

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipFile, file)
		return err
	})
}

func createCMS(signature []byte, passCert, wwdrCert *x509.Certificate) ([]byte, error) {
	// 创建一个临时目录保存中间文件
	tempDir, err := ioutil.TempDir("", "cms")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	// 保存签名内容到文件
	contentFile := filepath.Join(tempDir, "content.bin")
	if err := ioutil.WriteFile(contentFile, signature, 0600); err != nil {
		return nil, err
	}

	// 保存证书到 PEM 文件
	passCertFile := filepath.Join(tempDir, "pass.crt")
	if err := ioutil.WriteFile(passCertFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: passCert.Raw}), 0600); err != nil {
		return nil, err
	}

	wwdrCertFile := filepath.Join(tempDir, "wwdr.crt")
	if err := ioutil.WriteFile(wwdrCertFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: wwdrCert.Raw}), 0600); err != nil {
		return nil, err
	}

	// 使用 openssl 命令创建 PKCS7 签名
	cmd := exec.Command("openssl", "smime", "-sign",
		"-in", contentFile,
		"-inkey", "/Users/gaoyang/go/src/APT11_CRS/certificates/passkey.pem",
		"-signer", passCertFile,
		"-certfile", wwdrCertFile,
		"-outform", "DER")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("openssl 签名失败: %v", err)
	}

	return out.Bytes(), nil
}

// LoadSigner 从文件加载签名者
func LoadSigner(passCertPath, privateKeyPath, wwdrCertPath string) (*Signer, error) {
	// 加载 Apple Worldwide Developer Relations 证书
	wwdrCertData, err := ioutil.ReadFile(wwdrCertPath)
	if err != nil {
		return nil, fmt.Errorf("读取 WWDR 证书失败: %v", err)
	}

	wwdrBlock, _ := pem.Decode(wwdrCertData)
	if wwdrBlock == nil {
		return nil, fmt.Errorf("无法解码 WWDR 证书")
	}

	wwdrCert, err := x509.ParseCertificate(wwdrBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析 WWDR 证书失败: %v", err)
	}

	// 加载通行证证书
	passCertData, err := ioutil.ReadFile(passCertPath)
	if err != nil {
		return nil, fmt.Errorf("读取通行证证书失败: %v", err)
	}

	passBlock, _ := pem.Decode(passCertData)
	if passBlock == nil {
		return nil, fmt.Errorf("无法解码通行证证书")
	}

	passCert, err := x509.ParseCertificate(passBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析通行证证书失败: %v", err)
	}

	// 加载私钥
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("读取私钥失败: %v", err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyData)
	if privateKeyBlock == nil {
		return nil, fmt.Errorf("无法解码私钥")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析私钥失败: %v", err)
		}
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("私钥不是 RSA 类型")
	}

	return &Signer{
		WWDRCert:   wwdrCert,
		PassCert:   passCert,
		PrivateKey: rsaPrivateKey,
	}, nil
}
