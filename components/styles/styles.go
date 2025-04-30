package styles

// Cores base
const (
	BgDark        = "bg-gray-900"
	BgInput       = "bg-gray-800"
	TextPrimary   = "text-white"
	TextMuted     = "text-gray-400"
	BorderDark    = "border-gray-700"
	BorderInput   = "border-gray-600"
	Rounded       = "rounded"
	Shadow        = "shadow-sm"
	Danger        = "bg-red-500"
	DangerText    = "text-red-400"
	Success       = "bg-green-500"
	SuccessText   = "text-green-400"
	Warning       = "bg-yellow-500"
	WarningText   = "text-yellow-400"
	Info          = "bg-blue-500"
	InfoText      = "text-blue-400"
	Light         = "bg-gray-100"
	LightText     = "text-gray-700"
	Dark          = "bg-gray-800"
	DarkText      = "text-gray-200"
	Primary       = "bg-blue-600"
	PrimaryText   = "text-blue-400"
	Secondary     = "bg-gray-400"
	SecondaryText = "text-gray-800"
	Muted         = "bg-gray-500"
	MutedText     = "text-gray-600"
)

// Input
const (
	InputBase = BgInput + " " + TextPrimary + " " + BorderInput + " px-3 py-2 " + Rounded + " " + Shadow +
		" focus:outline-none focus:ring-2 focus:ring-blue-400"
	Label     = TextPrimary + " font-medium mb-1"
	ErrorText = "text-red-400 text-sm mt-1"
)

// Button
const (
	ButtonPrimary = "bg-blue-600 hover:bg-blue-700 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonDanger  = "bg-red-600 hover:bg-red-700 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonSuccess = "bg-green-600 hover:bg-green-700 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonWarning = "bg-yellow-600 hover:bg-yellow-700 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonInfo    = "bg-blue-600 hover:bg-blue-700 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonLight   = "bg-gray-100 hover:bg-gray-200 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonDark    = "bg-gray-800 hover:bg-gray-900 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
)

// Sidebar
const (
	SidebarBase = BgDark + " " + TextPrimary + " p-4 border-r " + BorderDark + " w-64 min-h-screen"
)

// VBox
const (
	VBoxBase = BgDark + " " + TextPrimary + " p-4 space-y-4 flex flex-col"
)
