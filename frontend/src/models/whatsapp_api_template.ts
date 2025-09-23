export interface WhatsappAPITemplate {
  category: string;
  components: WhatsappTemplateComponent[];
  id: string;
  language: string;
  name: string;
  parameter_format: string;
  status: string;
  sub_category: string;
}

export interface WhatsappTemplateComponent {
  example?: WhatsappTemplateExample;
  format?: string;
  type: string;
  text?: string;
}

export interface WhatsappTemplateExample {
  header_handle?: string[];
  body_text_named_params?: WhatsappTemplateBodyTextNamedParam[];
  body_text?: string[][];
}

export interface WhatsappTemplateBodyTextNamedParam {
  example: string;
  param_name: string;
}
