package CommonRouter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
)

func (suite *BaseCommonRouterTestSuite) TestGlobalChecklist() {

	suite.Run("A=1=GetGlobalChecklist", func() {
		globalChecklist := HitGlobalChecklistApi(suite.authToken)
		log.Println("Validating the response of GlobalChecklist API")
		myDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(myDir, "OssInstallationMode") {
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.Git)
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.HostUrl)
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.Docker)
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.Git)
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.Environment)
			assert.Equal(suite.T(), 1, globalChecklist.Result.ChartChecklist.Environment)
			assert.Equal(suite.T(), 1, globalChecklist.Result.ChartChecklist.Project)
		} else {
			assert.Equal(suite.T(), 1, globalChecklist.Result.AppChecklist.HostUrl)
		}
	})
}
