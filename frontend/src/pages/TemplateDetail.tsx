import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  addMessageTemplate,
  addProductTemplate,
  deleteMessageTemplate,
  getTemplateDetail,
  updateTemplate,
} from "../services/api/templateApi";
import { useParams } from "react-router-dom";
import { MessageTemplate, TemplateModel } from "../models/template";
import { Mention, MentionsInput } from "react-mentions";
import { Button, Label, Modal, Textarea } from "flowbite-react";
import toast from "react-hot-toast";
import { BsCamera, BsCart } from "react-icons/bs";
import { getMembers, uploadFile } from "../services/api/commonApi";
import { HiDocumentAdd, HiOutlineDocumentAdd } from "react-icons/hi";
import { RiAttachmentLine } from "react-icons/ri";
import { IoAttach, IoDocumentsOutline } from "react-icons/io5";
import ModalProductList from "../components/ModalProductList";
import { money } from "../utils/helper";
import { MemberContext } from "../contexts/ProfileContext";
import { MemberModel } from "../models/member";
import Select from "react-select";
import MessageMention from "../components/MessageMention";
import MessageTemplateField from "../components/MessageTemplateField";
import { FileModel } from "../models/file";

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
