import { useContext, useEffect, useState, type FC } from "react";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import { FileModel } from "../models/file";
import {
  WhatsappMessageModel,
  WhatsappMessageSessionModel,
} from "../models/whatsapp_message";
import Moment from "react-moment";
import { HiOutlineUserGroup } from "react-icons/hi";
import {
  getWhatsappMessages,
  getWhatsappSessionDetail,
} from "../services/api/whatsappApi";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";

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

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getWhatsappSessionDetail(sessionId).then((res: any) => {
        setSession(res.data);
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
      </div>
      <div
        id="channel-messages"
        className="messages h-[calc(100vh-260px)] overflow-y-auto p-4 bg-gray-50 "
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
                <img
                  src={msg.media_url}
                  alt=""
                  className={` rounded-md mb-2 ${
                    msg.is_from_me ? "ml-auto" : "mr-auto"
                  } w-[300px] h-[300px] object-cover`}
                />
              )}
              {msg.is_group && !msg.is_from_me && (
                <small>{msg.message_info?.PushName}</small>
              )}

              <Markdown remarkPlugins={[remarkGfm]}>{msg.message}</Markdown>
              <div className="text-[10px]">
                {msg.sent_at && (
                  <Moment fromNow>{msg.sent_at}</Moment>
                )}
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
export default WhatsappMessages;
