package chat

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

func (m ChatMessage) Class() string {
    switch m.Role {
        case ChatRoleUser:
            return "bg-blue-600 text-white"
        case ChatRoleSystem:
            return "bg-white/10 text-gray-100"
        case ChatRoleDummy:
            return "bg-orange-500/20 text-gray-100 border-2 border-orange-500"
        case ChatRoleError:
            return "bg-red-500/20 text-gray-100 border-2 border-red-500"
        default:
            panic("unknown chat role: " + string(m.Role))
    }
}
