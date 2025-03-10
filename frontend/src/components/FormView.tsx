import type { FC } from "react";
import { FormField, FormFieldType, FormSection } from "../models/form";
import {
    Button,
  Checkbox,
  Datepicker,
  FileInput,
  Label,
  Radio,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";

interface FormViewProps {
  sections: FormSection[];
}

const FormView: FC<FormViewProps> = ({ sections }) => {
  const renderField = (field: FormField) => {
    switch (field.type) {
      case FormFieldType.TextField:
        return (
          <div>
            <TextInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.PasswordField:
        return (
          <div>
            <TextInput
              type="password"
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.FileUpload:
        return (
          <div>
            <FileInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.EmailField:
        return (
          <div>
            <TextInput
              type={"email"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.NumberField:
        return (
          <div>
            <TextInput
              type={"number"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.Currency:
        return (
          <div>
            <TextInput
              type={"number"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.TextArea:
        return (
          <div>
            <Textarea
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DatePicker:
        return (
          <div>
            <Datepicker
              placeholder={field.placeholder}
              required={field.required}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DateRangePicker:
        return (
          <div className="">
            <div className="grid grid-cols-2 gap-2">
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
              />
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
              />
            </div>
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.RadioButton:
        return (
          <div>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}
              {field.options.map((option, i) => (
                <div className="flex items-center gap-2">
                  <Radio
                    id={`${field.label}-${i}`}
                    name={field.label}
                    value={option.value}
                  />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.Checkbox:
        return (
          <div>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}
              {field.options.map((option, i) => (
                <div className="flex items-center gap-2">
                  <Checkbox id={`${field.label}-${i}`} value={option.value} />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.ToggleSwitch:
        return (
          <div>
            <ToggleSwitch
              sizing="sm"
              checked={false}
              label={field.help_text}
              onChange={(val) => {
                console.log(val);
              }}
            />
          </div>
        );
    }

    return <div></div>;
  };
  return (
    <div className="flex flex-col justify-center w-1/2 space-y-4 ">
      {sections.map((section, index) => (
        <div
          className="bg-white p-4 rounded-lg border border-t-4 border-t-blue-400"
          key={section.id}
        >
          <h1 className="text-2xl font-bold">{section.section_title}</h1>
          <div className="text-md text-gray-600">{section?.description}</div>
          <div className="flex flex-col space-y-4 mt-4">
            {section.fields.map((field) => (
              <div key={field.id}>
                <Label className="font-bold text-md">{field.label}</Label>
                {renderField(field)}
              </div>
            ))}
          </div>
        </div>
      ))}
      <Button>Submit Form</Button>
    </div>
  );
};
export default FormView;
