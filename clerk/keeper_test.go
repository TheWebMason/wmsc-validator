package clerk_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/maticnetwork/heimdall/app"
	"github.com/maticnetwork/heimdall/clerk"
	"github.com/maticnetwork/heimdall/clerk/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

//
// Test suite
//

// KeeperTestSuite integrate test suite context object
type KeeperTestSuite struct {
	suite.Suite

	app *app.HeimdallApp
	ctx sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app, suite.ctx = createTestApp(false)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

//
// Tests
//

func (suite *KeeperTestSuite) TestHasGetSetEventRecord() {
	t, app, ctx := suite.T(), suite.app, suite.ctx

	hAddr := hmTypes.BytesToHeimdallAddress([]byte("some-address"))
	hHash := hmTypes.BytesToHeimdallHash([]byte("some-address"))
	testRecord1 := types.NewEventRecord(hHash, 1, 1, hAddr, make([]byte, 0), "1")

	// SetEventRecord
	ck := app.ClerkKeeper
	err := ck.SetEventRecord(ctx, testRecord1)
	require.Nil(t, err)

	err = ck.SetEventRecord(ctx, testRecord1)
	require.NotNil(t, err)

	// GetEventRecord
	respRecord, err := ck.GetEventRecord(ctx, testRecord1.ID)
	require.Nil(t, err)
	require.Equal(t, (*respRecord).ID, testRecord1.ID)

	respRecord, err = ck.GetEventRecord(ctx, testRecord1.ID+1)
	require.NotNil(t, err)

	// HasEventRecord
	recordPresent := ck.HasEventRecord(ctx, testRecord1.ID)
	require.True(t, recordPresent)

	recordPresent = ck.HasEventRecord(ctx, testRecord1.ID+1)
	require.False(t, recordPresent)

	recordList := ck.GetAllEventRecords(ctx)
	require.Len(t, recordList, 1)
}

func (suite *KeeperTestSuite) TestGetEventRecordList() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	var i uint64

	hAddr := hmTypes.BytesToHeimdallAddress([]byte("some-address"))
	hHash := hmTypes.BytesToHeimdallHash([]byte("some-address"))
	ck := app.ClerkKeeper
	for i = 0; i < 30; i++ {
		testRecord := types.NewEventRecord(hHash, i, i, hAddr, make([]byte, 0), "1")
		ck.SetEventRecord(ctx, testRecord)
	}

	recordList, _ := ck.GetEventRecordList(ctx, 1, 20)
	require.Len(t, recordList, 20)

	recordList, _ = ck.GetEventRecordList(ctx, 2, 20)
	require.Len(t, recordList, 10)

	recordList, _ = ck.GetEventRecordList(ctx, 3, 20)
	require.Len(t, recordList, 0)

	recordList, _ = ck.GetEventRecordList(ctx, 1, 30)
	require.Len(t, recordList, 20)
}

func (suite *KeeperTestSuite) TestGetEventRecordKey() {
	t, _, _ := suite.T(), suite.app, suite.ctx

	hAddr := hmTypes.BytesToHeimdallAddress([]byte("some-address"))
	hHash := hmTypes.BytesToHeimdallHash([]byte("some-address"))
	testRecord1 := types.NewEventRecord(hHash, 1, 1, hAddr, make([]byte, 0), "1")

	respKey := clerk.GetEventRecordKey(testRecord1.ID)
	require.Equal(t, respKey, []byte{17, 49})
}

func (suite *KeeperTestSuite) TestSetHasGetRecordSequence() {
	t, app, ctx := suite.T(), suite.app, suite.ctx

	testSeq := "testseq"
	ck := app.ClerkKeeper
	ck.SetRecordSequence(ctx, testSeq)
	found := ck.HasRecordSequence(ctx, testSeq)
	require.True(t, found)

	found = ck.HasRecordSequence(ctx, "testSeq")
	require.False(t, found)

	recordSequences := ck.GetRecordSequences(ctx)
	require.Len(t, recordSequences, 1)
}
