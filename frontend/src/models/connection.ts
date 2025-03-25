import { GeminiAgent } from "./gemini";

export interface ConnectionModel {
  id?: string;
  name?: string;
  description?: string;
  type?: string;
  username?: string;
  session_name?: string;
  password?: string;
  channel_id?: string;
  api_key?: string;
  api_secret?: string;
  access_token?: string;
  refresh_token?: string;
  status?: string;
  gemini_agent_id?: string;
  gemini_agent?: GeminiAgent;
  is_auto_pilot?: boolean;
  session_auth?: boolean;
  connected?: boolean;
}
