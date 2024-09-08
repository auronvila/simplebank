package token

import (
	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker_CreateToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrorExpireToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	invalidToken := "v2.local.sIgVm0es9uswZliPdyXOOi99czPbpl41KOUu45e62BvCaL5H3kHNibrbRZkM1-wW091ARzNexLY8g0GZA0-WCNsgs8GZLClEk5TJbgQjf__yExZRh2qMnqxfVr_KS9WoqKVlU-WrAG6TRUXZo43OSJQkeNBnB8Gq4rN2A8HYeA3ms20up80dgz2rpY79F9ILvPrAIzxNkDSE51vAxv50BTShuel3F3hXgReHsDv2PJCnMBnMyE_AfePxJ6WJ1obXSIUpSsOQX6wjwdQdOIcXZ853c-NPYMVU-abXJhhLVvvHyNZPi1wcEvjt.eyJraWQiOiAiMTIzNDUifQ"
	payload, err := maker.VerifyToken(invalidToken)
	require.Error(t, err)
	require.EqualError(t, err, "invalid token authentication")
	require.Nil(t, payload)
}
