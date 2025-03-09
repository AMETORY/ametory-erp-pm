import { useContext, useEffect, useState, type FC } from "react";
import {
  createMessage,
  getChannelDetail,
  getChannelMessages,
} from "../services/api/chatApi";
import toast from "react-hot-toast";
import { ChatChannelModel } from "../models/chat";
import { Mention, MentionsInput } from "react-mentions";
import { Console } from "console";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import { HiOutlineUser, HiOutlineUserGroup } from "react-icons/hi";
import { Button, Modal, TextInput } from "flowbite-react";
import MemberSelectModal from "./MemberSelectModal";
import MemberChatModal from "./MemberChat";

interface ChannelMessagesProps {
  channelId: string;
}

const ChannelMessages: FC<ChannelMessagesProps> = ({ channelId }) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [search, setSearch] = useState("");
  const [mounted, setMounted] = useState(false);
  const [messages, setMessages] = useState();
  const [channel, setChannel] = useState<ChatChannelModel>();
  const [content, setContent] = useState("");
  const [emojis, setEmojis] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

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
    if (wsMsg?.channel_id == channelId) {
      if (wsMsg.sender_id != profile?.id) {
        toast.success(wsMsg.data.message);
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

  const getMessages = () => {
    getChannelMessages(channelId, { page, size, search })
      .then((resp: any) => {
        setMessages(resp.data.items);
      })
      .catch(toast.error);
  };
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
    <div className="flex flex-col h-full">
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
      <div className="messages flex-grow overflow-y-auto p-4 bg-gray-50">
        {/* Chat messages go here */}
      </div>
      <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2">
        <MentionsInput
          value={content}
          onChange={(val: any) => {
            setContent(val.target.value);
          }}
          style={emojiStyle}
          placeholder={
            "Press ':' for emojis, mention people using '@' and shift+enter to send"
          }
          autoFocus
          onKeyDown={async (val: any) => {
            if (val.key === "Enter" && val.shiftKey) {
              try {
                await createMessage(channelId!, {
                  message: content,
                });
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
      </div>

      <Modal size="4xl" show={showModal} onClose={closeModal}>
        <Modal.Header>Add Member</Modal.Header>
        <Modal.Body>
          <MemberChatModal
            channel={channel!}
            onInvite={(val) => {
              closeModal();
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
    </div>
  );
};
export default ChannelMessages;
