import type { FC, ReactNode } from "react";
import {
  FormField,
  FormFieldType,
  formMenus,
  FormSection,
  FormTemplateModel,
} from "../models/form";
import {
  horizontalListSortingStrategy,
  SortableContext,
  useSortable,
} from "@dnd-kit/sortable";
import { DragOverlay, UniqueIdentifier } from "@dnd-kit/core";
import { BsTrash } from "react-icons/bs";
import { Dropdown, Label } from "flowbite-react";
import { IoMove } from "react-icons/io5";
import FormFieldComponent from "./FormFieldComponent";

interface FormSectionComponentProps {
  template: FormTemplateModel;
  section: FormSection;
  onChange: (template: FormTemplateModel) => void;
  onDeleteSection: (section: FormSection) => void;
  renderIcon: (key: FormFieldType, size: number) => ReactNode;
  addField: (section: FormSection, key: FormFieldType) => void;
  isDragged: boolean;
}

const FormSectionComponent: FC<FormSectionComponentProps> = ({
  template,
  section,
  onChange,
  renderIcon,
  addField,
  onDeleteSection,
  isDragged,
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
    id: section.id as UniqueIdentifier,
  });

  return (
    <div
      key={section.id}
      className="flex flex-col p-4 border-2 rounded-lg bg-white group/item"
      {...attributes}
      ref={setNodeRef}
    >
      <div className="flex justify-between  ">
        <input
          value={section.section_title}
          className="text-2xl font-bold"
          onChange={(el) => {
            section.section_title = el.target.value;
            onChange({
              ...template!,
              sections: template!.sections.map((s) =>
                s.id === section.id ? section : s
              ),
            });
          }}
        />
        <div className="flex items-center gap-2 group/edit invisible group-hover/item:visible">
          <Dropdown label="Fields" inline>
            {formMenus.map((menu) => (
              <Dropdown.Item
                key={menu.key}
                className="flex gap-2"
                onClick={() => addField(section, menu.key)}
              >
                {renderIcon(menu.key, 16)} {menu.text}
              </Dropdown.Item>
            ))}
          </Dropdown>
          <BsTrash
            className="text-red-400 hover:text-red-600"
            onClick={() => {
              onDeleteSection(section);
            }}
          />
          <IoMove {...listeners} />
        </div>
      </div>
      <textarea
        className="border-0 p-0 outline-none focus:ring-0"
        placeholder="add description"
        value={section.description}
        onChange={(el) => {
          section.description = el.target.value;
          onChange({
            ...template!,
            sections: template!.sections.map((s) =>
              s.id === section.id ? section : s
            ),
          });
        }}
      />
      {!isDragged && (
        <div className="fields flex flex-col space-y-4 my-4">
          <SortableContext
            id={section.id!}
            items={(section.fields ?? []).map((item) => ({
              id: item.id as UniqueIdentifier,
            }))}
            strategy={horizontalListSortingStrategy}
          >
            {section.fields.map((field) => (
              <FormFieldComponent
                field={field}
                template={template}
                onChange={onChange}
                key={field.id}
                renderIcon={renderIcon}
                section={section}
              />
            ))}
         
          </SortableContext>
        </div>
      )}
    </div>
  );
};
export default FormSectionComponent;
