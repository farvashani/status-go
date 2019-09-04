package mailservers

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/status-im/status-go/appdatabase"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*Database, func()) {
	tmpfile, err := ioutil.TempFile("", "mailservers-service")
	require.NoError(t, err)
	db, err := appdatabase.InitializeDB(tmpfile.Name(), "mailservers-tests")
	require.NoError(t, err)
	return NewDB(db), func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.Remove(tmpfile.Name()))
	}
}

func TestAddGetDeleteMailserver(t *testing.T) {
	db, close := setupTestDB(t)
	defer close()
	api := &API{db: db}
	testMailserver := Mailserver{
		ID:      "mailserver001",
		Name:    "My Mailserver",
		Address: "enode://...",
		Fleet:   "beta",
	}
	testMailserverWithPassword := testMailserver
	testMailserverWithPassword.ID = "mailserver002"
	testMailserverWithPassword.Password = "test-pass"

	err := api.AddMailserver(context.Background(), testMailserver)
	require.NoError(t, err)
	err = api.AddMailserver(context.Background(), testMailserverWithPassword)
	require.NoError(t, err)

	mailservers, err := api.GetMailservers(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, []Mailserver{testMailserver, testMailserverWithPassword}, mailservers)

	err = api.DeleteMailserver(context.Background(), testMailserver.ID)
	require.NoError(t, err)
	// Verify they was deleted.
	mailservers, err = api.GetMailservers(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, []Mailserver{testMailserverWithPassword}, mailservers)
	// Delete non-existing mailserver.
	err = api.DeleteMailserver(context.Background(), "other-id")
	require.NoError(t, err)
}

func TestAddGetDeleteMailserverRequestGap(t *testing.T) {
	db, close := setupTestDB(t)
	defer close()
	chatID1 := "chat-id-1"
	chatID2 := "chat-id-2"

	api := &API{db: db}
	gap1 := MailserverRequestGap{ID: "1", ChatID: chatID1, From: 1, To: 2}
	gap2 := MailserverRequestGap{ID: "2", ChatID: chatID2, From: 1, To: 2}
	gap3 := MailserverRequestGap{ID: "3", ChatID: chatID2, From: 1, To: 2}

	gaps := []MailserverRequestGap{
		gap1,
		gap2,
		gap3,
	}

	err := api.AddMailserverRequestGaps(context.Background(), gaps)
	require.NoError(t, err)

	actualGaps, err := api.GetMailserverRequestGaps(context.Background(), chatID1)
	require.NoError(t, err)
	require.EqualValues(t, []MailserverRequestGap{gap1}, actualGaps)

	actualGaps, err = api.GetMailserverRequestGaps(context.Background(), chatID2)
	require.NoError(t, err)
	require.EqualValues(t, []MailserverRequestGap{gap2, gap3}, actualGaps)

	err = api.DeleteMailserverRequestGaps(context.Background(), []string{gap1.ID, gap2.ID})
	require.NoError(t, err)

	// Verify it was deleted.
	actualGaps, err = api.GetMailserverRequestGaps(context.Background(), chatID1)
	require.NoError(t, err)
	require.Len(t, actualGaps, 0)

	actualGaps, err = api.GetMailserverRequestGaps(context.Background(), chatID2)
	require.NoError(t, err)
	require.Len(t, actualGaps, 1)

	err = api.DeleteMailserverRequestGapsByChatID(context.Background(), chatID2)
	require.NoError(t, err)

	// Verify it was deleted.
	actualGaps, err = api.GetMailserverRequestGaps(context.Background(), chatID2)
	require.NoError(t, err)
	require.Len(t, actualGaps, 0)
}
