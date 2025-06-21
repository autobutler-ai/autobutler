package llm

type ChatRole string

const (
    ChatRoleUser   ChatRole = "user"
    ChatRoleSystem ChatRole = "system"
    ChatRoleDummy  ChatRole = "dummy"
    ChatRoleError  ChatRole = "error"
)

type ChatMessage struct {
    ID        string   `json:"id"`
    Role      ChatRole `json:"role"`
    Content   string   `json:"content"`
    Timestamp string   `json:"timestamp"` // ISO 8601 format
}
