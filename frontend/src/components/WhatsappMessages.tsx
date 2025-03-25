import { Button, Popover, ToggleSwitch } from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
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
  updateWhatsappSession,
} from "../services/api/whatsappApi";

interface WhatsappMessagesProps {
  //   session: WhatsappMessageSessionModel;
  sessionId: string;
}

const WhatsappMessages: FC<WhatsappMessagesProps> = ({ sessionId }) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(100);
  const [search, setSearch] = useState("");
  const [mounted, setMounted] = useState(false);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);
  const [session, setSession] = useState<WhatsappMessageSessionModel>();
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);
  const [files, setFiles] = useState<FileModel[]>([]);
  const [openAttachment, setOpenAttachment] = useState(false);
  const [connection, setConnection] = useState<ConnectionModel>();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getWhatsappSessionDetail(sessionId).then((res: any) => {
        setSession(res.data);
        setConnection(res.connection);
      });
    }
  }, [mounted, sessionId]);

  useEffect(() => {
    getWhatsappMessages(sessionId, {
      page,
      size,
      search,
    }).then((res: any) => {
      setMessages(res.data.items);
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
  }, [wsMsg, profile, sessionId]);
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
            <Moment className="text-xs" fromNow>
              {session?.last_online_at}
            </Moment>
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
      >
        {messages.map((msg) => (
          <div
            key={msg.id}
            className={`flex flex-row items-end mb-2  ${
              msg.is_from_me ? "justify-end" : "justify-start"
            }`}
          >
            <div
              className={`min-w-[300px] max-w-[600px] ${
                !msg.is_from_me ? "bg-green-500 text-white" : "bg-gray-200"
              } p-2 rounded-md`}
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
              {!msg.is_from_me && <small>{msg.contact?.name}</small>}
              {msg.is_group && !msg.is_from_me && (
                <small>{msg.message_info?.PushName}</small>
              )}

              <Markdown remarkPlugins={[remarkGfm]}>{msg.message}</Markdown>
              <div className="text-[10px]">
                {msg.sent_at && <Moment fromNow>{msg.sent_at}</Moment>}
              </div>
            </div>
          </div>
        ))}
      </div>
      {files.length > 0 && (
        <div className="absolute bottom-[100px] flex w-full bg-red-50 p-4 z-50">
          {files.length} Attachments
        </div>
      )}
      <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2">
        <MentionsInput
          disabled={!session?.is_human_agent}
          value={content}
          onChange={(val: any) => {
            setContent(val.target.value);
          }}
          style={emojiStyle}
          placeholder={
            !session?.is_human_agent
              ? "Input disabled for auto pilot mode"
              : "Press ':' for emojis and shift+enter to send"
          }
          className="w-full"
          autoFocus
          onKeyDown={async (val: any) => {
            if (val.key === "Enter" && val.shiftKey) {
              try {
                await createWAMessage(sessionId!, {
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
        </MentionsInput>
        <Button color="gray" onClick={() => setOpenAttachment(true)}>
          <RiAttachment2 />
        </Button>
      </div>
    </div>
  );
};
export default WhatsappMessages;
