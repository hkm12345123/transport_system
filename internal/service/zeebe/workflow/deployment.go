package workflow

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/hkm12345123/transport_system/internal/model"
)

var (
	// ZbClient client to connect with zeebe engine
	zbClient zbc.Client
)

// ConnectZeebeEngine function
func ConnectZeebeEngine() error {
	gatewayAddress := os.Getenv("BROKER_ADDRESS")
	newZbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddress,
		UsePlaintextConnection: true,
	})

	if err != nil {
		return err
	}

	zbClient = newZbClient
	return nil
}

// DeployFullShipWorkflow function
func DeployFullShipWorkflow() error {

	ctx := context.Background()
	///// DEPLOY BPMN FILE TO ZEEEBE SERVER /////
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile(os.Getenv("FULL_SHIP_ZB_FILE_1")).Send(ctx)
	if err != nil {
		log.Printf("acc")
		return err
	}

	fmt.Println(response.String())
	return nil
}

// DeployLongShipWorkflow function
func DeployLongShipWorkflow() error {

	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile(os.Getenv("LONG_SHIP_ZB_FILE_1")).Send(ctx)
	if err != nil {
		return err
	}

	fmt.Println(response.String())
	return nil
}

// CreateFullShipInstance of workflow
func CreateFullShipInstance(orderWorkflowData *model.OrderWorkflowData) (string, uint, error) {

	// After the workflow is deployed.
	variables := make(map[string]interface{})
	variables["order_id"] = orderWorkflowData.OrderID
	variables["shipper_receive_money"] = orderWorkflowData.ShipperReceiveMoney
	variables["use_long_ship"] = orderWorkflowData.UseLongShip
	variables["customer_receive_id"] = orderWorkflowData.CustomerReceiveID

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId(os.Getenv("FULL_SHIP_ZB_ID_1")).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		return "", uint(0), err
	}

	ctx := context.Background()
	msg, err := request.Send(ctx)
	if err != nil {
		return "", uint(0), err
	}
	log.Println(msg.String())
	return strconv.Itoa(int(msg.ProcessDefinitionKey)), uint(msg.ProcessInstanceKey), nil
}

// CreateFullShipInstanceWithBug of workflow
func CreateFullShipInstanceWithBug(orderWorkflowData *model.OrderWorkflowData) (string, uint, error) {

	// After the workflow is deployed.
	variables := make(map[string]interface{})
	variables["shipper_receive_money"] = orderWorkflowData.ShipperReceiveMoney
	variables["use_long_ship"] = orderWorkflowData.UseLongShip
	variables["customer_receive_id"] = orderWorkflowData.CustomerReceiveID

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId(os.Getenv("FULL_SHIP_ZB_ID_1")).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		return "", uint(0), err
	}

	ctx := context.Background()
	msg, err := request.Send(ctx)
	if err != nil {
		return "", uint(0), err
	}
	log.Println(msg.String())
	return strconv.Itoa(int(msg.ProcessDefinitionKey)), uint(msg.ProcessInstanceKey), nil
}

// CreateLongShipInstance of workflow
func CreateLongShipInstance(longShipID uint) (string, uint, error) {

	// After the workflow is deployed.
	variables := make(map[string]interface{})
	variables["long_ship_id"] = longShipID

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId(os.Getenv("LONG_SHIP_ZB_ID_1")).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		return "", uint(0), err
	}

	ctx := context.Background()
	msg, err := request.Send(ctx)
	if err != nil {
		return "", uint(0), err
	}
	log.Println(msg.String())
	return strconv.Itoa(int(msg.ProcessDefinitionKey)), uint(msg.ProcessInstanceKey), nil
}

// CancelFullShipInstance of workflow
func CancelFullShipInstance(workflowInstanceKey uint) error {

	ctx := context.Background()
	_, err := zbClient.NewCancelInstanceCommand().ProcessInstanceKey(int64(workflowInstanceKey)).Send(ctx)
	if err != nil {
		return err
	}
	log.Println("Canceled instance with workflow key", workflowInstanceKey)
	return nil
}
