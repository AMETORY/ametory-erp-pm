import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Button, FileInput, Modal, Textarea, TextInput } from "flowbite-react";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { createChannel, getChannels } from "../services/api/chatApi";
import { FileModel } from "../models/file";
import { uploadFile } from "../services/api/commonApi";
import { PaginationResponse } from "../objects/pagination";
import { ChatChannelModel } from "../models/chat";
import { getPagination } from "../utils/helper";
import { useNavigate, useParams } from "react-router-dom";
import ChannelMessages from "../components/ChannelMessages";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_DEFAULT_CHANNEL } from "../utils/constants";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { MemberContext, ProfileContext } from "../contexts/ProfileContext";

interface ChatPageProps {}

const ChatPage: FC<ChatPageProps> = ({}) => {
  const [openChannelForm, setOpenChannelForm] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const { channelId } = useParams();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [file, setFile] = useState<FileModel>();
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [channels, setChannels] = useState<ChatChannelModel[]>([]);
  const nav = useNavigate();
  const { profile, setProfile } = useContext(ProfileContext);
  const { member, setMember } = useContext(MemberContext);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (!channelId) return;
    if (wsMsg?.command == "CHANNEL_RELOAD" && member?.id == wsMsg.recipient_id) {
      getAllChannels();
    }
  }, [wsMsg, profile, member, channelId]);

  useEffect(() => {
    if (mounted) {
      getAllChannels();
    }
  }, [mounted, page, size, search]);
  const processNew = async () => {
    try {
      setLoading(true);
      if (name.length == 0) {
        toast.error("Name is required");
        return;
      }

      let data: any = {
        name,
        description,
      };

      if (file) {
        data = { ...data, avatar: file };
      }
      await createChannel(data);
      setOpenChannelForm(false);
      getAllChannels();
    } catch (e) {
      toast.error(`${e}`);
    } finally {
      setLoading(false);
    }
  };

  const getAllChannels = async () => {
    try {
      setLoading(true);
      const resp: any = await getChannels({ page, size, search });
      setChannels(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (e) {
      toast.error(`${e}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Chat</h1>
          <div className="flex gap-2">
            <Button
              gradientDuoTone="purpleToBlue"
              pill
              onClick={() => {
                setOpenChannelForm(true);
              }}
            >
              + Channel
            </Button>
          </div>
        </div>
        <div className="flex flex-row w-full h-full flex-1 gap-2">
          <div className="w-[300px] h-full">
            <ul className="space-y-2">
              {channels.map((e) => (
                <li
                  className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                  key={e.id}
                  onClick={() => {
                    nav(`/chat/${e.id}`);
                    asyncStorage.setItem(LOCAL_STORAGE_DEFAULT_CHANNEL, e.id);
                  }}
                  style={{ background: channelId == e.id ? "#e5e7eb" : "" }}
                >
                  <div className="flex gap-2 items-center">
                    {e.avatar && (
                      <img
                        src={e.avatar.url}
                        className=" aspect-square rounded-full object-cover w-8 h-8"
                      />
                    )}
                    <div className="flex flex-col">
                      <span className="font-semibold">{e.name}</span>
                      <small className="">{e.description}</small>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          </div>
          <div className="w-full border-l relative">
            {channelId && <ChannelMessages channelId={channelId} />}
          </div>
        </div>
      </div>
      <Modal show={openChannelForm} onClose={() => setOpenChannelForm(false)}>
        <Modal.Header>New Channel</Modal.Header>
        <Modal.Body>
          <form className="flex flex-col space-y-4">
            <div>
              <label
                htmlFor="name"
                className="block text-sm font-medium text-gray-700"
              >
                Icon
              </label>
              <div className="relative">
                <FileInput
                  id="small-file-upload"
                  sizing="sm"
                  accept="image/*"
                  onChange={(el) => {
                    if (el.target.files) {
                      let f = el.target.files[0];
                      uploadFile(f, {}, (val) => {
                        console.log(val);
                      }).then((v: any) => {
                        setFile(v.data);
                      });
                    }
                  }}
                />

                {file && (
                  <div className="flex justify-center p-4">
                    {" "}
                    <img
                      src={file.url}
                      className="w-32 aspect-square bg-cover"
                    />
                  </div>
                )}
              </div>
            </div>
            <div>
              <label
                htmlFor="name"
                className="block text-sm font-medium text-gray-700"
              >
                Channel Name
              </label>
              <TextInput
                id="name"
                type="name"
                placeholder="Channel Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>
            <div>
              <label
                htmlFor="description"
                className="block text-sm font-medium text-gray-700"
              >
                Description
              </label>

              <Textarea
                id="description"
                className="mt-1 block w-full px-3 py-2 text-base text-gray-700 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                placeholder="Channel Description"
                rows={3}
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              />
            </div>
          </form>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button type="submit" color="blue" onClick={processNew}>
            Save
          </Button>
          <Button color="gray" onClick={() => setOpenChannelForm(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default ChatPage;
