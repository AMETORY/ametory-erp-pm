import { ConnectionModel } from "./connection";
import { ContactModel } from "./contact";

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
}
