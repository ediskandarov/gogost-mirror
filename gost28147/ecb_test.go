// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2022 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost28147

import (
	"bytes"
	"crypto/cipher"
	"testing"
)

func TestECBGCL3Vectors(t *testing.T) {
	key := []byte{
		0x04, 0x75, 0xf6, 0xe0, 0x50, 0x38, 0xfb, 0xfa,
		0xd2, 0xc7, 0xc3, 0x90, 0xed, 0xb3, 0xca, 0x3d,
		0x15, 0x47, 0x12, 0x42, 0x91, 0xae, 0x1e, 0x8a,
		0x2f, 0x79, 0xcd, 0x9e, 0xd2, 0xbc, 0xef, 0xbd,
	}
	c := NewCipher(key, &SboxIdGost2814789TestParamSet)
	plaintext := []byte{
		0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00,
		0x0f, 0x0e, 0x0d, 0x0c, 0x0b, 0x0a, 0x09, 0x08,
		0x17, 0x16, 0x15, 0x14, 0x13, 0x12, 0x11, 0x10,
		0x1f, 0x1e, 0x1d, 0x1c, 0x1b, 0x1a, 0x19, 0x18,
		0x27, 0x26, 0x25, 0x24, 0x23, 0x22, 0x21, 0x20,
		0x2f, 0x2e, 0x2d, 0x2c, 0x2b, 0x2a, 0x29, 0x28,
		0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x30,
		0x3f, 0x3e, 0x3d, 0x3c, 0x3b, 0x3a, 0x39, 0x38,
		0x47, 0x46, 0x45, 0x44, 0x43, 0x42, 0x41, 0x40,
		0x4f, 0x4e, 0x4d, 0x4c, 0x4b, 0x4a, 0x49, 0x48,
		0x57, 0x56, 0x55, 0x54, 0x53, 0x52, 0x51, 0x50,
		0x5f, 0x5e, 0x5d, 0x5c, 0x5b, 0x5a, 0x59, 0x58,
		0x67, 0x66, 0x65, 0x64, 0x63, 0x62, 0x61, 0x60,
		0x6f, 0x6e, 0x6d, 0x6c, 0x6b, 0x6a, 0x69, 0x68,
		0x77, 0x76, 0x75, 0x74, 0x73, 0x72, 0x71, 0x70,
		0x7f, 0x7e, 0x7d, 0x7c, 0x7b, 0x7a, 0x79, 0x78,
		0x87, 0x86, 0x85, 0x84, 0x83, 0x82, 0x81, 0x80,
		0x8f, 0x8e, 0x8d, 0x8c, 0x8b, 0x8a, 0x89, 0x88,
		0x97, 0x96, 0x95, 0x94, 0x93, 0x92, 0x91, 0x90,
		0x9f, 0x9e, 0x9d, 0x9c, 0x9b, 0x9a, 0x99, 0x98,
		0xa7, 0xa6, 0xa5, 0xa4, 0xa3, 0xa2, 0xa1, 0xa0,
		0xaf, 0xae, 0xad, 0xac, 0xab, 0xaa, 0xa9, 0xa8,
		0xb7, 0xb6, 0xb5, 0xb4, 0xb3, 0xb2, 0xb1, 0xb0,
		0xbf, 0xbe, 0xbd, 0xbc, 0xbb, 0xba, 0xb9, 0xb8,
		0xc7, 0xc6, 0xc5, 0xc4, 0xc3, 0xc2, 0xc1, 0xc0,
		0xcf, 0xce, 0xcd, 0xcc, 0xcb, 0xca, 0xc9, 0xc8,
		0xd7, 0xd6, 0xd5, 0xd4, 0xd3, 0xd2, 0xd1, 0xd0,
		0xdf, 0xde, 0xdd, 0xdc, 0xdb, 0xda, 0xd9, 0xd8,
		0xe7, 0xe6, 0xe5, 0xe4, 0xe3, 0xe2, 0xe1, 0xe0,
		0xef, 0xee, 0xed, 0xec, 0xeb, 0xea, 0xe9, 0xe8,
		0xf7, 0xf6, 0xf5, 0xf4, 0xf3, 0xf2, 0xf1, 0xf0,
		0xff, 0xfe, 0xfd, 0xfc, 0xfb, 0xfa, 0xf9, 0xf8,
	}
	ciphertext := []byte{
		0x4b, 0x8c, 0x4c, 0x98, 0x15, 0xf2, 0x4a, 0xea,
		0x1e, 0xc3, 0x57, 0x09, 0xb3, 0xbc, 0x2e, 0xd1,
		0xe0, 0xd1, 0xf2, 0x22, 0x65, 0x2d, 0x59, 0x18,
		0xf7, 0xdf, 0xfc, 0x80, 0x4b, 0xde, 0x5c, 0x68,
		0x46, 0x53, 0x75, 0x53, 0xa7, 0x46, 0x0d, 0xec,
		0x05, 0x1f, 0x1b, 0xd3, 0x0a, 0x63, 0x1a, 0xb7,
		0x78, 0xc4, 0x43, 0xe0, 0x5d, 0x3e, 0xa4, 0x0e,
		0x2d, 0x7e, 0x23, 0xa9, 0x1b, 0xc9, 0x02, 0xbc,
		0x21, 0x0c, 0x84, 0xcb, 0x0d, 0x0a, 0x07, 0xc8,
		0x7b, 0xd0, 0xfb, 0xb5, 0x1a, 0x14, 0x04, 0x5c,
		0xa2, 0x53, 0x97, 0x71, 0x2e, 0x5c, 0xc2, 0x8f,
		0x39, 0x3f, 0x6f, 0x52, 0xf2, 0x30, 0x26, 0x4e,
		0x8c, 0xe0, 0xd1, 0x01, 0x75, 0x6d, 0xdc, 0xd3,
		0x03, 0x79, 0x1e, 0xca, 0xd5, 0xc1, 0x0e, 0x12,
		0x53, 0x0a, 0x78, 0xe2, 0x0a, 0xb1, 0x1c, 0xea,
		0x3a, 0xf8, 0x55, 0xb9, 0x7c, 0xe1, 0x0b, 0xba,
		0xa0, 0xc8, 0x96, 0xeb, 0x50, 0x5a, 0xd3, 0x60,
		0x43, 0xa3, 0x0f, 0x98, 0xdb, 0xd9, 0x50, 0x6d,
		0x63, 0x91, 0xaf, 0x01, 0x40, 0xe9, 0x75, 0x5a,
		0x46, 0x5c, 0x1f, 0x19, 0x4a, 0x0b, 0x89, 0x9b,
		0xc4, 0xf6, 0xf8, 0xf5, 0x2f, 0x87, 0x3f, 0xfa,
		0x26, 0xd4, 0xf8, 0x25, 0xba, 0x1f, 0x98, 0x82,
		0xfc, 0x26, 0xaf, 0x2d, 0xc0, 0xf9, 0xc4, 0x58,
		0x49, 0xfa, 0x09, 0x80, 0x02, 0x62, 0xa4, 0x34,
		0x2d, 0xcb, 0x5a, 0x6b, 0xab, 0x61, 0x5d, 0x08,
		0xd4, 0x26, 0xe0, 0x08, 0x13, 0xd6, 0x2e, 0x02,
		0x2a, 0x37, 0xe8, 0xd0, 0xcf, 0x36, 0xf1, 0xc7,
		0xc0, 0x3f, 0x9b, 0x21, 0x60, 0xbd, 0x29, 0x2d,
		0x2e, 0x01, 0x48, 0x4e, 0xf8, 0x8f, 0x20, 0x16,
		0x8a, 0xbf, 0x82, 0xdc, 0x32, 0x7a, 0xa3, 0x18,
		0x69, 0xd1, 0x50, 0x59, 0x31, 0x91, 0xf2, 0x6c,
		0x5a, 0x5f, 0xca, 0x58, 0x9a, 0xb2, 0x2d, 0xb2,
	}
	e := c.NewECBEncrypter()
	tmp := make([]byte, len(plaintext))
	e.CryptBlocks(tmp, plaintext)
	if bytes.Compare(tmp, ciphertext) != 0 {
		t.Fatal("encryption failed")
	}
	d := c.NewECBDecrypter()
	d.CryptBlocks(tmp, tmp)
	if bytes.Compare(tmp, plaintext) != 0 {
		t.Fatal("decryption failed")
	}
}

// Crypto++ 5.6.2 test vectors
func TestECBCryptoPPVectors(t *testing.T) {
	sbox := &SboxAppliedCryptographyParamSet
	tmp := make([]byte, BlockSize)
	var key []byte
	var pt []byte
	var ct []byte
	var c *Cipher

	t.Run("1", func(t *testing.T) {
		key = []byte{
			0xBE, 0x5E, 0xC2, 0x00, 0x6C, 0xFF, 0x9D, 0xCF,
			0x52, 0x35, 0x49, 0x59, 0xF1, 0xFF, 0x0C, 0xBF,
			0xE9, 0x50, 0x61, 0xB5, 0xA6, 0x48, 0xC1, 0x03,
			0x87, 0x06, 0x9C, 0x25, 0x99, 0x7C, 0x06, 0x72,
		}
		pt = []byte{0x0D, 0xF8, 0x28, 0x02, 0xB7, 0x41, 0xA2, 0x92}
		ct = []byte{0x07, 0xF9, 0x02, 0x7D, 0xF7, 0xF7, 0xDF, 0x89}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("2", func(t *testing.T) {
		key = []byte{
			0xB3, 0x85, 0x27, 0x2A, 0xC8, 0xD7, 0x2A, 0x5A,
			0x8B, 0x34, 0x4B, 0xC8, 0x03, 0x63, 0xAC, 0x4D,
			0x09, 0xBF, 0x58, 0xF4, 0x1F, 0x54, 0x06, 0x24,
			0xCB, 0xCB, 0x8F, 0xDC, 0xF5, 0x53, 0x07, 0xD7,
		}
		pt = []byte{0x13, 0x54, 0xEE, 0x9C, 0x0A, 0x11, 0xCD, 0x4C}
		ct = []byte{0x4F, 0xB5, 0x05, 0x36, 0xF9, 0x60, 0xA7, 0xB1}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("3", func(t *testing.T) {
		key = []byte{
			0xAE, 0xE0, 0x2F, 0x60, 0x9A, 0x35, 0x66, 0x0E,
			0x40, 0x97, 0xE5, 0x46, 0xFD, 0x30, 0x26, 0xB0,
			0x32, 0xCD, 0x10, 0x7C, 0x7D, 0x45, 0x99, 0x77,
			0xAD, 0xF4, 0x89, 0xBE, 0xF2, 0x65, 0x22, 0x62,
		}
		pt = []byte{0x66, 0x93, 0xD4, 0x92, 0xC4, 0xB0, 0xCC, 0x39}
		ct = []byte{0x67, 0x00, 0x34, 0xAC, 0x0F, 0xA8, 0x11, 0xB5}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("4", func(t *testing.T) {
		key = []byte{
			0x32, 0x0E, 0x9D, 0x84, 0x22, 0x16, 0x5D, 0x58,
			0x91, 0x1D, 0xFC, 0x7D, 0x8B, 0xBB, 0x1F, 0x81,
			0xB0, 0xEC, 0xD9, 0x24, 0x02, 0x3B, 0xF9, 0x4D,
			0x9D, 0xF7, 0xDC, 0xF7, 0x80, 0x12, 0x40, 0xE0,
		}
		pt = []byte{0x99, 0xE2, 0xD1, 0x30, 0x80, 0x92, 0x8D, 0x79}
		ct = []byte{0x81, 0x18, 0xFF, 0x9D, 0x3B, 0x3C, 0xFE, 0x7D}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("5", func(t *testing.T) {
		key = []byte{
			0xC9, 0xF7, 0x03, 0xBB, 0xBF, 0xC6, 0x36, 0x91,
			0xBF, 0xA3, 0xB7, 0xB8, 0x7E, 0xA8, 0xFD, 0x5E,
			0x8E, 0x8E, 0xF3, 0x84, 0xEF, 0x73, 0x3F, 0x1A,
			0x61, 0xAE, 0xF6, 0x8C, 0x8F, 0xFA, 0x26, 0x5F,
		}
		pt = []byte{0xD1, 0xE7, 0x87, 0x74, 0x9C, 0x72, 0x81, 0x4C}
		ct = []byte{0xA0, 0x83, 0x82, 0x6A, 0x79, 0x0D, 0x3E, 0x0C}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("6", func(t *testing.T) {
		key = []byte{
			0x72, 0x8F, 0xEE, 0x32, 0xF0, 0x4B, 0x4C, 0x65,
			0x4A, 0xD7, 0xF6, 0x07, 0xD7, 0x1C, 0x66, 0x0C,
			0x2C, 0x26, 0x70, 0xD7, 0xC9, 0x99, 0x71, 0x32,
			0x33, 0x14, 0x9A, 0x1C, 0x0C, 0x17, 0xA1, 0xF0,
		}
		pt = []byte{0xD4, 0xC0, 0x53, 0x23, 0xA4, 0xF7, 0xA7, 0xB5}
		ct = []byte{0x4D, 0x1F, 0x2E, 0x6B, 0x0D, 0x9D, 0xE2, 0xCE}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("7", func(t *testing.T) {
		key = []byte{
			0x35, 0xFC, 0x96, 0x40, 0x22, 0x09, 0x50, 0x0F,
			0xCF, 0xDE, 0xF5, 0x35, 0x2D, 0x1A, 0xBB, 0x03,
			0x8F, 0xE3, 0x3F, 0xC0, 0xD9, 0xD5, 0x85, 0x12,
			0xE5, 0x63, 0x70, 0xB2, 0x2B, 0xAA, 0x13, 0x3B,
		}
		pt = []byte{0x87, 0x42, 0xD9, 0xA0, 0x5F, 0x6A, 0x3A, 0xF6}
		ct = []byte{0x2F, 0x3B, 0xB8, 0x48, 0x79, 0xD1, 0x1E, 0x52}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})

	t.Run("8", func(t *testing.T) {
		key = []byte{
			0xD4, 0x16, 0xF6, 0x30, 0xBE, 0x65, 0xB7, 0xFE,
			0x15, 0x06, 0x56, 0x18, 0x33, 0x70, 0xE0, 0x70,
			0x18, 0x23, 0x4E, 0xE5, 0xDA, 0x3D, 0x89, 0xC4,
			0xCE, 0x91, 0x52, 0xA0, 0x3E, 0x5B, 0xFB, 0x77,
		}
		pt = []byte{0xF8, 0x65, 0x06, 0xDA, 0x04, 0xE4, 0x1C, 0xB8}
		ct = []byte{0x96, 0xF0, 0xA5, 0xC7, 0x7A, 0x04, 0xF5, 0xCE}
		c = NewCipher(key, sbox)
		c.Encrypt(tmp, pt)
		if bytes.Compare(tmp, ct) != 0 {
			t.FailNow()
		}
	})
}

// http://cryptomanager.com/tv.html test vectors.
func TestECBCryptomanager(t *testing.T) {
	key := []byte{
		0x75, 0x71, 0x31, 0x34, 0xB6, 0x0F, 0xEC, 0x45,
		0xA6, 0x07, 0xBB, 0x83, 0xAA, 0x37, 0x46, 0xAF,
		0x4F, 0xF9, 0x9D, 0xA6, 0xD1, 0xB5, 0x3B, 0x5B,
		0x1B, 0x40, 0x2A, 0x1B, 0xAA, 0x03, 0x0D, 0x1B,
	}
	c := NewCipher(key, &SboxIdGostR341194TestParamSet)
	tmp := make([]byte, BlockSize)
	c.Encrypt(tmp, []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88})
	if bytes.Compare(tmp, []byte{0x03, 0x25, 0x1E, 0x14, 0xF9, 0xD2, 0x8A, 0xCB}) != 0 {
		t.FailNow()
	}
}

func TestECBInterface(t *testing.T) {
	var key [KeySize]byte
	c := NewCipher(key[:], SboxDefault)
	var _ cipher.BlockMode = c.NewECBEncrypter()
	var _ cipher.BlockMode = c.NewECBDecrypter()
}
