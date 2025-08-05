export interface AgentModel {
  id?: string;
  name?: string;
  api_key?: string;
  host?: string;
  active?: boolean;
  system_instruction?: string;
  agent_type?: string;
  model?: string;
  // set_temperature?: number;
  // set_top_k?: number;
  // set_top_p?: number;
  // set_max_output_tokens?: number;
  // response_mimetype?: string;
}

export interface AgentHistoryModel {
  id: string;
  created_at: string;
  input: string;
  output: string;
  file_url: string;
  mime_type: string;
  agent_id: string;
  is_model: boolean;
  session_code: string | null;
}
