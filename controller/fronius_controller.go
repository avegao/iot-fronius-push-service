package froniusController

import (
	"github.com/gin-gonic/gin"
	"github.com/avegao/gocondi"
	"io/ioutil"
	"encoding/json"
	"github.com/avegao/iot-fronius-push-service/entity/fronius/current_data/meter"
	"github.com/avegao/iot-fronius-push-service/entity/fronius/current_powerflow"
	"github.com/avegao/iot-fronius-push-service/service"
	"context"
	"github.com/avegao/iot-fronius-push-service/entity/fronius/current_data/inverter"
	"github.com/avegao/iot-fronius-push-service/entity/fronius/io_state"
)

func PostCurrentDataMeterAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Error()
	}

	var currentData froniusCurrentDataMeter.CurrentDataMeter
	if err := json.Unmarshal(body, &currentData); err != nil {
		logger.WithError(err).Error()
	}

	request := currentData.Body.ToGrpcRequest()

	connection, err := service.CreateConnection()
	if err != nil {
		logger.WithError(err).Error()
	}

	client := service.CreateClient(connection)
	ctx := context.Background()

	response, err := client.InsertCurrentDataMeter(ctx, request)
	if err != nil || !response.Success {
		logger.WithError(err).Error()
	}
}

func PostCurrentDataInverterAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Error()
	}

	logger.WithField("body", string(body)).Debug()

	var currentData froniusCurrentDataInverter.CurrentDataInverter
	json.Unmarshal(body, &currentData)

	request := currentData.ToGrpcRequest()

	connection, err := service.CreateConnection()
	if err != nil {
		logger.WithError(err).Error()
	}

	client := service.CreateClient(connection)
	ctx := context.Background()

	response, err := client.InsertCurrentDataInverter(ctx, request)
	if err != nil || !response.Success {
		logger.WithError(err).Error()
	}
}

func PostCurrentDataPowerflowAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)
	if err != nil {
		logger.WithError(err).Error()
	}

	var currentData froniusCurrentPowerflow.CurrentPowerflow
	if err = json.Unmarshal(body, &currentData); err != nil {
		logger.WithError(err).Error()
	}

	request := currentData.ToGrpcRequest()

	connection, err := service.CreateConnection()
	if err != nil {
		logger.WithError(err).Error()
	}

	client := service.CreateClient(connection)
	ctx := context.Background()

	response, err := client.InsertCurrentDataPowerflow(ctx, &request)
	if err != nil || !response.Success {
		logger.WithError(err).Error()
	}
}

func PostCurrentDataSensorCardAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Error()
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}

func PostCurrentDataStringControlAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Error()
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}

func PostCurrentDataIoStatesAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Error()
	}

	var ioState froniusIoState.IoState
	if err = json.Unmarshal(body, &ioState); err != nil {
		logger.WithError(err).Error()
	}

	logger.WithField("ioState", ioState).Debug()

	request := ioState.ToGrpcRequest()

	connection, err := service.CreateConnection()
	if err != nil {
		logger.WithError(err).Error()
	}

	client := service.CreateClient(connection)
	ctx := context.Background()

	response, err := client.InsertCurrentIoState(ctx, &request)
	if err != nil || !response.Success {
		logger.WithError(err).Error()
	}
}
