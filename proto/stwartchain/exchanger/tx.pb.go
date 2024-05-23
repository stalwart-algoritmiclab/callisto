// source: stwartchain/exchanger/tx.proto

package exchanger

import (
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
)


// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
}

func (*MsgUpdateParams) ProtoMessage() {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
}

	}
	return ""
}

	}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (*MsgUpdateParamsResponse) ProtoMessage() {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
}

type MsgExchange struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount  string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	DenomTo string `protobuf:"bytes,4,opt,name=denomTo,proto3" json:"denomTo,omitempty"`
}

func (*MsgExchange) ProtoMessage() {}
func (*MsgExchange) Descriptor() ([]byte, []int) {
}

	}
	return ""
}

	}
	return ""
}

	}
	return ""
}

	}
	return ""
}

type MsgExchangeResponse struct {
}

func (*MsgExchangeResponse) ProtoMessage() {}
func (*MsgExchangeResponse) Descriptor() ([]byte, []int) {
}


}


}

}
}

	}
			case 1:
			case 2:
			default:
				return nil
			}
		}
			case 1:
			case 2:
			default:
				return nil
			}
		}
			default:
				return nil
			}
		}
			case 0:
			case 1:
			case 2:
			default:
			}
		}
	}
}
