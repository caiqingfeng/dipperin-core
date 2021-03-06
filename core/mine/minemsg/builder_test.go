// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package minemsg

import (
	"github.com/dipperin/dipperin-core/common"
	"github.com/dipperin/dipperin-core/tests/factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeDefaultWorkBuilder(t *testing.T) {
	assert.NotNil(t, MakeDefaultWorkBuilder())
}

func TestDefaultWorkBuilder_BuildWorks(t *testing.T) {
	diff := common.HexToDiff("0x1effffff")
	block := factory.CreateBlock2(diff, 1)
	workBuilder := MakeDefaultWorkBuilder()
	code, works := workBuilder.BuildWorks(block, 10)
	assert.Equal(t, code, 16)
	assert.Equal(t, len(works), 10)

	code, works = workBuilder.BuildWorks(nil, 10)

	assert.Equal(t, code, 0)
	assert.Equal(t, len(works), 0)
}
