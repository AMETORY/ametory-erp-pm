import { ConnectionModel } from "./connection";
import { ContactModel } from "./contact";
import { TemplateModel } from "./template";

export interface BroadcastModel {
  id?: string;
  description?: string;
  message?: string;
  company_id?: string;
  status?: string;
  max_contacts_per_batch?: number;
  scheduled_at?: Date | null;
  connections: ConnectionModel[];
  contacts: ContactModel[];
  groups: any[];
  contact_count?: number;
  failed_count?: number;
  group_count?: number;
  success_count?: number;
  completed_count?: number;
  template_id?: string;
  template?: TemplateModel;
  delay_time?: number;
}

export interface MessageLog {
  id?: string;
  created_at?: Date | null;
  broadcast_id?: string;
  broadcast?: BroadcastModel;
  contact_id?: string;
  contact?: ContactModel;
  sender_id?: string;
  sender?: ConnectionModel;
  status?: string;
  error_message?: string;
  sent_at?: Date | null;
}


export interface MessageRetry {
  id?: string;
  created_at?: Date | null;
  broadcast_id?: string;
  broadcast?: BroadcastModel;
  contact_id?: string;
  contact?: ContactModel;
  sender_id?: string;
  sender?: ConnectionModel;
  attempt?: number;
  last_error?: string;
  last_tried_at?: Date | null;
}
