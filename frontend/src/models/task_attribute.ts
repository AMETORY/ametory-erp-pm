import { FormField, FormFieldType } from "./form";

export interface TaskAttributeModel {
  id: string;
  title: string;
  description: string;
  fields: FormField[];
}



export const attributeMenus = [
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
    key: FormFieldType.Price,
    text: "Price",
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
  {
    key: FormFieldType.Product,
    text: "Product",
  },
  {
    key: FormFieldType.Contact,
    text: "Contact",
  },
];