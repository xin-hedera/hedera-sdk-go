package hedera

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestSerializeContractCreateTransaction(t *testing.T) {
	mockClient, err := newMockClient()
	assert.NoError(t, err)

	privateKey, err := Ed25519PrivateKeyFromString(mockPrivateKey)
	assert.NoError(t, err)

	tx, err := NewContractCreateTransaction().
		SetAdminKey(privateKey.PublicKey()).
		SetInitialBalance(HbarFromTinybar(1e3)).
		SetBytecodeFileID(FileID{File: 4}).
		SetGas(100).
		SetProxyAccountID(AccountID{Account: 3}).
		SetAutoRenewPeriod(60 * 60 * 24 * 14 * time.Second).
		SetMaxTransactionFee(HbarFromTinybar(1e6)).
		SetTransactionID(testTransactionID).
		Build(mockClient)

	assert.NoError(t, err)

	tx.Sign(privateKey)

	assert.Equal(t, `bodyBytes:"\n\016\n\010\010\334\311\007\020\333\237\t\022\002\030\003\022\002\030\003\030\300\204=\"\002\010xB7\n\002\030\004\032\"\022\344\361\300\353L}\315\303\347\353\021p\263\010\212=\022\242\227\364\243\353\342\362\205\003\375g5F\355\216d(\350\0072\002\030\003B\004\010\200\352I"sigMap:<sigPair:<pubKeyPrefix:"\344\361\300\353L}\315\303\347\353\021p\263\010\212=\022\242\227\364\243\353\342\362\205\003\375g5F\355\216"ed25519:"\335\035\364\306\364\013\204\232\376\343\320b\270\310\017I\020\223\306K\316\264\277\363x\251\3418\004>\352\272\264\202\343\260\344\264\344<l\004H\200\253m\306\305)\304t\245\227z\305\267\270w\353\326\255\340\374\005">>transactionID:<transactionValidStart:<seconds:124124nanos:151515>accountID:<accountNum:3>>nodeAccountID:<accountNum:3>transactionFee:1000000transactionValidDuration:<seconds:120>contractCreateInstance:<fileID:<fileNum:4>adminKey:<ed25519:"\344\361\300\353L}\315\303\347\353\021p\263\010\212=\022\242\227\364\243\353\342\362\205\003\375g5F\355\216">gas:100initialBalance:1000proxyAccountID:<accountNum:3>autoRenewPeriod:<seconds:1209600>>`, strings.ReplaceAll(strings.ReplaceAll(tx.String(), " ", ""), "\n", ""))
}
