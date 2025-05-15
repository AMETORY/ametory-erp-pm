import { useContext, useEffect, useRef, useState, type FC } from "react";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import {
  TelegramMessage,
  TelegramMessageSessionModel,
} from "../models/telegram";
import { FileModel } from "../models/file";
import { ConnectionModel } from "../models/connection";
import { TemplateModel } from "../models/template";
import {
  createTelegramMessage,
  getTelegramMessages,
  getTelegramSessionDetail,
} from "../services/api/telegramApi";
import { Button, Modal, Popover } from "flowbite-react";
import Markdown from "react-markdown";
import { IoCheckmarkDone } from "react-icons/io5";
import Moment from "react-moment";
import remarkGfm from "remark-gfm";
import { Mention, MentionsInput } from "react-mentions";
import { TbTemplate } from "react-icons/tb";
import { BsSend } from "react-icons/bs";
import toast from "react-hot-toast";
interface TelegramMessagesProps {
  sessionId: string;
}

//TELEGRAM_RECEIVED
const TelegramMessages: FC<TelegramMessagesProps> = ({ sessionId }) => {
  const timeout = useRef<number | null>(null);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [page] = useState(1);
  const [size, setSize] = useState(100);
  const [search, setSearch] = useState("");
  const [mounted, setMounted] = useState(false);
  const [messages, setMessages] = useState<TelegramMessage[]>([]);
  const [session, setSession] = useState<TelegramMessageSessionModel>();
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);
  const [files, setFiles] = useState<FileModel[]>([]);
  const [openAttachment, setOpenAttachment] = useState(false);
  const [connection, setConnection] = useState<ConnectionModel>();
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const [modalEmojis, setModalEmojis] = useState(false);
  const [templates, setTemplates] = useState<TemplateModel[]>([]);
  const [modalTemplates, setModalTemplates] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);

  useEffect(() => {
    if (mounted) {
      getTelegramSessionDetail(sessionId)
        .then((res: any) => {
          setSession(res.data);
          setConnection(res.connection);
        })
        .catch((err) => {
          // toast.error(`${err}`);
          window.location.href = "/whatsapp";
        })
        .finally(() => {});

      //   getAllTemplates();
    }
  }, [mounted, sessionId]);

  useEffect(() => {
    getTelegramMessages(sessionId, {
      page,
      size,
      search,
    })
      .then((res: any) => {
        setMessages(res.data.items);
      })
      .catch((err) => {
        console.error(err);
        //   window.location.href = "/whatsapp";
      });
  }, [session, sessionId]);

  const sendMessage = async () => {
    try {
      if (!content) return;
      setContent("");
      await createTelegramMessage(sessionId!, {
        message: content,
        files: files,
      });
      setOpenAttachment(false);
      setFiles([]);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setTimeout(() => {
        setContent("");
      }, 300);
    }
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
    if (!sessionId) return;
    if (
      wsMsg?.session_id == sessionId &&
      wsMsg?.command == "TELEGRAM_RECEIVED"
    ) {
      setMessages([...messages, wsMsg.data]);
      setTimeout(() => {
        scrollToBottom();
        setSession({
          ...session!,
          last_online_at: new Date(),
        });
      }, 300);
    }

    if (
      wsMsg?.session_id == sessionId &&
      wsMsg?.command == "TELEGRAM_RECEIVED_READ"
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
      padding: 9,
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
                    // markAsRead(message!.id!);
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

  const renderMessages = () => (
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
            {!msg.is_from_me && (
              <small className="font-semibold">{msg.contact?.name}</small>
            )}
            {msg.is_from_me && (
              <small className="font-semibold">
                {msg.member?.user?.full_name}
              </small>
            )}
            {/* {msg.is_group && !msg.is_from_me && (
                <small className="font-semibold">
                  {msg.message_info?.PushName}
                </small>
              )} */}
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
  );
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
      </div>
      {renderMessages()}
      {files.length > 0 && (
        <div className="absolute bottom-[100px] flex w-full bg-red-50 p-4 z-50">
          {files.length} Attachments
        </div>
      )}
      <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2">
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
            className="w-full"
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
    </div>
  );
};
export default TelegramMessages;
