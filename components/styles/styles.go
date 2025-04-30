package styles

// Dracula theme colors
const (
	DraculaBackground  = "#282a36"
	DraculaCurrentLine = "#44475a"
	DraculaForeground  = "#f8f8f2"
	DraculaComment     = "#6272a4"
	DraculaPurple      = "#bd93f9"
	DraculaGreen       = "#50fa7b"
	DraculaOrange      = "#ffb86c"
	DraculaPink        = "#ff79c6"
	DraculaRed         = "#ff5555"
	DraculaYellow      = "#f1fa8c"
	DraculaCyan        = "#8be9fd"
)

// Cores base
const (
	BgDark        = "bg-dracula-background"
	BgInput       = "bg-dracula-current-line"
	TextPrimary   = "text-dracula-foreground"
	TextMuted     = "text-dracula-comment"
	BorderDark    = "border-dracula-comment"
	BorderInput   = "border-dracula-current-line"
	Rounded       = "rounded"
	Shadow        = "shadow-sm"
	Danger        = "bg-dracula-red"
	DangerText    = "text-dracula-red"
	Success       = "bg-dracula-green"
	SuccessText   = "text-dracula-green"
	Warning       = "bg-dracula-orange"
	WarningText   = "text-dracula-orange"
	Info          = "bg-dracula-cyan"
	InfoText      = "text-dracula-cyan"
	Light         = "bg-dracula-foreground"
	LightText     = "text-dracula-background"
	Dark          = "bg-dracula-current-line"
	DarkText      = "text-dracula-foreground"
	Primary       = "bg-dracula-purple"
	PrimaryText   = "text-dracula-purple"
	Secondary     = "bg-dracula-pink"
	SecondaryText = "text-dracula-pink"
	Muted         = "bg-dracula-comment"
	MutedText     = "text-dracula-comment"
)

// Input
const (
	InputBase = BgInput + " " + TextPrimary + " " + BorderInput + " px-3 py-2 " + Rounded + " " + Shadow +
		" focus:outline-none focus:ring-2 focus:ring-dracula-purple"
	Label     = TextPrimary + " font-medium mb-1"
	ErrorText = "text-dracula-red text-sm mt-1"
)

// Button
const (
	ButtonPrimary   = "bg-dracula-purple hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonDanger    = "bg-dracula-red hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonSuccess   = "bg-dracula-green hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonWarning   = "bg-dracula-orange hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonInfo      = "bg-dracula-cyan hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonLight     = "bg-dracula-foreground hover:bg-opacity-80 text-dracula-background font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonDark      = "bg-dracula-current-line hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonSecondary = "bg-dracula-pink hover:bg-opacity-80 " + TextPrimary + " font-bold px-4 py-2 " + Rounded + " block w-fit"
	ButtonYellow    = "bg-dracula-yellow hover:bg-opacity-80 text-dracula-background font-bold px-4 py-2 " + Rounded + " block w-fit"
)

// Sidebar
const (
	SidebarBase = BgDark + " " + TextPrimary + " p-4 border-r " + BorderDark + " w-64 min-h-screen"
)

// VBox
const (
	VBoxBase = BgDark + " " + TextPrimary + " p-4 space-y-4 flex flex-col"
)

// Select
const (
	SelectBase = BgInput + " " + TextPrimary + " " + BorderInput + " px-3 py-2 " + Rounded + " " + Shadow +
		" focus:outline-none focus:ring-2 focus:ring-dracula-purple w-full appearance-none bg-dracula-current-line" +
		" bg-no-repeat bg-right pr-8"
)

// Checkbox
const (
	CheckboxBase = "form-checkbox h-5 w-5 text-dracula-purple rounded border-dracula-comment focus:ring-dracula-purple"
)

// Radio
const (
	RadioBase = "form-radio h-5 w-5 text-dracula-purple border-dracula-comment focus:ring-dracula-purple"
)
