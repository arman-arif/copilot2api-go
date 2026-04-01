package copilot

// VSCodeVersion is the VSCode version reported to the GitHub Copilot API.
// Update this constant when a newer VSCode release is available.
const VSCodeVersion = "1.104.3"

// GetVSCodeVersion returns the VSCode version string used in Copilot API headers.
func GetVSCodeVersion() string {
	return VSCodeVersion
}
