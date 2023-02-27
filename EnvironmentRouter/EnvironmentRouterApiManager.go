package EnvironmentRouter

import (
	"automation-suite/EnvironmentRouter/RequestDTOs"
	"automation-suite/EnvironmentRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"net/http"
)

type EnvironmentsRouterStruct struct {
	createEnvironmentDTO         RequestDTOs.CreateEnvironmentRequestDTO
	createEnvironmentResponseDTO ResponseDTOs.CreateEnvironmentResponseDTO
	deleteEnvResponseDTO         ResponseDTOs.DeleteEnvResponseDTO
}

func GetCreateUpdateDeleteEnvRequestDto(EnvId int, envName string, clusterId int, isDefault bool) RequestDTOs.CreateEnvironmentRequestDTO {
	var createEnvironmentDTO RequestDTOs.CreateEnvironmentRequestDTO
	createEnvironmentDTO.Id = EnvId
	createEnvironmentDTO.Environment = envName
	createEnvironmentDTO.ClusterId = clusterId
	createEnvironmentDTO.Namespace = envName
	createEnvironmentDTO.Active = true
	createEnvironmentDTO.Default = isDefault
	return createEnvironmentDTO
}

func HitCreateEnvApi(byteValueOfStruct []byte, authToken string) ResponseDTOs.CreateEnvironmentResponseDTO {
	resp, err := Base.MakeApiCall(EnvRouterBaseUrl, http.MethodPost, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, CreateEnvApi)
	environmentsRouterStruct := EnvironmentsRouterStruct{}
	envRouter := environmentsRouterStruct.UnmarshalGivenResponseBody(resp.Body(), CreateEnvApi)
	return envRouter.createEnvironmentResponseDTO
}

func HitDeleteEnvApi(byteValueOfStruct []byte, authToken string) ResponseDTOs.DeleteEnvResponseDTO {
	resp, err := Base.MakeApiCall(EnvRouterBaseUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, CreateEnvApi)
	environmentsRouterStruct := EnvironmentsRouterStruct{}
	envRouter := environmentsRouterStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteEnvApi)
	return envRouter.deleteEnvResponseDTO
}

func (environmentsRouterStruct EnvironmentsRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) EnvironmentsRouterStruct {
	switch apiName {
	case CreateEnvApi:
		json.Unmarshal(response, &environmentsRouterStruct.createEnvironmentResponseDTO)
	case DeleteEnvApi:
		json.Unmarshal(response, &environmentsRouterStruct.deleteEnvResponseDTO)
	}
	return environmentsRouterStruct
}
