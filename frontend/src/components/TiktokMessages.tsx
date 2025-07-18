import { useContext, useEffect, useRef, useState, type FC } from "react";
import { ConnectionModel } from "../models/connection";
import { PaginationResponse } from "../objects/pagination";
import { connect } from "http2";
import {
  getTiktokSessionDetail,
  getTiktokSessionMessages,
  sendTiktokSessionFile,
  sendTiktokSessionMessage,
} from "../services/api/tiktokApi";
import {
  TiktokMessage,
  TiktokMessageSession,
  TiktokParticipant,
  TiktokSessionDetail,
} from "../models/tiktok";
import { Avatar, Button, Modal, Popover } from "flowbite-react";
import { initial } from "../utils/helper";
import Moment from "react-moment";
import { ScrollContext } from "../contexts/ScrollContext";
import { BsChevronDown, BsSend } from "react-icons/bs";
import { Mention, MentionsInput } from "react-mentions";
import { IoAttachOutline } from "react-icons/io5";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import toast from "react-hot-toast";
import { WebsocketContext } from "../contexts/WebsocketContext";
interface TiktokMessagesProps {
  sessionId: string;
  session?: TiktokMessageSession;
  connection: ConnectionModel;
}

const TiktokMessages: FC<TiktokMessagesProps> = ({ sessionId, connection }) => {
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [session, setSession] = useState<TiktokSessionDetail>();
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const [messages, setMessages] = useState<TiktokMessage[]>([]);
  const [showScrollBottom, setShowScrollButton] = useState(false);
  const { scrollPositions, setScrollPositions } = useContext(ScrollContext);
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [modalAttach, setModalAttach] = useState(false);
  const [modalEmojis, setModalEmojis] = useState(false);
  const [selectedCS, setSelectedCS] = useState<TiktokParticipant>();
  const fileRef = useRef<HTMLInputElement>(null);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);

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
    if (!sessionId) return;
    if (wsMsg?.session_id == sessionId && wsMsg?.command == "TIKTOK_RECEIVED") {
      // setMessages([...messages, wsMsg.data]);
      if (wsMsg.data) {
        setMessages([wsMsg.data, ...messages]);
        setTimeout(() => {
          scrollToBottom();
        }, 300);
      }
    }
  }, [wsMsg, sessionId]);

  useEffect(() => {
    setShowScrollButton(
      (chatContainerRef.current?.scrollHeight ?? 0) -
        (chatContainerRef.current?.scrollTop ?? 0) -
        (chatContainerRef.current?.clientHeight ?? 0) >
        50
    );

    return () => {};
  }, [scrollPositions]);

  useEffect(() => {
    if (sessionId && connection) {
      getTiktokSessionDetail(sessionId).then((resp: any) => {
        // console.log(resp.data);
        setSession(resp.data);
      });
      getTiktokSessionMessages(sessionId, {
        page: page,
        size: size,
        search: search,
        connection_session: connection.id,
      }).then((resp: any) => {
        setMessages(resp.data.messages);
        // setPagination(getPagination(resp.data));
      });
    }

    return () => {};
  }, [sessionId, connection]);
  useEffect(() => {
    for (const msg of messages) {
      if (msg.type == "TEXT" && msg.sender?.role == "CUSTOMER_SERVICE") {
        setSelectedCS(msg.sender);
      }
    }
  }, [messages]);

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
    const container = chatContainerRef.current;

    if (container && container.scrollHeight <= container.clientHeight) {
      //   markAllAsRead();
    }

    const handleScrollPosition = () => {
      if (sessionId) {
        setScrollPositions({
          ...scrollPositions,
          [sessionId]: container?.scrollTop ?? 0,
        });
        // const scrollHeight = chatContainerRef.current?.scrollHeight ?? 0;
        // setScrollHeight(scrollHeight);
      }
    };

    container?.addEventListener("scroll", handleScrollPosition);

    // Cleanup on unmount
    return () => {
      container?.removeEventListener("scroll", handleScrollPosition);
    };
  }, [messages, sessionId]); // atau [chatList], tergantung state kamu

  const renderShoutBox = () => (
    <div className="shoutbox border-t pt-2 min-h-[20px] max-h-[150px] px-2  flex justify-between items-start gap-2 absolute bottom-0 left-0 right-0 bg-white ">
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
              val.preventDefault();
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
          // setModalTemplates(true);
          // setModalAttach(true);
          fileRef.current?.click();
        }}
      >
        <IoAttachOutline />
      </Button>
      <Button color="gray" onClick={sendMessage}>
        <BsSend />
      </Button>
    </div>
  );

  const sendMessage = async () => {
    try {
      if (content === "") return;
      if (!content) return;
      setContent("");
      let data = {
        type: "TEXT",
        content: JSON.stringify({
          content,
        }),
        nickname: selectedCS?.nickname,
        role: selectedCS?.role,
        connection_id: connection.id,
      };
      let resp: any = await sendTiktokSessionMessage(sessionId!, data);
      // if (resp.data) {
      //   setMessages([resp.data, ...messages]);
      //   setTimeout(() => {
      //     scrollToBottom();
      //   }, 300);
      // }
    } catch (error) {
      toast.error(`${error}`);
    }
    //   setOpenAttachment(false);
    //   setFiles([]);
    //   setSelectedProducts([]);
  };

  const handleScroll = () => {
    const messageElements = document.querySelectorAll(".message");

    messageElements.forEach((el) => {
      const observer = new IntersectionObserver(
        (entries) => {
          entries.forEach((entry) => {});
        },
        { threshold: 0.3 } // Minimal 50% pesan terlihat
      );

      observer.observe(el);
    });
  };

  const renderMessages = () => {
    return (
      <div
        id="channel-messages"
        className="messages h-[calc(100vh-260px)] overflow-y-auto p-4 bg-gray-50 space-y-8"
        ref={chatContainerRef}
        onScroll={handleScroll}
      >
        {messages
          .slice()
          .reverse()
          .map((msg) => {
            let content = JSON.parse(msg.content);
            if (msg.type == "NOTIFICATION") {
              return (
                content.content && (
                  <div
                    className="flex flex-row items-center justify-center"
                    key={msg.id}
                  >
                    <div className="text-xs text-center bg-white px-4 py-2 rounded w-fit">
                      {content.content}
                    </div>
                  </div>
                )
              );
            }
            return (
              <div
                key={msg.id}
                className={`message flex flex-row items-end mb-2  ${
                  msg?.sender?.role == "CUSTOMER_SERVICE"
                    ? "justify-end"
                    : "justify-start"
                }`}
              >
                <div
                  className={`min-w-[300px] max-w-[600px] ${
                    msg?.sender?.role != "CUSTOMER_SERVICE"
                      ? "bg-green-500 text-white"
                      : "bg-gray-200"
                  } p-2 rounded-md`}
                  data-id={msg.id}
                >
                  {msg.type == "IMAGE" && (
                    <Popover
                      placement="bottom"
                      content={
                        <div className="bg-white p-4 rounded-md w-[600px]">
                          <img
                            src={content.url}
                            alt=""
                            className="w-full h-full object-cover rounded-md"
                          />
                        </div>
                      }
                    >
                      <img
                        src={content.url}
                        alt=""
                        className={` rounded-md mb-2 ${
                          msg?.sender?.role != "CUSTOMER_SERVICE"
                            ? "ml-auto"
                            : "mr-auto"
                        } w-[300px] h-[300px] object-cover`}
                      />
                    </Popover>
                  )}
                  <small className="font-semibold">
                    {msg.sender?.nickname}
                  </small>
                  <Markdown remarkPlugins={[remarkGfm]}>
                    {content.content}
                  </Markdown>
                  <div className="text-[10px] justify-between flex items-center">
                    {msg.create_time && (
                      <Moment fromNow>{msg.create_time * 1000}</Moment>
                    )}
                  </div>
                </div>
              </div>
            );
          })}
      </div>
    );
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

  return (
    <div className="flex flex-col h-full ">
      <div className="shoutbox border-b py-2 min-h-[40px] flex justify-between items-center">
        <div className="flex gap-2 items-center px-4">
          <Avatar
            size="md"
            img={session?.participant?.avatar}
            rounded
            stacked
            placeholderInitials={initial(session?.session_name)}
            className="cursor-pointer mt-2"
            // onClick={() => nav("/profile")}
          />
          <div className="flex flex-col">
            <span className="font-semibold">{session?.session_name}</span>
            <div className="flex justify-between">
              <Moment className="text-xs" fromNow>
                {session?.last_msg_time}
              </Moment>
            </div>
          </div>
        </div>
      </div>
      {renderMessages()}
      {showScrollBottom && (
        <button
          className="fixed bottom-[60px] right-6 p-2 rounded-full bg-gray-700 hover:bg-gray-300 z-50 text-white transition-colors"
          onClick={() => {
            chatContainerRef.current?.scrollTo({
              top: chatContainerRef?.current?.scrollHeight ?? 0,
              behavior: "smooth",
            });
          }}
        >
          <BsChevronDown />
        </button>
      )}
      {renderShoutBox()}
      <input
        accept=".png, .jpg, .jpeg, .doc"
        type="file"
        name="file"
        id=""
        ref={fileRef}
        className="hidden"
        onChange={async (e) => {
          if ((e.target.files ?? []).length > 0) {
            try {
              const resp: any = await sendTiktokSessionFile(
                sessionId!,
                connection.id!,
                e.target.files![0]
              );
              // console.log(resp);
              let data = {
                type: "IMAGE",
                content: JSON.stringify(resp.data),
                nickname: selectedCS?.nickname,
                role: selectedCS?.role,
                connection_id: connection.id,
              };
              await sendTiktokSessionMessage(sessionId!, data);
            } catch (err) {
              console.log(err);
            } finally {
              if (fileRef.current) {
                fileRef.current.value = "";
              }
            }
          }
        }}
      />
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
export default TiktokMessages;
