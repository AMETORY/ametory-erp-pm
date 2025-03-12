import { useContext, useEffect, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Link, useParams } from "react-router-dom";
import { getFormDetail, updateForm } from "../services/api/formApi";
import {
  FormModel,
  FormSection,
  FormTemplateModel,
  FormData,
  FormField,
  FormFieldType,
} from "../models/form";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import {
  Button,
  Label,
  Modal,
  Tabs,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import Select, { InputActionMeta } from "react-select";
import { ProjectModel } from "../models/project";
import { ColumnModel } from "../models/column";
import { getProjects } from "../services/api/projectApi";
import { uploadFile } from "../services/api/commonApi";
import { FileModel } from "../models/file";
import { IoImageOutline } from "react-icons/io5";
import FormView from "../components/FormView";
import {
  BsCheck2Circle,
  BsCode,
  BsCodeSlash,
  BsEye,
  BsShare,
} from "react-icons/bs";
import Moment from "react-moment";
import { money } from "../utils/helper";
interface FormDetailProps {}

const FormDetail: FC<FormDetailProps> = ({}) => {
  const { formId } = useParams();
  const [form, setForm] = useState<FormModel>();
  const { loading, setLoading } = useContext(LoadingContext);
  const [templates, setTemplates] = useState<FormTemplateModel[]>([]);
  const [projects, setProjects] = useState<ProjectModel[]>([]);
  const [columns, setColumns] = useState<ColumnModel[]>([]);
  const fileRef = useRef<HTMLInputElement>(null);
  const [fileCover, setFileCover] = useState<FileModel>();
  const [selectedResponse, setSelectedResponse] = useState<FormData>();

  useEffect(() => {
    getAllProjects("");
  }, []);
  useEffect(() => {
    if (formId) {
      setLoading(true);
      getFormDetail(formId)
        .then((e: any) => setForm(e.data))
        .catch(toast.error)
        .finally(() => {
          setLoading(false);
        });
    }
  }, [formId]);
  const getAllProjects = (s: string) => {
    getProjects({ page: 1, size: 10, search: s })
      .then((e: any) => setProjects(e.data.items))
      .catch(toast.error);
  };

  const renderValue = (fieldType: FormFieldType, val: any) => {
    switch (fieldType) {
      case FormFieldType.DateRangePicker:
        return (
          val && (
            <div>
              {val[0] && <Moment format="DD MMM YYYY">{val[0]}</Moment>} ~{" "}
              {val[1] && <Moment format="DD MMM YYYY">{val[1]}</Moment>}
            </div>
          )
        );
      case FormFieldType.DatePicker:
        return val && <Moment format="DD MMM YYYY">{val}</Moment>;
      case FormFieldType.PasswordField:
        return val && "* * * * * * *";
      case FormFieldType.ToggleSwitch:
        return val && <BsCheck2Circle />;
      case FormFieldType.FileUpload:
        return (
          val && (
            <Link to={val} target="_blank">
              {val}
            </Link>
          )
        );
      case FormFieldType.NumberField:
      case FormFieldType.Currency:
        return money(parseFloat(val));
      case FormFieldType.Checkbox:
        return (
          <ul className=" list-disc ml-4">
            {val.map((e: any) => (
              <li key={e}>{e}</li>
            ))}
          </ul>
        );

      default:
        break;
    }
    return val;
  };

  const renderResponses = (responses: FormData[]) => {
    if (responses.length == 0) return;
    let sections: any[] = [];
    for (const element of responses) {
      for (const section of element.sections) {
        sections.push({
          id: section.id,
          record_id: element.id,
          title: section.section_title,
          number_field: section.fields.length,
          sub_title: section.fields.map((e) => e.label),
          fields: section.fields,
        });
      }
    }

    return (
      <div className="overflow-x-auto w-[calc(100% - 300px)]">
        <table className="bg-white min-w-full">
          <thead>
            <tr>
              {sections.map((e) => (
                <td
                  className="px-4 py-2 bg-gray-50 border text-center font-semibold"
                  key={e.id}
                  colSpan={e.number_field}
                >
                  {e.title}
                </td>
              ))}
            </tr>
            <tr>
              {sections.map((section) =>
                section.sub_title.map((e: any) => (
                  <td
                    className="px-4 py-2 border bg-gray-50 font-semibold min-w-[160px]"
                    key={e}
                  >
                    {e}
                  </td>
                ))
              )}
            </tr>
          </thead>
          <tbody>
            <tr className="hover:bg-blue-50 ">
              {sections.map((section) => {
                return section.fields.map((e: FormField) => (
                  <td
                    className="px-4 py-2 border cursor-pointer "
                    key={e.id}
                    onClick={() => {
                      setSelectedResponse(
                        form?.responses?.find((e) => e.id == section.record_id)
                      );
                    }}
                  >
                    {renderValue(e.type, e.value)}
                  </td>
                ));
              })}
            </tr>
          </tbody>
        </table>
      </div>
    );
  };
  return (
    <AdminLayout>
      <div className="flex flex-row w-full h-full flex-1 gap-2">
        <div className="w-[300px] h-full p-4 space-y-4 flex flex-col overflow-y-auto">
          <div className="flex justify-between items-center">
            <h3 className="text-2xl font-bold">Form Detail</h3>
            <div className="flex gap-2 items-center">
              {form?.is_public && (
                <BsShare
                  size={12}
                  className="cursor-pointer"
                  onClick={() => {
                    let url = `${process.env.REACT_APP_BASE_URL}/public/form/${form?.code}`;
                    navigator.clipboard.writeText(url);
                    toast.success("URL copied to clipboard");
                    window.open(url);
                  }}
                />
              )}
              {form?.is_public && (
                <BsCodeSlash
                  className=" cursor-pointer"
                  size={12}
                  onClick={() => {
                    let code = `<iframe src="${process.env.REACT_APP_BASE_URL}/public/form/${form?.code}" width="100%" style="height: 100vh;"  frameborder="0" marginheight="0" marginwidth="0">Loading...</iframe>
<script>
    window.addEventListener('message', (e) => {
    if (JSON.parse(e.data).type === 'FORM_SUBMITTED') {
        console.log(JSON.parse(e.data).data);
        alert('Form submitted successfully!');
    }
    });
</script>`;
                    navigator.clipboard.writeText(code);
                    toast.success("Code copied to clipboard");
                  }}
                />
              )}
            </div>
          </div>
          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Cover
              </Label>
            </div>

            {form?.cover ? (
              <img
                src={form?.cover.url}
                className=" aspect-video w-full object-cover rounded-lg mt-4 bg-gray-50 cursor-pointer"
                onClick={() => {
                  fileRef.current?.click();
                }}
              />
            ) : (
              <div
                className="aspect-video w-full object-cover rounded-lg mt-4 bg-gray-100 hover:bg-gray-200 cursor-pointer transition-all flex justify-center items-center flex-row"
                onClick={() => {
                  fileRef.current?.click();
                }}
              >
                <IoImageOutline className="" size={24} />
              </div>
            )}
          </div>
          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Template
              </Label>
            </div>
            <Select
              value={
                form?.form_template
                  ? {
                      label: form?.form_template?.title,
                      value: form?.form_template?.id,
                    }
                  : null
              }
              options={templates.map((e) => ({
                label: e.title,
                value: e.id,
              }))}
              onChange={(val) => {
                setForm({
                  ...form,
                  form_template_id: val?.value,
                  form_template: {
                    ...form!.form_template!,
                    id: val!.value,
                    title: val!.label,
                  },
                });
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Title
              </Label>
            </div>
            <TextInput
              id="name"
              type="text"
              placeholder="Title"
              required={true}
              value={form?.title ?? ""}
              onChange={(e) => {
                setForm({
                  ...form,
                  title: e.target.value,
                });
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Description
              </Label>
            </div>
            <Textarea
              id="description"
              placeholder="Description"
              required={true}
              value={form?.description ?? ""}
              onChange={(e) => {
                setForm({
                  ...form,
                  description: e.target.value,
                });
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Project
              </Label>
            </div>
            <Select
              value={
                form?.project
                  ? { label: form?.project?.name, value: form?.project?.id }
                  : null
              }
              options={[
                { label: "Select Project", value: "" },
                ...projects.map((e) => ({
                  label: e.name,
                  value: e.id,
                })),
              ]}
              onChange={(val) => {
                setForm({
                  ...form,
                  project_id: val?.value,
                  column: undefined,
                  column_id: undefined,
                  project: {
                    ...form!.project!,
                    id: val!.value,
                    name: val!.label,
                    columns:
                      projects.find((e) => e.id == val?.value)?.columns ?? [],
                  },
                });
              }}
              onInputChange={(val) => {
                getAllProjects(val);
              }}
            />
          </div>
          {form?.project && (
            <div>
              <div className=" block">
                <Label className="font-bold" htmlFor="name">
                  Column
                </Label>
              </div>
              <Select
                value={
                  form?.column
                    ? { label: form?.column?.name, value: form?.column?.id }
                    : null
                }
                options={[
                  { label: "Select Column", value: "" },
                  ...(form.project?.columns ?? []).map((e) => ({
                    label: e.name!,
                    value: e.id!,
                  })),
                ]}
                onChange={(val) => {
                  setForm({
                    ...form,
                    column_id: val?.value!,
                    column: {
                      ...form!.column!,
                      id: val!.value,
                      name: val!.label,
                    },
                  });
                }}
              />
            </div>
          )}
          <div>
            <div className=" grid grid-cols-1">
              <ToggleSwitch
                checked={form?.is_public ?? false}
                onChange={(val) => {
                  setForm({
                    ...form,
                    is_public: val,
                  });
                }}
                label={"Is Public"}
              />
            </div>
          </div>

          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Webhook Submit URL
              </Label>
            </div>
            <TextInput
              name="url"
              id="url"
              placeholder="Webhook Submit URL"
              value={form?.submit_url}
              onChange={(e) => {
                setForm({
                  ...form,
                  submit_url: e.target.value,
                });
              }}
            />
          </div>

          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Webhook Method
              </Label>
            </div>
            <Select
              options={[
                { label: "POST", value: "POST" },
                { label: "GET", value: "GET" },
                { label: "PUT", value: "PUT" },
                { label: "DELETE", value: "DELETE" },
              ]}
              value={{ value: form?.method, label: form?.method }}
              onChange={(val) => {
                setForm({
                  ...form,
                  method: val?.value,
                });
              }}
            />
          </div>

          <div>
            <div className=" block">
              <Label className="font-bold" htmlFor="name">
                Webhook Headers
              </Label>
            </div>
            <Textarea
              id="header"
              placeholder={`{
        "token":"abcd"
}`}
              value={form?.headers}
              onChange={(e) => {
                setForm({
                  ...form,
                  headers: e.target.value,
                });
              }}
            />
          </div>
          <div>
            <div className="mt-4">
              <Button
                type="submit"
                onClick={() => {
                  setLoading(true);
                  updateForm(form?.id!, {
                    ...form,
                  })
                    .then((v) => {
                      toast.success("Form Template Updated");
                    })
                    .finally(() => {
                      setLoading(false);
                    })
                    .catch(toast.error);
                }}
              >
                Update
              </Button>
            </div>
          </div>
        </div>
        <div className="w-[calc(100%-300px)] border-l relative bg-gray-50">
          <Tabs>
            <Tabs.Item title="Preview">
              <div className="bg-gray-50 flex flex-col  items-center p-16 overflow-y-auto h-[calc(100vh-100px)] ">
                <div className="flex flex-col justify-center w-1/2 space-y-4 mb-4">
                  <div className="bg-white rounded-lg border border-t-4 border-t-blue-400">
                    {form?.cover && (
                      <img
                        src={form?.cover?.url}
                        className=" aspect-video w-full object-cover"
                      />
                    )}
                    <div className="p-4  ">
                      <h1 className="text-2xl font-semibold">{form?.title}</h1>
                      <p>{form?.description}</p>
                    </div>
                  </div>
                </div>
                <FormView
                  sections={form?.form_template?.sections ?? []}
                  onSubmit={(val) => {
                    console.log(val);
                  }}
                />
              </div>
            </Tabs.Item>
            <Tabs.Item title="Responses">
              <div className="px-4">
                {form?.responses && renderResponses(form.responses)}
              </div>
            </Tabs.Item>
          </Tabs>
        </div>
      </div>
      <input
        accept="image/*"
        ref={fileRef}
        className="hidden"
        type="file"
        onChange={(val) => {
          const files = val?.target.files;
          if (files) {
            try {
              uploadFile(files[0], {}, (val) => console.log).then((v: any) => {
                setFileCover(v.data);
                setForm({
                  ...form,
                  cover: v.data,
                });
              });
            } catch (error) {
              console.log(error);
            }
          }
        }}
      />
      <Modal
        show={selectedResponse != undefined}
        onClose={() => setSelectedResponse(undefined)}
      >
        <Modal.Header>
        Response Details
        </Modal.Header>
        <Modal.Body>
          {(selectedResponse?.sections ?? []).map((e) => (
            <div className="" key={e.id}>
              <h2 className="text-lg font-bold">{e.section_title}</h2>

              <table className="w-full mb-4">
                {e.fields.map((f) => (
                  <tr key={f.id} className="border">
                    <td className="px-2 py-1 font-semibold bg-gray-50 border " style={{ 
                      width: 200
                     }}>
                      {f.label}
                    </td>

                    <td className="px-2 py-1 border">
                      {renderValue(f.type, f.value)}
                    </td>
                  </tr>
                ))}
              </table>
            </div>
          ))}
        </Modal.Body>
      </Modal>
    </AdminLayout>
  );
};
export default FormDetail;
