import type { FC, ReactNode } from "react";
import { FormField, FormFieldType } from "../models/form";
import { BsTrash } from "react-icons/bs";
import { IoMove } from "react-icons/io5";
import { Button, Label, ToggleSwitch } from "flowbite-react";
import { useSortable } from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import { RxDrawingPin, RxDrawingPinFilled } from "react-icons/rx";

interface TaskAttributeFieldComponentProps {
  field: FormField;
  renderIcon: (key: FormFieldType, size: number) => ReactNode;
  onChange: (field: FormField) => void;
  removeField: (field: FormField) => void;
}

const TaskAttributeFieldComponent: FC<TaskAttributeFieldComponentProps> = ({
  field,
  renderIcon,
  onChange,
  removeField,
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

  return (
    <div
      className="flex gap-2 flex-col space-y-2 bg-gray-50 p-4 rounded-lg hover:bg-gray-100 cursor-pointer"
      {...attributes}
      ref={setNodeRef}
    >
      <div className="flex items-center gap-2 justify-between">
        <div className="flex flex-row gap-2 items-center">
          {renderIcon(field.type, 16)}{" "}
          <input
            className="min-w-[600px] bg-transparent"
            value={field.label}
            onChange={(el) => {
              field.label = el.target.value;
              onChange(field);
            }}
          />
        </div>
        <div className="flex gap-2">
          {field.is_pinned ? (
            <RxDrawingPinFilled
              style={{}}
              className=" cursor-pointer text-green-400"
              onClick={() => {
                field.is_pinned = !field.is_pinned;
                onChange(field);
              }}
            />
          ) : (
            <RxDrawingPin
              style={{}}
              className=" cursor-pointer"
              onClick={() => {
                field.is_pinned = !field.is_pinned;
                onChange(field);
              }}
            />
          )}

          <BsTrash
            className="cursor-pointer text-red-400 hover:text-red-600"
            onClick={() => {
              removeField(field);
            }}
          />
          <IoMove {...listeners} />
        </div>
      </div>
      <div className="grid grid-cols-3 gap-2">
        {(field.type === FormFieldType.TextField ||
          field.type === FormFieldType.TextArea ||
          field.type === FormFieldType.EmailField ||
          field.type === FormFieldType.Currency ||
          field.type === FormFieldType.PasswordField ||
          field.type === FormFieldType.NumberField) && (
          <div className="flex flex-col">
            <Label value="Placeholder" className="text-sm font-medium px-2" />
            <input
              placeholder="add placeholder"
              value={field.placeholder}
              onChange={(el) => {
                field.placeholder = el.target.value;
                onChange(field);
              }}
              className="px-2 py-1 border rounded-md"
            />
          </div>
        )}
        {(field.type === FormFieldType.TextField ||
          field.type === FormFieldType.TextArea ||
          field.type === FormFieldType.Currency ||
          field.type === FormFieldType.EmailField ||
          field.type === FormFieldType.NumberField) && (
          <div className="flex flex-col">
            <Label value="Default Value" className="text-sm font-medium px-2" />
            <input
              placeholder="add default value"
              value={field.default_value}
              onChange={(el) => {
                field.default_value = el.target.value;
                onChange(field);
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
              onChange(field);
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
                  className="px-2 py-1 border rounded-md "
                  onChange={(val) => {
                    e.label = val.target.value;
                    field.options[i] = e;
                    onChange(field);
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
                    onChange(field);
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
                onChange(field);
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
            onChange(field);
          }}
        />
        <ToggleSwitch
          sizing="sm"
          checked={!field.disabled}
          label="Enable"
          onChange={(val) => {
            field.disabled = !val;
            onChange(field);
          }}
        />
      </div>
    </div>
  );
};
export default TaskAttributeFieldComponent;
