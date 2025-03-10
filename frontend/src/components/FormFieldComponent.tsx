import type { FC, ReactNode } from "react";
import {
  FormField,
  FormFieldType,
  FormSection,
  FormTemplateModel,
} from "../models/form";
import { BsTrash } from "react-icons/bs";
import { Button, Label, ToggleSwitch } from "flowbite-react";
import { useSortable } from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import { IoMove } from "react-icons/io5";

interface FormFieldComponentProps {
  field: FormField;
  section: FormSection;
  template: FormTemplateModel;
  onChange: (template: FormTemplateModel) => void;
  renderIcon: (key: FormFieldType, size: number) => ReactNode;
}

const FormFieldComponent: FC<FormFieldComponentProps> = ({
  field,
  section,
  template,
  onChange,
  renderIcon,
}) => {
  const {
    isOver,
    setNodeRef,
    setActivatorNodeRef,
    attributes,
    listeners,
    transform,
    transition,
  } = useSortable({
    id: field.id as UniqueIdentifier,
  });

  const removeField = (section: FormSection, field: FormField) => {
    section.fields = section.fields.filter((f) => f.id !== field.id);
    onChange({
      ...template!,
      sections: template!.sections.map((s) =>
        s.id === section.id ? section : s
      ),
    });
  };

  return (
    <div
      key={field.id}
      className="flex gap-2 flex-col space-y-2 bg-gray-50 p-4 rounded-lg hover:bg-gray-100 cursor-pointer"
      {...attributes}
      ref={setNodeRef}
    >
      <div className="flex items-center gap-2 justify-between">
        <div className="flex flex-row gap-2 items-center">
          {renderIcon(field.type, 16)}{" "}
          <input
            value={field.label}
            onChange={(el) => {
              field.label = el.target.value;
              section.fields = section.fields.map((f) =>
                f.id === field.id ? field : f
              );
              onChange({
                ...template!,
                sections: template!.sections.map((s) =>
                  s.id === section.id ? section : s
                ),
              });
            }}
          />
        </div>
        <div className="flex gap-2">
          <BsTrash
            className="cursor-pointer text-red-400 hover:text-red-600"
            onClick={() => removeField(section, field)}
          />
          <IoMove {...listeners} />
        </div>
      </div>
      <div className="grid grid-cols-3 gap-2">
        {(field.type === FormFieldType.TextField ||
          field.type === FormFieldType.TextArea || 
          field.type === FormFieldType.EmailField || 
          field.type === FormFieldType.NumberField 
        ) && (
          <div className="flex flex-col">
            <Label value="Placeholder" className="text-sm font-medium px-2" />
            <input
              placeholder="add placeholder"
              value={field.placeholder}
              onChange={(el) => {
                field.placeholder = el.target.value;
                section.fields = section.fields.map((f) =>
                  f.id === field.id ? field : f
                );
                onChange({
                  ...template!,
                  sections: template!.sections.map((s) =>
                    s.id === section.id ? section : s
                  ),
                });
              }}
              className="px-2 py-1 border rounded-md"
            />
          </div>
        )}
        {(field.type === FormFieldType.TextField ||
          field.type === FormFieldType.TextArea || 
          field.type === FormFieldType.EmailField || 
          field.type === FormFieldType.NumberField ) && (
          <div className="flex flex-col">
            <Label value="Default Value" className="text-sm font-medium px-2" />
            <input
              placeholder="add default value"
              value={field.default_value}
              onChange={(el) => {
                field.default_value = el.target.value;
                section.fields = section.fields.map((f) =>
                  f.id === field.id ? field : f
                );
                onChange({
                  ...template!,
                  sections: template!.sections.map((s) =>
                    s.id === section.id ? section : s
                  ),
                });
              }}
              className="px-2 py-1 border rounded-md"
            />
          </div>
        )}
        <div className="flex flex-col">
          <Label value="Help Text" className="text-sm font-medium px-2" />
          <input
            placeholder="add help text"
            value={field.help_text}
            onChange={(el) => {
              field.help_text = el.target.value;
              section.fields = section.fields.map((f) =>
                f.id === field.id ? field : f
              );
              onChange({
                ...template!,
                sections: template!.sections.map((s) =>
                  s.id === section.id ? section : s
                ),
              });
            }}
            className="px-2 py-1 border rounded-md"
          />
        </div>
      </div>
      {(field.type === FormFieldType.Checkbox ||
        field.type === FormFieldType.RadioButton ||
        field.type === FormFieldType.Dropdown) && (
        <div className="flex flex-col p-2 border rounded-lg bg-white">
          <h3 className="font-semibold mb-2 px-2">Options</h3>
          {field.options.map((e, i) => (
            <div className="grid grid-cols-3  gap-2 mb-2 last:mb-0" key={i}>
              <div className="flex flex-col">
                <Label value={"Label"} className="text-sm font-medium px-2" />
                <input
                  value={e.label}
                  className="px-2 py-1 border rounded-md"
                  onChange={(val) => {
                    e.label = val.target.value;
                    field.options[i] = e;
                    section.fields = section.fields.map((f) =>
                      f.id === field.id ? field : f
                    );
                    onChange({
                      ...template!,
                      sections: template!.sections.map((s) =>
                        s.id === section.id ? section : s
                      ),
                    });
                  }}
                />
              </div>
              <div className="flex flex-col">
                <Label value={"Value"} className="text-sm font-medium px-2" />
                <input
                  value={e.value}
                  className="px-2 py-1 border rounded-md"
                  onChange={(val) => {
                    e.value = val.target.value;
                    field.options[i] = e;
                    section.fields = section.fields.map((f) =>
                      f.id === field.id ? field : f
                    );
                    onChange({
                      ...template!,
                      sections: template!.sections.map((s) =>
                        s.id === section.id ? section : s
                      ),
                    });
                  }}
                />
              </div>
            </div>
          ))}
          <div className="mt-4">
            <Button
              size="xs"
              onClick={() => {
                field.options.push({ label: "New Label", value: "New Value" });
                section.fields = section.fields.map((f) =>
                  f.id === field.id ? field : f
                );
                onChange({
                  ...template!,
                  sections: template!.sections.map((s) =>
                    s.id === section.id ? section : s
                  ),
                });
              }}
            >
              Add Option
            </Button>
          </div>
        </div>
      )}
      <div className="flex gap-8">
        <ToggleSwitch
          sizing="sm"
          checked={field.required}
          label="Required"
          onChange={(val) => {
            field.required = val;
            section.fields = section.fields.map((f) =>
              f.id === field.id ? field : f
            );
            onChange({
              ...template!,
              sections: template!.sections.map((s) =>
                s.id === section.id ? section : s
              ),
            });
          }}
        />
        <ToggleSwitch
          sizing="sm"
          checked={!field.disabled}
          label="Enable"
          onChange={(val) => {
            field.disabled = !val;
            section.fields = section.fields.map((f) =>
              f.id === field.id ? field : f
            );
            onChange({
              ...template!,
              sections: template!.sections.map((s) =>
                s.id === section.id ? section : s
              ),
            });
          }}
        />
      </div>
    </div>
  );
};
export default FormFieldComponent;
