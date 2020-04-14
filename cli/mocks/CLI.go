// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cli "github.com/monitoror/monitoror/cli"
	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/models"
)

// CLI is an autogenerated mock type for the CLI type
type CLI struct {
	mock.Mock
}

// PrintBanner provides a mock function with given fields:
func (_m *CLI) PrintBanner() {
	_m.Called()
}

// PrintDevMode provides a mock function with given fields:
func (_m *CLI) PrintDevMode() {
	_m.Called()
}

// PrintMonitorable provides a mock function with given fields: displayName, enabledVariants, erroredVariants
func (_m *CLI) PrintMonitorable(displayName string, enabledVariants []models.VariantName, erroredVariants []cli.ErroredVariant) {
	_m.Called(displayName, enabledVariants, erroredVariants)
}

// PrintMonitorableFooter provides a mock function with given fields: isProduction, nonEnabledMonitorableCount
func (_m *CLI) PrintMonitorableFooter(isProduction bool, nonEnabledMonitorableCount int) {
	_m.Called(isProduction, nonEnabledMonitorableCount)
}

// PrintMonitorableHeader provides a mock function with given fields:
func (_m *CLI) PrintMonitorableHeader() {
	_m.Called()
}

// PrintServerStartup provides a mock function with given fields: ip, port
func (_m *CLI) PrintServerStartup(ip string, port int) {
	_m.Called(ip, port)
}
