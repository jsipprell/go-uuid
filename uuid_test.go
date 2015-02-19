/* The MIT License (MIT)
 *
 * Copyright (c) 2015 Jesse Sipprell <jessesipprell@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Test hardness for uuid package

package uuid

import "testing"
import "strings"

func TestZeroUUID(t *testing.T) {
	var u UUID

	if !IsZero(u) {
		t.Log("IsZero() failed on default nil UUID")
		t.Fail()
	}

	u = uuid{}
	if !IsZero(u) || !u.IsZero() {
		t.Log("IsZero() failed on static initialized uuid")
		t.Fail()
	}

	t.Logf("zero uuid string = '%v'", u.String())
	if u.String() != "" {
		t.Fail()
	}
	t.Logf("zero uuid bytes = %v", u.Bytes())

	if u.Bytes() != nil {
		t.Fail()
	}
}

func testuuidHelper(t *testing.T, u UUID) {
	s := u.String()
	u2, err := Decode(s)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Logf("%v == %v", u, u2)
	}
	s2 := strings.ToUpper(u2.String())
	u2, err = Decode(s2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if !u.Equals(u2) || !u2.Equals(u) {
		t.Logf("%v != %v", s, s2)
		t.Fail()
	} else {
		t.Logf("%v == %v", s, s2)
	}
}

func TestRandomUUID(t *testing.T) {
	u := NewRandom()

	t.Logf("psuedo-random uuid = %v", u)
	testuuidHelper(t, u)
}

func TestSeqUUID(t *testing.T) {
	for i := 0; i < 12; i++ {
		testuuidHelper(t, New())
	}
}
