import {
  Button,
  Label,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Textarea,
  TextInput,
} from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import toast from "react-hot-toast";
import { useParams } from "react-router-dom";
import Select from "react-select";
import AdminLayout from "../components/layouts/admin";
import MessageTemplateField from "../components/MessageTemplateField";
import ModalProductList from "../components/ModalProductList";
import { MemberContext } from "../contexts/ProfileContext";
import { FileModel } from "../models/file";
import { MemberModel } from "../models/member";
import { MessageTemplate, TemplateModel } from "../models/template";
import { getMembers } from "../services/api/commonApi";
import {
  addMessageTemplate,
  addProductTemplate,
  createInteractiveTemplate,
  getInteractiveTemplate,
  getTemplateDetail,
  updateInteractiveTemplate,
  updateTemplate,
} from "../services/api/templateApi";
import {
  WhatsappInteractiveList,
  WhatsappInteractiveListRow,
  WhatsappInteractiveListSection,
  WhatsappInteractiveModel,
} from "../models/whatsapp_interactive_message";
import { interactiveTypes } from "../utils/constants";
import { HiOutlineTableCells } from "react-icons/hi2";
import { randomString } from "../utils/helper";
import { BsTrash } from "react-icons/bs";
import { title } from "process";

interface TemplateDetailProps {}

const TemplateDetail: FC<TemplateDetailProps> = ({}) => {
  const [loading, setLoading] = useState<boolean>(false);
  const { member, setMember } = useContext(MemberContext);

  const { templateId } = useParams();
  const [template, setTemplate] = useState<TemplateModel | null>(null);
  const [emojis, setEmojis] = useState<any[]>([]);
  const [modalEmojis, setModalEmojis] = useState(false);
  const [selectedMessage, setSelectedMessage] = useState<MessageTemplate>();
  const [modalProduct, setModalProduct] = useState(false);
  const [mounted, setMounted] = useState(false);
  const [members, setMembers] = useState<MemberModel[]>([]);
  const [interactive, setInteractive] = useState<WhatsappInteractiveModel>();
  const [modalInteractive, setModalInteractive] = useState(false);
  useEffect(() => {
    setMounted(true);
    return () => setMounted(false);
  }, []);

  useEffect(() => {
    if (mounted && templateId) {
      getDetail();
      getMembers({ page: 1, size: 10, search: "" })
        .then((res: any) => {
          setMembers(res.data.items);
        })
        .catch(toast.error);
    }
  }, [mounted, templateId]);
  const getDetail = async () => {
    try {
      setLoading(true);
      const resp: any = await getTemplateDetail(templateId!);
      setTemplate(resp.data);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    for (const msg of template?.messages ?? []) {
    }
  }, [template?.messages]);

  const save = async () => {
    setLoading(true);
    try {
      await updateTemplate(templateId!, template!);
      toast.success("update successfully");
    } catch (error) {
      console.log(error);
      toast.error("Save failed");
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8 h-[calc(100vh-80px)] overflow-y-auto">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-4xl font-bold">/{template?.title}</h1>
            <div className="text-sm">{template?.description}</div>
          </div>
          <Button onClick={save}>Save</Button>
        </div>
        <div className="grid grid-cols-4 gap-4">
          <div className="col-span-3">
            <div className="mt-8">
              <h3 className="text-2xl font-semibold">Messages</h3>
              {template?.messages?.map((message, i) => (
                <MessageTemplateField
                  key={i}
                  templateId={template?.id}
                  msgId={message?.id}
                  title={`#Message ${i + 1}`}
                  onChangeBody={(val: string) => {
                    setTemplate({
                      ...template!,
                      messages: (template?.messages ?? []).map((m, index) => {
                        if (index === i) {
                          return {
                            ...m,
                            body: val,
                          };
                        }
                        return m;
                      }),
                    });
                  }}
                  index={i}
                  body={message.body ?? ""}
                  onClickEmoji={() => {
                    setSelectedMessage(template?.messages?.[i]);
                  }}
                  files={message.files ?? []}
                  onUploadImage={(file: FileModel, index?: number) => {
                    setTemplate({
                      ...template!,
                      messages: (template?.messages ?? []).map((m, index) => {
                        if (index === i) {
                          if (!m.files) {
                            m.files = [];
                          }
                          if (
                            (m.files ?? []).filter((f) =>
                              f.mime_type.includes("image")
                            ).length === 0
                          ) {
                            m.files = [file];
                          } else {
                            m.files = m.files.map((f) => {
                              if (f.mime_type.includes("image")) {
                                return file;
                              }
                              return f;
                            });
                          }
                          return m;
                        }
                        return m;
                      }),
                    });
                  }}
                  onUploadFile={(file: FileModel, index?: number) => {
                    setTemplate({
                      ...template!,
                      messages: (template?.messages ?? []).map((m, index) => {
                        if (!m.files) {
                          m.files = [];
                        }
                        if (index === i) {
                          if (
                            (m.files ?? []).filter(
                              (f) => !f.mime_type.includes("image")
                            ).length === 0
                          ) {
                            m.files.push(file);
                          } else {
                            m.files = m.files.map((f) => {
                              if (!f.mime_type.includes("image")) {
                                return file;
                              }
                              return f;
                            });
                          }
                          return m;
                        }
                        return m;
                      }),
                    });
                  }}
                  onTapProduct={() => {
                    setSelectedMessage(message);
                    setModalProduct(true);
                  }}
                  onTapInteractive={() => {
                    setSelectedMessage(template?.messages?.[i]);
                    setModalInteractive(true);
                    setInteractive({
                      type: "list",
                      title: "",
                      data: {
                        type: "list",
                        header: {
                          type: "text",
                          title: "",
                        }
                      }
                    });
                  }}
                  onEditInteractive={(d: WhatsappInteractiveModel) => {
                    setSelectedMessage(template?.messages?.[i]);
                    setModalInteractive(true);
                    setInteractive(d);
                  }}
                  product={message.products && message.products[0]}
                  showDelete={i > 0}
                />
              ))}
              {(template?.messages ?? []).length < 3 && (
                <Button
                  className=""
                  color="gray"
                  onClick={() => {
                    setLoading(true);
                    addMessageTemplate(template!.id!, {})
                      .catch((err) => {
                        toast.error(`${err}`);
                      })
                      .then(() => {
                        getDetail();
                      })
                      .finally(() => {
                        setLoading(false);
                      });
                  }}
                >
                  + Add Message
                </Button>
              )}
            </div>
          </div>
          <div className="">
            <div className="mt-8">
              <h3 className="text-2xl font-semibold">Option</h3>
              <div className="mb-4">
                <Label>Description</Label>
                <Textarea
                  value={template?.description}
                  onChange={(e) => {
                    setTemplate({
                      ...template!,
                      description: e.target.value,
                    });
                  }}
                />
              </div>

              {member?.role?.is_super_admin && (
                <div className="mb-4">
                  <Label>Member</Label>
                  <Select
                    options={members.map((m) => ({
                      value: m.id!,
                      label: m.user?.full_name!,
                    }))}
                    value={{
                      value: template?.member_id ?? "",
                      label: template?.member?.user?.full_name,
                    }}
                    onChange={(e) => {
                      if (!e) {
                        setTemplate({
                          ...template!,
                          member_id: null,
                          member: null,
                        });
                        return;
                      }
                      setTemplate({
                        ...template!,
                        member_id: e!.value!,
                        member: members.find((m) => m.id === e!.value),
                      });
                    }}
                    placeholder="Select member"
                    isClearable
                  />
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
      <ModalProductList
        show={modalProduct}
        setShow={setModalProduct}
        selectProduct={(val) => {
          setLoading(true);
          addProductTemplate(template?.id!, selectedMessage?.id!, val)
            .then(() => {
              getDetail();
              setModalProduct(false);
            })
            .catch((err) => {
              toast.error(`${err}`);
            })
            .finally(() => {
              setLoading(false);
            });
        }}
      />
      <Modal
        size="7xl"
        className="modal"
        show={modalInteractive}
        onClose={() => setModalInteractive(false)}
      >
        <ModalHeader className="modal-header">Interactive Message</ModalHeader>
        <ModalBody className="modal-body">
          <div className="grid grid-cols-2 gap-4 m-h-[400px]">
            <div className=" flex flex-col space-y-2">
              <div>
                <Label>Title</Label>
                <TextInput
                  placeholder="Title"
                  value={interactive?.title}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      title: e.target.value,
                    });
                  }}
                />
              </div>
              <div>
                <Label>Description</Label>
                <Textarea
                  placeholder="Description"
                  value={interactive?.description}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      description: e.target.value,
                    });
                  }}
                />
              </div>
              <div>
                <Label>Type</Label>
                <Select
                  options={interactiveTypes}
                  value={interactiveTypes.find(
                    (i) => i.value === interactive?.type
                  )}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      type: e!.value!,
                    });
                  }}
                />
              </div>
            </div>
            <div className=" flex flex-col space-y-2">
              <div>
                <Label>Header</Label>
                <TextInput
                  placeholder="Header"
                  value={interactive?.data?.header?.text}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      data: {
                        ...interactive?.data!,
                        header: {
                          ...interactive?.data?.header!,
                          text: e.target.value,
                        },
                      },
                    });
                  }}
                />
              </div>
              <div>
                <Label>Body</Label>
                <TextInput
                  placeholder="Body"
                  value={interactive?.data?.body?.text}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      data: {
                        ...interactive?.data!,
                        body: {
                          ...interactive?.data?.body!,
                          text: e.target.value,
                        },
                      },
                    });
                  }}
                />
              </div>
              <div>
                <Label>Footer</Label>
                <TextInput
                  placeholder="Footer"
                  value={interactive?.data?.footer?.text}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      data: {
                        ...interactive?.data!,
                        footer: {
                          ...interactive?.data?.footer!,
                          text: e.target.value,
                        },
                      },
                    });
                  }}
                />
              </div>
              <div>
                <Label>Action</Label>
                <TextInput
                  placeholder="Action"
                  value={interactive?.data?.action?.button}
                  onChange={(e) => {
                    setInteractive({
                      ...interactive!,
                      data: {
                        ...interactive?.data!,
                        action: {
                          ...interactive?.data?.action!,
                          button: e.target.value,
                        },
                      },
                    });
                  }}
                />
              </div>
              <div className="p-4  bg-gray-50 border-gray-200 rounded-lg">
                <div>
                  <div className="flex flex-row justify-between mb-4">
                    <Label>Sections</Label>
                    <div
                      className="text-sm cursor-pointer"
                      onClick={() => {
                        setInteractive({
                          ...interactive!,
                          data: {
                            ...interactive?.data!,
                            action: {
                              ...interactive?.data?.action!,
                              sections: [
                                ...(interactive?.data?.action?.sections ?? []),
                                {
                                  title: "",
                                  rows: [],
                                },
                              ],
                            },
                          },
                        });
                      }}
                    >
                      + Section
                    </div>
                  </div>
                </div>
                <div>
                  {interactive?.data?.action?.sections?.map(
                    (s: WhatsappInteractiveListSection, i: number) => (
                      <div key={i} className="mb-4">
                        <Label>Section #{i + 1}</Label>
                        <div className="flex flex-row justify-center items-center ">
                          <TextInput
                            placeholder="Section Title"
                            value={s.title}
                            className="flex-1"
                            onChange={(e) => {
                              let section =
                                interactive?.data?.action?.sections[i];
                              section.title = e.target.value;
                              interactive.data.action.sections[i] = section;
                              setInteractive({
                                ...interactive!,
                              });
                            }}
                          />
                          <div
                            className="cursor-pointer px-2"
                            onClick={() => {
                              let rows =
                                interactive?.data?.action?.sections[i].rows;
                              rows.push({
                                id: randomString(10),
                                title: "",
                                description: "",
                              });
                              setInteractive({
                                ...interactive!,
                              });
                            }}
                          >
                            <HiOutlineTableCells className="" size={20} />
                          </div>
                        </div>
                        <table className="w-full mt-4">
                          <thead>
                            <tr>
                              <th
                                className="p-2  border border-gray-200 text-sm"
                                style={{ width: "30%" }}
                              >
                                Title
                              </th>
                              <th
                                className="p-2  border border-gray-200 text-sm"
                                style={{ width: "60%" }}
                              >
                                Description
                              </th>
                              <th
                                className="p-2  border border-gray-200 text-sm "
                                style={{ width: "10%" }}
                              ></th>
                            </tr>
                          </thead>
                          <tbody>
                            {s.rows.length === 0 && (
                              <tr>
                                <td
                                  className="p-2 w-1/2 border border-gray-200 text-center"
                                  colSpan={3}
                                >
                                  No Rows
                                </td>
                              </tr>
                            )}
                            {s.rows.map(
                              (r: WhatsappInteractiveListRow, j: number) => (
                                <tr key={j}>
                                  <td className="p-2 w-1/2 border border-gray-200">
                                    <TextInput
                                      placeholder="Title"
                                      value={r.title}
                                      onChange={(e) => {
                                        let row =
                                          interactive?.data?.action?.sections[i]
                                            .rows[j];
                                        row.title = e.target.value;
                                        interactive.data.action.sections[
                                          i
                                        ].rows[j] = row;
                                        setInteractive({
                                          ...interactive!,
                                        });
                                      }}
                                    />
                                  </td>
                                  <td className="p-2 w-1/2 border border-gray-200">
                                    <TextInput
                                      placeholder="Description"
                                      value={r.description}
                                      onChange={(e) => {
                                        let row =
                                          interactive?.data?.action?.sections[i]
                                            .rows[j];
                                        row.description = e.target.value;
                                        interactive.data.action.sections[
                                          i
                                        ].rows[j] = row;
                                        setInteractive({
                                          ...interactive!,
                                        });
                                      }}
                                    />
                                  </td>
                                  <td className="p-2 w-1/2 border border-gray-200">
                                    <BsTrash
                                      className="text-red-400 hover:text-red-600 cursor-pointer del"
                                      onClick={() => {
                                        let rows =
                                          interactive?.data?.action?.sections[i]
                                            .rows;
                                        rows.splice(j, 1);
                                        setInteractive({
                                          ...interactive!,
                                        });
                                      }}
                                    />
                                  </td>
                                </tr>
                              )
                            )}
                          </tbody>
                        </table>
                      </div>
                    )
                  )}
                </div>
              </div>
            </div>
          </div>
        </ModalBody>
        <ModalFooter className="modal-footer ">
          <div className="flex flex-row justify-end gap-2 w-full">
            <Button color="gray" onClick={() => setModalInteractive(false)}>
              Close
            </Button>
            <Button
              onClick={async () => {
                if (!interactive?.title) {
                  toast.error("Title is required");
                  return;
                }

                if (interactive.type === "list") {
                  if (!interactive.data?.header?.text) {
                    toast.error("Header text is required");
                    return;
                  }
                  if (!interactive.data?.body?.text) {
                    toast.error("Header text is required");
                    return;
                  }
                }

                try {
                  setLoading(true);
                  if (interactive?.id) {
                    await updateInteractiveTemplate(
                      templateId!,
                      interactive?.id!,
                      interactive
                    );
                  } else {
                    await createInteractiveTemplate(
                      templateId!,
                      selectedMessage?.id!,
                      interactive
                    );
                  }

                  toast.success("Save successfully");
                  setModalInteractive(false);
                  setTemplate(null)
                  getDetail();

                } catch (err) {
                  console.log(err);
                } finally {
                  setLoading(false);
                }
              }}
            >
              Save
            </Button>
          </div>
        </ModalFooter>
      </Modal>
    </AdminLayout>
  );
};
export default TemplateDetail;

const emojiStyle = {
  control: {
    fontSize: 16,
    lineHeight: 1.2,
    minHeight: 160,
  },

  highlighter: {
    padding: 9,
    border: "1px solid transparent",
  },

  input: {
    fontSize: 16,
    lineHeight: 1.2,
    padding: 9,
    border: "1px solid silver",
    borderRadius: 10,
  },

  suggestions: {
    list: {
      backgroundColor: "white",
      border: "1px solid rgba(0,0,0,0.15)",
      fontSize: 16,
    },

    item: {
      padding: "5px 15px",
      borderBottom: "1px solid rgba(0,0,0,0.15)",

      "&focused": {
        backgroundColor: "#cee4e5",
      },
    },
  },
};
