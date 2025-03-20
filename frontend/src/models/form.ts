import { UserModel } from "./user";
import { MemberModel } from "./member";
import { ColumnModel } from "./column";
import { ProjectModel } from "./project";
import { FileModel } from "./file";

export interface FormTemplateModel {
  id: string;
  title: string;
  company_id: string;
  created_by_id: string;
  created_by: UserModel;
  created_by_member_id: string;
  created_by_member: MemberModel;
  sections: FormSection[];
  data?: any;
}

export const enum FormFieldType {
  TextField = "text",
  TextArea = "textarea",
  RadioButton = "radio",
  Checkbox = "checkbox",
  Dropdown = "dropdown",
  DatePicker = "date",
  DateRangePicker = "date_range",
  NumberField = "number",
  Currency = "currency",
  Price = "price",
  EmailField = "email",
  PasswordField = "password",
  FileUpload = "file",
  Product = "product",
  Contact = "contact",
  ToggleSwitch = "toggle",
}

export interface FormFieldOption {
  label: string;
  value: string;
}

export interface FormField {
  id: string;
  label: string;
  type: FormFieldType;
  options: FormFieldOption[];
  required: boolean;
  is_multi: boolean;
  is_pinned?: boolean;
  placeholder: string;
  default_value: string;
  help_text: string;
  disabled: boolean;
  value?: any;
}

export interface FormSection {
  id: string;
  section_title: string;
  description: string;
  fields: FormField[];
}

export const formMenus = [
  {
    key: FormFieldType.TextField,
    text: "Text Field",
  },
  {
    key: FormFieldType.TextArea,
    text: "Text Area",
  },
  {
    key: FormFieldType.RadioButton,
    text: "Radio Button",
  },
  {
    key: FormFieldType.Checkbox,
    text: "Checkbox",
  },
  {
    key: FormFieldType.Dropdown,
    text: "Dropdown",
  },
  {
    key: FormFieldType.DatePicker,
    text: "Date Picker",
  },
  {
    key: FormFieldType.DateRangePicker,
    text: "Date Range Picker",
  },
  {
    key: FormFieldType.NumberField,
    text: "Number Field",
  },
  {
    key: FormFieldType.Currency,
    text: "Currency",
  },
  {
    key: FormFieldType.EmailField,
    text: "Email Field",
  },
  {
    key: FormFieldType.PasswordField,
    text: "Password Field",
  },
  {
    key: FormFieldType.FileUpload,
    text: "File Upload",
  },
  {
    key: FormFieldType.ToggleSwitch,
    text: "Toggle Switch",
  },
];

export interface FormModel {
  id?: string;
  code?: string;
  title?: string;
  cover?: FileModel;
  description?: string;
  submit_url?: string;
  method?: string;
  headers?: string;
  is_public?: boolean;
  status?: string;
  form_template_id?: string;
  form_template?: FormTemplateModel;
  created_by_id?: string;
  created_by?: UserModel;
  created_by_member_id?: string;
  created_by_member?: MemberModel;
  company_id?: string;
  column_id?: string;
  column?: ColumnModel;
  project_id?: string;
  project?: ProjectModel;
  responses?: FormData[];
}

export interface FormData {
  id: string;
  form_id: string;
  form: object;
  sections: FormSection[];
  metadata: string;
  ref_id: string;
}
