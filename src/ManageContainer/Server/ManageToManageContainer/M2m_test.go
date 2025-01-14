package m2mserver

import (
	"context"
	"testing"

	m2db "github.com/acompany-develop/QuickMPC/src/ManageContainer/Client/ManageToDbGate"
	utils "github.com/acompany-develop/QuickMPC/src/ManageContainer/Utils"
	pb "github.com/acompany-develop/QuickMPC/src/Proto/ManageToManageContainer"
)

// Test用のDbGとCCのmock
type localDbGate struct{}

func (localDbGate) InsertShares(string, []string, int32, string, string) error {
	return nil
}
func (localDbGate) DeleteShares([]string) error {
	return nil
}
func (localDbGate) GetSchema(string) ([]string, error) {
	return []string{""}, nil
}
func (localDbGate) GetComputationResult(string) (*m2db.ComputationResult, error) {
	return &m2db.ComputationResult{}, nil
}
func (localDbGate) InsertModelParams(string, string) error {
	return nil
}
func (localDbGate) GetDataList() (string, error) {
	return "result", nil
}

// Test用のサーバを起動(MC)
var s *utils.TestServer

func init() {
	s = &utils.TestServer{}
	pb.RegisterManageToManageServer(s.GetServer(), &server{m2dbclient: localDbGate{}})
	s.Serve()
}

// Share削除リクエストが受け取れるかのTest
func TestGetSchema(t *testing.T) {
	conn := s.GetConn()
	defer conn.Close()

	client := pb.NewManageToManageClient(conn)

	_, err := client.DeleteShares(context.Background(), &pb.DeleteSharesRequest{
		DataId: "id"})

	if err != nil {
		t.Fatal(err)
	}
}

// Syncリクエストが受け取れるかのTest
func TestSync(t *testing.T) {
	conn := s.GetConn()
	defer conn.Close()

	client := pb.NewManageToManageClient(conn)

	_, err := client.Sync(context.Background(), &pb.SyncRequest{SyncId: "SyncID"})
	if err != nil {
		t.Fatal(err)
	}
}
