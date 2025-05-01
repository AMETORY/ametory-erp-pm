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
  useEffect(() => {
    fetch(
      "https://gist.githubusercontent.com/oliveratgithub/0bf11a9aff0d6da7b46f1490f86a71eb/raw/d8e4b78cfe66862cf3809443c1dba017f37b61db/emojis.json"
    )
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);

  const neverMatchingRegex = /($a)/;
  const queryEmojis = (query: any, callback: (emojis: any) => void) => {
    if (query.length === 0) return;

    const matches = emojis
      .filter((emoji: any) => {
        return emoji.name.indexOf(query.toLowerCase()) > -1;
      })
      .slice(0, 10);
    return matches.map(({ emoji }) => ({ id: emoji }));
  };

  const groupBy = (emojis: any[], category: string): { [s: string]: any[] } => {
    return emojis.reduce((acc, curr) => {
      const key = curr[category];
      if (!acc[key]) {
        acc[key] = [];
      }
      acc[key].push(curr);
      return acc;
    }, {});
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
                <div
                  className="bg-gray-50 rounded-lg p-4 flex flex-col mb-8"
                  key={i}
                >
                  <h4 className="font-semibold">#Message {i + 1}</h4>
                  <div className="relative w-full">
                    <MentionsInput
                      value={message.body}
                      onChange={(val: any) => {
                        setTemplate({
                          ...template!,
                          messages: (template?.messages ?? []).map(
                            (m, index) => {
                              if (index === i) {
                                return {
                                  ...m,
                                  body: val.target.value,
                                };
                              }
                              return m;
                            }
                          ),
                        });
                      }}
                      style={emojiStyle}
                      placeholder={
                        "Press ':' for emojis and shift+enter for new line"
                      }
                      className="w-full bg-white"
                      autoFocus
                    >
                      <Mention
                        trigger="@"
                        data={[
                          { id: "{{user}}", display: "Full Name" },
                          { id: "{{phone}}", display: "Phone Number" },
                          { id: "{{agent}}", display: "Agent Name" },
                        ]}
                        style={{
                          backgroundColor: "#cee4e5",
                        }}
                        appendSpaceOnAdd
                      />
                      <Mention
                        trigger=":"
                        markup="__id__"
                        regex={neverMatchingRegex}
                        data={queryEmojis}
                      />
                    </MentionsInput>
                    <div
                      className="absolute bottom-2 right-2 cursor-pointer"
                      onClick={() => {
                        setModalEmojis(true);
                        setSelectedMessage(template?.messages?.[i]);
                      }}
                    >
                      😀
                    </div>
                  </div>
                  <div className="mt-8">
                    <h4 className="font-semibold">Files</h4>
                    <div className="grid grid-cols-2 gap-4">
                      <div
                        className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100"
                        onClick={() => {
                          document.getElementById(`image-${i}`)?.click();
                        }}
                      >
                        {(message.files ?? []).filter((f) =>
                          f.mime_type.includes("image")
                        ).length === 0 ? (
                          <div className="flex flex-col items-center">
                            <span>Add Photo to message</span>
                            <BsCamera />
                          </div>
                        ) : (
                          <img
                            className="w-32 h-32 object-cover"
                            src={
                              (message.files ?? []).find((f) =>
                                f.mime_type.includes("image")
                              )?.url
                            }
                          />
                        )}
                        <input
                          type="file"
                          className="hidden"
                          accept="image/*"
                          id={`image-${i}`}
                          onChange={(e) => {
                            const file = e.target.files?.[0];
                            if (file) {
                              uploadFile(file, {}, () => {}).then(
                                (resp: any) => {
                                  setTemplate({
                                    ...template!,
                                    messages: (template?.messages ?? []).map(
                                      (m, index) => {
                                        if (index === i) {
                                          if (!m.files) {
                                            m.files = [];
                                          }
                                          if (
                                            (m.files ?? []).filter((f) =>
                                              f.mime_type.includes("image")
                                            ).length === 0
                                          ) {
                                            m.files = [resp.data];
                                          } else {
                                            m.files = m.files.map((f) => {
                                              if (
                                                f.mime_type.includes("image")
                                              ) {
                                                return resp.data;
                                              }
                                              return f;
                                            });
                                          }
                                          console.log(m);
                                          return m;
                                        }
                                        return m;
                                      }
                                    ),
                                  });
                                }
                              );
                            }
                          }}
                        />
                      </div>
                      <div
                        className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100"
                        onClick={() => {
                          document.getElementById(`image-${i}-file`)?.click();
                        }}
                      >
                        {(message.files ?? []).filter(
                          (f) => !f.mime_type.includes("image")
                        ).length === 0 ? (
                          <div className="flex flex-col items-center">
                            <span>Add File to message</span>
                            <HiOutlineDocumentAdd size={32} />
                          </div>
                        ) : (
                          // <IoAttach className="rotate-[30deg]" size={32}/>
                          <div className="flex items-center flex-col px-8">
                            <IoDocumentsOutline size={32} />
                            <small className="text-center mt-4">
                              {
                                (message.files ?? []).find(
                                  (f) => !f.mime_type.includes("image")
                                )?.file_name
                              }
                            </small>
                          </div>
                        )}
                        <input
                          type="file"
                          className="hidden"
                          accept=".doc,.docx,.pdf,.xls,.xlsx,.txt"
                          id={`image-${i}-file`}
                          onChange={(e) => {
                            const file = e.target.files?.[0];
                            if (file) {
                              uploadFile(file, {}, () => {}).then(
                                (resp: any) => {
                                  setTemplate({
                                    ...template!,
                                    messages: (template?.messages ?? []).map(
                                      (m, index) => {
                                        if (!m.files) {
                                          m.files = [];
                                        }
                                        if (index === i) {
                                          if (
                                            (m.files ?? []).filter(
                                              (f) =>
                                                !f.mime_type.includes("image")
                                            ).length === 0
                                          ) {
                                            m.files.push(resp.data);
                                          } else {
                                            m.files = m.files.map((f) => {
                                              if (
                                                !f.mime_type.includes("image")
                                              ) {
                                                return resp.data;
                                              }
                                              return f;
                                            });
                                          }
                                          // console.log(m);
                                          return m;
                                        }
                                        return m;
                                      }
                                    ),
                                  });
                                }
                              );
                            }
                          }}
                        />
                      </div>
                    </div>
                  </div>
                  <div className="mt-8">
                    <h4 className="font-semibold">Product</h4>
                    <div className="grid grid-cols-2 gap-4 ">
                      <div
                        className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100"
                        onClick={() => {
                          setSelectedMessage(message);
                          setModalProduct(true);
                        }}
                      >
                        {(message.products ?? []).length === 0 ? (
                          <div className="flex flex-col items-center">
                            <span>Add Product</span>
                            <BsCart size={32} />
                          </div>
                        ) : (
                          <div className="flex items-center flex-col  px-8">
                            {(message.products![0].product_images ?? [])
                              .length > 0 && (
                              <img
                                src={
                                  message.products![0].product_images![0].url
                                }
                                alt="product"
                                className="w-32 h-32 rounded-lg"
                              />
                            )}
                            <h3 className="font-semibold mt-2 text-center">
                              {message.products![0].name}
                            </h3>
                            <small>{money(message.products![0].price)}</small>
                          </div>
                        )}
                      </div>
                    </div>
                  </div>
                  <div className="mt-4">
                    {i > 0 && (
                      <Button
                        className=""
                        color="red"
                        onClick={() => {
                          setLoading(true);
                          deleteMessageTemplate(template!.id!, message.id!)
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
                        + Delete Message
                      </Button>
                    )}
                  </div>
                </div>
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
      <Modal
        dismissible
        show={modalEmojis}
        onClose={() => setModalEmojis(false)}
      >
        <Modal.Header>Emojis</Modal.Header>
        <Modal.Body>
          <div>
            {Object.entries(groupBy(emojis, "category")).map(
              ([category, emojis]) => (
                <div
                  className="mb-4 hover:bg-gray-100 rounded-lg p-2"
                  key={category}
                >
                  <h3 className="font-bold">{category}</h3>
                  <div className=" flex flex-wrap gap-1">
                    {emojis.map((e: any, index: number) => (
                      <div
                        key={index}
                        className="cursor-pointer text-lg"
                        onClick={() => {
                          setTemplate({
                            ...template!,
                            messages: (template?.messages ?? []).map(
                              (m, index) => {
                                if (selectedMessage?.id !== m.id) {
                                  return m;
                                }
                                return {
                                  ...selectedMessage,
                                  body: m.body + e.emoji,
                                };
                              }
                            ),
                          });
                          // setModalEmojis(false);
                        }}
                      >
                        {e.emoji}
                      </div>
                    ))}
                  </div>
                </div>
              )
            )}
          </div>
        </Modal.Body>
      </Modal>
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
