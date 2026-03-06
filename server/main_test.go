package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetServer() {
	server = Server{}
}

// --- randSeq tests ---

func TestRandSeq_Length(t *testing.T) {
	// given
	length := 10

	// when
	result := randSeq(length)

	// then
	assert.Equal(t, length, len(result))
}

func TestRandSeq_OnlyLetters(t *testing.T) {
	// given / when
	result := randSeq(100)

	// then
	for _, r := range result {
		assert.Contains(t, string(letters), string(r))
	}
}

func TestRandSeq_Unique(t *testing.T) {
	// given / when
	a := randSeq(20)
	b := randSeq(20)

	// then
	assert.NotEqual(t, a, b)
}

// --- assignSide tests ---

func TestAssignSide_FirstClient(t *testing.T) {
	// given / when
	s := assignSide(0)

	// then
	assert.Equal(t, LEFT, s)
}

func TestAssignSide_SecondClient(t *testing.T) {
	// given / when
	s := assignSide(1)

	// then
	assert.Equal(t, RIGHT, s)
}

func TestAssignSide_ThirdClient(t *testing.T) {
	// given / when
	s := assignSide(2)

	// then
	assert.Equal(t, NONE, s)
}

// --- createClient tests ---

func TestCreateClient_FirstClient(t *testing.T) {
	// given
	resetServer()

	// when
	client := createClient()

	// then
	assert.Equal(t, LEFT, client.Side)
	assert.Equal(t, 5, len(client.ID))
	assert.NotNil(t, client.Ingress)
	assert.NotNil(t, client.Egress)
}

func TestCreateClient_SecondClient(t *testing.T) {
	// given
	resetServer()
	server.Clients = append(server.Clients, createClient())

	// when
	client := createClient()

	// then
	assert.Equal(t, RIGHT, client.Side)
}

func TestCreateClient_UniqueIDs(t *testing.T) {
	// given
	resetServer()

	// when
	a := createClient()
	b := createClient()

	// then
	assert.NotEqual(t, a.ID, b.ID)
}

// --- deregisterClient tests ---

func TestDeregisterClient_RemovesClient(t *testing.T) {
	// given
	resetServer()
	client := createClient()
	server.Clients = append(server.Clients, client)
	assert.Equal(t, 1, len(server.Clients))

	// when
	deregisterClient(client.ID)

	// then
	assert.Equal(t, 0, len(server.Clients))
}

func TestDeregisterClient_KeepsOtherClients(t *testing.T) {
	// given
	resetServer()
	a := createClient()
	b := createClient()
	server.Clients = append(server.Clients, a, b)

	// when
	deregisterClient(a.ID)

	// then
	assert.Equal(t, 1, len(server.Clients))
	assert.Equal(t, b.ID, server.Clients[0].ID)
}

func TestDeregisterClient_NonexistentID(t *testing.T) {
	// given
	resetServer()
	client := createClient()
	server.Clients = append(server.Clients, client)

	// when
	deregisterClient("nonexistent")

	// then
	assert.Equal(t, 1, len(server.Clients))
}

// --- handleMessage tests ---

func TestHandleMessage_MovementMessage(t *testing.T) {
	// given / when / then — should not panic
	assert.NotPanics(t, func() {
		handleMessage("m abc123 1")
	})
}

func TestHandleMessage_MalformedMovement(t *testing.T) {
	// given / when / then — should not panic on short message
	assert.NotPanics(t, func() {
		handleMessage("m")
	})
}

func TestHandleMessage_UnknownMessage(t *testing.T) {
	// given / when / then — should not panic
	assert.NotPanics(t, func() {
		handleMessage("x something")
	})
}

// --- side constants tests ---

func TestSideConstants_AreDistinct(t *testing.T) {
	// given / when / then
	assert.NotEqual(t, LEFT, RIGHT)
	assert.NotEqual(t, LEFT, NONE)
	assert.NotEqual(t, RIGHT, NONE)
}
