export interface GeminiAgent {
  id?: string;
  name?: string;
  api_key?: string;
  active?: boolean;
  system_instruction?: string;
  model?: string;
  set_temperature?: number;
  set_top_k?: number;
  set_top_p?: number;
  set_max_output_tokens?: number;
  response_mimetype?: string;
}


