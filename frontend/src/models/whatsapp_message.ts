import { CompanyModel } from "./company";
import { ContactModel } from "./contact";

export interface WhatsappMessageModel {
  id?: string;
  jid?: string;
  sender?: string;
  receiver?: string;
  message?: string;
  media_url?: string;
  mime_type?: string;
  session?: string;
  info?: string;
  message_id?: string;
  message_info?: any;
  contact_id?: string;
  contact?: ContactModel;
  company_id?: string;
  company?: CompanyModel;
  is_from_me?: boolean;
  is_group?: boolean;
  sent_at?: Date;
  is_read?: boolean;
}

export interface WhatsappMessageSessionModel {
  id?: string;
  jid?: string;
  session?: string;
  session_name?: string;
  last_online_at?: Date;
  last_message?: string;
  company_id?: string;
  company?: CompanyModel;
  contact_id?: string;
  contact?: ContactModel;
  is_human_agent?: boolean;
  count_unread?: number;
}
