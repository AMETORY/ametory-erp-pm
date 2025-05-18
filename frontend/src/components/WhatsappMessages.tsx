import {
  Button,
  Dropdown,
  Modal,
  Popover,
  Table,
  Tabs,
  ToggleSwitch,
} from "flowbite-react";
import { useContext, useEffect, useRef, useState, type FC } from "react";
import toast from "react-hot-toast";
import { RiAttachment2 } from "react-icons/ri";
import Markdown from "react-markdown";
import { Mention, MentionsInput } from "react-mentions";
import Moment from "react-moment";
import remarkGfm from "remark-gfm";
import { ProfileContext } from "../contexts/ProfileContext";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ConnectionModel } from "../models/connection";
import { FileModel } from "../models/file";
import {
  WhatsappMessageModel,
  WhatsappMessageSessionModel,
} from "../models/whatsapp_message";
import { updateConnection } from "../services/api/connectionApi";
import {
  createWAMessage,
  getWhatsappMessages,
  getWhatsappSessionDetail,
  markAsRead,
  updateWhatsappSession,
} from "../services/api/whatsappApi";
import { debounce } from "../utils/helper";
import { IoCheckmarkDone } from "react-icons/io5";
import {
  BsFileEarmark,
  BsImage,
  BsPlusCircle,
  BsSend,
  BsTag,
} from "react-icons/bs";
import { MessageTemplate, TemplateModel } from "../models/template";
import { getTemplates } from "../services/api/templateApi";
import { TbTemplate } from "react-icons/tb";
import { uploadFile } from "../services/api/commonApi";
import { FaXmark } from "react-icons/fa6";
import { ProductModel } from "../models/product";
import ModalProduct from "./ModalProduct";
import { getProducts } from "../services/api/productApi";
import { HiMagnifyingGlass } from "react-icons/hi2";
import { SearchContext } from "../contexts/SearchContext";

interface WhatsappMessagesProps {
  //   session: WhatsappMessageSessionModel;
  sessionId: string;
}

const WhatsappMessages: FC<WhatsappMessagesProps> = ({ sessionId }) => {
  const timeout = useRef<number | null>(null);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(100);
  const { search, setSearch } = useContext(SearchContext);
  const [mounted, setMounted] = useState(false);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);
  const [session, setSession] = useState<WhatsappMessageSessionModel>();
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);
  const [files, setFiles] = useState<FileModel[]>([]);
  const [products, setProducts] = useState<ProductModel[]>([]);
  const [openAttachment, setOpenAttachment] = useState(false);
  const [connection, setConnection] = useState<ConnectionModel>();
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const [modalEmojis, setModalEmojis] = useState(false);
  const [templates, setTemplates] = useState<TemplateModel[]>([]);
  const [modalTemplates, setModalTemplates] = useState(false);
  const fileRef = useRef<HTMLInputElement>(null);
  const [modalProduct, setModalProduct] = useState(false);
  const [selectedProducts, setSelectedProducts] = useState<ProductModel[]>([]);
  const [isCaption, setIsCaption] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getWhatsappSessionDetail(sessionId)
        .then((res: any) => {
          setSession(res.data);
          setConnection(res.connection);
        })
        .catch((err) => {
          // toast.error(`${err}`);
          window.location.href = "/whatsapp";
        })
        .finally(() => {});

      getAllTemplates();
      getProducts({ page: 1, size: 100 }).then((res: any) => {
        setProducts(res.data.items);
      });
    }
  }, [mounted, sessionId]);

  useEffect(() => {
    const container = chatContainerRef.current;

    if (container && container.scrollHeight <= container.clientHeight) {
      markAllAsRead();
    }
  }, [messages]); // atau [chatList], tergantung state kamu

  const markAllAsRead = () => {
    for (const msg of messages) {
      if (!msg.is_read) {
        if (timeout.current) {
          window.clearTimeout(timeout.current);
        }

        timeout.current = window.setTimeout(() => {
          if (!msg.is_from_me) {
            markAsRead(msg!.id!);
          }
        }, 500);
      }
    }
  };

  const getAllTemplates = async () => {
    try {
      let resp: any = await getTemplates({ page: 1, size: 10 });
      setTemplates(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
    }
  };

  const handleScroll = () => {
    const messageElements = document.querySelectorAll(".message");

    messageElements.forEach((el) => {
      const observer = new IntersectionObserver(
        (entries) => {
          entries.forEach((entry) => {
            if (entry.isIntersecting) {
              // const messageId = parseInt(entry.target.dataset.id);
              // markAsRead(messageId);
              let message = messages.find(
                (m) => m.id == entry.target.getAttribute("id")
              );
              // console.log(message?.message)
              if (
                message &&
                !message.is_read &&
                !(message?.is_from_me ?? false)
              ) {
                setMessages([
                  ...messages.map((m) => {
                    if (m.id == message?.id) {
                      return { ...m, is_read: true };
                    }
                    return m;
                  }),
                ]);

                timeout.current = window.setTimeout(() => {
                  if (!message?.is_from_me) {
                    markAsRead(message!.id!);
                  }
                }, 500);
              }
            }
          });
        },
        { threshold: 0.3 } // Minimal 50% pesan terlihat
      );

      observer.observe(el);
    });
  };

  useEffect(() => {
    getWhatsappMessages(sessionId, {
      page,
      size,
      search,
    })
      .then((res: any) => {
        setMessages(res.data.items);
      })
      .catch((err) => {
        console.error(err);
        window.location.href = "/whatsapp";
      });
  }, [session, sessionId]);

  useEffect(() => {
    if (!sessionId) return;
    if (
      wsMsg?.session_id == sessionId &&
      wsMsg?.command == "WHATSAPP_RECEIVED"
    ) {
      setMessages([...messages, wsMsg.data]);
      setTimeout(() => {
        scrollToBottom();
        setSession({
          ...session,
          last_online_at: new Date(),
        });
      }, 300);
    }

    if (
      wsMsg?.session_id == sessionId &&
      wsMsg?.command == "WHATSAPP_MESSAGE_READ"
    ) {
      // console.log(wsMsg.message_ids);
      setMessages([
        ...messages.map((m) => {
          // console.log(m.message_id, wsMsg.message_ids.includes(m.message_id));
          if (wsMsg.message_ids.includes(m.message_id)) {
            return { ...m, is_read: true };
          }
          return m;
        }),
      ]);
    }
    if (
      wsMsg?.session_id == sessionId &&
      wsMsg?.command == "WHATSAPP_CLEAR_MESSAGE"
    ) {
      // console.log(wsMsg.message_ids);
      setMessages([]);
    }
  }, [wsMsg, profile, sessionId]);
  useEffect(() => {
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);
  const emojiStyle = {
    control: {
      fontSize: 16,
      lineHeight: 1.2,
      minHeight: 30,
      maxHeight: 80,
    },

    highlighter: {
      padding: 9,
      border: "1px solid transparent",
    },

    input: {
      fontSize: 16,
      lineHeight: 1.2,
      paddingLeft: 9,
      paddingRight: 9,
      paddingTop: 9,
      paddingBottom: 9,
      border: "1px solid silver",
      borderRadius: 10,
      backgroundColor: !session?.is_human_agent ? "#f0f0f0" : "white",
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

  const scrollToBottom = () => {
    const element = document.getElementById("channel-messages");
    if (element) {
      element.scrollTo({
        top: element.scrollHeight,
        behavior: "smooth",
      });
    }
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  useEffect(() => {
    if (connection?.id) {
      updateConnection(connection!.id!, {
        ...connection,
      });
    }
  }, [connection?.is_auto_pilot]);
  useEffect(() => {
    if (session?.id) {
    }
  }, [session?.is_human_agent]);

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

  const sendMessage = async () => {
    try {
      if (!content) return;
      setContent("");
      setOpenAttachment(false);
      setFiles([]);
      setSelectedProducts([]);
      await createWAMessage(sessionId!, {
        message: content,
        files: files,
        products: selectedProducts,
        is_caption: isCaption,
      });
      setIsCaption(false);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setTimeout(() => {
        setContent("");
      }, 300);
    }
  };

  return (
    <div className="flex flex-col h-full ">
      <div className="shoutbox border-b py-2 min-h-[40px] flex justify-between items-center">
        <div className="flex gap-2 items-center px-4">
          {session?.contact?.avatar && (
            <img
              src={session?.contact?.avatar.url}
              className=" aspect-square rounded-full object-cover w-8 h-8"
            />
          )}
          <div className="flex flex-col">
            <span className="font-semibold">{session?.contact?.name}</span>
            <div className="flex justify-between">
              <Moment className="text-xs" fromNow>
                {session?.last_online_at}
              </Moment>
            </div>
          </div>
        </div>
        {/* <HiOutlineUserGroup
              className=" text-gray-300 hover:text-gray-600 cursor-pointer"
              size={24}
              onClick={openModal}
            /> */}
        {connection?.is_auto_pilot && (
          <ToggleSwitch
            checked={session?.is_human_agent ?? false}
            label="Human Agent"
            onChange={(val) => {
              setSession({
                ...session,
                is_human_agent: val,
              });

              updateWhatsappSession(session!.id!, {
                ...session,
                is_human_agent: val,
              });
            }}
          />
        )}
      </div>
      <div
        id="channel-messages"
        className="messages h-[calc(100vh-260px)] overflow-y-auto p-4 bg-gray-50 space-y-8"
        ref={chatContainerRef}
        onScroll={handleScroll}
      >
        {messages.map((msg) => (
          <div
            key={msg.id}
            className={`message flex flex-row items-end mb-2  ${
              msg.is_from_me ? "justify-end" : "justify-start"
            }`}
            id={msg.id}
          >
            <div
              className={`min-w-[300px] max-w-[600px] ${
                !msg.is_from_me ? "bg-green-500 text-white" : "bg-gray-200"
              } p-2 rounded-md`}
              data-id={msg.id}
            >
              {msg.media_url && msg.mime_type?.includes("video") && (
                <video
                  controls
                  src={msg.media_url}
                  className={`rounded-md mb-2 ${
                    msg.is_from_me ? "ml-auto" : "mr-auto"
                  } w-[300px] h-[300px] object-cover`}
                />
              )}
              {msg.media_url && msg.mime_type?.includes("audio") && (
                <audio
                  controls
                  src={msg.media_url}
                  className={`rounded-md mb-2 ${
                    msg.is_from_me ? "ml-auto" : "mr-auto"
                  } w-[300px]`}
                />
              )}

              {msg.media_url && msg.mime_type?.includes("image") && (
                <Popover
                  placement="bottom"
                  content={
                    <div className="bg-white p-4 rounded-md w-[600px]">
                      <img
                        src={msg.media_url}
                        alt=""
                        className="w-full h-full object-cover rounded-md"
                      />
                    </div>
                  }
                >
                  <img
                    src={msg.media_url}
                    alt=""
                    className={` rounded-md mb-2 ${
                      msg.is_from_me ? "ml-auto" : "mr-auto"
                    } w-[300px] h-[300px] object-cover`}
                  />
                </Popover>
              )}
              {!msg.is_from_me && !msg.is_group && (
                <small className="font-semibold">{msg.contact?.name}</small>
              )}
              {msg.is_from_me && (
                <small className="font-semibold">
                  {msg.member?.user?.full_name}
                </small>
              )}
              {msg.is_group && !msg.is_from_me && (
                <small className="font-semibold">
                  {msg.message_info?.PushName}
                </small>
              )}
              {msg.media_url &&
                !msg.mime_type?.includes("image") &&
                !msg.mime_type?.includes("video") &&
                !msg.mime_type?.includes("audio") && (
                  <div
                    className="flex items-center cursor-pointer"
                    onClick={() => {
                      const url = msg.media_url;
                      window.open(url, "_blank");
                    }}
                  >
                    <RiAttachment2 /> File Attachment
                  </div>
                )}
              <Markdown remarkPlugins={[remarkGfm]}>{msg.message}</Markdown>
              <div className="text-[10px] justify-between flex items-center">
                {msg.sent_at && <Moment fromNow>{msg.sent_at}</Moment>}
                {msg.is_read && (
                  <IoCheckmarkDone
                    size={16}
                    style={{
                      color: msg.is_from_me ? "#0e9f6e" : "white",
                    }}
                  />
                )}
              </div>
            </div>
          </div>
        ))}
      </div>
      {(files.length > 0 || selectedProducts.length > 0) && (
        <div className="absolute bottom-[50px] flex w-full bg-red-50 p-4 justify-between z-0">
          <div className="flex flex-col">
            {files.length > 0 && <span>{files.length} Attachments</span>}
            {selectedProducts.length > 0 && (
              <span>{selectedProducts.length} Products</span>
            )}
          </div>
          <button
            className="text-gray-400 hover:text-gray-600 cursor-pointer"
            onClick={() => {
              setFiles([]);
              setSelectedProducts([]);
            }}
          >
            <FaXmark />
          </button>
        </div>
      )}
      <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2">
        <Dropdown
          label={<BsPlusCircle />}
          inline
          placement="top"
          arrowIcon={false}
        >
          <Dropdown.Item
            className="flex gap-2"
            onClick={() => {
              fileRef.current?.click();
            }}
            icon={BsFileEarmark}
          >
            File
          </Dropdown.Item>
          <Dropdown.Item
            className="flex gap-2"
            onClick={() => {
              getProducts({ page: 1, size: 10 }).then((res: any) => {
                setProducts(res.data.items);
              });
              setModalProduct(true);
            }}
            icon={BsTag}
          >
            Product
          </Dropdown.Item>
        </Dropdown>

        <div className="relative w-full">
          <MentionsInput
            disabled={!session?.is_human_agent && connection?.is_auto_pilot}
            value={content}
            onChange={(val: any) => {
              setContent(val.target.value);
            }}
            style={emojiStyle}
            placeholder={
              !session?.is_human_agent && connection?.is_auto_pilot
                ? "Input disabled for auto pilot mode"
                : "Press ':' for emojis, '/' for templates and shift+enter for new line"
            }
            className="w-full pl-8"
            autoFocus
            onKeyDown={async (val: any) => {
              if (val.key === "Enter" && val.shiftKey) {
                setContent((prev) => prev + "\n");
                return;
              }
              if (val.key === "Enter") {
                sendMessage();
                return;
              }
            }}
          >
            <Mention
              trigger=":"
              markup="__id__"
              regex={neverMatchingRegex}
              data={queryEmojis}
            />
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
              trigger="#"
              data={products.map((p: any) => ({
                id: p.id,
                display: p.display_name,
              }))}
              onAdd={(e: any) => {
                // console.log(products.find((p: any) => p.id === e));
                // let selected = products.find((p: any) => p.id === e);
                // if (selected && (selected?.product_images ?? []).length) {
                //   let file: FileModel = {
                //     file_name: selected.display_name!,
                //     url: selected.product_images![0].url,
                //     mime_type: selected.product_images![0].mime_type,
                //     path: selected.product_images![0].path,
                //   };
                //   setFiles([file]);
                //   setIsCaption(true);
                // }
              }}
              markup="*__display__*"
              regex={neverMatchingRegex}
            />
            <Mention
              trigger="/"
              data={templates.map((t: any) => ({
                id: t.id,
                display: t.title,
              }))}
              appendSpaceOnAdd
              onAdd={(e: any) => {
                console.log(e);
              }}
              markup="@@[__display__](__id__)"
            />
          </MentionsInput>

          <div
            className="absolute top-2 right-2 cursor-pointer"
            onClick={() => setModalEmojis(true)}
          >
            😀
          </div>
        </div>
        <Button
          color="gray"
          onClick={() => {
            setModalTemplates(true);
          }}
        >
          <TbTemplate />
        </Button>
        <Button color="gray" onClick={sendMessage}>
          <BsSend />
        </Button>
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
              ([category, emojis], i) => (
                <div className="mb-4 hover:bg-gray-100 rounded-lg p-2" key={i}>
                  <h3 className="font-bold">{category}</h3>
                  <div className=" flex flex-wrap gap-1">
                    {emojis.map((e: any, index: number) => (
                      <div
                        key={index}
                        className="cursor-pointer text-lg"
                        onClick={() => setContent((prev) => prev + e.emoji)}
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
      <Modal
        show={modalTemplates}
        onClose={() => setModalTemplates(false)}
        dismissible
      >
        <Modal.Header>Templates</Modal.Header>
        <Modal.Body>
          <Table striped>
            <Table.Head>
              <Table.HeadCell>Title</Table.HeadCell>
              <Table.HeadCell>Description</Table.HeadCell>
              <Table.HeadCell>Message</Table.HeadCell>
            </Table.Head>
            <Table.Body className="bg-white">
              {templates.length === 0 && (
                <Table.Row>
                  <Table.Cell colSpan={3} className="text-center">
                    No template found.
                  </Table.Cell>
                </Table.Row>
              )}
              {templates.map((template) => (
                <Table.Row
                  key={template.id}
                  className="bg-white dark:border-gray-700 dark:bg-gray-800 cursor-pointer hover:bg-gray-100"
                  onClick={async () => {
                    let content = `@@[${template.title}](${template.id})`;
                    setModalTemplates(false);
                    await createWAMessage(sessionId!, {
                      message: content,
                      files: files,
                    });
                  }}
                >
                  <Table.Cell>
                    <span className="font-medium">{template.title}</span>
                  </Table.Cell>
                  <Table.Cell>
                    <span className="font-medium">{template.description}</span>
                  </Table.Cell>
                  <Table.Cell>
                    {(template.messages ?? []).map(
                      (message: MessageTemplate, index: number) => (
                        <div key={index} className="mb-2">
                          <h3 className="font-semibold">#Msg {index + 1}</h3>
                          <div>{message.body}</div>
                        </div>
                      )
                    )}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        </Modal.Body>
      </Modal>
      <input
        multiple
        accept=".png, .jpg, .jpeg, .doc, .docx, .xls, .xlsx, .pdf"
        type="file"
        name="file"
        id=""
        ref={fileRef}
        className="hidden"
        onChange={async (e) => {
          if ((e.target.files ?? []).length > 0) {
            for (
              let index = 0;
              index < (e.target.files ?? []).length;
              index++
            ) {
              const element = (e.target.files ?? [])[index];
              let resp: any = await uploadFile(element, {}, console.log);
              setFiles((prev) => [...prev, resp.data]);
            }
          }
        }}
      />
      <Modal show={modalProduct} onClose={() => setModalProduct(false)}>
        <Modal.Header>Product</Modal.Header>
        <Modal.Body>
          <div className="relative w-full mb-8 mr-6 focus-within:text-purple-500">
            <div className="absolute inset-y-0 left-0 flex items-center pl-3">
              <HiMagnifyingGlass />
            </div>
            <input
              type="text"
              className="w-full py-2 pl-10 text-sm text-gray-700 bg-white border border-gray-300 rounded-2xl shadow-sm focus:outline-none focus:ring focus:ring-indigo-200 focus:border-indigo-500"
              placeholder="Search"
              onChange={(e) => {
                getProducts({
                  page: 1,
                  search: e.target.value,
                  size: 10,
                }).then((res: any) => {
                  setProducts(res.data.items);
                });
              }}
            />
          </div>
          {products.length === 0 && (
            <div className="text-center">No product found.</div>
          )}
          <div className="flex flex-col gap-2">
            {products.map((product) => (
              <div
                key={product.id}
                className="flex flex-row gap-2 items-center cursor-pointer hover:bg-gray-100 p-2"
                onClick={() => {
                  setSelectedProducts((prev) => [...prev, product]);
                  setModalProduct(false);
                }}
              >
                {" "}
                {(product.product_images??[]).length !== 0 ? (
                  <img
                    src={product.product_images![0].url}
                    className="w-10 h-10 rounded-full"
                  />
                ) : (
                  <div className="rounded-full w-10 h-10 bg-gray-200 flex justify-center items-center">
                    <BsImage className="w-4 h-4 " />
                  </div>
                )}
                <div className="flex flex-col">
                  <span className="font-semibold">{product.name}</span>
                  <small>{product.description}</small>
                </div>
              </div>
            ))}
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};
export default WhatsappMessages;
