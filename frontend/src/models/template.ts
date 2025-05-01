import { FileModel } from "./file";
import { MemberModel } from "./member";
import { ProductModel } from "./product";

export interface TemplateModel {
  id?: string;
  title?: string;
  description?: string;
  company_id?: string;
  user_id?: string;
  member_id?: string | null;
  member?: MemberModel | null;
  messages?: MessageTemplate[];
}



export interface MessageTemplate {
  id?: string;
  whatsapp_message_template_id?: string;
  type?: string;
  header?: string;
  body?: string;
  footer?: string;
  button_text?: string;
  button_url?: string;
  files?: FileModel[];
  products?: ProductModel[];
}