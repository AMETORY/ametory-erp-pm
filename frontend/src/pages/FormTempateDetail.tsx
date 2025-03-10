import {
  DndContext,
  DragEndEvent,
  DragOverEvent,
  DragOverlay,
  DragStartEvent,
  KeyboardSensor,
  PointerSensor,
  UniqueIdentifier,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  SortableContext,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { Button, Dropdown, Label, Tabs } from "flowbite-react";
import { ReactNode, useEffect, useRef, useState, type FC } from "react";
import toast from "react-hot-toast";
import { BsBuilding, BsCurrencyDollar, BsEye, BsTrash } from "react-icons/bs";
import { MdOutlineAlternateEmail } from "react-icons/md";
import {
  PiCalendar,
  PiCalendarPlus,
  PiPassword,
  PiRadioButtonFill,
} from "react-icons/pi";
import {
  RiDropdownList,
  RiFileUploadLine,
  RiText,
  RiTextWrap,
} from "react-icons/ri";
import { RxSwitch } from "react-icons/rx";
import { TbCheckbox, TbNumber12Small } from "react-icons/tb";
import { useParams } from "react-router-dom";
import AdminLayout from "../components/layouts/admin";
import {
  FormField,
  FormFieldType,
  formMenus,
  FormSection,
  FormTemplateModel,
} from "../models/form";
import { getFormTemplate, updateFormTemplate } from "../services/api/formApi";
import { IoMove } from "react-icons/io5";
import FormSectionComponent from "../components/FormSectionComponent";
import FormFieldComponent from "../components/FormFieldComponent";
import FormView from "../components/FormView";

interface FormTempateDetailProps {}

const FormTempateDetail: FC<FormTempateDetailProps> = ({}) => {
  const { templateId } = useParams();
  const [mounted, setMounted] = useState(false);
  const [template, setTemplate] = useState<FormTemplateModel>();
  const [activeTab, setActiveTab] = useState(0);
  const [activeSection, setActiveSection] = useState<FormSection>();
  const [activeField, setActiveField] = useState<FormField>();
  const timeout = useRef<number | null>(null);
  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getFormTemplate(templateId!)
        .then((resp: any) => {
          setTemplate(resp.data);
        })
        .catch(toast.error);
    }
  }, [mounted]);

  const renderIcon = (e: FormFieldType, size: number): ReactNode => {
    switch (e) {
      case FormFieldType.TextField:
        return <RiText size={size} />;
      case FormFieldType.TextArea:
        return <RiTextWrap size={size} />;
      case FormFieldType.RadioButton:
        return <PiRadioButtonFill size={size} />;
      case FormFieldType.Checkbox:
        return <TbCheckbox size={size} />;
      case FormFieldType.DatePicker:
        return <PiCalendar size={size} />;
      case FormFieldType.DateRangePicker:
        return <PiCalendarPlus size={size} />;
      case FormFieldType.NumberField:
        return <TbNumber12Small size={size} />;
      case FormFieldType.EmailField:
        return <MdOutlineAlternateEmail size={size} />;
      case FormFieldType.PasswordField:
        return <PiPassword size={size} />;
      case FormFieldType.Currency:
        return <BsCurrencyDollar size={size} />;
      case FormFieldType.FileUpload:
        return <RiFileUploadLine size={size} />;
      case FormFieldType.ToggleSwitch:
        return <RxSwitch size={size} />;
      case FormFieldType.Dropdown:
        return <RiDropdownList size={size} />;
    }

    return null;
  };

  const addField = (section: FormSection, e: FormFieldType) => {
    let data: FormField = {
      id: crypto.randomUUID(),
      label: "New Field",
      type: e,
      options: [],
      required: false,
      is_multi: false,
      placeholder: "",
      default_value: "",
      help_text: "",
      disabled: false,
    };

    if (
      e == FormFieldType.Checkbox ||
      e == FormFieldType.RadioButton ||
      e == FormFieldType.Dropdown
    ) {
      data.options = [{ label: "New Label", value: "New Value" }];
    }
    section.fields.push(data);

    setTemplate({
      ...template!,
      sections: template!.sections.map((s) =>
        s.id === section.id ? section : s
      ),
    });
  };

  useEffect(() => {
    if (template) {
      if ((template.sections ?? []).length > 0) {
        if (timeout.current) {
          window.clearTimeout(timeout.current);
        }
        timeout.current = window.setTimeout(() => {
          updateFormTemplate(templateId!, {
            ...template,
            data: JSON.stringify(template.sections),
          });
        }, 1000);
      }
    }
  }, [template?.sections]);

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        delay: 0,
        tolerance: 0,
      },
    }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );
  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;
    const { id } = active;
    let isSection: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");

    if (event.over) {
      if (isSection) {
        const activeIndex = template!.sections.findIndex(
          (item) => item.id == id
        );
        const overIndex = template!.sections.findIndex(
          (item) => item.id == over?.id
        );

        if (
          activeIndex !== -1 &&
          overIndex !== -1 &&
          activeIndex !== overIndex
        ) {
          let columnsBefore = [...template!.sections];
          const movedColumn = template!.sections.splice(activeIndex, 1)[0];
          template!.sections.splice(overIndex, 0, movedColumn);
          setTemplate({
            ...template!,
            sections: template!.sections,
          });

          console.log(activeIndex);

          updateFormTemplate(templateId!, {
            ...template,
            data: JSON.stringify(template!.sections),
          }).then(() => {
            setActiveSection(undefined);
            setActiveField(undefined);
          });
        }
      } else {
        const activeSection = template!.sections.find(
          (section) => section.id === active.data.current?.sortable?.containerId
        );

        const overSection = template!.sections.find(
          (section) => section.id === over?.data.current?.sortable?.containerId
        );

        if (activeSection && overSection) {
          const activeIndex = (activeSection.fields ?? []).findIndex(
            (item) => item.id === id
          );
          const overIndex = (overSection.fields ?? []).findIndex(
            (item) => item.id === over?.id
          );
          const item = (activeSection.fields ?? []).splice(activeIndex, 1)[0];
          (overSection.fields ?? []).splice(overIndex, 0, item);
          // console.log(item);
          // Reload the columns to trigger a re-render

          setTemplate({
            ...template!,
            sections: [
              ...template!.sections.slice(
                0,
                template!.sections.indexOf(activeSection)
              ),
              activeSection,
              ...template!.sections.slice(
                template!.sections.indexOf(activeSection) + 1
              ),
            ],
          });

          if (overSection.id != activeSection.id) {
            console.log("ID SECTION BERBEDA");
          } else {
            console.log("ID SECTION SAMA");
          }
        }
        setActiveField(undefined);
        setActiveSection(undefined);
      }
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    const { id } = active;
  };
  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const { id } = active;
    let isSection: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");
    if (isSection) {
      const activeIndex = template!.sections.findIndex((item) => item.id == id);
      setActiveSection(template!.sections[activeIndex]);
    } else {
      let sectionId = event.active.data.current?.sortable
        ?.containerId as string;
      let itemId = id as string;
      setActiveSection(
        template!.sections.find((section) => section.id == sectionId)
      );
      setActiveField(
        template!.sections
          .find((section) => section.id == sectionId)!
          .fields.find((field) => field.id == itemId)
      );
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4 ">
          <h1 className="text-3xl font-bold ">{template?.title}</h1>
        </div>
        <Tabs
          aria-label="Default tabs"
          variant="default"
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
          className="overflow-y-auto"
        >
          <Tabs.Item title="Builder" active={activeTab == 0} icon={BsBuilding}>
            <div className="flex justify-end ">
              <Button
                gradientDuoTone="purpleToBlue"
                pill
                onClick={() => {
                  setTemplate({
                    ...template!,
                    sections: [
                      ...(template?.sections ?? []),
                      {
                        id: crypto.randomUUID(),
                        section_title: "New Section",
                        description: "",
                        fields: [],
                      },
                    ],
                  });
                }}
              >
                + Create new section
              </Button>
            </div>
            <div className=" overflow-y-auto h-[calc(100vh-300px)]">
              <DndContext
                onDragEnd={handleDragEnd}
                onDragStart={handleDragStart}
                onDragOver={handleDragOver}
                sensors={sensors}
              >
                <SortableContext
                  items={(template?.sections ?? []).map((val) => ({
                    id: val.id as UniqueIdentifier,
                  }))}
                  strategy={verticalListSortingStrategy}
                >
                  <div className="flex flex-col space-y-4 my-4">
                    {(template?.sections ?? []).map((section) => (
                      <FormSectionComponent
                        isDragged={false}
                        template={template!}
                        section={section}
                        key={section.id}
                        onChange={(val) => setTemplate(val)}
                        renderIcon={renderIcon}
                        addField={addField}
                        onDeleteSection={(val) => {
                          setTemplate({
                            ...template!,
                            sections: template!.sections.filter(
                              (s) => s.id !== val.id
                            ),
                          });
                        }}
                      />
                    ))}
                  </div>
                  <DragOverlay>
                    {activeSection && !activeField && (
                      <FormSectionComponent
                        isDragged={true}
                        template={template!}
                        section={activeSection}
                        key={activeSection.id}
                        onChange={(val) => setTemplate(val)}
                        renderIcon={renderIcon}
                        addField={addField}
                        onDeleteSection={(val) => {
                          setTemplate({
                            ...template!,
                            sections: template!.sections.filter(
                              (s) => s.id !== val.id
                            ),
                          });
                        }}
                      />
                    )}
                    {activeSection && activeField && (
                      <FormFieldComponent
                        field={activeField}
                        template={template!}
                        onChange={(val) => {}}
                        key={activeField.id}
                        renderIcon={renderIcon}
                        section={activeSection}
                      />
                    )}
                  </DragOverlay>
                </SortableContext>
              </DndContext>
            </div>
          </Tabs.Item>
          <Tabs.Item title="Preview" active={activeTab == 1} icon={BsEye}>
            <div className="bg-gray-100 flex flex-col  items-center p-16 overflow-y-auto h-[calc(100vh-240px)]">
              <FormView sections={template?.sections ?? []} />
            </div>
          </Tabs.Item>
        </Tabs>
      </div>
    </AdminLayout>
  );
};
export default FormTempateDetail;
