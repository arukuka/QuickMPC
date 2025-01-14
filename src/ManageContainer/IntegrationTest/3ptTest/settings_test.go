package integration_3pt_test

import (
	"os"
	"testing"
	"time"

	m2m "github.com/acompany-develop/QuickMPC/src/ManageContainer/Client/ManageToManageContainer"
	l2mserver "github.com/acompany-develop/QuickMPC/src/ManageContainer/Server/LibcToManageContainer"
	m2mserver "github.com/acompany-develop/QuickMPC/src/ManageContainer/Server/ManageToManageContainer"
)

func TestMain(m *testing.M) {

	// サーバを起動する
	go m2mserver.RunServer()
	go l2mserver.RunServer()
	time.Sleep(time.Second)

	// Test実行タイミングを全てのPTで同期させる
	client := m2m.Client{}
	client.Sync("start")

	// Testを実行
	code := m.Run()

	// 他のPTがサーバが使うのを少し待機
	// (自分だけTestを終えて早々にサーバを閉じることがあるため)
	time.Sleep(3 * time.Second)

	os.Exit(code)
}
