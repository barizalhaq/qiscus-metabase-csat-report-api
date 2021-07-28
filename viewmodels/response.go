package viewmodels

type MultichannelAuthResponse struct {
	Data struct {
		User struct {
			ID                  int         `json:"id"`
			Name                string      `json:"name"`
			Email               string      `json:"email"`
			AuthenticationToken string      `json:"authentication_token"`
			CreatedAt           string      `json:"created_at"`
			UpdatedAt           string      `json:"updated_at"`
			SDKEmail            string      `json:"sdk_email"`
			SDKKey              string      `json:"sdk_key"`
			IsAvailable         bool        `json:"is_available"`
			Type                int         `json:"type"`
			AvatarURL           string      `json:"avatar_url"`
			AppID               int         `json:"app_id"`
			IsVerified          bool        `json:"is_verified"`
			NotificationsRoomID string      `json:"notifications_room_id"`
			BubbleColor         string      `json:"bubble_color"`
			QismoKey            string      `json:"qismo_key"`
			DirectLoginToken    interface{} `json:"direct_login_token"`
			LastLogin           string      `json:"last_login"`
			ForceOffline        bool        `json:"force_offline"`
			DeletedAt           string      `json:"deleted_at"`
			TypeAsString        string      `json:"type_as_string"`
			AssignedRules       interface{} `json:"assigned_rules"`
			App                 struct {
				ID                             int         `json:"id"`
				Name                           string      `json:"name"`
				AppCode                        string      `json:"app_code"`
				SecretKey                      string      `json:"secret_key"`
				CreatedAt                      string      `json:"created_at"`
				UpdatedAt                      string      `json:"updated_at"`
				BotWebhookURL                  string      `json:"bot_webhook_url"`
				IsBotEnabled                   bool        `json:"is_bot_enabled"`
				AllocateAgentWebhookURL        string      `json:"allocate_agent_webhook_url"`
				IsAllocateAgentWebhookEnabled  bool        `json:"is_allocate_agent_webhook_enabled"`
				MarkAsResolvedWebhookURL       string      `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool        `json:"is_mark_as_resolved_webhook_enabled"`
				IsMobilePnEnabled              bool        `json:"is_mobile_pn_enabled"`
				IsActive                       bool        `json:"is_active"`
				IsSessional                    bool        `json:"is_sessional"`
				IsAgentAllocationEnabled       bool        `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool        `json:"is_agent_takeover_enabled"`
				IsTokenExpiring                bool        `json:"is_token_expiring"`
				PaidChannelApproved            interface{} `json:"paid_channel_approved"`
				UseLatest                      bool        `json:"use_latest"`
				AppConfig                      interface{} `json:"app_config"`
				AgentRoles                     interface{} `json:"agent_roles"`
			} `json:"app"`
		} `json:"user"`
		Details        interface{} `json:"details"`
		LongLivedToken string      `json:"long_lived_token"`
		UserConfigs    interface{} `json:"user_configs"`
	} `json:"data"`
}

type MetabaseAuthResponse struct {
	ID string
}

type MetabaseDataResponse struct {
	Data struct {
		Rows [][]string `json:"rows"`
		Cols []struct {
			DisplayName string        `json:"display_name"`
			Source      string        `json:"source"`
			FieldRef    []interface{} `json:"field_ref"`
			Name        string        `json:"name"`
			BaseType    string        `json:"base_type"`
		} `json:"cols"`
		NativeForm struct {
			Query  string      `json:"query"`
			params interface{} `json:"params"`
		} `json:"native_form"`
		ResultsTimezone string `json:"results_timezone"`
		ResultsMetadata struct {
			Checksum string `json:"checksum"`
			Columns  []struct {
				Name         string      `json:"name"`
				DisplayName  string      `json:"display_name"`
				BaseType     string      `json:"base_type"`
				FieldRef     interface{} `json:"field_ref"`
				SemanticType interface{} `json:"semantic_type"`
				Fingerprint  interface{} `json:"fingerprint"`
			} `json:"columns"`
		} `json:"results_metadata"`
		Insights interface{} `json:"insights"`
	} `json:"data"`
	DatabaseID           int         `json:"database_id"`
	StartedAt            string      `json:"started_at"`
	JsonQuery            interface{} `json:"json_query"`
	AverageExecutionTime interface{} `json:"average_execution_time"`
	Status               string      `json:"status"`
	Context              string      `json:"context"`
	RowCount             int         `json:"row_count"`
	RunningTime          int         `json:"running_time"`
}
