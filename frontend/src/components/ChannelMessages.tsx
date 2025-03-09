import { useContext, useEffect, useState, type FC } from "react";
import {
  createMessage,
  getChannelDetail,
  getChannelMessages,
} from "../services/api/chatApi";
import toast from "react-hot-toast";
import { ChatChannelModel, ChatMessageModel } from "../models/chat";
import { Mention, MentionsInput } from "react-mentions";
import { Console } from "console";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import { HiOutlineUser, HiOutlineUserGroup } from "react-icons/hi";
import {
  Avatar,
  Button,
  FileInput,
  Label,
  Modal,
  Popover,
  TextInput,
} from "flowbite-react";
import MemberSelectModal from "./MemberSelectModal";
import MemberChatModal from "./MemberChatModal";
import { MemberModel } from "../models/member";
import { initial } from "../utils/helper";
import Moment from "react-moment";
import { parseMentions } from "../utils/helper-ui";
import { FileModel } from "../models/file";
import { MdAttachFile } from "react-icons/md";
import { RiAttachment2 } from "react-icons/ri";
import { uploadFile } from "../services/api/commonApi";
import { BsFile } from "react-icons/bs";

interface ChannelMessagesProps {
  channelId: string;
}

const ChannelMessages: FC<ChannelMessagesProps> = ({ channelId }) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(100);
  const [search, setSearch] = useState("");
  const [mounted, setMounted] = useState(false);
  const [messages, setMessages] = useState<ChatMessageModel[]>([]);
  const [channel, setChannel] = useState<ChatChannelModel>();
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);
  const [participants, setParticipants] = useState<MemberModel[]>([]);
  const [files, setFiles] = useState<FileModel[]>([]);
  const [openAttachment, setOpenAttachment] = useState(false);

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

  useEffect(() => {
    if (!channelId) return;
    if (wsMsg?.channel_id == channelId && wsMsg?.command == "RECEIVE_MESSAGE") {
      setMessages([...messages, wsMsg.data]);
      setTimeout(() => {
        scrollToBottom()
      }, 300);
      if (wsMsg.sender_id != profile?.id) {
        toast.success(
          <div className="flex flex-col">
            <span className="text-sm font-bold">
              {" "}
              {wsMsg.sender_name} ({wsMsg.channel_name}){" "}
            </span>
            <span className="text-sm"> {wsMsg.data.message} </span>
          </div>,
          {}
        );
      }
    }
  }, [wsMsg, profile, channelId]);

  useEffect(() => {
    setMounted(true);
  }, []);
  useEffect(() => {
    if (channelId && mounted) {
      getChannelDetail(channelId)
        .then((resp: any) => setChannel(resp.data))
        .catch(toast.error);
      getMessages();
    }
  }, [mounted, channelId, page, size, search]);

  useEffect(() => {
    if (!channel) return;
    if (channel.participant_members) {
      setParticipants(channel.participant_members);
    }
  }, [channel]);

  const getMessages = () => {
    getChannelMessages(channelId, { page, size, search })
      .then((resp: any) => {
        resp.data.items.reverse();
        setMessages(resp.data.items);
      })
      .catch(toast.error);
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

  return (
    <div className="flex flex-col h-full ">
      <div className="shoutbox border-b py-2 min-h-[40px] flex justify-between items-center">
        <div className="flex gap-2 items-center px-4">
          {channel?.avatar && (
            <img
              src={channel?.avatar.url}
              className=" aspect-square rounded-full object-cover w-8 h-8"
            />
          )}
          <div className="flex flex-col">
            <span className="font-semibold">{channel?.name}</span>
            <small className="">{channel?.description}</small>
          </div>
        </div>
        <HiOutlineUserGroup
          className=" text-gray-300 hover:text-gray-600 cursor-pointer"
          size={24}
          onClick={openModal}
        />
      </div>
      <div id="channel-messages" className="messages h-[calc(100vh-260px)] overflow-y-auto p-4 bg-gray-50 ">
        {(messages ?? []).map((message) => {
          let isMe = message.sender_member?.user?.id == profile?.id;
          return (
            <div
              key={message.id}
              className={`flex items-start gap-2 mb-2 ${
                isMe ? "justify-end" : ""
              } md w-full`}
            >
              <div className="w-fit max-w-[80%] min-w-[300px] bg-white p-4 rounded-xl hover:shadow-md flex flex-row gap-2 items-start">
                <Avatar
                  size="xs"
                  img={message.sender_member?.user?.profile_picture?.url}
                  rounded
                  stacked
                  placeholderInitials={initial(
                    message.sender_member?.user?.full_name
                  )}
                />
                <div className="flex flex-col">
                  <span className="font-semibold -mt-1">
                    {message.sender_member?.user?.full_name}
                  </span>
                  <small>
                    <Moment fromNow>{message.date}</Moment>
                  </small>
                  {(message.files ?? []).length > 0 && (
                    <div className="flex gap-2 cursor-pointer items-center mt-4">
                      {(message.files ?? []).map((file) => {
                        if (file.mime_type.includes("image")) {
                          return (
                            <Popover
                              trigger="hover"
                              content={<div className=" aspect-square rounded-full object-cover w-80 h-80"><img src={file.url} /></div>}
                            >
                              
                              <img
                                src={file.url}
                                className=" aspect-square rounded-lg object-cover w-24 h-24"
                                onClick={() => window.open(file.url)}
                              />
                              
                            </Popover>
                          );
                        }
                        return (
                          <div className=" aspect-square rounded-lg object-cover w-24 h-24 flex justify-center items-center bg-gray-200" onClick={() => window.open(file.url)}>
                            <RiAttachment2 size={24} />
                          </div>
                        );
                      })}
                    </div>
                  )}
                  <span className="mt-8">
                    {" "}
                    {parseMentions(message.message ?? "", (type, id) => {
                      console.log(type, id);
                    })}
                  </span>
                </div>
              </div>
            </div>
          );
        })}
      </div>
      {files.length > 0 && (
        <div className="absolute bottom-[100px] flex w-full bg-red-50 p-4 z-50">
          {files.length} Attachments
        </div>
      )}
      <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2">
        <MentionsInput
          value={content}
          onChange={(val: any) => {
            setContent(val.target.value);
          }}
          style={emojiStyle}
          placeholder={
            "Press ':' for emojis, mention people using '@' and shift+enter to send"
          }
          className="w-full"
          autoFocus
          onKeyDown={async (val: any) => {
            if (val.key === "Enter" && val.shiftKey) {
              try {
                await createMessage(channelId!, {
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
            trigger="@"
            data={(channel?.participant_members ?? []).map((member) => ({
              id: member.id!,
              display: member.user?.full_name!,
            }))}
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
        <Button color="gray" onClick={() => setOpenAttachment(true)}>
          <RiAttachment2 />
        </Button>
      </div>

      <Modal size="4xl" show={showModal} onClose={closeModal}>
        <Modal.Header>Add Member</Modal.Header>
        <Modal.Body>
          <MemberChatModal
            participants={participants}
            channel={channel!}
            onInvite={() => {
              if (channelId) {
                setParticipants([]);
                getChannelDetail(channelId)
                  .then((resp: any) => setChannel(resp.data))
                  .catch(toast.error);
              }
              // closeModal();
              // setInviteEmail(val);
              // setInviteModal(true);
            }}
            // onClose={closeModal}
          />
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button color="gray" onClick={closeModal}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal show={openAttachment} onClose={() => setOpenAttachment(false)}>
        <Modal.Header>Attachment</Modal.Header>
        <Modal.Body>
          <div className="flex w-full items-center justify-center">
            <Label
              htmlFor="dropzone-file"
              className="flex h-64 w-full cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed border-gray-300 bg-gray-50 hover:bg-gray-100 dark:border-gray-600 dark:bg-gray-700 dark:hover:border-gray-500 dark:hover:bg-gray-600"
            >
              <div className="flex flex-col items-center justify-center pb-6 pt-5">
                <svg
                  className="mb-4 h-8 w-8 text-gray-500 dark:text-gray-400"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 20 16"
                >
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
                  />
                </svg>
                <p className="mb-2 text-sm text-gray-500 dark:text-gray-400">
                  <span className="font-semibold">Click to upload</span> or drag
                  and drop
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-400">
                  Upload Image/Document
                </p>
              </div>
              <FileInput
                multiple
                id="dropzone-file"
                className="hidden"
                accept=".jpg, .jpeg, .png, .gif, .bmp, .doc, .docx, .xls, .xlsx, .pdf"
                onChange={(val) => {
                  const files = val?.target.files;
                  if (files) {
                    for (let index = 0; index < files.length; index++) {
                      const file = files[index];
                      try {
                        uploadFile(file, {}, (val) => console.log).then(
                          (v: any) => {
                            setFiles((files) => [...files, v.data]);
                          }
                        );
                      } catch (error) {
                        console.log(error);
                      }
                    }
                    setOpenAttachment(false);
                  }
                }}
              />
            </Label>
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};
export default ChannelMessages;
