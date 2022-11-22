package UrlsRouter

import (
	userRouter "automation-suite/UserRouter"
	"automation-suite/UserRouter/ResponseDTOs"
)

func (suite *UrlsTestSuite) TestGetUrlsForHelmApp() {
	testGetUrlsForHelmApp(suite, suite.authToken)
}
func CreateRoleFilterWithDevtronAppsOnly(role string) ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter = userRouter.CreateRoleFilter("", "devtron-demo", "", role, "")
	return roleFilter
}
func (suite *UrlsTestSuite) TestUrlsdataWithViewOnlyAccess() {
	testUrlsdataWithRoleAccess(suite, "view")
}
func (suite *UrlsTestSuite) TestUrlsdataWithAdminAccess() {
	testUrlsdataWithRoleAccess(suite, "admin")
}
func (suite *UrlsTestSuite) TestUrlsdataWithManagerAccess() {
	testUrlsdataWithRoleAccess(suite, "manager")
}
func (suite *UrlsTestSuite) TestGetUrlsForHelmAppWithIncorrectAppId() {
	testGetUrlsForHelmAppWithIncorrectAppId(suite, suite.authToken)
}
func (suite *UrlsTestSuite) TestGetUrlsForDevtronApp() {
	testGetUrlsForDevtronApp(suite, suite.authToken)
}
func (suite *UrlsTestSuite) TestGetUrlsForDevtronAppWithIncorrectAppId() {
	testGetUrlsForDevtronAppWithIncorrectAppId(suite, suite.authToken)
}
func (suite *UrlsTestSuite) TestGetUrlsForInstalledApp() {
	testGetUrlsForInstalledApp(suite, suite.authToken)
}
func (suite *UrlsTestSuite) TestGetUrlsForInstalledAppWithIncorrectAppId() {
	testGetUrlsForInstalledAppWithIncorrectAppId(suite, suite.authToken)
}
func (suite *UrlsTestSuite) TestGetUrlsdata() {
	testGetUrlsdata(suite, suite.authToken)
}
