package txlogs

import (
	"encoding/json"
	"fmt"

	txlModel "github.com/TanmoySG/wunderDB/internal/txlogs/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
)

type DotTxLog struct {
	txLogFilepath   string
	transactionLogs TransactionLogs
}

func UseDotTxLog(wunderDbBasePath string) DotTxLog {
	wdbLogBasePath := fmt.Sprintf(WDB_DOT_TX_LOG_BASEPATH, wunderDbBasePath)
	return DotTxLog{
		txLogFilepath: fmt.Sprintf(WDB_DOT_TX_LOG_FILEPATH, wdbLogBasePath, wdbDotTxLogFilename),
	}
}

func (dotTxL *DotTxLog) Commit() error {
	preCommitTxLogBytes, err := dotTxL.transactionLogs.Marshal()
	if err != nil {
		return err
	}
	return fs.WriteToFile(dotTxL.txLogFilepath, preCommitTxLogBytes)

}

func (dotTxL *DotTxLog) Log(newLog txlModel.TxlogSchemaJson) error {
	dotTxLFilepath := dotTxL.txLogFilepath
	if fs.CheckFileExists(dotTxLFilepath) {
		previousTxLogBytes, err := fs.ReadFile(dotTxLFilepath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(previousTxLogBytes, &dotTxL.transactionLogs)
		if err != nil {
			return err
		}

		dotTxL.transactionLogs.Logs = append(dotTxL.transactionLogs.Logs, &newLog)
	}
	return nil
}
