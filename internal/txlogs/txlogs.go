package txlogs

import (
	"fmt"
	"time"

	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	txlModel "github.com/TanmoySG/wunderDB/internal/txlogs/model"
	"github.com/gofiber/fiber/v2"
)

type TxnEntityType string

type TransactionLogs struct {
	Logs []*txlModel.TxlogSchemaJson
}

var (
	DatabaseTxnEntity   TxnEntityType = "database"
	CollectionTxnEntity TxnEntityType = "collection"
)

func Use() TransactionLogs {
	return TransactionLogs{}
}

func (tl *TransactionLogs) Log(c *fiber.Ctx) {
	entityPath := c.Path()
	fmt.Println(entityPath)
	newLog := txlModel.TxlogSchemaJson{}
	tl.Logs = append(tl.Logs, &newLog)
}

func CreateTxLog(txnAction, txnActor string, txnRequestStatus string, txnEntities txlModel.TxlogSchemaJsonEntityPath, txnDetails txlModel.TxlogSchemaJsonTransactionDetails) txlModel.TxlogSchemaJson {
	return txlModel.TxlogSchemaJson{
		Action:             txnAction,
		Actor:              &txnActor,
		EntityPath:         txnEntities,
		EntityType:         getEntityType(txnEntities),
		Status:             getTxnStatus(txnRequestStatus),
		Timestamp:          float64(time.Now().Unix()),
		TransactionDetails: txnDetails,
	}
}

func GetTxnHttpDetails(c fiber.Ctx) txlModel.TxlogSchemaJsonTransactionDetails {
	txnHttpUrl, txnUserAgent, txnRequestIP := c.Path(), c.Get("User-Agent"), c.IP()

	txnRequestHttpMethod := txlModel.TxlogSchemaJsonTransactionDetailsMethod(c.Method())
	txnRequestPayload := string(c.Body())

	txnResponseHttpStatusCode := c.Response().StatusCode()
	txnResponsePayload := string(c.Response().Body())

	return txlModel.TxlogSchemaJsonTransactionDetails{
		UrlEndpoint: txnHttpUrl,
		Method:      txnRequestHttpMethod,
		Request: txlModel.TxlogSchemaJsonTransactionDetailsRequest{
			IsAuthenticated: true, // make it dynamic based on auth
			Payload:         &txnRequestPayload,
			UserAgent: txlModel.TxlogSchemaJsonTransactionDetailsRequestUserAgent{
				"userAgent": txnUserAgent,
				"requestIP": txnRequestIP,
			},
		},
		Response: txlModel.TxlogSchemaJsonTransactionDetailsResponse{
			HttpStatus:   txnResponseHttpStatusCode,
			ResponseBody: txnResponsePayload,
		},
	}
}

// only write actions are loggable
func IsTxnLoggable(txnAction string) bool {
	if txnType := privileges.GetPrivilegeType(txnAction); txnType == privileges.WritePrivilege {
		return true
	}
	return false
}

func getEntityType(txnEntities txlModel.TxlogSchemaJsonEntityPath) txlModel.TxlogSchemaJsonEntityType {
	if txnEntities.Database != "" {
		if txnEntities.Collection != nil || *(txnEntities.Collection) != "" {
			return txlModel.TxlogSchemaJsonEntityTypeCOLLECTION
		}
		return txlModel.TxlogSchemaJsonEntityTypeDATABASE
	}
	return ""
}

func getTxnStatus(txnRequestStatus string) txlModel.TxlogSchemaJsonStatus {
	switch txnRequestStatus {
	case response.StatusFailure:
		return txlModel.TxlogSchemaJsonStatusFAILED
	case response.StatusSuccess:
		return txlModel.TxlogSchemaJsonStatusSUCCESS
	default:
		return ""
	}
}
