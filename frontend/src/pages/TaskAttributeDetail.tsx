import {
  ReactNode,
  useContext,
  useEffect,
  useRef,
  useState,
  type FC,
} from "react";
import AdminLayout from "../components/layouts/admin";
import { useParams } from "react-router-dom";
import {
  getTaskAttributeDetail,
  getTaskAttributes,
  updateTaskAttribute,
} from "../services/api/taskAttributeApi";
import { attributeMenus, TaskAttributeModel } from "../models/task_attribute";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import {
  Button,
  Dropdown,
  Label,
  Modal,
  Textarea,
  TextInput,
} from "flowbite-react";
import { FormField, FormFieldType, formMenus } from "../models/form";
import {
  RiDropdownList,
  RiFileUploadLine,
  RiText,
  RiTextWrap,
} from "react-icons/ri";
import {
  PiCalendar,
  PiCalendarPlus,
  PiPassword,
  PiRadioButtonFill,
} from "react-icons/pi";
import { TbCheckbox, TbNumber12Small } from "react-icons/tb";
import { MdOutlineAlternateEmail } from "react-icons/md";
import { BsCart, BsCurrencyDollar } from "react-icons/bs";
import { RxSwitch } from "react-icons/rx";
import { LuContact2 } from "react-icons/lu";
import { randomString } from "../utils/helper";
import FormFieldComponent from "../components/FormFieldComponent";
import TaskAttributeFieldComponent from "../components/TaskAttributeFieldComponent";
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
import { error } from "console";

interface TaskAttributeDetailProps {}

const TaskAttributeDetail: FC<TaskAttributeDetailProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const { attributeId } = useParams();
  const [mounted, setMounted] = useState(false);
  const [taskAttribute, setTaskAttribute] = useState<TaskAttributeModel>();
  const [uniqeKey, setUniqeKey] = useState("");
  const timeout = useRef<number | null>(null);
  const [activeField, setActiveField] = useState<FormField>();
  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted && attributeId) {
      setLoading(true);
      getTaskAttributeDetail(attributeId)
        .then((e: any) => setTaskAttribute(e.data))
        .catch(toast.error)
        .finally(() => setLoading(false));
    }
  }, [mounted, attributeId]);

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
      case FormFieldType.Price:
      case FormFieldType.Currency:
        return <BsCurrencyDollar size={size} />;
      case FormFieldType.FileUpload:
        return <RiFileUploadLine size={size} />;
      case FormFieldType.ToggleSwitch:
        return <RxSwitch size={size} />;
      case FormFieldType.Product:
        return <BsCart size={size} />;
      case FormFieldType.Contact:
        return <LuContact2 size={size} />;
      case FormFieldType.Dropdown:
        return <RiDropdownList size={size} />;
    }

    return null;
  };

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

  const addField = (e: FormFieldType) => {
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

    setTaskAttribute({
      ...taskAttribute!,
      fields: [...taskAttribute!.fields, data],
    });
    setUniqeKey(randomString(5));
  };

  useEffect(() => {
    if (taskAttribute) {
      if (timeout.current) {
        window.clearTimeout(timeout.current);
      }
      timeout.current = window.setTimeout(() => {
        updateTaskAttribute(attributeId!, {
          ...taskAttribute,
          data: JSON.stringify(taskAttribute.fields),
        });
      }, 1000);
    }
  }, [uniqeKey]);
  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;
    const { id } = active;
    let isSection: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");

    const activeSection = taskAttribute?.fields.find(
      (field) => field.id === id
    );

    const overSection = taskAttribute?.fields.find(
      (field) => field.id === over?.id
    );

    if (activeSection && overSection) {
      const activeIndex = (taskAttribute?.fields ?? []).findIndex(
        (item) => item.id === id
      );
      const overIndex = (taskAttribute?.fields ?? []).findIndex(
        (item) => item.id === over?.id
      );
      const item = (taskAttribute?.fields ?? []).splice(activeIndex, 1)[0];
      (taskAttribute?.fields ?? []).splice(overIndex, 0, item);

      // console.log(taskAttribute)
      setTaskAttribute(taskAttribute);
      setUniqeKey(randomString(5));
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    const { id } = active;
  };
  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const { id } = active;

    setActiveField(taskAttribute?.fields.find((e) => e.id == id));
  };

  return (
    <AdminLayout>
      <div className="p-8 h-full overflow-y-auto">
        <div className="flex justify-between items-center mb-4 ">
          <div>
            <h1 className="text-3xl font-bold ">{taskAttribute?.title}</h1>
            <small>{taskAttribute?.description}</small>
          </div>
          <div className="flex gap-2">
            <Dropdown label="Fields" inline>
              {attributeMenus.map((menu) => (
                <Dropdown.Item
                  key={menu.key}
                  className="flex gap-2"
                  onClick={() => {
                    addField(menu.key);
                  }}
                >
                  {renderIcon(menu.key, 16)} {menu.text}
                </Dropdown.Item>
              ))}
            </Dropdown>
            <Button
              color="gray"
              className=""
              size="xs"
              onClick={() => {
                setShowModal(true);
              }}
            >
              Edit
            </Button>
          </div>
        </div>
        <div className=" h-[calc(100vh - 160px)] ">
          <DndContext
            onDragEnd={handleDragEnd}
            onDragStart={handleDragStart}
            onDragOver={handleDragOver}
            sensors={sensors}
          >
            <SortableContext
              items={(taskAttribute?.fields ?? []).map((val) => ({
                id: val.id as UniqueIdentifier,
              }))}
              strategy={verticalListSortingStrategy}
            >
              <div className="flex flex-col space-y-4 my-4">
                {(taskAttribute?.fields ?? []).map((field) => (
                  <TaskAttributeFieldComponent
                    field={field}
                    key={field.id}
                    renderIcon={renderIcon}
                    removeField={(val) => {
                      setTaskAttribute({
                        ...taskAttribute!,
                        fields: (taskAttribute!.fields ?? []).filter(
                          (f) => f.id !== val.id
                        ),
                      });
                      setUniqeKey(randomString(5));
                    }}
                    onChange={(val) => {
                      setTaskAttribute({
                        ...taskAttribute!,
                        fields:
                          taskAttribute?.fields?.map((e) =>
                            e.id === field.id ? { ...e, ...val } : e
                          ) ?? [],
                      });
                      setUniqeKey(randomString(5));
                    }}
                  />
                ))}
              </div>
              <DragOverlay>
                {activeField && (
                  <TaskAttributeFieldComponent
                    field={activeField}
                    key={activeField?.id}
                    renderIcon={renderIcon}
                    removeField={(val) => {}}
                    onChange={(val) => {}}
                  />
                )}
              </DragOverlay>
            </SortableContext>
          </DndContext>
        </div>
      </div>
      <Modal show={showModal} onClose={() => setShowModal(false)}>
        <Modal.Header>Update attribute</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col space-y-4">
            <div>
              <Label htmlFor="title" value="Title" className="mb-1" />
              <TextInput
                id="title"
                type="text"
                value={taskAttribute?.title ?? ""}
                onChange={(e) => {
                  setTaskAttribute({
                    ...taskAttribute!,
                    title: e.target.value,
                  });
                }}
                placeholder="Title"
                className="mb-4"
              />
            </div>
            <div>
              <Label
                htmlFor="description"
                value="Description"
                className="mb-1"
              />
              <Textarea
                id="description"
                value={taskAttribute?.description ?? ""}
                onChange={(e) => {
                  setTaskAttribute({
                    ...taskAttribute!,
                    description: e.target.value,
                  });
                }}
                placeholder="Description"
                className="mb-4"
              />
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex flex-row justify-end w-full">
            <Button
              onClick={() => {
                setLoading(true);
                updateTaskAttribute(attributeId!, {
                  ...taskAttribute!,
                  data: JSON.stringify(taskAttribute?.fields ?? []),
                })
                  .then(() => {
                    toast.success("Attribute updated successfully");
                  })
                  .catch(toast.error)
                  .finally(() => {
                    setLoading(false);
                  });
              }}
            >
              Update
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default TaskAttributeDetail;
