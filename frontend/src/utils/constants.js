export const LOCAL_STORAGE_KEY = "projects";
export const LOCAL_STORAGE_TOKEN = "token";
export const LOCAL_STORAGE_REMEMBER_TOKEN = "rememberToken";
export const LOCAL_STORAGE_COMPANIES = "companies";
export const LOCAL_STORAGE_COMPANY = "company";
export const LOCAL_STORAGE_COMPANY_ID = "companyID";
export const LOCAL_STORAGE_COLLAPSED = "collapsed";
export const LOCAL_STORAGE_DEFAULT_CHANNEL = "defaultChannel";
export const LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION = "defaultWhatsappSession";
export const LOCAL_STORAGE_DEFAULT_TELEGRAM_SESSION = "defaultTelegramSession";
export const LOCAL_STORAGE_DEFAULT_INSTAGRAM_SESSION = "defaultInstagramSession";
export const LOCAL_STORAGE_DEFAULT_TIKTOK_SESSION = "defaultTiktokSession";


export const severityOptions = [
    { value: "LOW", label: "Low", color: "#8BC34A" },
    { value: "MEDIUM", label: "Medium", color: "#F7DC6F" },
    { value: "HIGH", label: "High", color: "#FFC107" },
    { value: "CRITICAL", label: "Critical", color: "#F44336" },
];

export const priorityOptions = [
    { value: "LOW", label: "Low", color: "#8BC34A" },
    { value: "MEDIUM", label: "Medium", color: "#F7DC6F" },
    { value: "HIGH", label: "High", color: "#FFC107" },
    { value: "URGENT", label: "Urgent", color: "#F44336" },
];


export const llmModel = {
    "ollama": [
        { value: "ollama-13b", label: "ollama-13b" },
        { value: "ollama-7b", label: "ollama-7b" },
        { value: "ollama-3b", label: "ollama-3b" },
        { value: "ollama-1.3b", label: "ollama-1.3b" },
        { value: "ollama-1.3b-16k", label: "ollama-1.3b-16k" },
        { value: "gemma3:1b", label: "Gemma 3:1b" },
        { value: "gemma3:3b", label: "Gemma 3:3b" },
    ],
    "openai": [
        { value: "gpt-4o", label: "GPT-4" },
    ],
    "deepseek": [
        { value: "deepseek-chat", label: "Deepseek Chat" },
    ],
    "gemini": [
        { value: "gemini-1.5-flash", label: "Gemini 1.5 Flash" },
        { value: "gemini-2.0-flash", label: "Gemini 2.0 Flash" },
        { value: "gemini-2.5-flash", label: "Gemini 2.5 Flash" },
        { value: "gemini-2.5-pro", label: "Gemini 2.5 Pro" },
    ]
}

export const interactiveTypes = [
    { value: "list", label: "List" },
]