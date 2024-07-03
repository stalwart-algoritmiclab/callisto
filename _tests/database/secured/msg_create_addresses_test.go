/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"

	d "github.com/stalwart-algoritmiclab/callisto/_tests/database"
)

func TestRepository_InsertMsgCreateAddresses(t *testing.T) {
	type args struct {
		msg  []*types.MsgCreateAddresses
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertToCreateAddresses single",
			args: args{
				msg: []*types.MsgCreateAddresses{
					{
						Address: []string{"stwart14pduuu5szx6dp9zwwj33fwkevtw638pnrguh5z"},
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
		{
			name: "[success] InsertToCreateAddresses multiple",
			args: args{
				msg: []*types.MsgCreateAddresses{
					{
						Address: []string{
							"stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							"stwart1nkga4azkfczlysh8re8xg3lnlrq88435dwpx9n",
						},
						Creator: d.TestAddressCreator,
					},
					{
						Address: []string{
							"stwart1y3n5h0r5nwvw6n6dfk0xn3xxh4thph9nda6t98",
						},
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.Datastore.Secured.InsertMsgCreateAddresses(tt.args.hash, tt.args.msg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertToCreateAddresses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
