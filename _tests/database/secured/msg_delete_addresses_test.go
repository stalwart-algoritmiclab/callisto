/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"

	d "github.com/stalwart-algoritmiclab/callisto/_tests/database"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

func TestRepository_InsertMsgDeleteAddresses(t *testing.T) {
	type args struct {
		msg  []*secured.MsgDeleteAddresses
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertToDeleteAddresses single",
			args: args{
				msg: []*secured.MsgDeleteAddresses{
					{
						Id:      1,
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
		{
			name: "[success] InsertToDeleteAddresses multiple",
			args: args{
				msg: []*secured.MsgDeleteAddresses{
					{
						Id:      2,
						Creator: d.TestAddressCreator,
					},
					{
						Id:      3,
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.Datastore.Secured.InsertMsgDeleteAddresses(tt.args.hash, tt.args.msg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertToDeleteAddresses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
