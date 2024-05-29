package secured

import (
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"

	d "github.com/stalwart-algoritmiclab/callisto/_tests/database"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

func TestRepository_InsertMsgUpdateAddresses(t *testing.T) {
	type args struct {
		msg  []*secured.MsgUpdateAddresses
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertToUpdateAddresses single",
			args: args{
				msg: []*secured.MsgUpdateAddresses{
					{
						Id:      1,
						Address: []string{"stwart14pduuu5szx6dp9zwwj33fwkevtw638pnrguh5z"},
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
		{
			name: "[success] InsertToUpdateAddresses multiple",
			args: args{
				msg: []*secured.MsgUpdateAddresses{
					{
						Id: 2,
						Address: []string{
							"stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							"stwart1nkga4azkfczlysh8re8xg3lnlrq88435dwpx9n",
						},
						Creator: d.TestAddressCreator,
					},
					{
						Id: 3,
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
			err := d.Datastore.Secured.InsertMsgUpdateAddresses(tt.args.hash, tt.args.msg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertToUpdateAddresses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}