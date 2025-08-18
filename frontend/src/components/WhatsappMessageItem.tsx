import { useEffect, useRef, useState, type FC } from "react";
import { WhatsappMessageModel } from "../models/whatsapp_message";
import { RiAttachment2, RiReplyFill } from "react-icons/ri";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import Moment from "react-moment";
import { IoCheckmarkDone } from "react-icons/io5";
import { Dropdown, Popover } from "flowbite-react";
import { markAsRead } from "../services/api/whatsappApi";
interface WhatsappMessageItemProps {
  sessionId: string;
  msg: WhatsappMessageModel;
  selectMessage: (msg: WhatsappMessageModel) => void;
  spottedMsgId: string | undefined;
}

const WhatsappMessageItem: FC<WhatsappMessageItemProps> = ({
  sessionId,
  msg,
  selectMessage,
  spottedMsgId,
}) => {
  const [message, setMessage] = useState<WhatsappMessageModel | null>(null);
  const timeout = useRef<number | null>(null);

  useEffect(() => {
    setMessage({ ...msg });
  }, [msg]);

  useEffect(() => {
    if (spottedMsgId) {
      if (message?.id == spottedMsgId) {
        if (!message.is_read && !(message?.is_from_me ?? false)) {
          setMessage((prev) => ({
            ...prev,
            is_read: true,
          }));
        }

        if (timeout.current) {
          window.clearTimeout(timeout.current);
        }
        timeout.current = window.setTimeout(() => {
          if (!message?.is_from_me && !message?.is_read) {
            // console.log(message);
            markAsRead(message!.id!, sessionId);
          }
        }, 500);
      }
    }
  }, [spottedMsgId]);
  if (!message) return null;
  return (
    <div
      key={msg.id}
      className={`group/item message flex flex-row items-end mb-2  ${
        msg.is_from_me ? "justify-end" : "justify-start"
      }`}
      id={msg.id}
    >
      <div
        className={`relative min-w-[300px] max-w-[600px] ${
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
          <small className="font-semibold">{msg.member?.user?.full_name}</small>
        )}
        {msg.is_group && !msg.is_from_me && (
          <small className="font-semibold">{msg.message_info?.PushName}</small>
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
        {msg.quoted_message && (
          <div className="text-sm p-4 rounded-lg bg-[rgb(255,255,255,0.3)]">
            {msg.quoted_message}
          </div>
        )}
        <Markdown remarkPlugins={[remarkGfm]}>{msg.message}</Markdown>
        {(msg.whatsapp_message_reactions ?? []).length > 0 && (
          <div className="flex mt-2 absolute -bottom-3 right-2 flex-row gap-1">
            {(msg.whatsapp_message_reactions ?? []).map((reaction) => (
              <span key={reaction.id} className=" cursor-pointer ">
                {reaction.reaction}
              </span>
            ))}
          </div>
        )}
        <div className="text-[10px] justify-between flex items-center">
          {msg.created_at && <Moment fromNow>{msg.created_at}</Moment>}
          {msg.is_read && (
            <IoCheckmarkDone
              size={16}
              style={{
                color: msg.is_from_me ? "#0e9f6e" : "white",
              }}
            />
          )}
        </div>
        <div className="group/edit invisible group-hover/item:visible absolute top-2 right-2">
          <Dropdown label="" inline className="context-menu">
            <Dropdown.Item
              className="flex flex-row gap-1"
              icon={RiReplyFill}
              onClick={() => {
                selectMessage(msg);
              }}
            >
              Reply
            </Dropdown.Item>
          </Dropdown>
        </div>
      </div>
    </div>
  );
};
export default WhatsappMessageItem;
